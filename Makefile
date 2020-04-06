.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./aws-sam-go/app
	
build:
    GOOS=linux GOARCH=amd64 go build -o aws-sam-go/app ./main.go 