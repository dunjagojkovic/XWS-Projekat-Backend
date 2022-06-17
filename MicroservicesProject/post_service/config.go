package main

/*import "os"*/

type Config struct {
	Port       string
	PostDBHost string
	PostDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:  "8000",   /* os.Getenv("POST_SERVICE_PORT"),*/
		PostDBHost:"localhost", /* os.Getenv("POST_DB_HOST"),*/
		PostDBPort:"27017",    /* os.Getenv("POST_DB_PORT"),*/
	}
}
