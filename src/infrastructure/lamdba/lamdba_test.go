package lamdba

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"tkc/go-excelize-sandbox/src/infrastructure/param"
	"tkc/go-excelize-sandbox/src/usecase"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {

	var (
		excelUsecase     = usecase.NewExcelUsecase()
		excelParamParser = param.NewExcelParamParser()
	)

	excelLamdba := NewlamdbaInfrastructure(excelUsecase, excelParamParser)

	t.Run("Unable to get IP", func(t *testing.T) {
		res, err := excelLamdba.handler(events.APIGatewayProxyRequest{})
		if err != nil {
			t.Fatal("Everything should be ok")
		}
		log.Print(res)
	})

	t.Run("Successful Request", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
		}))
		defer ts.Close()
		_, err := excelLamdba.handler(events.APIGatewayProxyRequest{})
		if err != nil {
			t.Fatal("Everything should be ok")
		}
	})
}
