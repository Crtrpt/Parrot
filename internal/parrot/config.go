package parrot

import (
	"fmt"

	"github.com/spf13/viper"
)

func ConfigInit() {
	viper.SetConfigName("parrot")        // name of config file (without extension)
	viper.SetConfigType("yaml")          // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/parrot/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.parrot") // call multiple times to add many search paths
	viper.AddConfigPath(".")             // optionally look for config in the working directory
	viper.AddConfigPath("./configs")     // optionally look for config in the working directory
	err := viper.ReadInConfig()          // Find and read the config file
	if err != nil {                      // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

}
