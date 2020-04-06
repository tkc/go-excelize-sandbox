package http

import (
	"fmt"
	"net/http"
	"time"
	"tkc/go-excelize-sandbox/src/usecase"
)

type httpInfrastructure struct {
	excelUsecase *usecase.ExcelUsecase
}

type HttpInfrastructure interface {
	Serve()
}

func NewHttpInfrastructure(excelUsecase *usecase.ExcelUsecase) HttpInfrastructure {
	return &httpInfrastructure{
		excelUsecase: excelUsecase,
	}
}

func (h *httpInfrastructure) Serve() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		excel, err := h.excelUsecase.CreateExcelByte(r.Body)
		if err != nil {

		}
		t := time.Now().In(time.FixedZone("Asia/Tokyo", 9*60*60))
		downloadName := fmt.Sprintf("%d%02d%02d%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", "attachment; filename="+downloadName)
		c.Data(http.StatusOK, "application/octet-stream", excel)
	})
	http.ListenAndServe(":8080", nil)
}
