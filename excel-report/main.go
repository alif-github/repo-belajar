package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Buat file Excel baru
		f := excelize.NewFile()

		// Buat sheet baru
		sheetName := "Sheet1"
		index := f.NewSheet(sheetName)

		// Tambahkan data ke sheet
		data := [][]interface{}{
			{"ID", "Name", "Age"},
			{1, "Alice", 28},
			{2, "Bob", 32},
			{3, "Charlie", 22},
		}

		for rowIdx, rowData := range data {
			for colIdx, cellData := range rowData {
				cellName, _ := excelize.CoordinatesToCellName(colIdx+1, rowIdx+1)
				f.SetCellValue(sheetName, cellName, cellData)
			}
		}

		// Simpan file Excel ke buffer
		fileBuffer, err := f.WriteToBuffer()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set header HTTP
		w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		w.Header().Set("Content-Disposition", "attachment; filename=example.xlsx")

		// Salin isi file Excel ke response HTTP
		_, err = w.Write(fileBuffer.Bytes())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	port := 8080
	fmt.Printf("Server started on port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
