package main

import (
	"fmt"

	"github.com/spf13/viper"
)

var Level *viper.Viper

var Tiles map[string]interface{}

func LoadMap(levelName string) {

	Level = viper.New()

	Level.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	Level.AddConfigPath("levels")
	Level.SetConfigName(levelName)

	err := Level.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	LevelMap = Level.GetString("level")
	fmt.Println(levelName + " is loaded")

}

var LevelMap string
