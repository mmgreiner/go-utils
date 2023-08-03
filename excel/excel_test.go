package excel

import (
	"testing"

	"time"

	"github.com/mmgreiner/go-utils/str2"
	"gotest.tools/assert"
)

type TestStruct struct {
	IntegerValue int
	DateValue    time.Time
	FloatValue   float32
	StringValue  string
}

func toRow(row []string) TestStruct {
	return TestStruct{
		IntegerValue: str2.ToInt(row[0]),
		DateValue:    str2.ToTime(row[1]),
		FloatValue:   float32(str2.ToFloat(row[2])),
		StringValue:  row[3],
	}
}
func TestRead(t *testing.T) {
	data := []TestStruct{}
	err := ReadExcel("./test.xlsx", true, toRow, &data)
	assert.Assert(t, err)
	assert.Assert(t, len(data) > 0)
	data = append(data, TestStruct{})
}
