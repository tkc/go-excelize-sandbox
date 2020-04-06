.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./aws-sam-go/app

serve:
    GOOS=linux GOARCH=amd64 go build -o aws-sam-go/app ./main.go
	sam local start-api