package main

import (
	"tkc/go-excelize-sandbox/src/infrastructure/http"
	"tkc/go-excelize-sandbox/src/infrastructure/lamdba"
	"tkc/go-excelize-sandbox/src/infrastructure/param"
	"tkc/go-excelize-sandbox/src/usecase"
)

var isHttp = false

func main() {
	var (
		excelUsecase     = usecase.NewExcelUsecase()
		excelParamParser = param.NewExcelParamParser()
	)
	if isHttp {
		excel := lamdba.NewlamdbaInfrastructure(excelUsecase, excelParamParser)
		excel.Start()
	} else {
		http := http.NewHttpInfrastructure(excelUsecase, excelParamParser)
		http.Start()
	}
}
