package lamdba

import (
	"net/http"
	"tkc/go-excelize-sandbox/src/infrastructure/param"
	"tkc/go-excelize-sandbox/src/usecase"
	"unsafe"

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
	var data = []byte("はろー")

	excelRequestType, err := h.excelParamParser.DecodeJsonParam(request.Body)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusConflict,
		}, nil
	}

	data, err = h.excelUsecase.CreateExcelByte(*excelRequestType)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusConflict,
		}, nil
	}

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
	lambda.Start(h.handler)
}
