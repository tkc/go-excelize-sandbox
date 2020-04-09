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
	excelInteractor  usecase.ExcelInteractor
	excelParamParser param.ExcelParamParser
	lamdbaLogger     logger.LamdbaLogger
}

type Infrastructure interface {
	Start()
	handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}

func NewlamdbaInfrastructure(
	excelInteractor usecase.ExcelInteractor,
	excelParamParser param.ExcelParamParser,
	lamdbaLogger logger.LamdbaLogger,
) Infrastructure {
	return &lamdbaInfrastructure{
		excelInteractor:  excelInteractor,
		excelParamParser: excelParamParser,
		lamdbaLogger:     lamdbaLogger,
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
		h.lamdbaLogger.Capture(request, err)
		return events.APIGatewayProxyResponse{
			Body:       "Error DecodeJsonParam ",
			StatusCode: http.StatusConflict,
		}, nil
	}

	data, err := h.excelInteractor.CreateExcelByte(*excelRequestType)
	if err != nil {
		h.lamdbaLogger.Capture(request, err)
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
