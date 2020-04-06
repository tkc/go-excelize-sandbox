package main

import (
	"tkc/go-excelize-sandbox/infrastructure/lamdba"
)

func main() {
	excel := lamdba.NewlamdbaInfrastructure()
	excel.Start()
}
