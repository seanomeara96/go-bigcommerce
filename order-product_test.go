package bigcommerce

import (
	"encoding/json"
	"testing"
)

func TestStringOrNumberUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      StringOrNumber
		wantError bool
	}{
		{name: "string", input: `"1.2500"`, want: "1.2500"},
		{name: "number", input: `1.25`, want: "1.25"},
		{name: "null", input: `null`, wantError: true},
		{name: "object", input: `{}`, wantError: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var got StringOrNumber
			err := json.Unmarshal([]byte(test.input), &got)
			if test.wantError {
				if err == nil {
					t.Fatal("expected an error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != test.want {
				t.Fatalf("got %q, want %q", got, test.want)
			}
		})
	}
}
