package main

import (
	"testing"
)

func TestCheckNumber(t *testing.T) {
	tt := []struct {
		description string
		value       string
		message     string
		e           error
	}{
		{description: "empty value", value: "", message: "", e: ErrInvalidNumber},
		{description: "random string", value: "foo", message: "", e: ErrInvalidNumber},
		{description: "boolean", value: "false", message: "", e: ErrInvalidNumber},
		{description: "integer", value: "1", message: "The value <1.000000> is a valid number", e: nil},
		{description: "float", value: "50.5", message: "The value <50.500000> is a valid number", e: nil},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			msg, err := checkNumber(test.value)

			if err != test.e {
				t.Fatalf("expected error %v; got %v", test.e, err)
			}

			if msg != test.message {
				t.Fatalf("expected message %v; got %v", test.message, msg)
			}
		})
	}
}
