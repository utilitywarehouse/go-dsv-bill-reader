package billdsv

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

// A ParseError is returned for parsing errors.
// Line numbers are 1-indexed and columns are 0-indexed.
type ParseError struct {
	StartLine int   // Line   where the record starts
	Line      int   // Line where the error occurred
	Column    int   // Column (rune index) where the error occurred
	Err       error // The actual error
}

func (e *ParseError) Error() string {
	if e.Err == ErrFieldCount {
		return fmt.Sprintf("record on line %d: %v", e.Line, e.Err)
	}
	if e.StartLine != e.Line {
		return fmt.Sprintf("record on line %d; parse error on line %d, column %d: %v", e.StartLine, e.Line, e.Column, e.Err)
	}
	return fmt.Sprintf("parse error on line %d, column %d: %v", e.Line, e.Column, e.Err)
}

// These are the errors that can be returned in ParseError.Err.
var (
	ErrTrailingComma = errors.New("extra delimiter at end of line") // Deprecated: No longer used.
	ErrBareQuote     = errors.New("bare \" in non-quoted-field")
	ErrQuote         = errors.New("extraneous or missing \" in quoted-field")
	ErrFieldCount    = errors.New("wrong number of fields")
)

// Reader implements a DSV reader that reads the pipe separated values
// that Bill outputs.
type Reader struct {
	// Comma is the field delimiter.
	// It is set to comma (',') by NewReader.
	// Comma must be a valid rune and must not be \r, \n,
	// or the Unicode replacement character (0xFFFD).
	Comma rune

	// FieldsPerRecord is the number of expected fields per record.
	// If FieldsPerRecord is positive, Read requires each record to
	// have the given number of fields. If FieldsPerRecord is 0, Read sets it to
	// the number of fields in the first record, so that future records must
	// have the same field count.
	FieldsPerRecord int

	// Comment is unused in this implementation
	Comment rune
	// ReuseRecord is unused in this implementation
	ReuseRecord bool
	// LazyQuotes is unused in this implementation
	LazyQuotes bool
	// TrimLeadingSpace is unused in this implementation
	TrimLeadingSpace bool

	s *bufio.Scanner

	// numLine is the current line being read in the CSV file.
	// nolint
	numLine int

	// rawBuffer is a line buffer only used by the readLine method.
	// nolint
	rawBuffer []byte

	// recordBuffer holds the unescaped fields, one after another.
	// The fields can be accessed by using the indexes in fieldIndexes.
	// E.g., For the row `a,"b","c""d",e`, recordBuffer will contain `abc"de`
	// and fieldIndexes will contain the indexes [1, 2, 5, 6].
	// nolint
	recordBuffer []byte

	// fieldIndexes is an index of fields inside recordBuffer.
	// The i'th field ends at offset fieldIndexes[i] in recordBuffer.
	// nolint
	fieldIndexes []int

	// lastRecord is a record cache and only used when ReuseRecord == true.
	// nolint
	lastRecord []string
}

// NewReader returns a new Reader that reads from r.
func NewReader(r io.Reader) *Reader {
	return &Reader{
		s:     bufio.NewScanner(r),
		Comma: ',',
	}
}

// Read reads one record (a slice of fields) from r.
// If the record has an unexpected number of fields,
// Read returns the record along with the error ErrFieldCount.
// Except for that case, Read always returns either a non-nil
// record or a non-nil error, but not both.
// If there is no data left to be read, Read returns nil, io.EOF.
// If ReuseRecord is true, the returned slice may be shared
// between multiple calls to Read.
func (r *Reader) Read() (record []string, err error) {
	if r.ReuseRecord {
		record, err = r.readRecord(r.lastRecord)
		r.lastRecord = record
	} else {
		record, err = r.readRecord(nil)
	}
	return
}

// ReadAll reads all the remaining records from r.
// Each record is a slice of fields.
// A successful call returns err == nil, not err == io.EOF. Because ReadAll is
// defined to read until EOF, it does not treat end of file as an error to be
// reported.
func (r *Reader) ReadAll() (records [][]string, err error) {
	for {
		record, err := r.readRecord(nil)
		if err == io.EOF {
			return records, nil
		}
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
}

// readRecord reads a single record from the file. If the line it read was not
// long enough to satisfy the FieldsPerRecord value, the next line will be read
// and its contents appended to the end of the last field of the previously read
// line.
// nolint:unparam
func (r *Reader) readRecord(dst []string) ([]string, error) {
	ok := r.s.Scan()
	if !ok {
		return nil, io.EOF
	}

	dst = strings.Split(r.s.Text(), string(r.Comma))

	if r.FieldsPerRecord < 0 {
		r.FieldsPerRecord = len(dst)
	} else {
		for len(dst) < r.FieldsPerRecord {
			r.s.Scan()
			overflow := strings.Split(r.s.Text(), string(r.Comma))
			dst[len(dst)-1] += fmt.Sprintf("\n%s", overflow[0])
			dst = append(dst, overflow[1:]...)
		}
	}

	r.numLine++
	return dst, nil
}
