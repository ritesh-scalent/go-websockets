package config

type ServiceConfig struct {
	HttpPort    string `yaml:"http_port"`
	ServiceName string `yaml:"service_name"`
}
