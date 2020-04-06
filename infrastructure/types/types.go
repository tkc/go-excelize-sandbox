package types

import (
	"time"
	"tkc/go-excelize-sandbox/domain/model"
)

type ExcelParamRequestJson struct {
	ClientName string
	ExcelData  map[int]map[int]map[int]*model.Excel
	JoinUser   []*model.JoinUser
	StartJST   time.Time
}
