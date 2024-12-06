package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headerValue   string
		expectedToken string
		expectedError bool
	}{
		{
			name:          "valid api key",
			headerValue:   "ApiKey 923u10d01dj1dh1rh",
			expectedToken: "923u10d01dj1dh1rh",
			expectedError: false,
		},
		{
			name:          "invalid api key",
			headerValue:   "Bearer mytoken",
			expectedToken: "",
			expectedError: true,
		},
		{
			name:          "extra arguments provided",
			headerValue:   "ApiKey mytoken extravalue",
			expectedToken: "mytoken",
			expectedError: false,
		},
		{
			name:          "too little arguments provided",
			headerValue:   "ApiKey",
			expectedToken: "",
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}
			if tt.headerValue != "" {
				headers.Add("Authorization", tt.headerValue)
			}
			token, err := GetAPIKey(headers)

			if tt.expectedError && err == nil {
				t.Error("expected error but got none")
			}

			if !tt.expectedError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if token != tt.expectedToken {
				t.Errorf("got key %q, wanted %q", token, tt.expectedToken)
			}
		})
	}
}
