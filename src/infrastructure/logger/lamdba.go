package logger

import (
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/getsentry/sentry-go"
)

type lamdbaLogger struct{}

type LamdbaLogger interface {
	Capture(request events.APIGatewayProxyRequest, err error)
}

func NewLamdbaLogger(sentryDsn string) (LamdbaLogger, error) {
	err := sentry.Init(sentry.ClientOptions{Dsn: sentryDsn})
	if err != nil {
		log.Printf("Sentry Initialize Error: %s", err.Error())
		return &lamdbaLogger{}, err
	}
	return &lamdbaLogger{}, nil
}

func (h *lamdbaLogger) Capture(
	event events.APIGatewayProxyRequest,
	err error,
) {
	// Sentry
	sentry.CaptureException(err)
	// CloudWatch
	eventJSON, _ := json.MarshalIndent(event, "", "  ")
	log.Printf("EVENT: %s", eventJSON)
	log.Printf("Error: %s", err.Error())
	log.Printf("REGION: %s", os.Getenv("AWS_REGION"))

	log.Println("ALL ENV VARS:")
	for _, element := range os.Environ() {
		log.Println(element)
	}
}
