package billdsv

import (
	"bytes"
	"io"

	"github.com/pkg/errors"
)

// Reader implements a DSV reader that reads the pipe separated values
// that Bill outputs.
type Reader struct {
	Separator   byte
	SkipHeading bool
	BufferSize  int

	r         io.Reader
	fields    int
	rdBuffer  []byte
	wrBuffer  []byte
	rowBuffer [][]byte
}

var defaultBufferSize = 1024

// NewReader returns a new Reader that reads from r. The number of expected
// fields per row can be specified so the parser can safely deal with fields
// containing line breaks. The buffer size may be specified post-instantiate
// but the default should be fine for most cases. If fields is left at zero, the
// first line will be used to set the expected field count for the rest of the
// document. This means that if the CSV is malformed
func NewReader(r io.Reader, fields int) *Reader {
	return &Reader{
		Separator:  '|',
		BufferSize: defaultBufferSize,

		// all allocations happen at this stage, this includes a buffer to
		// stream chunks of data into, a buffer to stage field data into and a
		// buffer of fields to stage rows before calling the callback.
		r:         r,
		fields:    fields,
		rdBuffer:  make([]byte, defaultBufferSize),
		wrBuffer:  make([]byte, 1024),
		rowBuffer: make([][]byte, fields),
	}
}

// ReadAll reads all records and passes them to the specified function. This
// function will make no heap allocations in best case scenarios. The only time
// this function will allocate is if a field exceeds the default field buffer
// size of 1024, in which case the struct field `wrBuffer` will be resized to
// 1.5x the size. The other two potential allocation spots, are the two `append`
// calls in the switch blocks, these are allocated lazily as well as if the
// `rowBuffer` cell is at capacity and requires resizing to fit the new data.
func (r *Reader) ReadAll(function func([][]byte)) (err error) {
	// the buffer size must be able to accommodate at least `n` fields as well as
	// `n-1` field separators.
	if r.fields > r.BufferSize/2 {
		return errors.New("buffer size isn't large enough for the amount of specified fields")
	}

	var (
		rdBufferLen int
		rdIdx       int
		wrIdx       int
		fields      int
		rows        int
	)

	for {
		if rdBufferLen, err = r.r.Read(r.rdBuffer); err == io.EOF {
			if rdBufferLen == 0 {
				break
			}
		} else if err != nil {
			return err
		}
		rdIdx = 0

		if rows == 0 && r.SkipHeading {
			// consume the heading bytes and discard
			// counting the field headings and using the count if necessary
			for ; rdIdx < rdBufferLen; rdIdx++ {
				if r.rdBuffer[rdIdx] == '\n' {
					headings := len(bytes.Split(r.rdBuffer[:rdIdx], []byte{r.Separator}))
					if r.fields == 0 {
						// since the field count was calculated at "runtime"
						// it needs to allocate the row buffer because the
						// NewReader function would have allocated it with 0
						r.fields = headings
						r.rowBuffer = make([][]byte, headings)
					} else {
						if headings != r.fields {
							return errors.New("declared fields does not match headings")
						}
					}
					rows = 1
					rdIdx++
					break
				}
			}
		}

		for ; rdIdx < rdBufferLen; rdIdx++ {
			switch r.rdBuffer[rdIdx] {
			case r.Separator:
				if fields >= len(r.rowBuffer) {
					return errors.Errorf("on row %d, expected %d fields but read an extra field", rows, len(r.rowBuffer))
				}
				r.rowBuffer[fields] = r.rowBuffer[fields][:0]
				r.rowBuffer[fields] = append(r.rowBuffer[fields], r.wrBuffer...)
				r.rowBuffer[fields] = r.rowBuffer[fields][0:wrIdx]
				wrIdx = 0
				fields++

			case '\r':
				continue

			case '\n':
				if fields == r.fields-1 {
					r.rowBuffer[fields] = r.rowBuffer[fields][:0]
					r.rowBuffer[fields] = append(r.rowBuffer[fields], r.wrBuffer...)
					r.rowBuffer[fields] = r.rowBuffer[fields][0:wrIdx]
					wrIdx = 0
					fields = 0

					function(r.rowBuffer)
					rows++
					continue
				}

				fallthrough

			default:
				if wrIdx >= len(r.wrBuffer) {
					r.wrBuffer = append(r.wrBuffer, make([]byte, int(float64(len(r.wrBuffer))*1.5))...)
				}
				r.wrBuffer[wrIdx] = r.rdBuffer[rdIdx]
				wrIdx++
			}
		}
	}

	if (wrIdx != 0 && r.wrBuffer != nil) || (fields == r.fields-1 && wrIdx == 0) {
		r.rowBuffer[fields] = r.rowBuffer[fields][:0]
		r.rowBuffer[fields] = append(r.rowBuffer[fields], r.wrBuffer...)
		r.rowBuffer[fields] = r.rowBuffer[fields][0:wrIdx]

		function(r.rowBuffer)
	}

	return nil
}
