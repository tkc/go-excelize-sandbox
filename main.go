package main

import (
	"tkc/go-excelize-sandbox/src/infrastructure/lamdba"
)

func main() {
	excel := lamdba.NewlamdbaInfrastructure()
	excel.Start()
}
