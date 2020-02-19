package config

type Configuration struct {
	HttpListen struct {
		Ip   string
		Port int
	}
	LogFile  string
	LogLevel string
}
