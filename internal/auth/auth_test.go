package auth

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		key       string
		value     string
		expect    string
		expectErr string
	}{
		{
			expectErr: "no authorization header",
		},
		{
			key:       "Authorization",
			expectErr: "no authorization header",
		},
		{
			key:       "Authorization",
			value:     "-",
			expectErr: "malformed authorization header",
		},
		{
			key:       "Authorization",
			value:     "    ",
			expectErr: "malformed authorization header",
		},
		{
			key:       "Authorization",
			value:     "APIKEY xxxxxx",
			expectErr: "malformed authorization header",
		},
		{
			key:       "Authorization",
			value:     "ApiKey xxxxxx",
			expect:    "xxxxxx",
			expectErr: "not expecting an error",
		},
	}

	for index, test := range tests {
		t.Run(fmt.Sprintf("TestGetAPIKey - Case #%v:", index), func(t *testing.T) {

			header := http.Header{}
			header.Add(test.key, test.value)

			actual, err := GetAPIKey(header)
			if err != nil {
				if strings.Contains(err.Error(), test.expectErr) {
					return
				}

				t.Errorf("Unexpected err: %v\n", err)
				return
			}

			if test.expect != actual {
				t.Errorf("Unexpected result: %s", actual)
				return
			}

		})
	}
}
