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

func TestReader3(t *testing.T) {
	f := strings.NewReader(`7860442|1041729918|2017-08-28|GN9100|28-08-2017  - Customer has requested to book later|SYSTEM|General|||Appointment Cancelled (Cust)|16-30|006|Mr S & Mrs A Titterington|07510044409, 01768885022|titteringtonamy@gmail.com||YES||1||1899-12-30||0|2017-11-21|||JCOLEMAN/21-11-2017/13:39:16|1|0
7860462|1041731165|2017-08-28|GN9100|^2Àõ^X@¢Ðr^FJ^E^H<90>N)S   [d1aV<90>]s~l<88>ý>9Á3¿[Y<92><98>a<9c>o¿ð<8b>[Æ[Uæ³^O^T¸d<83>^Z<9a>^N\pÅ­[Z]¢Ggã^H7Ó8<81>-9<85>ûJ<9d>¥<95>\^^MôinRßå<8d>
°I^E¾÷¯¡êÝÝ<8e>.^HUÎ^Q^BD»f^\ä¤C  04-09-2017  - Customer has cancelled their appointment. 04-10-2017 - Install Date - 04-11-2017, Installation Time - 12:00-16:00, Status - Appointment Booked, EST Bulbs - ]<Àö^P 04-10-2017 - EGN9100_Booked email sent 04-10-2017 - Install Date - 04-11-2017 , Installation Time - 12:00-16:00, Status - Appointment Booked, EST Bulbs - ]<Àö^P  19-10-2017  - Customer has cancelled their appointment.|SYSTEM|General|||Appointment Cancelled (UW)|]<Àö^P|\:Ø|$^?<8a>­^@=ñ<94>7T^X[G·|\=Øð^S^¥×p^B^^D^Hé^Qn^QY^E8$^]^NÒ|^D^?<8a>­^RZ Ör^H\^\h±O.J ^d?FXÊM8||YES||1||2017-10-19|JKNIGHT|0|2017-10-19|||JKNIGHT/19-10-2017/11:17:58|1|0
`)

	want := [][]string{
		{
			`7860442`,
			`1041729918`,
			`2017-08-28`,
			`GN9100`,
			`28-08-2017  - Customer has requested to book later`,
			`SYSTEM`,
			`General`,
			``,
			``,
			`Appointment Cancelled (Cust)`,
			`16-30`,
			`006`,
			`Mr S & Mrs A Titterington`,
			`07510044409, 01768885022`,
			`titteringtonamy@gmail.com`,
			``,
			`YES`,
			``,
			`1`,
			``,
			`1899-12-30`,
			``,
			`0`,
			`2017-11-21`,
			``,
			``,
			`JCOLEMAN/21-11-2017/13:39:16`,
			`1`,
			`0`,
		},
		{
			`7860462`,
			`1041731165`,
			`2017-08-28`,
			`GN9100`,
			`^2Àõ^X@¢Ðr^FJ^E^H<90>N)S   [d1aV<90>]s~l<88>ý>9Á3¿[Y<92><98>a<9c>o¿ð<8b>[Æ[Uæ³^O^T¸d<83>^Z<9a>^N\pÅ­[Z]¢Ggã^H7Ó8<81>-9<85>ûJ<9d>¥<95>\^^MôinRßå<8d>
°I^E¾÷¯¡êÝÝ<8e>.^HUÎ^Q^BD»f^\ä¤C  04-09-2017  - Customer has cancelled their appointment. 04-10-2017 - Install Date - 04-11-2017, Installation Time - 12:00-16:00, Status - Appointment Booked, EST Bulbs - ]<Àö^P 04-10-2017 - EGN9100_Booked email sent 04-10-2017 - Install Date - 04-11-2017 , Installation Time - 12:00-16:00, Status - Appointment Booked, EST Bulbs - ]<Àö^P  19-10-2017  - Customer has cancelled their appointment.`,
			`SYSTEM`,
			`General`,
			``,
			``,
			`Appointment Cancelled (UW)`,
			`]<Àö^P`,
			`\:Ø`,
			`$^?<8a>­^@=ñ<94>7T^X[G·`,
			`\=Øð^S^¥×p^B^^D^Hé^Qn^QY^E8$^]^NÒ`,
			`^D^?<8a>­^RZ Ör^H\^\h±O.J ^d?FXÊM8`,
			``,
			`YES`,
			``,
			`1`,
			``,
			`2017-10-19`,
			`JKNIGHT`,
			`0`,
			`2017-10-19`,
			``,
			``,
			`JKNIGHT/19-10-2017/11:17:58`,
			`1`,
			`0`,
		},
	}

	got := [][]string{}

	cr := NewReader(f)
	cr.Comma = '|'
	cr.FieldsPerRecord = -1
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
