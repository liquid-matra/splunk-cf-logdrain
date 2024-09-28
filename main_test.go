package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListenString(t *testing.T) {
	cfg := NewConfiguration()
	defer func() {
		t.Setenv("LISTEN_PORT", "8080")

	}()
	_ = os.Setenv("LISTEN_PORT", "")
	s := ":" + cfg.ListenPort
	assert.Equal(t, s, ":8080")
	t.Setenv("LISTEN_PORT", "1028")
	s = ":" + os.Getenv("LISTEN_PORT")
	t.Logf("port (s) is: %s", ":"+cfg.ListenPort)
	assert.Equal(t, s, ":1028")
}
