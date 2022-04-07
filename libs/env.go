package libs

import (
	"golangapi/app/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("MONGO_URI")
}

func Environment() (c models.EnvConfig, err error) {
	viper.SetConfigName("env") // name of config file (without extension)
	viper.AddConfigPath(".")   // optionally look for config in the working directory
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Config file not found; ignore error if desired")
		} else {
			log.Println("Config file was found but another error was produced")
		}
	}

	viper.AutomaticEnv()

	// Config file found and successfully parsed
	err = viper.Unmarshal(&c)
	if err != nil {
		return
	}
	return
}
