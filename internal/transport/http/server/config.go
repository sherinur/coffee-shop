package server

type Config struct {
	Env            string
	port           string
	data_directory string

	menu_file      string
	inventory_file string
	order_file     string
	report_file    string

	read_timeout  string
	write_timeout string
	idle_timout   string

	Log_file string
	cfg_file string

	allow_overwrite bool
}

func NewConfig(configPath, port, dir string) *Config {
	if configPath != "configs/server.yaml" {
		// TODO: Parse .yaml config file
		return &Config{}
	}

	return &Config{
		Env:            "local",
		port:           port,
		data_directory: dir,

		menu_file:      dir + "/menu_items.json",
		inventory_file: dir + "/inventory.json",
		order_file:     dir + "/orders.json",
		report_file:    dir + "/report.json",

		read_timeout:  "4s",
		write_timeout: "4s",
		idle_timout:   "60s",

		Log_file: "./logs/all.log",
		cfg_file: "./configs/server.yaml",

		allow_overwrite: true,
	}
}

func (cfg *Config) GetPort() string {
	return cfg.port
}
