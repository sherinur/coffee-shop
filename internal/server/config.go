package server

type Config struct {
	env            string
	port           string
	data_directory string

	read_timeout  string
	write_timeout string
	idle_timout   string

	log_file string
	cfg_file string

	max_file_size   string
	allow_overwrite bool
}

func NewConfig(configPath, port, dir string) *Config {
	if configPath != "configs/server.yaml" {
		// TODO: Parse .yaml config file
		return &Config{}
	}

	return &Config{
		env:            "local",
		port:           port,
		data_directory: dir,
		read_timeout:   "4s",
		write_timeout:  "4s",
		idle_timout:    "60s",

		log_file: "./logs/triple-s.log",
		cfg_file: "./configs/server.yaml",

		max_file_size:   "10MB",
		allow_overwrite: true,
	}
}

func (cfg *Config) GetPort() string {
	return cfg.port
}
