package lamdba

import (
	"context"
	"tkc/go-excelize-sandbox/domain/model"
	"tkc/go-excelize-sandbox/usecase"

	"tkc/go-excelize-sandbox/usecase"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type lamdbaInfrastructure struct {
	excelUsecase *usecase.ExcelUsecase
}

type LamdbaInfrastructure interface {
	Serve()
}

func NewlamdbaInfrastructure(excelUsecase *usecase.ExcelUsecase) LamdbaInfrastructure {
	return &lamdbaInfrastructure{
		excelUsecase: excelUsecase,
	}
}

func (h *httpInfrastructure) Serve() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, param model.ExcelParam) ([]byte, error) {
	return events.APIGatewayProxyResponse{
		Body:       []byte{},
		StatusCode: 200,
	}, nil
}
