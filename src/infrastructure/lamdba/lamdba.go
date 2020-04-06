package lamdba

import (
	"fmt"
	"tkc/go-excelize-sandbox/src/usecase"

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
	return &lamdbaInfrastructure{
		excelUsecase: excelUsecase,
	}
}

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("OK"),
		StatusCode: 200,
	}, nil
}

func (h *lamdbaInfrastructure) Start() {
	lambda.Start(HandleRequest)
}