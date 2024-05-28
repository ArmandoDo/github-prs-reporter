package reporter

import (
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
)

// Excel File Object
type Excel struct {
	file *excelize.File // Excel file
	Name string         // Name of file
	Path string         // Path of file
}

// Object for style of excel cells
type Style struct {
	aligment    string   // Text aligment
	fillColor   []string // Color of cell
	fillPattern int      // Pattern of fill
	fillType    string   // Type of fill
	fontBold    bool     // Font bold
	fontColor   string   // Font color
}

// Object for excel table based on the PR list
type Table struct {
	Base      string     // Base branch
	ClosedAt  *time.Time // PR closed at
	CreatedAt *time.Time // PR created at
	Head      string     // Head branch
	Id        int64      // PR Id
	MergedAt  *time.Time // PR merged at
	Number    int        // Number of PR
	State     string     // State of PR
	Title     string     // PR title
	Url       string     // PR url
	UserId    int64      // User Id for PR
	UserName  string     // User name for PR
}

// GenerateTable creates an excel table for the PRs
func (excel *Excel) GenerateTable(table []*Table) error {
	defer func() {
		excel.file.Close()
	}()

	// Configure headers
	excel.setHeaders()
	// Configure style for pairs
	pair, err := boxColor("FFFFFF", excel)
	// Check for errors
	if err != nil {
		return err
	}
	// Configure style for odds
	odd, err := boxColor("EBF5FB", excel)
	// Check for errors
	if err != nil {
		return err
	}

	// Iterate table
	for i, row := range table {
		// Get current excel row
		e := strconv.Itoa(i + 2)
		// Print data
		excel.file.SetCellStyle("Sheet1", "A"+e, "M"+e, getStyle(i, pair, odd))
		excel.file.SetCellValue("Sheet1", "A"+e, i+1)
		excel.file.SetCellValue("Sheet1", "B"+e, row.Id)
		excel.file.SetCellValue("Sheet1", "C"+e, row.Number)
		excel.file.SetCellValue("Sheet1", "D"+e, row.State)
		excel.file.SetCellValue("Sheet1", "E"+e, row.Title)
		excel.file.SetCellValue("Sheet1", "F"+e, row.CreatedAt)
		excel.file.SetCellValue("Sheet1", "G"+e, row.ClosedAt)
		excel.file.SetCellValue("Sheet1", "H"+e, row.MergedAt)
		excel.file.SetCellValue("Sheet1", "I"+e, row.Head)
		excel.file.SetCellValue("Sheet1", "J"+e, row.Base)
		excel.file.SetCellValue("Sheet1", "K"+e, row.UserId)
		excel.file.SetCellValue("Sheet1", "L"+e, row.UserName)
		excel.file.SetCellValue("Sheet1", "M"+e, row.Url)
	}

	// Save spreadsheet by the given path.
	err = excel.file.SaveAs(excel.Path + excel.Name)
	// Check for errors
	if err != nil {
		return err
	}
	return nil
}

// NewClient creates a new Excel object
func NewExcelFile() *Excel {
	return &Excel{file: excelize.NewFile()}
}
