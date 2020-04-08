package lamdba

import (
	"encoding/base64"
	"net/http"
	"tkc/go-excelize-sandbox/src/infrastructure/param"
	"tkc/go-excelize-sandbox/src/usecase"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type lamdbaInfrastructure struct {
	excelUsecase     usecase.ExcelUsecase
	excelParamParser param.ExcelParamParser
}

type LamdbaInfrastructure interface {
	Start()
	handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}

func NewlamdbaInfrastructure(excelUsecase usecase.ExcelUsecase, excelParamParser param.ExcelParamParser) LamdbaInfrastructure {
	return &lamdbaInfrastructure{
		excelUsecase:     excelUsecase,
		excelParamParser: excelParamParser,
	}
}

func (h *lamdbaInfrastructure) handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if request.HTTPMethod != "POST" {
		return events.APIGatewayProxyResponse{
			Body:       "Bad Method",
			StatusCode: http.StatusForbidden,
		}, nil
	}

	excelRequestType, err := h.excelParamParser.DecodeJsonParam(request.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "DecodeJsonParam Error",
			StatusCode: http.StatusConflict,
		}, nil
	}

	data, err := h.excelUsecase.CreateExcelByte(*excelRequestType)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "CreateExcelByte Error",
			StatusCode: http.StatusConflict,
		}, nil
	}

	encoded := base64.StdEncoding.EncodeToString(data)
	return events.APIGatewayProxyResponse{
		Body:            encoded,
		StatusCode:      200,
		IsBase64Encoded: false,
	}, nil
}

func (h *lamdbaInfrastructure) Start() {
	lambda.Start(h.handler)
}
