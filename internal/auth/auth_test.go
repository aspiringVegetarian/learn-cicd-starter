package auth

import (
	"errors"
	"net/http"
	"testing"
)

// GetAPIKey -
func TestGetAPIKey(t *testing.T) {
	testHeaderOne := http.Header{}
	testHeaderOne.Add("Authorization", "ApiKey Pass")

	testHeaderTwo := http.Header{}
	testHeaderTwo.Add("Authorization", "ApiKe Fail")

	tests := map[string]struct {
		input http.Header
		want  string
		err   error
	}{
		"simple":  {input: testHeaderOne, want: "Pass", err: nil},
		"simple1": {input: testHeaderTwo, want: "", err: ErrMalformedHeader},
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
