package param

import (
	"testing"
	"time"
	"tkc/go-excelize-sandbox/src/domain/model"
	"tkc/go-excelize-sandbox/src/infrastructure/types"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func Test_encode_decode_json(t *testing.T) {
	fakeDate := time.Now()
	excelParamParser := NewExcelParamParser()

	dummyID := 1
	excel := model.Excel{
		UserID:           dummyID,
		UserName:         faker.Name(),
		StartedDate:      &fakeDate,
		StartedDatetime:  &fakeDate,
		EndedDatetime:    &fakeDate,
		ConstructionName: faker.Sentence(),
		Memo:             faker.Sentence(),
		Address:          faker.Sentence(),
		SalesUserName:    faker.Sentence(),
	}

	joinUser := model.JoinUser{
		ID:        &dummyID,
		UserID:    &dummyID,
		CreatedAt: &fakeDate,
	}

	JoinUsers := []*model.JoinUser{&joinUser}

	excelData := make(map[int]map[int]map[int]*model.Excel)
	excelData[0] = make(map[int]map[int]*model.Excel)
	excelData[0][0] = make(map[int]*model.Excel)
	excelData[0][0][0] = &excel

	excelParam := types.ExcelRequestType{
		ClientName: faker.Name(),
		StartJST:   fakeDate,
		ExcelData:  excelData,
		JoinUser:   JoinUsers,
	}

	json, err := excelParamParser.EncodeJSONParam(excelParam)
	assert.NoError(t, err)

	generatedParam, err := excelParamParser.DecodeJSONParam(*json)
	assert.NoError(t, err)

	assert.Equal(t, generatedParam.ClientName, excelParam.ClientName)
	assert.Equal(t, generatedParam.JoinUser[0].ID, excelParam.JoinUser[0].ID)
	assert.Equal(t, generatedParam.JoinUser[0].UserID, excelParam.JoinUser[0].UserID)

	assert.Equal(t,
		generatedParam.JoinUser[0].CreatedAt.Format(time.UnixDate),
		excelParam.JoinUser[0].CreatedAt.Format(time.UnixDate))

	assert.Equal(t,
		generatedParam.ExcelData[0][0][0].UserID,
		excelParam.ExcelData[0][0][0].UserID)

	assert.Equal(t,
		generatedParam.ExcelData[0][0][0].UserName,
		excelParam.ExcelData[0][0][0].UserName)

	assert.Equal(t,
		generatedParam.ExcelData[0][0][0].Memo,
		excelParam.ExcelData[0][0][0].Memo)

	assert.Equal(t,
		generatedParam.ExcelData[0][0][0].Address,
		excelParam.ExcelData[0][0][0].Address)

	assert.Equal(t,
		generatedParam.ExcelData[0][0][0].ConstructionName,
		excelParam.ExcelData[0][0][0].ConstructionName)

	assert.Equal(t,
		generatedParam.ExcelData[0][0][0].SalesUserName,
		excelParam.ExcelData[0][0][0].SalesUserName)

	assert.Equal(t,
		generatedParam.ExcelData[0][0][0].StartedDate.Format(time.RFC1123),
		excelParam.ExcelData[0][0][0].StartedDate.Format(time.RFC1123))

	assert.Equal(t,
		generatedParam.ExcelData[0][0][0].StartedDatetime.Format(time.RFC1123),
		excelParam.ExcelData[0][0][0].StartedDatetime.Format(time.RFC1123))

	assert.Equal(t,
		generatedParam.ExcelData[0][0][0].EndedDatetime.Format(time.RFC1123),
		excelParam.ExcelData[0][0][0].EndedDatetime.Format(time.RFC1123))

	assert.Equal(t,
		generatedParam.StartJST.Format(time.UnixDate),
		excelParam.StartJST.Format(time.UnixDate))
}
