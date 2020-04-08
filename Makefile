.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./aws-sam-go/app

sam:
	GOOS=linux GOARCH=amd64 go build -o sam-go/app ./main.go
	cp format.xlsx sam-go/format.xlsx
	sam local start-api
