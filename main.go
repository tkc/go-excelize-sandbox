package main

import (
	"tkc/go-excelize-sandbox/src/infrastructure/http"
	"tkc/go-excelize-sandbox/src/infrastructure/lamdba"
	"tkc/go-excelize-sandbox/src/usecase"
)

var isHttp = false

func main() {
	var (
		excelUsecase = usecase.NewExcelUsecase()
	)
	if isHttp {
		excel := lamdba.NewlamdbaInfrastructure(&excelUsecase)
		excel.Start()
	} else {
		http := http.NewHttpInfrastructure(&excelUsecase)
		http.Start()
	}
}
