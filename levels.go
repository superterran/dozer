package main

import (
	"fmt"
	"strings"

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

	LevelString = Level.GetString("level")

	fmt.Println(levelName + " is loaded")

	rows := strings.Split(LevelString, "\n")
	for y, row := range rows {
		col := strings.Split(row, "")
		for x, char := range col {
			LevelMap[y][x] = char
			fmt.Println(x, y, char)
			width = x
			height = y
		}
	}

}

var LevelString string

var LevelMap [255][255]string

var width int = 0
var height int = 0
