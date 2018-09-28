package billdsv

import (
	"fmt"
	"strings"
	"testing"

	"github.com/bmizerany/assert"
)

func TestReader1(t *testing.T) {
	f := strings.NewReader(`1a0|first string|final string
2b1|second string
that is multi-line
|final string
3c2|third string|final string
`)

	want := [][]string{
		{"1a0", "first string", "final string"},
		{"2b1", "second string\nthat is multi-line\n", "final string"},
		{"3c2", "third string", "final string"},
	}

	got := [][]string{}

	cr := NewReader(f, 3)

	err := cr.ReadAll(func(row [][]byte) {
		fmt.Println(truncateStrings(20, row))
		rowStrings := make([]string, 3)
		for i, c := range row {
			rowStrings[i] = string(c)
		}
		got = append(got, rowStrings)
	})
	if err != nil {
		t.Error(err)
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

	cr := NewReader(f, 3)

	err := cr.ReadAll(func(row [][]byte) {
		fmt.Println(truncateStrings(20, row))
		rowStrings := make([]string, 3)
		for i, c := range row {
			rowStrings[i] = string(c)
		}
		got = append(got, rowStrings)
	})
	if err != nil {
		t.Error(err)
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

	cr := NewReader(f, 29)
	err := cr.ReadAll(func(row [][]byte) {
		fmt.Println(truncateStrings(20, row))
		rowStrings := make([]string, 29)
		for i, c := range row {
			rowStrings[i] = string(c)
		}
		got = append(got, rowStrings)
	})
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, want, got)
}

func TestReader4(t *testing.T) {
	f := strings.NewReader(`A|B|C
str1|123|str2
str3|456|str"4
`)

	want := [][]string{
		{"str1", "123", "str2"},
		{"str3", "456", "str\"4"},
	}

	got := [][]string{}

	cr := NewReader(f, 3)
	cr.SkipHeading = true

	err := cr.ReadAll(func(row [][]byte) {
		fmt.Println(truncateStrings(20, row))
		rowStrings := make([]string, 3)
		for i, c := range row {
			rowStrings[i] = string(c)
		}
		got = append(got, rowStrings)
	})
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, want, got)
}

func TestReader5(t *testing.T) {
	f := strings.NewReader(`A|B|C
str1|123|str2
str3|456|str"4`)

	want := [][]string{
		{"str1", "123", "str2"},
		{"str3", "456", "str\"4"},
	}

	got := [][]string{}

	cr := NewReader(f, 3)
	cr.SkipHeading = true

	err := cr.ReadAll(func(row [][]byte) {
		fmt.Println(truncateStrings(20, row))
		rowStrings := make([]string, 3)
		for i, c := range row {
			rowStrings[i] = string(c)
		}
		got = append(got, rowStrings)
	})
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, want, got)
}

