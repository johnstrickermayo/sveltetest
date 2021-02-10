package env

import (
	"log"
	"os"
)

func envRead(value string) string {
	return os.Getenv(value)
}

// Go Get GO_ENV envrionment string
func Go() string {
	return envRead("GO_ENV")
}

/*Port get PORT environment string
If not set and GO_ENV is 'production' log fatal error */
func Port() string {
	port := envRead("PORT")
	if port == "" {
		if Go() == "Production" {
			log.Fatal("$PORT must be set")
		} else {
			port = "3000"
		}
	}
	return port
}
