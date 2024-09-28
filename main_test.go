package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvironmentVariables(t *testing.T) {

	type TestCase struct {
		description string
		input       string
		want        string
	}

	testCases := []TestCase{
		{description: "the default port of 8080 is used if nothing specified", input: "", want: "8080"},
		{description: "the provided port is used", input: "1028", want: "1028"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			t.Setenv("FLUENTBIT_PORT", testCase.input)
			cfg := NewConfiguration()
			got := cfg.FluentBitPort
			t.Logf("VARIABLE IS %s but expected %s", cfg.FluentBitPort, testCase.want)
			assert.Equal(t, testCase.want, got)
		})
	}

}
