package config

import (
	"bufio"
	"errors"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/Dataman-Cloud/rolex/src/util"
	log "github.com/Sirupsen/logrus"
)

const (
	defaultDockerEntryPort = "2375"
)

type Config struct {
	RolexAddr         string
	SwarmManagerIP    string
	DockerEntryScheme string
	DockerEntryPort   string
	DockerTlsVerify   bool
	DockerCertPath    string
	DbDriver          string
	DbDSN             string
	FeatureFlags      []string

	// registry
	RegistryPrivateKeyPath string
	RegistryAddr           string

	// account
	AccountAuthenticator string
	AccountTokenStore    string

	AccountEmailDefault    string
	AccountPasswordDefault string

	CatalogPath string

	SearchLoadDataInterval uint16
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
	ROLEX_DOCKER_TLS_VERIFY bool   `required:"true"`
	ROLEX_SWARM_MANAGER_IP  string `required:"true"`
	ROLEX_DOCKER_ENTRY_PORT string `required:"false"`
	ROLEX_DOCKER_CERT_PATH  string `required:"true"`

	ROLEX_DB_DRIVER                 string `required:"true"`
	ROLEX_DB_DSN                    string `required:"true"`
	ROLEX_FEATURE_FLAGS             string `required:"false"`
	ROLEX_REGISTRY_PRIVATE_KEY_PATH string `required:"false"`
	ROLEX_REGISTRY_ADDR             string `required:"false"`

	ROLEX_ACCOUNT_TOKEN_STORE   string `required:"false"`
	ROLEX_ACCOUNT_AUTHENTICATOR string `required:"false"`

	ROLEX_ACCOUNT_EMAIL_DEFAULT    string `required:"false"`
	ROLEX_ACCOUNT_PASSWORD_DEFAULT string `required:"false"`

	ROLEX_CATALOG_PATH string `required:"false"`

	ROLEX_SEARCH_LOAD_DATA_INTERVAL uint16
}

func InitConfig(envFile string) *Config {
	loadEnvFile(envFile)

	envEntry := NewEnvEntry()
	config.RolexAddr = envEntry.ROLEX_ADDR
	config.SwarmManagerIP = envEntry.ROLEX_SWARM_MANAGER_IP
	config.DockerEntryPort = envEntry.ROLEX_DOCKER_ENTRY_PORT
	if config.DockerEntryPort == "" {
		config.DockerEntryPort = defaultDockerEntryPort
	}
	config.DockerTlsVerify = envEntry.ROLEX_DOCKER_TLS_VERIFY
	if config.DockerTlsVerify {
		config.DockerEntryScheme = "https"
	} else {
		config.DockerEntryScheme = "http"
	}
	config.DockerCertPath = envEntry.ROLEX_DOCKER_CERT_PATH

	config.DbDriver = envEntry.ROLEX_DB_DRIVER
	config.DbDSN = envEntry.ROLEX_DB_DSN
	config.FeatureFlags = strings.SplitN(envEntry.ROLEX_FEATURE_FLAGS, ",", -1)

	config.RegistryPrivateKeyPath = envEntry.ROLEX_REGISTRY_PRIVATE_KEY_PATH
	config.RegistryAddr = envEntry.ROLEX_REGISTRY_ADDR

	config.AccountAuthenticator = envEntry.ROLEX_ACCOUNT_AUTHENTICATOR
	config.AccountTokenStore = envEntry.ROLEX_ACCOUNT_TOKEN_STORE

	config.AccountEmailDefault = envEntry.ROLEX_ACCOUNT_EMAIL_DEFAULT
	config.AccountPasswordDefault = envEntry.ROLEX_ACCOUNT_PASSWORD_DEFAULT

	config.CatalogPath = envEntry.ROLEX_CATALOG_PATH

	config.SearchLoadDataInterval = envEntry.ROLEX_SEARCH_LOAD_DATA_INTERVAL

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
	if err != nil {
		log.Infof("Failed to open config file %s: %s", envfile, err.Error())
		return
	}
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
			log.Errorf("Parse info %s got error: %s", line, err.Error())
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
func removeComments(s string) string {
	if len(s) == 0 || string(s[0]) == "#" {
		return ""
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
	log.Errorf("Check env %s got error: %s", env, err.Error())
}
