package models

type (
	EnvConfig struct {
		Host     string         `mapstructure:"host"`
		Port     int            `mapstructure:"port"`
		Database DatabaseConfig `mapstructure:"database"`
	}
	DatabaseConfig struct {
		Timeout int      `mapstructure:"timeout"`
		DBname  string   `mapstructure:"mongo_db_name"`
		URI     []string `mapstructure:"mongo_uri"`
	}
)
