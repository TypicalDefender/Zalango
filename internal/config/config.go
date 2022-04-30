package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func Init() {
	if os.Getenv("ENVIRONMENT") == "test" {
		viper.SetConfigName("test")
	} else {
		viper.SetConfigName("application")
	}

	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("./../../configs")
	viper.AddConfigPath("./../../../configs")
	viper.AddConfigPath("./../../../../configs")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("could not read in config %s", err.Error()))
	}
	viper.AutomaticEnv()

	initAppConfig()
	initRESTServerConfig()
	initGRPCServerConfig()
	initFlagrConfig()
}
