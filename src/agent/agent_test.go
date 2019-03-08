package agent

import (
	"testing"
)

func TestAgent(t *testing.T) {

	testCases := []struct {
		Description   string
		GivenURL      string
		ExpectedError bool
	}{
		{
			Description:   "happy path",
			GivenURL:      "https://httpbin.org/base64/SFRUUEJJTiBpcyBhd2Vzb21l",
			ExpectedError: false,
		},
		{
			Description:   "happy path/no scheme",
			GivenURL:      "httpbin.org/base64/SFRUUEJJTiBpcyBhd2Vzb21l",
			ExpectedError: false,
		},
		{
			Description:   "failed path/invalid url",
			GivenURL:      "invalid\nURL",
			ExpectedError: true,
		},
		{
			Description:   "failed path/get failed",
			GivenURL:      "localhost",
			ExpectedError: true,
		},
		{
			Description:   "failed path/nil url",
			GivenURL:      "",
			ExpectedError: true,
		},
	}

	for _, testCase := range testCases {
		agent := NewAgent()
		t.Run(testCase.Description, func(t *testing.T) {
			_, err := agent.MakeRequest(testCase.GivenURL)

			if !testCase.ExpectedError && err != nil {
				t.Error("Assertion fails, actual error is not nil")
			} else if testCase.ExpectedError && err == nil {
				t.Error("Assertion fails, actual error is nil")
			}
		})
	}
}
