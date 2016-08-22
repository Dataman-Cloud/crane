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

func LoadEnvFile(envfile string) {
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

		key, val, err := Parseln(string(line))
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
func Parseln(line string) (string, string, error) {
	splits := strings.SplitN(removeComments(line), "=", 2)

	if len(splits) < 2 {
		return "", "", errors.New("missing delimiter '='")
	}

	key := strings.Trim(splits[0], " ")
	val := strings.Trim(splits[1], ` "'`)
	return key, val, nil

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

func LoadConfig(configEntry interface{}) error {
	val := reflect.ValueOf(configEntry).Elem()

	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		required := typeField.Tag.Get("required")
		envKey := typeField.Tag.Get("env")

		env := os.Getenv(envKey)

		if env == "" && required == "true" {
			exitMissingEnv(envKey)
		}

		var configEntryValue interface{}
		var err error
		valueFiled := val.Field(i).Interface()
		value := val.Field(i)
		switch valueFiled.(type) {
		case int64:
			configEntryValue, err = strconv.ParseInt(env, 10, 64)
		case int16:
			configEntryValue, err = strconv.ParseInt(env, 10, 16)
			_, ok := configEntryValue.(int64)
			if !ok {
				exitCheckEnv(typeField.Name, err)
			}
			configEntryValue = int16(configEntryValue.(int64))
		case uint16:
			configEntryValue, err = strconv.ParseUint(env, 10, 16)

			_, ok := configEntryValue.(uint64)
			if !ok {
				exitCheckEnv(typeField.Name, err)
			}
			configEntryValue = uint16(configEntryValue.(uint64))
		case uint64:
			configEntryValue, err = strconv.ParseUint(env, 10, 64)
		case bool:
			configEntryValue, err = strconv.ParseBool(env)
		case []string:
			configEntryValue = strings.SplitN(env, ",", -1)
		default:
			configEntryValue = env
		}

		if err != nil {
			exitCheckEnv(typeField.Name, err)
		}
		value.Set(reflect.ValueOf(configEntryValue))
	}

	return nil
}
