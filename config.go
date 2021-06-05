package main

import "os"

//Config represents a struct that contains app configuration data
type Config struct {
	ListenPort string
	CryptoURL  string
}

//ParseEnv this is a function that get local env and return an error
func (conf *Config) ParseEnv() error {

	InitStringEnv("CryptoURL", &conf.CryptoURL, "https://api.hitbtc.com/api/2/public", false)
	InitStringEnv("ListenPort", &conf.ListenPort, "10000", false)

	return nil
}

//Initialize is a function to initialize config with environmental variables.
func (conf *Config) Initialize() error {
	err := conf.ParseEnv()
	if err != nil {
		return err
	}
	return nil
}

func InitStringEnv(envName string, loc *string, defaultValOrErr string, exitOnErr bool) {
	if os.Getenv(envName) != "" {
		*loc = os.Getenv(envName)
	} else {
		if exitOnErr {
			panic(defaultValOrErr)
		}
		*loc = defaultValOrErr
	}
}
