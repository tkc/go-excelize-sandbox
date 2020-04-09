package usecase

import (
	"testing"
	"time"
	"tkc/go-excelize-sandbox/src/domain"
	"tkc/go-excelize-sandbox/src/infrastructure/types"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func Test_create_excel_byte(t *testing.T) {
	p := NewExcelInteractor()
	testDate := time.Now()
	testid := 1

	excel := domain.Excel{
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

	excelData := make(map[int]map[int]map[int]*domain.Excel)
	excelData[1] = make(map[int]map[int]*domain.Excel)
	excelData[1][0] = make(map[int]*domain.Excel)
	excelData[1][0][1] = &excel

	joinUser := domain.JoinUser{
		ID:        &testid,
		UserID:    &testid,
		CreatedAt: &testDate,
	}

	var (
		startJST   = time.Now()
		clientName = faker.Name()
	)

	JoinUsers := []*domain.JoinUser{&joinUser, &joinUser, &joinUser, &joinUser, &joinUser}

	excelParam := types.ExcelRequestType{
		StartJST:   startJST,
		ExcelData:  excelData,
		JoinUser:   JoinUsers,
		ClientName: clientName,
	}

	_, err := p.CreateExcelByte(excelParam)
	assert.NoError(t, err)

	err = p.SaveExcelFile(excelParam)
	assert.NoError(t, err)
}
