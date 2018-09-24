package billdsv

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

var e [][]string

func handler(row []string) {
	e = append(e, row)
}

func BenchmarkReader(b *testing.B) {
	b.ReportAllocs()

	csv := getCsv(10000, 134)
	r := NewReader(csv)
	r.ReuseRecord = true

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 10000; i++ {
			s, err := r.Read()
			if err != nil {
				if err == io.EOF {
					return
				}
				b.Error(err)
			}
			handler(s)
		}
	}
}

func getCsv(rows int, cols int) (r io.Reader) {
	r, w := io.Pipe()
	go func() {
		for i := 0; i < rows; i++ {
			cells := make([]string, cols)
			for j := 0; j < cols; j++ {
				cells[j] = fmt.Sprintf("field %d", j)
			}
			_, err := w.Write([]byte(strings.Join(cells, ",") + "\n"))
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
