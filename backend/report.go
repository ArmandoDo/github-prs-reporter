package main

import "backend/reporter"

// generateExcelReport generates an excel document based on the data of the table
func generateExcelReport(path, name string, table []*reporter.Table) error {
	// Create excel file
	f := reporter.NewExcelFile()
	f.Name = name
	f.Path = path
	// Generate table
	err := f.GenerateTable(table)
	if err != nil {
		return err
	}

	return nil
}
