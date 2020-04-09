.PHONY: deps clean build

deps:
	go mod tidy -v
	go get -u ./...

clean: 
	rm -rf ./aws-sam-go/app

sam:
	GOOS=linux GOARCH=amd64 go build -o excelize/excelize ./main.go
	cp ./format.xlsx excelize/format.xlsx
	sam local start-api
