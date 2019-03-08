package coordinator

import (
	"net/http"
	"reflect"
	"testing"
)

func TestHashedBody(t *testing.T) {
	testCases := []struct {
		Description  string
		GivenBody    []byte
		ExpectedHash string
	}{
		{
			Description:  "happy path",
			GivenBody:    []byte("HTTPBIN is awesome"),
			ExpectedHash: "d93bf0bc80a7de7e8968579fc7aafb8a",
		},
	}

	for _, testCase := range testCases {
		hash := hashResponse(testCase.GivenBody)

		if hash != testCase.ExpectedHash {
			t.Errorf("Assertion fails, expected hash: %v actual hash: %v", testCase.ExpectedHash, hash)
		}
	}
}

func TestGetResponseBody(t *testing.T) {
	testCases := []struct {
		Description   string
		GivenURL      string
		ExpectedBody  []byte
		ExpectedError error
	}{
		{
			Description:   "happy path",
			GivenURL:      "https://httpbin.org/base64/SFRUUEJJTiBpcyBhd2Vzb21l",
			ExpectedBody:  []byte("HTTPBIN is awesome"),
			ExpectedError: nil,
		},
	}

	for _, testCase := range testCases {
		resp, err := http.DefaultClient.Get(testCase.GivenURL)
		if testCase.ExpectedError == nil {
			if err != nil {
				t.Error("Assertion fails, actual error is not nil")
			}

			body, err := getResponseBody(resp)
			if err != nil {
				t.Error("Assertion fails, actual error is not nil")
			}

			if !reflect.DeepEqual(body, testCase.ExpectedBody) {
				t.Errorf("Assertion fails, expected body: %v actual body: %v", testCase.ExpectedBody, body)
			}
		}
	}
}
