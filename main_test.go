package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListenString(t *testing.T) {
	cfg := NewConfiguration()
	defer func() {
		t.Setenv("FLUENTBIT_PORT", "8080")
	}()
	_ = os.Setenv("FLUENTBIT_PORT", "")
	s := ":" + cfg.FluentBitPort
	assert.Equal(t, s, ":8080")
	t.Setenv("FLUENTBIT_PORT", "1028")
	s = ":" + os.Getenv("FLUENTBIT_PORT")
	assert.Equal(t, s, ":1028")
}
