package types

import (
	"time"
	"tkc/go-excelize-sandbox/src/domain/model"
)

type ExcelRequestType struct {
	ClientName string
	ExcelData  map[int]map[int]map[int]*model.Excel
	JoinUser   []*model.JoinUser
	StartJST   time.Time
}
