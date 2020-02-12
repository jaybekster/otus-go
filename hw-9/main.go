package main

import (
	"log"

	"github.com/jaybekster/otus-go/hw-9/config"
	pflag "github.com/spf13/flag"
	"github.com/spf13/viper"
)

func main() {
	var configPath *string = pflag.String("config", "./config.yml", "path to config file")

	pflag.Parse()

	viper.SetConfigFile(*configPath)

	var configuration config.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	log.Println(configuration)
}
