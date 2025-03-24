package jsonclean_test

import (
	"reflect"
	"testing"

	"github.com/mrumyantsev/pastebin-app/internal/jsonclean"
)

func TestClean(t *testing.T) {
	testTable := []struct {
		Input, Expected []byte
	}{
		{
			Input: []byte(`{
	"name": "Gordon F.",
	"pass": "dr.  k l e i n",
	"email": "crow@bar.com"
}`),
			Expected: []byte(`{"name":"Gordon F.","pass":"dr.  k l e i n","email":"crow@bar.com"}`),
		},
	}

	for _, test := range testTable {
		res := jsonclean.Clean(test.Input)

		if !reflect.DeepEqual(res, test.Expected) {
			t.Fatalf("got: %s, expected: %s", res, test.Expected)
		}
	}
}
