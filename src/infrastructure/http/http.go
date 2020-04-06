package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"tkc/go-excelize-sandbox/src/usecase"
)

type httpInfrastructure struct {
	excelUsecase *usecase.ExcelUsecase
}

type HttpInfrastructure interface {
	Start()
}

func NewHttpInfrastructure(excelUsecase *usecase.ExcelUsecase) HttpInfrastructure {
	return &httpInfrastructure{
		excelUsecase: excelUsecase,
	}
}

func (h *httpInfrastructure) Start() {
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		url := "http://localhost:8080"
		resp, _ := http.Get(url)
		defer resp.Body.Close()
		byteArray, _ := ioutil.ReadAll(resp.Body)
		t := time.Now().In(time.FixedZone("Asia/Tokyo", 9*60*60))
		downloadName := fmt.Sprintf("%d%02d%02d%02d%02d%02d.xlsx", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
		w.Header().Set("Content-Description", "File Transfer")
		w.Header().Set("Content-Transfer-Encoding", "binary")
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename="+downloadName)
		w.Write(byteArray)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// excel, err := h.excelUsecase.
		// if err != nil {
		// }
		var data = []byte("hellow")
		w.Header().Set("Content-Description", "File Transfer")
		w.Header().Set("Content-Transfer-Encoding", "binary")
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(data)
	})
	http.ListenAndServe(":8080", nil)
}
