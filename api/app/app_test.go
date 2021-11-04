package app

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApp(t *testing.T) {
	tests := []struct {
		method        string
		description   string
		route         string
		expectedCode  int
		expectedError bool
		expectedBody  string
	}{
		{
			method:        "POST",
			description:   "hello route",
			route:         "/public/auth",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "{\"message\":\"Hello world!\",\"success\":true}",
		},
		{
			method:        "GET",
			description:   "error 404",
			route:         "/go-rest",
			expectedError: false,
			expectedCode:  404,
			expectedBody:  "Cannot GET /go-rest",
		},
	}

	app := Setup()

	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.route, nil)
		res, err := app.Test(req, 1)

		assert.Equalf(t, test.expectedError, err != nil, test.description)
		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		body, err := ioutil.ReadAll(res.Body)
		assert.Nilf(t, err, test.description)
		assert.Equalf(t, test.expectedBody, string(body), test.description)
	}
}
