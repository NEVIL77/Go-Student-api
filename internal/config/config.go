package config

type HTTPServer struct {
	addr string
}

type Config struct {
	Env          string `yaml:"env"  env-required:"true" env-default:"production"`
	Storage_path string `yaml:"storage_path"`
	HTTPServer   `yaml:"http_server"`
}
