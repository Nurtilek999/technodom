package usecase

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"merchant/internal/entity"
	"strconv"
)

func Parse(fileName string, id int) (map[int][]entity.Product, error) {

	rowsMap := make(map[int][]entity.Product)
	file, err := excelize.OpenFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("Error in OpenFile: ", err.Error())
	}
	//name, err := file.GetCellValue("File", "A2")
	rows, err := file.Rows("File")
	if err != nil {
		return nil, err
	}
	counter := 0
	for rows.Next() {
		if counter == 0 {
			counter++
			continue
		}
		row, err := rows.Columns()
		if err != nil {
			fmt.Println(err)
		}

		prod := entity.Product{}
		offerID, err := strconv.Atoi(row[0])
		if err != nil {
			return nil, err
		}
		prod.OfferID = offerID
		prod.Name = row[1]
		price, err := strconv.ParseFloat(row[2], 32)
		if err != nil {
			return nil, err
		}
		prod.Price = float32(price)
		quantity, err := strconv.Atoi(row[3])
		if err != nil {
			return nil, err
		}
		prod.Quantity = quantity
		if row[4] == "TRUE" {
			prod.Available = true
		} else {
			prod.Available = false
		}
		rowsMap[id] = append(rowsMap[id], prod)
	}

	return rowsMap, nil
}
