package config

import (
	"os"
	"time"
)

type Config struct {
	ListenPort         string
	Salutation         string
	ServerWriteTimeout time.Duration
	ServerReadTimeout  time.Duration
	ServerIdleTimeout  time.Duration
}

func NewConfig() *Config {
	serverWriteTimeout, err := getDurationFromEnv("SERVER_WRITE_TIMEOUT", "15s")
	if err != nil {
		panic(err)
	}
	serverReadTimeout, err := getDurationFromEnv("SERVER_READ_TIMEOUT", "15s")
	if err != nil {
		panic(err)
	}
	serverIdleTimeout, err := getDurationFromEnv("SERVER_IDLE_TIMEOUT", "30s")
	if err != nil {
		panic(err)
	}

	cfg := Config{
		ListenPort:         getEnvOrDefault("SERVER_PORT", "5000"),
		Salutation:         getEnvOrDefault("SALUTATION", "Hi"),
		ServerWriteTimeout: serverWriteTimeout,
		ServerReadTimeout:  serverReadTimeout,
		ServerIdleTimeout:  serverIdleTimeout,
	}
	return &cfg
}

func getEnvOrDefault(name, value string) string {
	v := os.Getenv(name)
	if v == "" {
		v = value
	}
	return v
}

func getDurationFromEnv(key, value string) (time.Duration, error) {
	v := getEnvOrDefault(key, value)
	return time.ParseDuration(v)
}
