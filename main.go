package main

import (
	"tkc/go-excelize-sandbox/src/infrastructure/lamdba"
	"tkc/go-excelize-sandbox/src/usecase"
)

func main() {
	excelUsecase := usecase.NewExcelUsecase()
	excel := lamdba.NewlamdbaInfrastructure(&excelUsecase)
	excel.Start()
}
