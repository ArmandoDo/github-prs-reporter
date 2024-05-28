package reporter

import "github.com/xuri/excelize/v2"

// boxColor returns the style for the cell using a specific color
func boxColor(color string, excel *Excel) (int, error) {
	return excel.createStyle(&Style{
		aligment:    "center",
		fillColor:   []string{color},
		fillPattern: 1,
		fillType:    "pattern",
		fontBold:    false,
		fontColor:   "000000",
	})
}

// createStyle returns an Excel file
func (excel *Excel) createStyle(style *Style) (int, error) {
	s, err := excel.file.NewStyle(
		&excelize.Style{
			Alignment: &excelize.Alignment{Horizontal: style.aligment},
			Font:      &excelize.Font{Bold: style.fontBold, Color: style.fontColor},
			Border: []excelize.Border{
				{Type: "left", Color: "000000", Style: 1},
				{Type: "right", Color: "000000", Style: 1},
				{Type: "top", Color: "000000", Style: 1},
				{Type: "bottom", Color: "000000", Style: 1},
			},
			Fill: excelize.Fill{
				Type:    style.fillType,
				Color:   style.fillColor,
				Pattern: style.fillPattern,
			},
		},
	)
	// Check for errors
	if err != nil {
		return s, err
	}

	return s, nil
}

// getStyle returns the style for pair and odd rows
func getStyle(row, pair, odd int) int {
	if row%2 == 0 {
		return pair
	}

	return odd
}

// setHeaders prints the headers of the table
func (excel *Excel) setHeaders() error {
	headerstyle, err := excel.createStyle(&Style{
		aligment:    "center",
		fillColor:   []string{"EBF5FB"},
		fillPattern: 1,
		fillType:    "pattern",
		fontBold:    true,
		fontColor:   "000000",
	})
	// Check for errors
	if err != nil {
		return err
	}
	// Set value of a cell.
	excel.file.SetCellStyle("Sheet1", "A1", "M1", headerstyle)
	excel.file.SetCellValue("Sheet1", "A1", "#")
	excel.file.SetCellValue("Sheet1", "B1", "PR ID")
	excel.file.SetCellValue("Sheet1", "C1", "PR Number")
	excel.file.SetCellValue("Sheet1", "D1", "PR State")
	excel.file.SetCellValue("Sheet1", "E1", "PR Title")
	excel.file.SetCellValue("Sheet1", "F1", "PR Created At")
	excel.file.SetCellValue("Sheet1", "G1", "PR Closed At")
	excel.file.SetCellValue("Sheet1", "H1", "PR Merged At")
	excel.file.SetCellValue("Sheet1", "I1", "PR Head")
	excel.file.SetCellValue("Sheet1", "J1", "PR Base")
	excel.file.SetCellValue("Sheet1", "K1", "PR User ID")
	excel.file.SetCellValue("Sheet1", "L1", "PR User Login")
	excel.file.SetCellValue("Sheet1", "M1", "PR URL")

	return nil
}
