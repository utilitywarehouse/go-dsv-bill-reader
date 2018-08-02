package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/utilitywarehouse/go-dsv-bill-reader"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: billdsv [filename]")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	r := billdsv.NewReader(f)
	r.Comma = '|'
	for {
		record, err := r.Read()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(truncateStrings(24, record))
	}
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
