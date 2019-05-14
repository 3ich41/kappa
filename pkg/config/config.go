package config

import (
	"os"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	MqUsername string
	MqPassword string
	MqHostname string
	MqPort     string
	LogLevel   string
}

// mustGetEnv returns an env variable value if present and fails othwewise
func mustGetEnv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("%s environment variable not set.", k)
	}
	return v
}

// InitConfig populates config variable and supposed to be called when application started
func NewConfig() *Config {
	m := mustGetEnv
	c := &Config{
		MqUsername: m("MQ_USERNAME"),
		MqPassword: m("MQ_PASSWORD"),
		MqHostname: m("MQ_HOSTNAME"),
		MqPort:     m("MQ_PORT"),
		LogLevel:   m("LOGLEVEL"),
	}

	return c
}
