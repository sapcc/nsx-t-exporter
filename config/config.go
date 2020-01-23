package config

import (
	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"
)

// NSXv3Configuration holds the configuration of the NSXv3 Manager
type NSXv3Configuration struct {
	//LoginPort            int    `env:"NSXV3_LOGIN_PORT" envDefault:"443"`
	LoginHost            string `env:"NSXV3_LOGIN_HOST,required"`
	LoginUser            string `env:"NSXV3_LOGIN_USER,required"`
	LoginPassword        string `env:"NSXV3_LOGIN_PASSWORD,required"`
	ScrapPort            int    `env:"SCRAP_PORT"  envDefault:"9999"`
	RequestsPerSecond    int    `env:"NSXV3_REQUESTS_PER_SECOND" envDefault:"10"`
	RequestsConnPoolSize int    `env:"NSXV3_CONNECTION_POOL_SIZE" envDefault:"0"`
	RequestTimeout       int    `env:"NSXV3_REQUEST_TIMEOUT_SECONDS" envDefault:"0"`
	SuppressSslWarnings  bool   `env:"NSXV3_SUPPRESS_SSL_WARNINGS" envDefault:"false"`
	ScrapScheduleSeconds int    `env:"SCRAP_SCHEDULE_SECONDS" envDefault:"0"`
}



// Init Loads NSXv3Configuration value from the OS environment variables and validate them
func Init() NSXv3Configuration {

	nsxv3Config := NSXv3Configuration{}

	if err := env.Parse(&nsxv3Config); err != nil {
		log.Fatalf("%+v", err)
	}

	// Variables that cannot be empty strings
	variablesToCheck := map [string]string {
		"NSXV3_LOGIN_HOST": nsxv3Config.LoginHost,
	}

	ensureNotEmpty(variablesToCheck)
	return nsxv3Config
}

func ensureNotEmpty(variablesToCheck map[string]string) {
	for key, value := range variablesToCheck {
		if value == "" {
			log.Fatalf("Variable %v cannot be an empty string.", key)
		}
	}
}
