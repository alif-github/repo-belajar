package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	var (
		data       [][]string
		headerData []string
	)

	//--- Data
	headerData = []string{"Name", "City", "Skills"}
	data = append(data, headerData)
	data = append(data, []string{
		"Smith", "Newyork", "Java",
	})

	//--- CSV
	csvFile, errorS := os.Create("employee.csv")
	if errorS != nil {
		log.Fatalf("failed with error: %s", errorS)
	}

	defer csvFile.Close()

	//--- Write CSV
	csvWriter := csv.NewWriter(csvFile)
	for _, empRow := range data {
		_ = csvWriter.Write(empRow)
	}

	defer csvWriter.Flush()
}
