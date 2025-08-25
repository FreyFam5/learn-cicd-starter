package auth

import (
	"net/http"
	"reflect"
	"testing"
)

var randomKey = "ThisIsARandomKey"

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		headerKeyName string
		headerKeyValue string
		want string
	}{
		"get key normal": {"Authorization", "ApiKey " + randomKey, randomKey},
		"get key with more options": {"Authorization", "ApiKey " + randomKey + " This is extra", randomKey},
		"get key broken": {"Authorization", "ApiKey ", randomKey},
	}

	for name, curTest := range tests {
		t.Run(name, func(t *testing.T) {
			header := http.Header{}
			header.Add(curTest.headerKeyName, curTest.headerKeyValue)
			got, err := GetAPIKey(header)
			if err != nil {
				t.Fatalf("%s got an internal error: %s", name, err)
			}

			if !reflect.DeepEqual(curTest.want, got) {
				t.Fatalf("expected: %v, got: %v", curTest.want, got)
			}
		})
	}
}