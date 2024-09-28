package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListenString(t *testing.T) {
	cfg := NewConfiguration()
	port := os.Getenv("PORT")
	defer func() {
		_ = os.Setenv("PORT", port)
	}()
	_ = os.Setenv("PORT", "")
	s := ":" + cfg.ListenPort
	assert.Equal(t, s, ":8080")
	_ = os.Setenv("PORT", "1028")
	s = ":" + cfg.ListenPort
	assert.Equal(t, s, ":1028")
}
