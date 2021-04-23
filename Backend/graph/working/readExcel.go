package working

import (
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func check() {
	f, err := excelize.OpenFile("C:\\blogger\\src\\Backend\\Master.xlsx")
	if err != nil {
		log.Println(err)
	}
	UserName := make(map[string]string)
	var data []string
	// OneData := make(map[int]string)
	cell := f.GetCellValue("User Sheet", "A1")
	log.Println(cell)
	rows := f.GetRows("User Sheet")
	for key, rowsData := range rows {
		// log.Println("rows--->", rows)
		// UserName[rows] = rows
		data = append(data, rowsData[key])
		for _, colsCells := range rowsData {
			if rowsData[key] != "" {
				UserName[rowsData[key]] = colsCells
				key += 1
			}
			// check := colsCells[key]
			// log.Println(check)
			log.Println("cols--->", colsCells)
		}
		log.Println("Data Map--->", data)
	}
}
