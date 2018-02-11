package config

import (
	"fmt"
	"github.com/spf13/viper"
)


func init() {
	LaodConfigs()
}

func LaodConfigs() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/go/src/games/connectfour/")
	viper.SetConfigName("applicationConfig")
	err := viper.ReadInConfig()
	if err != nil { 
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}


	//viper.SetDefault("game", "connectfour")
	//viper.SetDefault("game-type", "oneplayer")
	//viper.SetDefault("difficulty-level", 3)
	//viper.SetDefault("algorithm", "minmax")
}

func GetString(propertyName string) string {
	val := viper.GetString(propertyName)
	return val
}

func GetInt(propertyName string) int {
	val := viper.GetInt(propertyName)
	return val
}

func GetBool(propertyName string) bool {
	val := viper.GetBool(propertyName)
	return val
}