package agent

import (
	"agent/errors"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestAgent(t *testing.T) {

	testCases := []struct {
		Description      string
		GivenURL         string
		ExpectedError    error
		ExpectedResponse *http.Response
	}{
		{
			Description:   "happy path",
			GivenURL:      "https://httpbin.org/base64/SFRUUEJJTiBpcyBhd2Vzb21l",
			ExpectedError: nil,
			ExpectedResponse: &http.Response{
				Body: ioutil.NopCloser(strings.NewReader("HTTPBIN is awesome")),
			},
		},
		{
			Description:   "happy path/no scheme",
			GivenURL:      "httpbin.org/base64/SFRUUEJJTiBpcyBhd2Vzb21l",
			ExpectedError: nil,
			ExpectedResponse: &http.Response{
				Body: ioutil.NopCloser(strings.NewReader("HTTPBIN is awesome")),
			},
		},
		{
			Description:      "failed path/invalid url",
			GivenURL:         "invalidurl\n",
			ExpectedError:    errors.ErrInvalidURL,
			ExpectedResponse: nil,
		},
		{
			Description:      "failed path/get failed",
			GivenURL:         "localhost",
			ExpectedError:    errors.ErrConnectionFailed,
			ExpectedResponse: nil,
		},
		{
			Description:      "failed path/nil url",
			GivenURL:         "",
			ExpectedError:    errors.ErrConnectionFailed,
			ExpectedResponse: nil,
		},
	}

	for _, testCase := range testCases {
		agent := NewAgent()
		t.Run(testCase.Description, func(t *testing.T) {
			resp, err := agent.MakeRequest(testCase.GivenURL)

			if testCase.ExpectedError == nil && err != nil {
				t.Error("Assertion fails, actual error is not nil")
			}
			if err != testCase.ExpectedError {
				t.Errorf("Assertion fails, expected error: %v actual error: %v", testCase.ExpectedError, err)
			}
			if testCase.ExpectedResponse != nil && resp == nil {
				t.Errorf("Assertion fails, expected response: %v actual response: %v", testCase.ExpectedResponse, nil)
			}

			if testCase.ExpectedResponse != nil && resp != nil && testCase.ExpectedResponse.Body != resp.Body {
				t.Errorf("Assertion fails, expected body: %v actual body: %v", testCase.ExpectedResponse, resp.Body)
			}
		})
	}
}
