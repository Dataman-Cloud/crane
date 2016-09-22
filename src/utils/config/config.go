package config

import (
	"flag"
	"os"

	"github.com/Dataman-Cloud/crane/src/utils"

	log "github.com/Sirupsen/logrus"
)

const (
	defaultDockerEntryPort = "2375"
)

type Config struct {
	CraneAddr         string   `env:"CRANE_ADDR",required:"true"`
	SwarmManagerIP    string   `env:"CRANE_SWARM_MANAGER_IP",required:"true"`
	DockerEntryScheme string   `required:"false"`
	DockerEntryPort   string   `env:"CRANE_DOCKER_ENTRY_PORT",required:"false"`
	DockerTlsVerify   bool     `env:"CRANE_DOCKER_TLS_VERIFY",required:"true"`
	DockerCertPath    string   `env:"CRANE_DOCKER_CERT_PATH",required:"true"`
	DockerApiVersion  string   `env:"CRANE_DOCKER_API_VERSION",required:"false"`
	DbDriver          string   `env:"CRANE_DB_DRIVER",required:"true"`
	DbDSN             string   `env:"CRANE_DB_DSN",required:"true"`
	FeatureFlags      []string `env:"CRANE_FEATURE_FLAGS",required:"true"`

	// registry
	RegistryPrivateKeyPath string `env:"CRANE_REGISTRY_PRIVATE_KEY_PATH",required:"true"`
	RegistryAddr           string `env:"CRANE_REGISTRY_ADDR",required:"true"`

	// account
	AccountAuthenticator string `env:"CRANE_ACCOUNT_AUTHENTICATOR",required:"false"`
	AccountTokenStore    string `env:"CRANE_ACCOUNT_TOKEN_STORE",required:"false"`

	AccountEmailDefault    string `env:"CRANE_ACCOUNT_EMAIL_DEFAULT",required:"false"`
	AccountPasswordDefault string `env:"CRANE_ACCOUNT_PASSWORD_DEFAULT",required:"false"`

	CatalogPath string `env:"CRANE_CATALOG_PATH",required:"false"`

	SearchLoadDataInterval uint16 `env:"CRANE_SEARCH_LOAD_DATA_INTERVAL",required:"false"`
}

func (c *Config) FeatureEnabled(feature string) bool {
	return utils.StringInSlice(feature, c.FeatureFlags)
}

var config *Config

func GetConfig() *Config {
	if config == nil {
		Init()
	}
	return config
}

func Init() {
	envFile := flag.String("config", "env_file", "")
	LoadEnvFile(*envFile)

	config = new(Config)
	if err := LoadConfig(config); err != nil {
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
