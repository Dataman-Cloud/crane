package config

import (
	"bufio"
	"errors"
	"os"
	"reflect"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
)

type Config struct {
	RolexAddr       string
	DockerHost      string
	DockerTlsVerify string
	DockerCertPath  string
	DbDriver        string
	DbDSN           string

	HOST string
	PORT uint64
}

var config Config

func GetConfig() *Config {
	return &config
}

type EnvEntry struct {
	ROLEX_ADDR        string `required:"true"`
	DOCKER_TLS_VERIFY string `required:"true"`
	DOCKER_HOST       string `required:"true"`
	DOCKER_CERT_PATH  string `required:"true"`

	ROLEX_DB_DRIVER string `required:"true"`
	ROLEX_DB_DSN    string `required:"true"`
}

func InitConfig(envFile string) *Config {
	loadEnvFile(envFile)

	envEntry := NewEnvEntry()
	config.RolexAddr = envEntry.ROLEX_ADDR
	config.DockerHost = envEntry.DOCKER_HOST
	config.DockerTlsVerify = envEntry.DOCKER_TLS_VERIFY
	config.DockerCertPath = envEntry.DOCKER_CERT_PATH

	config.DbDriver = envEntry.ROLEX_DB_DRIVER
	config.DbDSN = envEntry.ROLEX_DB_DSN
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
