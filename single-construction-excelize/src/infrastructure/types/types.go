package types

import (
	"time"
	"tkc/go-excelize-sandbox/src/domain"
)

type ExcelRequestType struct {
	ClientName string
	ExcelData  map[int]map[int]map[int]*domain.Excel
	JoinUser   []*domain.JoinUser
	StartJST   time.Time
}
