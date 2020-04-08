package web

import (
	"log"
	"testing"
	"tkc/go-excelize-sandbox/src/infrastructure/param"
	"tkc/go-excelize-sandbox/src/usecase"
)

func Test_serve(t *testing.T) {
	var (
		excelUsecase     = usecase.NewExcelUsecase()
		excelParamParser = param.NewExcelParamParser()
	)
	s := NewHTTPInfrastructure(excelUsecase, excelParamParser)
	log.Print(s)
}
