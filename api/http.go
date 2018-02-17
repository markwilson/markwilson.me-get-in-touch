package api

import (
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"regexp"
)

var (
	ErrorNoBody       = errors.New("No HTTP body")
	ErrorInvalidBody  = errors.New("Invalid JSON body")
	ErrorInvalidEmail = errors.New("Invalid email provided")

	emailRegExp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

type Request struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func (r Request) IsValid() bool {
	return !(len(r.Email) < 1 || len(r.Message) < 1)
}

func Response() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Body:       "{success:true}",
		StatusCode: 200,
	}
}

func RequestFromBody(body string) (Request, error) {
	var request Request
	err := json.Unmarshal([]byte(body), &request)
	if err != nil || !request.IsValid() {
		return Request{}, ErrorInvalidBody
	}

	if !emailRegExp.MatchString(request.Email) {
		return Request{}, ErrorInvalidEmail
	}

	return request, nil
}
