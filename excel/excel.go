package excel

import (
	"github.com/xuri/excelize/v2"
)

// reads the given excel file and converts each row into a slice of strings.
// uses the 'rowser' function to convert this slice into the actual data struct.
func ReadExcel[T any](fn string, header bool, rower func([]string) T) ([]T, error) {
	file, err := excelize.OpenFile(fn)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	sheet := file.GetSheetName(0)
	rows, _ := file.Rows(sheet)

	if header {
		// skip header line
		rows.Next()
	}

	data := []T{}
	for rows.Next() {
		row, _ := rows.Columns()
		dataRow := rower(row)
		data = append(data, dataRow)
	}
	return data, nil
}
