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
		excelInteractor  = usecase.NewExcelInteractor()
		excelParamParser = param.NewExcelParamParser()
	)

	if len(os.Getenv("AWS_REGION")) > 0 {
		sentryDNS := os.Getenv("SENTY_DNS")
		lamdbaLogger, err := logger.NewLamdbaLogger(sentryDNS)
		if err != nil {
			panic("SENTY_DNS Not found")
		}
		app := lamdba.NewlamdbaInfrastructure(
			excelInteractor,
			excelParamParser,
			lamdbaLogger,
		)
		app.Start()
	} else {
		app := web.NewHTTPInfrastructure(excelInteractor, excelParamParser)
		app.Start()
	}
}
