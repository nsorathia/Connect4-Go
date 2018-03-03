package config

import (
	"fmt"
	"github.com/spf13/viper"
)


func init() {
	loadConfigs()
}

func loadConfigs() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/go/src/games/connectfour/")
	viper.SetConfigName("applicationConfig")
	err := viper.ReadInConfig()
	if err != nil { 
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	//viper.SetDefault("game", "connectfour")
	//viper.SetDefault("game-type", "oneplayer")
	//viper.SetDefault("difficulty-level", 3)
	//viper.SetDefault("algorithm", "minmax")
}

//GetString returns a string given the Config propertyNmae
func GetString(propertyName string) string {
	return viper.GetString(propertyName)
}

//GetInt returns a int given the Config propertyNmae
func GetInt(propertyName string) int {
	return viper.GetInt(propertyName)
}

//GetBool returns a bool given the Config propertyNmae
func GetBool(propertyName string) bool {
	return viper.GetBool(propertyName)
}
