package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const alfa = 1
const nonAlfaNumeric = 2

type configurationSetUp struct {
	fileName       string
	delimiter      rune
	lazyQuotes     bool
	idxRowStart    int
	idxColumnStart int
	idxColumnLast  int
	query          string
	iclIdxRow      map[int]fieldInfo
	additional     string
	sampleData     int
}

type fieldInfo struct {
	typeData     int
	defaultValue string
}

func setUp() []configurationSetUp {
	idxRowAccProvince := make(map[int]fieldInfo)
	idxRowAccProvince[0] = fieldInfo{typeData: nonAlfaNumeric}
	idxRowAccProvince[1] = fieldInfo{typeData: nonAlfaNumeric}

	return []configurationSetUp{
		{
			fileName:       fmt.Sprintf(`C:\repo-belajar\repo-eksperimen-code\src\read-csv-query\db\site_202305221400.csv`),
			delimiter:      '|',
			lazyQuotes:     true,
			idxRowStart:    1,
			idxColumnLast:  1,
			idxColumnStart: 0,
			query:          fmt.Sprintf("insert into customer_site(parent_customer_id, customer_id, created_by, created_client, updated_by, updated_client) values \n"),
			iclIdxRow:      idxRowAccProvince,
			sampleData:     3,
		},
	}
}

func main() {
	//-- Variable
	var (
		setting []configurationSetUp
		final   string
		err     error
	)

	//-- Set Up
	setting = setUp()
	final += fmt.Sprintf("-- +migrate Up\n")
	fmt.Println("-- +migrate Up")
	final += fmt.Sprintf("-- +migrate StatementBegin\n")
	fmt.Println("-- +migrate StatementBegin")
	final += fmt.Sprintf("\n")
	fmt.Println()

	for idx, itemSetting := range setting {
		final += doGenerate(itemSetting)
		final += fmt.Sprintf("\n")
		final += fmt.Sprintf("\n")

		fmt.Println()

		if len(setting)-(idx+1) == 0 {
			final += fmt.Sprintf("\n-- +migrate StatementEnd")
			fmt.Println("-- +migrate StatementEnd")
		}
	}

	err = ioutil.WriteFile("migration-site.sql", []byte(final), 0600)
	if err != nil {
		panic(err)
	}
}

func doGenerate(setting configurationSetUp) (resultQuery string) {
	var (
		file    *os.File
		err     error
		reader  *csv.Reader
		records [][]string
	)

	//-- Buka file CSV
	file, err = os.Open(setting.fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	//-- Baca file CSV
	reader = csv.NewReader(file)

	//-- Set pemisah koma
	reader.Comma = setting.delimiter
	reader.LazyQuotes = setting.lazyQuotes
	records, err = reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	//-- Tulisan pertama
	fmt.Println(setting.query)

	//-- Iterasi seluruh baris
	for idxRow, record := range records {

		if idxRow >= setting.idxRowStart {

			//-- Iterasi seluruh kolom
			for idxColumn, field := range record {
				v, ok := setting.iclIdxRow[idxColumn]
				if !ok {
					continue
				}

				if v.typeData == alfa {
					if v.defaultValue != "" {
						field = fmt.Sprintf(`'%s'`, v.defaultValue)
					} else {
						if strings.Contains(field, "'") {
							field = strings.ReplaceAll(field, "'", " ")
						}

						if field == "" {
							field = "null"
						} else {
							if idxColumn == 7 {
								field = fmt.Sprintf(`(select id from product where product_id = '%s')`, field)
							} else if idxColumn == 8 || idxColumn == 9 || idxColumn == 10 {
								date, errorS := time.Parse("1/2/2006", field)
								if errorS != nil {
									log.Fatalf(errorS.Error())
								}

								field = fmt.Sprintf(`'%s'`, date.Format("02-01-2006"))
							} else {
								field = fmt.Sprintf(`'%s'`, field)
							}
						}

					}
				} else {
					if field != "" {
						angka, _ := strconv.Atoi(field)
						if angka < 1 {
							log.Fatalln("Stop error")
							//field = fmt.Sprintf(`%s`, v.defaultValue)
						} else {
							//if idxColumn == 3 || idxColumn == 4 {
							//	field = fmt.Sprintf(`(select id from customer where mdb_company_profile_id = %s)`, field)
							//} else {
							//	field = fmt.Sprintf(`%s`, field)
							//}
							field = fmt.Sprintf(`(select id from customer where mdb_company_profile_id = %s)`, field)
						}
					} else {
						field = fmt.Sprintf(`%s`, v.defaultValue)
					}
				}

				if idxColumn == setting.idxColumnStart {
					setting.query += fmt.Sprintf(`(%s, `, field)
					fmt.Print(fmt.Sprintf(`(%s, `, field))
				} else {
					if setting.idxColumnLast == idxColumn {
						if setting.additional != "" {
							newField, _ := strconv.Atoi(setting.additional)
							setting.query += fmt.Sprintf(`%s, '%s')`, field, setting.iclIdxRow[newField].defaultValue)
							fmt.Print(fmt.Sprintf(`%s, '%s')`, field, setting.iclIdxRow[newField].defaultValue))
						} else {
							setting.query += fmt.Sprintf(`%s, 1, 'SYSTEM', 1, 'SYSTEM')`, field)
							fmt.Print(fmt.Sprintf(`%s, 1, 'SYSTEM', 1, 'SYSTEM')`, field))
						}
						break
					}

					setting.query += fmt.Sprintf(`%s`, field)
					fmt.Print(fmt.Sprintf(`%s`, field))
					setting.query += fmt.Sprintf(`, `)
					fmt.Print(fmt.Sprintf(`, `))
				}
			}

			if len(records)-(idxRow+1) == 0 {
				setting.query += fmt.Sprintf(";")
				fmt.Println(";")
			} else {
				setting.query += fmt.Sprintf(", \n")
				fmt.Println(", ")
			}
		}
	}

	resultQuery = setting.query
	return
}
