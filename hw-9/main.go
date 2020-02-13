package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	logrusmiddleware "github.com/bakins/logrus-middleware"
	"github.com/jaybekster/otus-go/hw-9/config"
	"github.com/sirupsen/logrus"
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

	logger := logrus.New()

	switch configuration.Log_level {
	case "error":
		logger.Level = logrus.ErrorLevel
	case "warn":
		logger.Level = logrus.WarnLevel
	case "info":
		logger.Level = logrus.InfoLevel
	case "debug":
		logger.Level = logrus.DebugLevel
	}

	logger.Formatter = &logrus.JSONFormatter{}

	l := logrusmiddleware.Middleware{
		Name:   "example",
		Logger: logger,
	}

	http.Handle("/hello", l.Handler(http.HandlerFunc(helloController), "hello page"))

	done := make(chan bool)

	go func() {
		err = http.ListenAndServe(
			strings.Join([]string{configuration.Http_listen.Ip, strconv.Itoa(configuration.Http_listen.Port)}, ":"),
			nil,
		)

		if err != nil {
			log.Fatalf("server can not start, %v", err)
		}
	}()

	log.Printf("Server start lisening on ip %s and port %d", configuration.Http_listen.Ip, configuration.Http_listen.Port)

	<-done

}

func helloController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}
