package excel

import (
	"log"
	"testing"
	"time"
)

func Test_create_excel(t *testing.T) {

	p := NewExcelPresenter()
	log.Print(p)
	date := time.Now()

	excel := Excel{
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

	excelData := make(map[int]map[int]map[int]*Excel)
	excelData[1] = make(map[int]map[int]*Excel)
	excelData[1][1] = make(map[int]*Excel)
	excelData[1][1][1] = &excel

	id := 1
	joinUser := JoinUser{
		ID:        &id,
		UserID:    &id,
		CreatedAt: &date,
	}

	var (
		startJST   = time.Now()
		clientName = "clientName"
	)

	JoinUsers := []*JoinUser{&joinUser, &joinUser, &joinUser, &joinUser, &joinUser}
	excelParam := ExcelParams{
		startJST:   startJST,
		excelData:  excelData,
		joinUser:   JoinUsers,
		clientName: clientName,
	}

	// log.Print(excelParam)
	// _, err := CreateExcel(startJST, clientName, excelData, []*JoinUser{&joinUser, &joinUser, &joinUser, &joinUser, &joinUser})
	_, err := CreateExcel(excelParam)

	if err != nil {
		log.Print(err)
	} else {
		log.Print("ok!!")
	}
}
