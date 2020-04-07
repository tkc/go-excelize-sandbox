package lamdba

import (
	"tkc/go-excelize-sandbox/src/usecase"
	"unsafe"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type lamdbaInfrastructure struct {
	excelUsecase *usecase.ExcelUsecase
}

type LamdbaInfrastructure interface {
	Start()
}

func NewlamdbaInfrastructure(excelUsecase *usecase.ExcelUsecase) LamdbaInfrastructure {
	return &lamdbaInfrastructure{excelUsecase: excelUsecase}
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var data = []byte("はろー")
	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Description":       "File Transfer",
			"Content-Transfer-Encoding": "binary",
			"Content-Type":              "application/octet-stream",
		},
		Body:            *(*string)(unsafe.Pointer(&data)),
		StatusCode:      200,
		IsBase64Encoded: true,
	}, nil
}

func (h *lamdbaInfrastructure) Start() {
	lambda.Start(handler)
}
