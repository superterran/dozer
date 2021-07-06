package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

var LevelString string

var LevelMap [255][255]string

var width int = 0

var height int = 0

var Level *viper.Viper

func LoadLevel(levelName string) {

	Level = viper.New()

	if _, err := os.Stat(levelName); os.IsNotExist(err) {
		// path/to/whatever does not exist
		Level.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
		Level.AddConfigPath("levels")
		Level.SetConfigName(levelName)
	} else {

		fmt.Println(levelName)
		dir := filepath.Dir(levelName)
		Level.AddConfigPath(dir)
		Level.SetConfigFile(levelName)
	}

	err := Level.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	holes = 0

	currentLevelName = levelName

	LevelString = Level.GetString("level")

	rows := strings.Split(LevelString, "\n")
	for y, row := range rows {
		col := strings.Split(row, "")
		for x, char := range col {
			LevelMap[y][x] = char

			if char == "o" {
				holes++
			}

			width = x
			height = y
		}
	}

	isSpawned = false

	fmt.Println(levelName + " is loaded")

}
