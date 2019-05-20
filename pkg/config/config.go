package config

import (
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	LdapBase               string
	LdapBindDN             string
	LdapBindPassword       string
	LdapHost               string
	LdapServerName         string
	LdapUserFilter         string // e.g. "(uid=%s)"
	LdapPort               int
	LdapInsecureSkipVerify bool
	LdapUseSSL             bool
	LdapSkipTLS            bool
	GrpcPort               int
	LogLevel               string
}

// mustGetEnv returns an env variable value if present and fails othwewise
func getStrEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Errorf("%s environment variable not set.", key)
		os.Exit(1)
	}
	return val
}

func getIntEnv(key string) int {
	val := getStrEnv(key)
	ret, err := strconv.Atoi(val)
	if err != nil {
		log.WithField("err", err.Error()).Errorf("Error while reading %v environment variable", key)
		os.Exit(1)
	}
	return ret
}

func getBoolEnv(key string) bool {
	val := getStrEnv(key)
	ret, err := strconv.ParseBool(val)
	if err != nil {
		log.WithField("err", err.Error()).Errorf("Error while reading %v environment variable", key)
		os.Exit(1)
	}
	return ret
}

// InitConfig populates config variable and supposed to be called when application started
func NewConfig() *Config {
	s := getStrEnv
	i := getIntEnv
	b := getBoolEnv

	c := &Config{
		LdapBase:               s("LDAP_BASE"),
		LdapBindDN:             s("LDAP_BINDDN"),
		LdapBindPassword:       s("LDAP_BINDPASSWORD"),
		LdapHost:               s("LDAP_HOST"),
		LdapServerName:         s("LDAP_SERVERNAME"),
		LdapUserFilter:         s("LDAP_USERFILTER"), // e.g. "(uid=%s)"
		LdapPort:               i("LDAP_PORT"),
		LdapInsecureSkipVerify: b("LDAP_INSECURESKIPVERIFY"),
		LdapUseSSL:             b("LDAP_USESSL"),
		LdapSkipTLS:            b("LDAP_SKIPTLS"),
		GrpcPort:               i("GRPC_PORT"),
		LogLevel:               s("LOGLEVEL"),
	}

	return c
}
