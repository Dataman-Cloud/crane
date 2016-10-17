package config

import (
	"github.com/Dataman-Cloud/crane/src/utils"
	log "github.com/Sirupsen/logrus"
)

type Config struct {
	CraneAddr         string   `env:"CRANE_ADDR,required"`
	SwarmManagerIP    string   `env:"CRANE_SWARM_MANAGER_IP,required"`
	DockerEntryScheme string   `env:"CRANE_DOCKER_ENTRY_SCHEME" envDefault:"http"`
	DockerEntryPort   string   `env:"CRANE_DOCKER_ENTRY_PORT" envDefault:"2375"`
	DockerCertPath    string   `env:"CRANE_DOCKER_CERT_PATH,required"`
	DockerApiVersion  string   `env:"CRANE_DOCKER_API_VERSION"`
	DbDriver          string   `env:"CRANE_DB_DRIVER,required"`
	DbDSN             string   `env:"CRANE_DB_DSN,required"`
	FeatureFlags      []string `env:"CRANE_FEATURE_FLAGS,required"`
	DockerTlsVerify   bool     `env:"CRANE_DOCKER_TLS_VERIFY envDefault:"false"`

	// registry
	RegistryPrivateKeyPath string `env:"CRANE_REGISTRY_PRIVATE_KEY_PATH,required"`
	RegistryAddr           string `env:"CRANE_REGISTRY_ADDR,required"`

	// account
	AccountAuthenticator   string `env:"CRANE_ACCOUNT_AUTHENTICATOR,required"`
	AccountTokenStore      string `env:"CRANE_ACCOUNT_TOKEN_STORE"`
	AccountEmailDefault    string `env:"CRANE_ACCOUNT_EMAIL_DEFAULT"`
	AccountPasswordDefault string `env:"CRANE_ACCOUNT_PASSWORD_DEFAULT"`

	CatalogPath            string `env:"CRANE_CATALOG_PATH"`
	SearchLoadDataInterval int    `env:"CRANE_SEARCH_LOAD_DATA_INTERVAL"`
}

var config *Config

func (c *Config) FeatureEnabled(feature string) bool {
	return utils.StringInSlice(feature, c.FeatureFlags)
}

func InitConfig() *Config {
	cfg := Config{}
	if err := Parse(&cfg); err != nil {
		log.Fatalf("Parse Env into config got error: ", err)
	}

	return &cfg
}
