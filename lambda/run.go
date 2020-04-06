package lambda

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
	"tkc/go-excelize-sandbox/domain/model"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

var (
	sheetName = "Sheet1"
	fileName  = "./format.xlsx"
	jweek     = [7]string{"日", "月", "火", "水", "木", "金", "土"}
	weeks     = map[int]int{0: 5, 1: 10, 2: 15, 3: 20, 4: 25, 5: 30, 6: 35}
)

type excelPresenter struct{}

type ExcelPresenter interface {
	CreateExcel(param model.ExcelParam) ([]byte, error)
}

func NewExcelPresenter() ExcelPresenter {
	return &excelPresenter{}
}

func (excelPresenter *excelPresenter) CreateExcel(param model.ExcelParam) ([]byte, error) {
	rows := 8
	pageNum := 1
	f := excelize.NewFile()
	maxRows := rows

	for c := 0; c < len(param.JoinUser); c++ {
		if maxRows >= 48 {
			pageNum = pageNum + 1
			sheetName = "Sheet" + strconv.Itoa(pageNum)
			maxRows = rows
			index := f.NewSheet(sheetName)
			err := f.CopySheet(1, index)
			if err != nil {
				return nil, err
			}
		}

		// 営業所の設定 D2
		err := f.SetCellValue(sheetName, "D2", param.ClientName)
		if err != nil {
			return nil, err
		}

		// 更新日の設定出力日 AR2
		t := time.Now().In(time.FixedZone("Asia/Tokyo", 9*60*60))
		err = f.SetCellValue(sheetName, "AK2", t.Format("2006/01/02 15:04:05"))
		if err != nil {
			return nil, err
		}

		RowsCount := 0
		// 一人の一週間の予定を出力するループ
		for i, dayData := range param.ExcelData[*param.JoinUser[c].UserID] {
			row := maxRows
			flag := true
			if i == 7 {
				flag = false
			}
			if flag {
				day := param.StartJST.AddDate(0, 0, i)
				//日付の設定
				err = f.SetCellValue(sheetName, ColumnNumberToName(weeks[i], 4), day.Format("01/02"))
				if err != nil {
					return nil, err
				}
				//曜日の設定
				err = f.SetCellValue(sheetName, ColumnNumberToName(weeks[i], 5), jweek[day.Weekday()])
				if err != nil {
					return nil, err
				}
			}
			if len(dayData) == 0 {
				continue
			}
			for d := 1; d <= len(dayData); d++ {
				if !flag {
					//工事者名
					err = f.SetCellValue(sheetName, "D"+strconv.Itoa(row), dayData[d].UserName)
					if err != nil {
						return nil, err
					}
					break
				}
				columnNum := weeks[i]

				//工事者名
				err = f.SetCellValue(sheetName, "D"+strconv.Itoa(row), dayData[d].UserName)
				if err != nil {
					return nil, err
				}

				//時間開始
				err = f.SetCellValue(sheetName, ColumnNumberToName(columnNum, row), TimeFormat(dayData[d].StartedDatetime))
				if err != nil {
					return nil, err
				}
				columnNum += 1

				//時間終了
				err = f.SetCellValue(sheetName, ColumnNumberToName(columnNum, row), TimeFormat(dayData[d].EndedDatetime))
				if err != nil {
					return nil, err
				}
				columnNum += 1

				//現場名
				err = f.SetCellValue(sheetName, ColumnNumberToName(columnNum, row), ConstructionName(dayData[d].ConstructionName))
				if err != nil {
					return nil, err
				}
				columnNum += 1

				//工事内容
				err = f.SetCellValue(sheetName, ColumnNumberToName(columnNum, row), dayData[d].Memo)
				if err != nil {
					return nil, err
				}
				columnNum += 1

				//住所
				err = f.SetCellValue(sheetName, ColumnNumberToName(columnNum, row), dayData[d].Address)
				if err != nil {
					return nil, err
				}
				row += 1
			}

			if RowsCount < len(dayData) {
				RowsCount = len(dayData)
			}
		}

		// 書き込みがあった場合のみ、セル結合
		if len(param.ExcelData[*param.JoinUser[c].UserID]) > 0 {
			err = f.MergeCell(sheetName, "D"+strconv.Itoa(maxRows), "D"+strconv.Itoa((maxRows+RowsCount)-1))
			if err != nil {
				return nil, err
			}
		}
		maxRows += RowsCount
	}

	f.SaveAs("./Book1.xlsx")
	f.DeleteSheet("temp")

	var b bytes.Buffer
	if err := f.Write(&b); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func ConstructionName(s string) string {
	if s == "" {
		return "案件名未入力"
	}
	return s
}

func TimeFormat(i *time.Time) string {
	if i == nil {
		return "時刻未定"
	}
	t := i.In(time.FixedZone("Asia/Tokyo", 9*60*60))
	return fmt.Sprintf("%02d:%02d", t.Hour(), t.Minute())
}

func ColumnNumberToName(i int, n int) string {
	j, _ := excelize.ColumnNumberToName(i)
	return j + strconv.Itoa(n)
}
