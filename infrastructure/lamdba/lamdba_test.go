package lamdba

import (
	"log"
	"testing"
	"tkc/go-excelize-sandbox/usecase"
)

func Test_serve(t *testing.T) {
	var (
		excelUsecase = usecase.NewExcelUsecase()
	)
	l := NewlamdbaInfrastructure(&excelUsecase)
	log.Print(l)
}
