package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Host     string
	Username string
	Password string
	Database string
	Port     int
	Timezone string
	SSLmode  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Host:     "host.docker.internal",
			Username: "postgres",
			Password: "developer",
			Database: "cars",
			Port:     5432,
			Timezone: "Europe/Lisbon",
			SSLmode:  "disable",
		},
	}
}
