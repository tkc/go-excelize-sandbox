package lambda

import (
	"bytes"
	"encoding/json"
	"tkc/go-excelize-sandbox/domain/model"
)

func EncodeJsonParam(param model.ExcelParam) (*string, error) {
	jsonBytes, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	out := new(bytes.Buffer)
	json.Indent(out, jsonBytes, "", "    ")
	jsonStr := out.String()
	return &jsonStr, nil
}

func DecodeJsonParam(jsonStr string) (*model.ExcelParam, error) {
	jsonBytes := ([]byte)(jsonStr)
	data := new(model.ExcelParam)
	if err := json.Unmarshal(jsonBytes, data); err != nil {
		return nil, err
	}
	return data, nil
}
