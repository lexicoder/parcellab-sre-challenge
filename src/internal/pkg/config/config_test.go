// Tests for the config package.
package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConfigDefaultConfig(t *testing.T) {
	cfg := NewConfig()
	assert.Equal(t, cfg.ListenPort, "5000")
	assert.Equal(t, cfg.ServerIdleTimeout, 30*time.Second)
	assert.Equal(t, cfg.ServerReadTimeout, 15*time.Second)
	assert.Equal(t, cfg.ServerWriteTimeout, 15*time.Second)
}

func TestConfigListenPortEnv(t *testing.T) {
	t.Setenv("SERVER_PORT", "8080")
	cfg := NewConfig()
	assert.Equal(t, cfg.ListenPort, "8080")
}

func TestConfigServerIdleTimeoutEnv(t *testing.T) {
	t.Setenv("SERVER_IDLE_TIMEOUT", "1m")
	cfg := NewConfig()
	assert.Equal(t, cfg.ServerIdleTimeout, 1*time.Minute)
}

func TestConfigServerReadTimeoutEnv(t *testing.T) {
	t.Setenv("SERVER_READ_TIMEOUT", "2m")
	cfg := NewConfig()
	assert.Equal(t, cfg.ServerReadTimeout, 2*time.Minute)
}

func TestConfigServerWriteTimeoutEnv(t *testing.T) {
	t.Setenv("SERVER_WRITE_TIMEOUT", "3m")
	cfg := NewConfig()
	assert.Equal(t, cfg.ServerWriteTimeout, 3*time.Minute)
}

func TestConfigSalutationEnv(t *testing.T) {
	t.Setenv("SALUTATION", "Random Salutation")
	cfg := NewConfig()
	assert.Equal(t, cfg.Salutation, "Random Salutation")
}
