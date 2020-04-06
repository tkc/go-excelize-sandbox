package main

import (
	"tkc/go-excelize-sandbox/infrastructure/lamdba"
)

// var (
// 	// DefaultHTTPGetAddress Default Address
// 	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

// 	// ErrNoIP No IP found in response
// 	ErrNoIP = errors.New("No IP in HTTP response")

// 	// ErrNon200Response non 200 status code in response
// 	ErrNon200Response = errors.New("Non 200 Response found")
// )

// func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
// 	resp, err := http.Get(DefaultHTTPGetAddress)
// 	if err != nil {
// 		return events.APIGatewayProxyResponse{}, err
// 	}

// 	if resp.StatusCode != 200 {
// 		return events.APIGatewayProxyResponse{}, ErrNon200Response
// 	}

// 	ip, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return events.APIGatewayProxyResponse{}, err
// 	}

// 	if len(ip) == 0 {
// 		return events.APIGatewayProxyResponse{}, ErrNoIP
// 	}

// 	f := excelize.NewFile()
// 	index := f.NewSheet("Sheet1")
// 	f.SetCellValue("Sheet1", "B2", 100)
// 	f.SetActiveSheet(index)
// 	err = f.SaveAs("./Book1.xlsx")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	return events.APIGatewayProxyResponse{
// 		Body:       fmt.Sprintf("OK, %v", string(ip)),
// 		StatusCode: 200,
// 	}, nil
// }

func main() {
	// a := model.JoinUser{}
	excel := lamdba.NewlamdbaInfrastructure()
	excel.Start()
}
