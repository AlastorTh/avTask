package store

type Config struct {
	ConnectString string `yaml:"conn_string"`
}

func NewConfig() *Config {
	return &Config{ConnectString: "host=localhost user = postgres password = pass port=5432 dbname=postgres sslmode=disable"}
}
