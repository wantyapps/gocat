package main

import (
	"bufio"
	"os"
	c "github.com/wantyapps/gocat/config"
	"github.com/spf13/viper"
	"fmt"
)

func main() {
	viper.SetConfigName(".gocat.yml")
	homeDirectory, _ := os.UserHomeDir()
	viper.AddConfigPath(homeDirectory)
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	var config c.Config

	viperReadErr := viper.ReadInConfig()
	if viperReadErr != nil {
		fmt.Println(viperReadErr)
	}

	unmErr := viper.Unmarshal(&config)
	if unmErr != nil {
		fmt.Println(unmErr)
	}

	logfile, err := os.Create(homeDirectory + "/" + config.LogfileDir + config.LogfileName)
	if err != nil {
		fmt.Println(err)
	}

	writer := bufio.NewWriter(logfile)
	if len(os.Args) >= 2 {
		file, fileErr := os.Open(os.Args[1])

		if fileErr != nil {
			fmt.Println("Error: FileErr.")
			writer.WriteString("[LOG]: Error: FileErr.\n")
			writer.Flush()
		}

		defer file.Close()
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	} else {
		fmt.Println("Usage: " + os.Args[0] + " [filename]")
	}
}
