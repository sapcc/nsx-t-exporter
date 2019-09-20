package config

import (
	"os"
	"strconv"
)

// NSXv3Configuration holds the configuration of the NSXv3 Manager
type NSXv3Configuration struct {
	LoginHost           string
	LoginPort           string
	LoginUser           string
	LoginPassword       string
	RequestsPerSecond   int
	RequestTimeout      int
	SuppressSslWornings bool
	ScrapPort           int
}

// Init Loads NSXv3Configuration value from the OS environment variables
func Init() NSXv3Configuration {

	// TODO - Add validate method
	requestsPerSecond, _ := strconv.Atoi(os.Getenv("NSXV3_REQUESTS_PER_SECOND"))
	requestTimeoutSeconds, _ := strconv.Atoi(os.Getenv("NSXV3_REQUEST_TIMEOUT_SECONDS"))
	suppressSslWornings, _ := strconv.ParseBool(os.Getenv("NSXV3_SUPPRESS_SSL_WORNINGS"))
	scrapPort, _ := strconv.Atoi(os.Getenv("SCRAP_PORT"))

	nsxv3Config := NSXv3Configuration{
		LoginHost:           os.Getenv("NSXV3_LOGIN_HOST"),
		LoginPort:           os.Getenv("NSXV3_LOGIN_PORT"),
		LoginUser:           os.Getenv("NSXV3_LOGIN_USER"),
		LoginPassword:       os.Getenv("NSXV3_LOGIN_PASSWORD"),
		RequestsPerSecond:   int(requestsPerSecond),
		RequestTimeout:      int(requestTimeoutSeconds),
		SuppressSslWornings: suppressSslWornings,
		ScrapPort:           int(scrapPort),
	}

	return nsxv3Config

}
