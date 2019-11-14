package billdsv

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

var e [][]byte

func BenchmarkReader(b *testing.B) {
	r := make([]*Reader, b.N)
	for i := 0; i < b.N; i++ {
		r[i] = NewReader(generateCSV(1000000, 134), 134, DefaultBufferSize)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r[i].ReadAll(func(row [][]byte) {
			e = row
		})
	}
}

// returns a reader that produces a CSV of the specified rows and columns filled
// with placeholder data.
func generateCSV(rows int, cols int) (r io.Reader) {
	r, w := io.Pipe()
	go func() {
		for i := 0; i < rows; i++ {
			cells := make([]string, cols)
			for j := 0; j < cols; j++ {
				cells[j] = fmt.Sprintf("field %d", j)
			}
			_, err := w.Write([]byte(strings.Join(cells, "|") + "\n"))
			if err != nil {
				panic(err)
			}
		}
		err := w.Close()
		if err != nil {
			panic(err)
		}
	}()
	return r
}
