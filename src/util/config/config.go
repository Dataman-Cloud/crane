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

	// To be removed, temp
	NodePort string

	// account
	AccountAuthenticator string
	AccountTokenStore    string

	AccountEmailDefault    string
	AccountPasswordDefault string

	CatalogPath string

	LoadDataInterval uint16
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

	// To be removed
	ROLEX_NODE_PORT string `required:"true"`

	ROLEX_ACCOUNT_TOKEN_STORE   string `required:"false"`
	ROLEX_ACCOUNT_AUTHENTICATOR string `required:"false"`

	ROLEX_ACCOUNT_EMAIL_DEFAULT    string `required:"false"`
	ROLEX_ACCOUNT_PASSWORD_DEFAULT string `required:"false"`

	ROLEX_CATALOG_PATH string `required:"false"`

	ROLEX_LOADDATA_INTERVAL uint16
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

	config.RegistryPrivateKeyPath = envEntry.ROLEX_REGISTRY_PRIVATE_KEY_PATH
	config.RegistryAddr = envEntry.ROLEX_REGISTRY_ADDR

	config.NodePort = envEntry.ROLEX_NODE_PORT

	config.AccountAuthenticator = envEntry.ROLEX_ACCOUNT_AUTHENTICATOR
	config.AccountTokenStore = envEntry.ROLEX_ACCOUNT_TOKEN_STORE

	config.AccountEmailDefault = envEntry.ROLEX_ACCOUNT_EMAIL_DEFAULT
	config.AccountPasswordDefault = envEntry.ROLEX_ACCOUNT_PASSWORD_DEFAULT

	config.CatalogPath = envEntry.ROLEX_CATALOG_PATH

	config.LoadDataInterval = envEntry.ROLEX_LOADDATA_INTERVAL

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
		log.Errorf("Read node addr info from %s got error: %s", envfile, err.Error())
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

	//TODO get node addr form swam API
	ReadNodeAddrFromFile(NodeAddrInfoFilePath)
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
	log.Errorf("Check env %s got error: %s", env, err.Error())
}

//////////////////////////////////////////////////////////////////////////////////
// Temporary solution: get map of node ip to node id(name) from file. Remove in future
var NodeAddrMap map[string]string
var NodeAddrInfoFilePath = "node_addr_file"

func ReadNodeAddrFromFile(filePath string) {
	log.Debug("Node addr info file path: ", filePath)
	NodeAddrMap = make(map[string]string)
	f, err := os.Open(filePath)
	if err != nil {
		log.Errorf("Read node addr info from %s got error: %s", filePath, err.Error())
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

		NodeAddrMap[key] = val
	}
}

//////////////////////////////////////////////////////////////////////////////////
