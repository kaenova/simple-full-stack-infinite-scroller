package config

type Config struct {
	ServicePort string
	Database    DatabaseConfig
	Secret      string
}

type DatabaseConfig struct {
	Host     string `env:"DATABASE_HOST,default=localhost"`
	Port     string `env:"DATABASE_PORT,default=5432"`
	Username string `env:"DATABASE_USERNAME,required"`
	Password string `env:"DATABASE_PASSWORD,required"`
	Name     string `env:"DATABASE_NAME,required"`
}

func GetConfig() Config {
	return Config{
		Database: DatabaseConfig{
			Host:     "127.0.0.1",
			Port:     "5432",
			Username: "postgres",
			Password: "kaenova",
			Name:     "postgres",
		},
	}
}
