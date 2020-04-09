package lamdba

import (
	"encoding/base64"
	"net/http"
	"tkc/go-excelize-sandbox/src/infrastructure/logger"
	"tkc/go-excelize-sandbox/src/infrastructure/param"
	"tkc/go-excelize-sandbox/src/usecase"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type lamdbaInfrastructure struct {
	excelUsecase     usecase.ExcelUsecase
	excelParamParser param.ExcelParamParser
	LamdbaLogger     logger.LamdbaLogger
}

type Infrastructure interface {
	Start()
	handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}

func NewlamdbaInfrastructure(
	excelUsecase usecase.ExcelUsecase,
	excelParamParser param.ExcelParamParser,
	LamdbaLogger logger.LamdbaLogger,
) Infrastructure {
	return &lamdbaInfrastructure{
		excelUsecase:     excelUsecase,
		excelParamParser: excelParamParser,
		LamdbaLogger:     LamdbaLogger,
	}
}

func (h *lamdbaInfrastructure) handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if request.HTTPMethod != "POST" {
		return events.APIGatewayProxyResponse{
			Body:       "Error Bad Method",
			StatusCode: http.StatusForbidden,
		}, nil
	}

	excelRequestType, err := h.excelParamParser.DecodeJSONParam(request.Body)
	if err != nil {
		h.LamdbaLogger.Capture(request, err)
		return events.APIGatewayProxyResponse{
			Body:       "Error DecodeJsonParam ",
			StatusCode: http.StatusConflict,
		}, nil
	}

	data, err := h.excelUsecase.CreateExcelByte(*excelRequestType)
	if err != nil {
		h.LamdbaLogger.Capture(request, err)
		return events.APIGatewayProxyResponse{
			Body:       "Error CreateExcelByte",
			StatusCode: http.StatusConflict,
		}, nil
	}

	encoded := base64.StdEncoding.EncodeToString(data)
	return events.APIGatewayProxyResponse{
		Body:            encoded,
		StatusCode:      http.StatusOK,
		IsBase64Encoded: false,
	}, nil
}

func (h *lamdbaInfrastructure) Start() {
	lambda.Start(h.handler)
}
