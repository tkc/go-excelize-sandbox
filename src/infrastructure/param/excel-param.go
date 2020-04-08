package param

import (
	"bytes"
	"encoding/json"
	"tkc/go-excelize-sandbox/src/infrastructure/types"
)

type excelParamParser struct{}

type ExcelParamParser interface {
	EncodeJSONParam(param types.ExcelRequestType) (*string, error)
	DecodeJSONParam(jsonStr string) (*types.ExcelRequestType, error)
}

func NewExcelParamParser() ExcelParamParser {
	return &excelParamParser{}
}

func (h *excelParamParser) EncodeJSONParam(param types.ExcelRequestType) (*string, error) {
	jsonBytes, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	out := new(bytes.Buffer)
	json.Indent(out, jsonBytes, "", "    ")
	jsonStr := out.String()
	return &jsonStr, nil
}

func (h *excelParamParser) DecodeJSONParam(jsonStr string) (*types.ExcelRequestType, error) {
	jsonBytes := ([]byte)(jsonStr)
	data := new(types.ExcelRequestType)
	if err := json.Unmarshal(jsonBytes, data); err != nil {
		return nil, err
	}
	return data, nil
}
