package lambda

import (
	"log"
	"testing"
	"time"
	"tkc/go-excelize-sandbox/domain/model"
)

func Test_create_excel(t *testing.T) {

	p := NewExcelPresenter()
	log.Print(p)
	date := time.Now()

	excel := model.Excel{
		UserID:           1,
		UserName:         "UserName",
		StartedDate:      &date,
		StartedDatetime:  &date,
		EndedDatetime:    &date,
		ConstructionName: "ConstructionName",
		Memo:             "Memo",
		Address:          "Address",
		SalesUserName:    "SalesUserName",
	}

	excelData := make(map[int]map[int]map[int]*model.Excel)
	excelData[1] = make(map[int]map[int]*model.Excel)
	excelData[1][1] = make(map[int]*model.Excel)
	excelData[1][1][1] = &excel

	id := 1
	joinUser := model.JoinUser{
		ID:        &id,
		UserID:    &id,
		CreatedAt: &date,
	}

	var (
		startJST   = time.Now()
		clientName = "clientName"
	)

	JoinUsers := []*model.JoinUser{&joinUser, &joinUser, &joinUser, &joinUser, &joinUser}
	excelParam := model.ExcelParam{
		StartJST:   startJST,
		ExcelData:  excelData,
		JoinUser:   JoinUsers,
		ClientName: clientName,
	}

	_, err := p.CreateExcel(excelParam)

	if err != nil {
		log.Print(err)
	} else {
		log.Print("ok!!")
	}
}
