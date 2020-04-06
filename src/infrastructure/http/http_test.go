package http

import (
	"log"
	"testing"
	"tkc/go-excelize-sandbox/src/usecase"
)

func Test_serve(t *testing.T) {
	var (
		excelUsecase = usecase.NewExcelUsecase()
	)
	s := NewHttpInfrastructure(&excelUsecase)
	log.Print(s)
}
