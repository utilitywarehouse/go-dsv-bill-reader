package billdsv

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReader1(t *testing.T) {
	f := strings.NewReader(`1000|first string|final string
1001|second string
that is multi-line
|final string
1002|third string|final string
`)

	want := [][]string{
		{"1000", "first string", "final string"},
		{"1001", "second string\nthat is multi-line\n", "final string"},
		{"1002", "third string", "final string"},
	}

	got := [][]string{}

	cr := NewReader(f)
	cr.Comma = '|'
	cr.FieldsPerRecord = 3
	for {
		row, err := cr.Read()
		if err != nil {
			fmt.Println(err)
			break
		}
		got = append(got, row)
		fmt.Println(truncateStrings(20, row))
	}

	assert.Equal(t, want, got)
}

func TestReader2(t *testing.T) {
	f := strings.NewReader(`1000|first string|final string
1001|second string
that is multi-line|final string
1002|third string|final string
`)

	want := [][]string{
		{"1000", "first string", "final string"},
		{"1001", "second string\nthat is multi-line", "final string"},
		{"1002", "third string", "final string"},
	}

	got := [][]string{}

	cr := NewReader(f)
	cr.Comma = '|'
	cr.FieldsPerRecord = 3
	for {
		row, err := cr.Read()
		if err != nil {
			fmt.Println(err)
			break
		}
		got = append(got, row)
		fmt.Println(truncateStrings(20, row))
	}

	assert.Equal(t, want, got)
}

func truncateStrings(limit int, in []string) string {
	sb := strings.Builder{}
	sb.WriteString("[")
	for i, s := range in {
		s = strings.Replace(s, "\n", `\n`, -1)
		// sb.WriteString("\t")
		sb.WriteString(fmt.Sprintf(`%d:"`, i))
		if len(s) > limit {
			sb.WriteString(s[:limit] + "...")
		} else {
			sb.WriteString(s)
		}
		sb.WriteRune('"')

		if i < len(in)-1 {
			sb.WriteString(", ")
		}

		// sb.WriteString("\n")
	}
	sb.WriteString("]")
	return sb.String()
}
