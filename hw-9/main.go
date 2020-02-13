package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

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

	http.HandleFunc("/hello", helloController)

	err = http.ListenAndServe(
		strings.Join([]string{configuration.Http_listen.Ip, strconv.Itoa(configuration.Http_listen.Port)}, ":"),
		nil,
	)

	log.Printf("Server start lisening on port: %d", configuration.Http_listen.Port)

	if err != nil {
		log.Fatalf("server can not start, %v", err)
	}
}

func helloController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}
