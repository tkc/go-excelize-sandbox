package usecase

import (
	"log"
	"testing"
	"time"
	"tkc/go-excelize-sandbox/src/domain/model"
	"tkc/go-excelize-sandbox/src/infrastructure/types"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func Test_create_excel_byte(t *testing.T) {

	p := NewExcelUsecase()
	testDate := time.Now()
	testid := 1

	excel := model.Excel{
		UserID:           testid,
		UserName:         faker.Name(),
		StartedDate:      &testDate,
		StartedDatetime:  &testDate,
		EndedDatetime:    &testDate,
		ConstructionName: faker.Sentence(),
		Memo:             faker.Sentence(),
		Address:          faker.Sentence(),
		SalesUserName:    faker.Sentence(),
	}

	excelData := make(map[int]map[int]map[int]*model.Excel)
	// excelData[0] = make(map[int]map[int]*model.Excel)
	// excelData[0][0] = make(map[int]*model.Excel)
	// excelData[0][0][0] = &excel

	excelData[1] = make(map[int]map[int]*model.Excel)
	excelData[1][0] = make(map[int]*model.Excel)
	excelData[1][0][1] = &excel

	// excelData[2] = make(map[int]map[int]*model.Excel)
	// excelData[2][0] = make(map[int]*model.Excel)
	// excelData[2][0][0] = &excel

	joinUser := model.JoinUser{
		ID:        &testid,
		UserID:    &testid,
		CreatedAt: &testDate,
	}

	var (
		startJST   = time.Now()
		clientName = "clientName"
	)

	JoinUsers := []*model.JoinUser{&joinUser, &joinUser, &joinUser, &joinUser, &joinUser}

	excelParam := types.ExcelRequestType{
		StartJST:   startJST,
		ExcelData:  excelData,
		JoinUser:   JoinUsers,
		ClientName: clientName,
	}
	_, err := p.CreateExcelByte(excelParam)
	log.Print(err)
	assert.NoError(t, err)

	err = p.SaveExcelFile(excelParam)
	log.Print(err)
	assert.NoError(t, err)
}
