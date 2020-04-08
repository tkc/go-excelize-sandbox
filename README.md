# go-excelize-sandbox

## Requirements

### Python 3.x

```
$ brew install python3
$ pip3 install virtualenv
$ python3 -m venv penv
$ source penv/bin/activate 
```

### Go 1.3.x

```
% go version  
=> go version go1.14 darwin/amd64
```

## Install AWS SAM

```
$ pip install awscli
$ pip install aws-sam-cli
```

## Test

```
$ go test ./...
```

### Serve Local Http Server

```
$ go run main.go
$ curl http://localhost:8080/test --output test_http.xlsx
```

### Serve Local Lamdba Server

```
$ python --version         
=> Python 3.7.3
$ make sam
@ curl -d @excel.json -H "Content-Type: application/json" http://localhost:3000/gen
```

## Appendix AWS SAM
create go lamdba workspace.

```
sam init --runtime go1.x --name aws-sam-golang1
```