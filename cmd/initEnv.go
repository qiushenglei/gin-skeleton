package cmd

import (
	"fmt"
	"github.com/qiushenglei/gin-skeleton/internal/app/configs"
	"github.com/spf13/viper"
)

func parseEnv() error {
	viper.SetConfigName(configs.EnvFile)  // name of config file (without extension)
	viper.SetConfigType("env")            // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(configs.BasePath) // path to look for the config file in
	//viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	//viper.AddConfigPath(".")              // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	configs.EnvConfig = viper.GetViper()
	return nil
}
