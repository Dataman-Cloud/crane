package config

import (
	"bufio"
	"errors"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/Dataman-Cloud/rolex/util"
	log "github.com/Sirupsen/logrus"
)

type Config struct {
	RolexAddr       string
	DockerHost      string
	DockerTlsVerify string
	DockerCertPath  string
	DbDriver        string
	DbDSN           string
	FeatureFlags    []string

	HOST string
	PORT uint64

	// registry
	RegistryPrivateKeyPath string
	RegistryAddr           string

	// swarm cluster
	RolexSecret string
	RolexCaHash string

	// To be removed, temp
	NodeIP   string
	NodePort string

	// account
	AccountAuthenticator string
	AccountTokenStore    string

	AccountEmailDefault    string
	AccountPasswordDefault string
}

func (c *Config) FeatureEnabled(feature string) bool {
	return util.StringInSlice(feature, c.FeatureFlags)
}

var config Config

func GetConfig() *Config {
	return &config
}

type EnvEntry struct {
	ROLEX_ADDR              string `required:"true"`
	ROLEX_DOCKER_TLS_VERIFY string `required:"true"`
	ROLEX_DOCKER_HOST       string `required:"true"`
	ROLEX_DOCKER_CERT_PATH  string `required:"true"`

	ROLEX_DB_DRIVER                 string `required:"true"`
	ROLEX_DB_DSN                    string `required:"true"`
	ROLEX_FEATURE_FLAGS             string `required:"false"`
	ROLEX_REGISTRY_PRIVATE_KEY_PATH string `required:"false"`
	ROLEX_REGISTRY_ADDR             string `required:"false"`

	ROLEX_SECRET  string `required:"true"`
	ROLEX_CA_HASH string `required:"true"`
	// To be removed
	ROLEX_NODE_IP   string `required:"true"`
	ROLEX_NODE_PORT string `required:"true"`

	ROLEX_ACCOUNT_TOKEN_STORE   string `required:"false"`
	ROLEX_ACCOUNT_AUTHENTICATOR string `required:"false"`

	ROLEX_ACCOUNT_EMAIL_DEFAULT    string `required:"false"`
	ROLEX_ACCOUNT_PASSWORD_DEFAULT string `required:"false"`
}

func InitConfig(envFile string) *Config {
	loadEnvFile(envFile)

	envEntry := NewEnvEntry()
	config.RolexAddr = envEntry.ROLEX_ADDR
	config.DockerHost = envEntry.ROLEX_DOCKER_HOST
	config.DockerTlsVerify = envEntry.ROLEX_DOCKER_TLS_VERIFY
	config.DockerCertPath = envEntry.ROLEX_DOCKER_CERT_PATH

	config.DbDriver = envEntry.ROLEX_DB_DRIVER
	config.DbDSN = envEntry.ROLEX_DB_DSN
	config.FeatureFlags = strings.SplitN(envEntry.ROLEX_FEATURE_FLAGS, ",", -1)

	config.RolexSecret = envEntry.ROLEX_SECRET
	config.RolexCaHash = envEntry.ROLEX_CA_HASH

	config.RegistryPrivateKeyPath = envEntry.ROLEX_REGISTRY_PRIVATE_KEY_PATH
	config.RegistryAddr = envEntry.ROLEX_REGISTRY_ADDR

	config.NodeIP = envEntry.ROLEX_NODE_IP
	config.NodePort = envEntry.ROLEX_NODE_PORT

	config.AccountAuthenticator = envEntry.ROLEX_ACCOUNT_AUTHENTICATOR
	config.AccountTokenStore = envEntry.ROLEX_ACCOUNT_TOKEN_STORE

	config.AccountEmailDefault = envEntry.ROLEX_ACCOUNT_EMAIL_DEFAULT
	config.AccountPasswordDefault = envEntry.ROLEX_ACCOUNT_PASSWORD_DEFAULT

	return &config
}

func NewEnvEntry() *EnvEntry {
	envEntry := &EnvEntry{}

	val := reflect.ValueOf(envEntry).Elem()

	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		required := typeField.Tag.Get("required")

		env := os.Getenv(typeField.Name)

		if env == "" && required == "true" {
			exitMissingEnv(typeField.Name)
		}

		var envEntryValue interface{}
		var err error
		valueFiled := val.Field(i).Interface()
		value := val.Field(i)
		switch valueFiled.(type) {
		case int64:
			envEntryValue, err = strconv.ParseInt(env, 10, 64)

		case int16:
			envEntryValue, err = strconv.ParseInt(env, 10, 16)
			_, ok := envEntryValue.(int64)
			if !ok {
				exitCheckEnv(typeField.Name, err)
			}
			envEntryValue = int16(envEntryValue.(int64))
		case uint16:
			envEntryValue, err = strconv.ParseUint(env, 10, 16)

			_, ok := envEntryValue.(uint64)
			if !ok {
				exitCheckEnv(typeField.Name, err)
			}
			envEntryValue = uint16(envEntryValue.(uint64))
		case uint64:
			envEntryValue, err = strconv.ParseUint(env, 10, 64)
		case bool:
			envEntryValue, err = strconv.ParseBool(env)
		default:
			envEntryValue = env
		}

		if err != nil {
			exitCheckEnv(typeField.Name, err)
		}
		value.Set(reflect.ValueOf(envEntryValue))
	}

	return envEntry
}

func loadEnvFile(envfile string) {
	// load the environment file
	log.Debug("envfile: ", envfile)
	f, err := os.Open(envfile)
	if err == nil {
		defer f.Close()

		r := bufio.NewReader(f)
		for {
			line, _, err := r.ReadLine()
			if err != nil {
				break
			}

			if len(line) == 0 {
				continue
			}

			key, val, err := parseln(string(line))
			if err != nil {
				continue
			}

			if len(os.Getenv(strings.ToUpper(key))) == 0 {
				err1 := os.Setenv(strings.ToUpper(key), val)
				if err1 != nil {
					log.Error(err1.Error())
				}
			}
		}
	}
}

// helper function to parse a "key=value" environment variable string.
func parseln(line string) (key string, val string, err error) {
	line = removeComments(line)
	if len(line) == 0 {
		return
	}
	splits := strings.SplitN(line, "=", 2)

	if len(splits) < 2 {
		err = errors.New("missing delimiter '='")
		return
	}

	key = strings.Trim(splits[0], " ")
	val = strings.Trim(splits[1], ` "'`)
	return

}

// helper function to trim comments and whitespace from a string.
func removeComments(s string) (_ string) {
	if len(s) == 0 || string(s[0]) == "#" {
		return
	} else {
		index := strings.Index(s, " #")
		if index > -1 {
			s = strings.TrimSpace(s[0:index])
		}
	}
	return s
}

func exitMissingEnv(env string) {
	log.Errorf("program exit missing config for env %s", env)
	os.Exit(1)
}

func exitCheckEnv(env string, err error) {
	log.Errorf("Check env %s, %s", env, err.Error())

}
