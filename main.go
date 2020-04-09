package main

import (
	"os"

	"tkc/go-excelize-sandbox/src/infrastructure/lamdba"
	"tkc/go-excelize-sandbox/src/infrastructure/logger"
	"tkc/go-excelize-sandbox/src/infrastructure/param"
	"tkc/go-excelize-sandbox/src/infrastructure/web"
	"tkc/go-excelize-sandbox/src/usecase"
)

func main() {
	var (
		excelUsecase     = usecase.NewExcelUsecase()
		excelParamParser = param.NewExcelParamParser()
	)

	if len(os.Getenv("AWS_REGION")) > 0 {
		sentryDNS := os.Getenv("SENTY_DNS")
		lamdbaLogger, err := logger.NewLamdbaLogger(sentryDNS)
		if err != nil {
			panic("SENTY_DNS Not found")
		}
		app := lamdba.NewlamdbaInfrastructure(
			excelUsecase,
			excelParamParser,
			lamdbaLogger,
		)
		app.Start()
	} else {
		app := web.NewHTTPInfrastructure(excelUsecase, excelParamParser)
		app.Start()
	}
}
