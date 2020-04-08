package http

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
	"tkc/go-excelize-sandbox/src/domain/model"
	"tkc/go-excelize-sandbox/src/infrastructure/param"
	"tkc/go-excelize-sandbox/src/infrastructure/types"
	"tkc/go-excelize-sandbox/src/usecase"
	"unsafe"

	"github.com/bxcodec/faker/v3"
)

type httpInfrastructure struct {
	excelUsecase     usecase.ExcelUsecase
	excelParamParser param.ExcelParamParser
}

type HttpInfrastructure interface {
	Start()
}

func NewHttpInfrastructure(excelUsecase usecase.ExcelUsecase, excelParamParser param.ExcelParamParser) HttpInfrastructure {
	return &httpInfrastructure{
		excelUsecase:     excelUsecase,
		excelParamParser: excelParamParser,
	}
}

func CreateDummyParam() (*string, error) {
	fakeDate := time.Now()
	excelParamParser := param.NewExcelParamParser()

	dummyId := 1
	excel := model.Excel{
		UserID:           dummyId,
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
		ID:        &dummyId,
		UserID:    &dummyId,
		CreatedAt: &fakeDate,
	}

	JoinUsers := []*model.JoinUser{&joinUser}
	excelData := make(map[int]map[int]map[int]*model.Excel)
	excelData[1] = make(map[int]map[int]*model.Excel)
	excelData[1][0] = make(map[int]*model.Excel)
	excelData[1][0][1] = &excel

	excelParam := types.ExcelRequestType{
		ClientName: faker.Name(),
		StartJST:   fakeDate,
		ExcelData:  excelData,
		JoinUser:   JoinUsers,
	}

	json, err := excelParamParser.EncodeJsonParam(excelParam)
	if err != nil {
		errStr := ""
		return &errStr, err
	}
	return json, nil
}

func (h *httpInfrastructure) Start() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		url := "http://localhost:8080/gen"
		json, err := CreateDummyParam()
		if err != nil {
			http.Error(w, "Error", http.StatusConflict)
		}
		var jsonStr = []byte(*json)
		req, err := http.NewRequest(
			"POST",
			url,
			bytes.NewBuffer(jsonStr),
		)
		if err != nil {
			http.Error(w, "Error", http.StatusConflict)
		}
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		byteArray, _ := ioutil.ReadAll(resp.Body)
		decoded, _ := base64.StdEncoding.DecodeString(*(*string)(unsafe.Pointer(&byteArray)))
		t := time.Now().In(time.FixedZone("Asia/Tokyo", 9*60*60))
		downloadName := fmt.Sprintf("%d%02d%02d%02d%02d%02d.xlsx", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
		w.Header().Set("Content-Description", "File Transfer")
		w.Header().Set("Content-Transfer-Encoding", "binary")
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename="+downloadName)
		w.Write(decoded)
	})

	http.HandleFunc("/lamdba_tset", func(w http.ResponseWriter, r *http.Request) {
		url := "http://localhost:3000/gen"
		json, err := CreateDummyParam()
		if err != nil {
			http.Error(w, "Error", http.StatusConflict)
		}
		var jsonStr = []byte(*json)
		req, err := http.NewRequest(
			"POST",
			url,
			bytes.NewBuffer(jsonStr),
		)
		if err != nil {
			http.Error(w, "Error", http.StatusConflict)
		}
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		decoded, _ := base64.StdEncoding.DecodeString(*(*string)(unsafe.Pointer(&byteArray)))
		t := time.Now().In(time.FixedZone("Asia/Tokyo", 9*60*60))
		downloadName := fmt.Sprintf("%d%02d%02d%02d%02d%02d.xlsx", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
		w.Header().Set("Content-Description", "File Transfer")
		w.Header().Set("Content-Transfer-Encoding", "binary")
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename="+downloadName)
		w.Write(decoded)
	})

	http.HandleFunc("/gen", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Error Method", http.StatusForbidden)
		}
		length, err := strconv.Atoi(r.Header.Get("Content-Length"))
		if err != nil {
			log.Print(err)
			http.Error(w, "Error Content-Length", http.StatusConflict)
		}
		body := make([]byte, length)

		length, err = r.Body.Read(body)
		if err != nil && err != io.EOF {
			http.Error(w, "Error Body.Read", http.StatusConflict)
		}

		var jsonBody map[string]interface{}
		err = json.Unmarshal(body[:length], &jsonBody)
		if err != nil {
			http.Error(w, "Error json.Unmarshal", http.StatusConflict)
		}

		excelRequestType, err := h.excelParamParser.DecodeJsonParam(string(body))
		if err != nil {
			http.Error(w, "Error DecodeJsonParam", http.StatusConflict)
		}

		data, err := h.excelUsecase.CreateExcelByte(*excelRequestType)
		if err != nil {
			http.Error(w, "Error CreateExcelByte", http.StatusConflict)
		}
		encoded := base64.StdEncoding.EncodeToString(data)
		fmt.Fprintf(w, encoded)
	})

	log.Print("http serve start")
	http.ListenAndServe(":8080", nil)
}
