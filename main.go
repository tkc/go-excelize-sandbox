package main

import (
	"os"

	"tkc/go-excelize-sandbox/src/infrastructure/http"
	"tkc/go-excelize-sandbox/src/infrastructure/lamdba"
	"tkc/go-excelize-sandbox/src/infrastructure/param"
	"tkc/go-excelize-sandbox/src/usecase"
)

func main() {
	var (
		excelUsecase     = usecase.NewExcelUsecase()
		excelParamParser = param.NewExcelParamParser()
	)
	if len(os.Getenv("AWS_REGION")) > 0 {
		app := lamdba.NewlamdbaInfrastructure(excelUsecase, excelParamParser)
		app.Start()
	} else {
		app := http.NewHttpInfrastructure(excelUsecase, excelParamParser)
		app.Start()
	}
}
