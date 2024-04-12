package config

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/labstack/gommon/log"
	utils "github.com/tashanemclean/genai-rest-api/util"

	"github.com/spf13/viper"
)

const LOCAL = "local"

type config struct {
	Env 	string `mapstructure:"ENVIRONMENT"`
	AppPort string `mapstructure:"PORT"`
	ApiBaseUrl string `mapstructure:"API_BASE_URL"`
}

var Config config

// Load os layer config
func setEnv() {
	viper.AutomaticEnv()

	// Need to bind keys manual using BindEnv
	// https://github.com/spf13/viper/issues/761
	keys := utils.GetStructTagVals("mapstucture", Config)
	for _, item := range keys {
		viper.BindEnv(item)
	}

	// Unmarshal keys into config map 
	err := viper.Unmarshal(&Config)

	if err != nil {
		log.Fatalf("Cannot decode into map struct, %v", err)
	}
}

func Load() {
	switch os.Getenv("ENVIRONMENT") {
	case "production":
		setEnv()
	case "stagging":
		setEnv()
	default:
		loadLocalConfig()
	}
}

// Load local config file
func loadLocalConfig() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	configPath := filepath.Join(basepath, "..", "..", "configs")

	viper.AddConfigPath(configPath)
	viper.SetConfigType("json")
	viper.SetConfigName("dev")

	// Read in config from file
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Could not load dev.json", err)
	}

	err = viper.Unmarshal(&Config)

	if err != nil {
		log.Fatalf("unable to decode into map struct, %v", err)
	}
}