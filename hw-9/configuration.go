package config

type Configuration struct {
	Http_listen struct {
		Ip   string
		Port int
	}
	Log_file  string
	Log_level string
}
