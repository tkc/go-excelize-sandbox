package param

import (
	"testing"
	"time"
	"tkc/go-excelize-sandbox/src/domain/model"
	"tkc/go-excelize-sandbox/src/infrastructure/types"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"github.com/wawandco/fako"
)

func Test_encode_decode_json(t *testing.T) {

	excelParamParser := NewExcelParamParser()

	var excel model.Excel
	fako.Fill(&excel)

	var joinUser model.JoinUser
	fako.Fill(&joinUser)

	JoinUsers := []*model.JoinUser{&joinUser}

	fakeData := time.Now()
	excelData := make(map[int]map[int]map[int]*model.Excel)
	excelData[1] = make(map[int]map[int]*model.Excel)
	excelData[1][1] = make(map[int]*model.Excel)
	excelData[1][1][1] = &excel

	excelParam := types.ExcelRequestType{
		ClientName: faker.Name(),
		StartJST:   fakeData,
		ExcelData:  excelData,
		JoinUser:   JoinUsers,
	}

	json, err := excelParamParser.EncodeJsonParam(excelParam)
	assert.NoError(t, err)

	generatedParam, err := excelParamParser.DecodeJsonParam(*json)
	assert.NoError(t, err)

	assert.Equal(t, generatedParam.ClientName, excelParam.ClientName)
	assert.Equal(t, generatedParam.JoinUser, excelParam.JoinUser)
	assert.Equal(t, generatedParam.ExcelData, excelParam.ExcelData)
	assert.Equal(t, generatedParam.StartJST.Format(time.UnixDate), excelParam.StartJST.Format(time.UnixDate))
}
