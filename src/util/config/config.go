package config

import (
	"flag"
	"os"

	"github.com/Dataman-Cloud/rolex/src/util"

	dmConfig "github.com/Dataman-Cloud/go-component/utils/config"
	log "github.com/Sirupsen/logrus"
)

const (
	defaultDockerEntryPort = "2375"
)

type Config struct {
	RolexAddr         string   `env:"ROLEX_ADDR", required:"true"`
	SwarmManagerIP    string   `env:"ROLEX_SWARM_MANAGER_IP", required:"true"`
	DockerEntryScheme string   `required:"false"`
	DockerEntryPort   string   `env:"ROLEX_DOCKER_ENTRY_PORT", required:"false"`
	DockerTlsVerify   bool     `env:"ROLEX_DOCKER_TLS_VERIFY", required:"true"`
	DockerCertPath    string   `env:"ROLEX_DOCKER_CERT_PATH", required:"true"`
	DbDriver          string   `env:"ROLEX_DB_DRIVER", required:"true"`
	DbDSN             string   `env:"ROLEX_DB_DSN", required:"true"`
	FeatureFlags      []string `env:"ROLEX_FEATURE_FLAGS", required:"true"`

	// registry
	RegistryPrivateKeyPath string `env:"ROLEX_REGISTRY_PRIVATE_KEY_PATH", required:"true"`
	RegistryAddr           string `env:"ROLEX_REGISTRY_ADDR", required:"true"`

	// account
	AccountAuthenticator string `env:"ROLEX_ACCOUNT_AUTHENTICATOR", required:"false"`
	AccountTokenStore    string `env:"ROLEX_ACCOUNT_TOKEN_STORE", required:"false"`

	AccountEmailDefault    string `env:"ROLEX_ACCOUNT_EMAIL_DEFAULT", required:"false"`
	AccountPasswordDefault string `env:"ROLEX_ACCOUNT_PASSWORD_DEFAULT", required:"false"`

	CatalogPath string `env:"ROLEX_CATALOG_PATH", required:"false"`

	SearchLoadDataInterval uint16 `env:"ROLEX_SEARCH_LOAD_DATA_INTERVAL", required:"false"`
}

func (c *Config) FeatureEnabled(feature string) bool {
	return util.StringInSlice(feature, c.FeatureFlags)
}

var config Config

func GetConfig() *Config {
	return &config
}

func init() {
	envFile := flag.String("config", "env_file", "")
	dmConfig.LoadEnvFile(*envFile)

	if err := dmConfig.LoadConfig(&config); err != nil {
		log.Error("LoadConfig got error: ", err)
		os.Exit(1)
	}
	if config.DockerEntryPort == "" {
		config.DockerEntryPort = defaultDockerEntryPort
	}
	if config.DockerTlsVerify {
		config.DockerEntryScheme = "https"
	} else {
		config.DockerEntryScheme = "http"
	}

	return
}