func TestReader6(t *testing.T) {
	f := strings.NewReader(`CrNumber|CrPeriod|CrCarBonusID|CrCBExecID|CrCBRepayment|CrCBRepaymentFee|CrCBBonusEligibl|CrCBBonusAmount|CrCBCommitted|CrCBNotes|CrCBScheme|CrCBSpareC2|CrCBBalance|CrCBSpareNum1|CrCBSpareNum2|CrCBSpareDate1|CrCBSpareDate2
328|2015/09|1097684|006308|Yes|-150|Yes|0|2015-09-11||Scheme3||4750|0|0|1899-12-30|1899-12-30
333|2015/11|1155246|006308|Yes|-150|Yes|0|2015-11-12||Scheme3||13795|0|0|1899-12-30|1899-12-30
335|2015/12|1184135|006308|Yes|-150|Yes|0|2015-12-11||Scheme3||13645|0|0|1899-12-30|1899-12-30
337|2016/01|1213806|006308|Yes|-150|Yes|0|2016-01-13||Scheme3||13495|0|0|1899-12-30|1899-12-30
339|2016/02|1243935|006308|Yes|-150|Yes|0|2016-02-11||Scheme3||13345|0|0|1899-12-30|1899-12-30
341|2016/03|1274518|006308|Yes|-150|Yes|0|2016-03-11||Scheme3||13195|0|0|1899-12-30|1899-12-30
343|2016/04|1305552|006308|Yes|-150|Yes|0|2016-04-13||Scheme3||13045|0|0|1899-12-30|1899-12-30
345|2016/05|1337043|006308|Yes|-150|Yes|0|2016-05-12||Scheme3||12895|0|0|1899-12-30|1899-12-30
346|2016/06|1400146|006308|Yes|-150|Yes|0|2016-06-13||Scheme3||12745|0|0|1899-12-30|1899-12-30
348|2016/07|1432564|006308|Yes|-150|Yes|0|2016-07-12||Scheme3||12595|0|0|1899-12-30|1899-12-30
352|2016/08|1465444|006308|Yes|-150|Yes|0|2016-08-11||Scheme3||12445|0|0|1899-12-30|1899-12-30
353|2016/09|1498796|006308|Yes|-150|Yes|0|2016-09-15||Scheme3||12295|0|0|1899-12-30|1899-12-30
355|2016/10|1532610|006308|Yes|-150|Yes|0|2016-10-12||Scheme3||12145|0|0|1899-12-30|1899-12-30
356|2016/11|1566908|006308|Yes|-150|Yes|0|2016-11-11||Scheme3||11995|0|0|1899-12-30|1899-12-30
358|2016/12|1601683|006308|Yes|-150|Yes|0|2016-12-12||Scheme3||11845|0|0|1899-12-30|1899-12-30
360|2017/01|1636937|006308|Yes|-150|Yes|0|2017-01-12||Scheme3||11695|0|0|1899-12-30|1899-12-30
362|2017/02|1672673|006308|Yes|-150|Yes|0|2017-02-13||Scheme3||11545|0|0|1899-12-30|1899-12-30
364|2017/03|1708895|006308|Yes|-150|Yes|0|2017-03-13||Scheme3||11395|0|0|1899-12-30|1899-12-30
366|2017/04|1745601|006308|Yes|-150|Yes|0|2017-04-11||Scheme3||11245|0|0|1899-12-30|1899-12-30
369|2017/05|1782786|006308|Yes|-150|Yes|0|2017-05-11||Scheme3||11095|0|0|1899-12-30|1899-12-30
370|2017/06|1820460|006308|Yes|-150|Yes|0|2017-06-13||Scheme3||10945|0|0|1899-12-30|1899-12-30
373|2017/07|1858638|006308|Yes|-150|Yes|0|2017-07-11||Scheme3||10795|0|0|1899-12-30|1899-12-30
374|2017/08|1897324|006308|Yes|-150|No|0|2017-08-11||Scheme3||10645|0|0|1899-12-30|1899-12-30
376|2017/09|1936512|006308|Yes|-150|No|0|2017-09-13||Scheme3||10495|0|0|1899-12-30|1899-12-30
376|2017/10|9876543|678901|Yes|-150|Yes|0|2017-10-13||Scheme3||99999|0|0|1899-12-30|1899-12-30
376|2017/11|9876543|789012|Yes|-150|Yes|0|2017-11-13||Scheme3||99999|0|0|1899-12-30|1899-12-30`)

	want := [][]string{
		{"328", "2015/09", "1097684", "006308", "Yes", "-150", "Yes", "0", "2015-09-11", "", "Scheme3", "", "4750", "0", "0", "1899-12-30", "1899-12-30"},
		{"333", "2015/11", "1155246", "006308", "Yes", "-150", "Yes", "0", "2015-11-12", "", "Scheme3", "", "13795", "0", "0", "1899-12-30", "1899-12-30"},
		{"335", "2015/12", "1184135", "006308", "Yes", "-150", "Yes", "0", "2015-12-11", "", "Scheme3", "", "13645", "0", "0", "1899-12-30", "1899-12-30"},
		{"337", "2016/01", "1213806", "006308", "Yes", "-150", "Yes", "0", "2016-01-13", "", "Scheme3", "", "13495", "0", "0", "1899-12-30", "1899-12-30"},
		{"339", "2016/02", "1243935", "006308", "Yes", "-150", "Yes", "0", "2016-02-11", "", "Scheme3", "", "13345", "0", "0", "1899-12-30", "1899-12-30"},
		{"341", "2016/03", "1274518", "006308", "Yes", "-150", "Yes", "0", "2016-03-11", "", "Scheme3", "", "13195", "0", "0", "1899-12-30", "1899-12-30"},
		{"343", "2016/04", "1305552", "006308", "Yes", "-150", "Yes", "0", "2016-04-13", "", "Scheme3", "", "13045", "0", "0", "1899-12-30", "1899-12-30"},
		{"345", "2016/05", "1337043", "006308", "Yes", "-150", "Yes", "0", "2016-05-12", "", "Scheme3", "", "12895", "0", "0", "1899-12-30", "1899-12-30"},
		{"346", "2016/06", "1400146", "006308", "Yes", "-150", "Yes", "0", "2016-06-13", "", "Scheme3", "", "12745", "0", "0", "1899-12-30", "1899-12-30"},
		{"348", "2016/07", "1432564", "006308", "Yes", "-150", "Yes", "0", "2016-07-12", "", "Scheme3", "", "12595", "0", "0", "1899-12-30", "1899-12-30"},
		{"352", "2016/08", "1465444", "006308", "Yes", "-150", "Yes", "0", "2016-08-11", "", "Scheme3", "", "12445", "0", "0", "1899-12-30", "1899-12-30"},
		{"353", "2016/09", "1498796", "006308", "Yes", "-150", "Yes", "0", "2016-09-15", "", "Scheme3", "", "12295", "0", "0", "1899-12-30", "1899-12-30"},
		{"355", "2016/10", "1532610", "006308", "Yes", "-150", "Yes", "0", "2016-10-12", "", "Scheme3", "", "12145", "0", "0", "1899-12-30", "1899-12-30"},
		{"356", "2016/11", "1566908", "006308", "Yes", "-150", "Yes", "0", "2016-11-11", "", "Scheme3", "", "11995", "0", "0", "1899-12-30", "1899-12-30"},
		{"358", "2016/12", "1601683", "006308", "Yes", "-150", "Yes", "0", "2016-12-12", "", "Scheme3", "", "11845", "0", "0", "1899-12-30", "1899-12-30"},
		{"360", "2017/01", "1636937", "006308", "Yes", "-150", "Yes", "0", "2017-01-12", "", "Scheme3", "", "11695", "0", "0", "1899-12-30", "1899-12-30"},
		{"362", "2017/02", "1672673", "006308", "Yes", "-150", "Yes", "0", "2017-02-13", "", "Scheme3", "", "11545", "0", "0", "1899-12-30", "1899-12-30"},
		{"364", "2017/03", "1708895", "006308", "Yes", "-150", "Yes", "0", "2017-03-13", "", "Scheme3", "", "11395", "0", "0", "1899-12-30", "1899-12-30"},
		{"366", "2017/04", "1745601", "006308", "Yes", "-150", "Yes", "0", "2017-04-11", "", "Scheme3", "", "11245", "0", "0", "1899-12-30", "1899-12-30"},
		{"369", "2017/05", "1782786", "006308", "Yes", "-150", "Yes", "0", "2017-05-11", "", "Scheme3", "", "11095", "0", "0", "1899-12-30", "1899-12-30"},
		{"370", "2017/06", "1820460", "006308", "Yes", "-150", "Yes", "0", "2017-06-13", "", "Scheme3", "", "10945", "0", "0", "1899-12-30", "1899-12-30"},
		{"373", "2017/07", "1858638", "006308", "Yes", "-150", "Yes", "0", "2017-07-11", "", "Scheme3", "", "10795", "0", "0", "1899-12-30", "1899-12-30"},
		{"374", "2017/08", "1897324", "006308", "Yes", "-150", "No", "0", "2017-08-11", "", "Scheme3", "", "10645", "0", "0", "1899-12-30", "1899-12-30"},
		{"376", "2017/09", "1936512", "006308", "Yes", "-150", "No", "0", "2017-09-13", "", "Scheme3", "", "10495", "0", "0", "1899-12-30", "1899-12-30"},
		{"376", "2017/10", "9876543", "678901", "Yes", "-150", "Yes", "0", "2017-10-13", "", "Scheme3", "", "99999", "0", "0", "1899-12-30", "1899-12-30"},
		{"376", "2017/11", "9876543", "789012", "Yes", "-150", "Yes", "0", "2017-11-13", "", "Scheme3", "", "99999", "0", "0", "1899-12-30", "1899-12-30"},
	}

	got := [][]string{}

	cr := NewReader(f, 17)
	cr.SkipHeading = true

	err := cr.ReadAll(func(row [][]byte) {
		fmt.Println(truncateStrings(20, row))
		rowStrings := make([]string, 17)
		for i, c := range row {
			rowStrings[i] = string(c)
		}
		got = append(got, rowStrings)
	})
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, want, got)
}

func TestReader7(t *testing.T) {
	f := strings.NewReader(`Content ID,Content name,User ID,Email,Username,Last login,Last activity,Registration date,Role,Name,Given name,Family name,PartnerID,Date course started,SCORM course status,Test score,Session time,Time spent on test,Date course completed,Approval date
,,,,,,,,,,,,,,not started,,,,,`)

	want := [][]string{
		{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "not started", "", "", "", "", ""},
	}

	got := [][]string{}

	cr := NewReader(f, 20)
	cr.Separator = ','
	cr.SkipHeading = true

	err := cr.ReadAll(func(row [][]byte) {
		fmt.Println(truncateStrings(20, row))
		rowStrings := make([]string, 20)
		for i, c := range row {
			rowStrings[i] = string(c)
		}
		got = append(got, rowStrings)
	})
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, want, got)
}

func truncateStrings(limit int, in [][]byte) string {
	sb := strings.Builder{}
	sb.WriteString("[")
	for i, b := range in {
		s := strings.Replace(string(b), "\n", `\n`, -1)
		sb.WriteString("\t")
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

		sb.WriteString("\n")
	}
	sb.WriteString("]")
	return sb.String()
}
