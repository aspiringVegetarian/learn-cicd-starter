package auth

import (
	"errors"
	"net/http"
	"testing"
)

// GetAPIKey -
func TestGetAPIKey(t *testing.T) {
	testHeaderPass := http.Header{}
	testHeaderPass.Add("Authorization", "ApiKey Pass")

	testHeaderNoAuth := http.Header{}

	testHeaderMalformedFirstElement := http.Header{}
	testHeaderMalformedFirstElement.Add("Authorization", "ApiKe Fail")

	testHeaderMalformedLength := http.Header{}
	testHeaderMalformedLength.Add("Authorization", "ApiKeyFail")

	tests := map[string]struct {
		input http.Header
		want  string
		err   error
	}{
		"Header Auth Pass":                    {input: testHeaderPass, want: "Pass", err: nil},
		"Header Auth Empty":                   {input: testHeaderNoAuth, want: "", err: ErrNoAuthHeaderIncluded},
		"Header Auth Malformed First Element": {input: testHeaderMalformedFirstElement, want: "", err: ErrMalformedHeader},
		"Header Auth Malformed Length":        {input: testHeaderMalformedLength, want: "", err: ErrMalformedHeader},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.input)
			if tc.want != got || !errors.Is(err, tc.err) {
				t.Fatalf("We received %s with %v, when we expected %s with %v", got, err, tc.want, tc.err)
			}
		})
	}
}
