package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type configJSON struct {
	Normal    string `json:"normal"`
	SnakeCase string `json:"snake_case"`
}

func main() {

	// Unmarshalling from viper config:
	config := viper.New()
	config.SetConfigFile("./config.json")
	err := config.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	c := configJSON{}
	config.UnmarshalKey("data", &c)

	// Print out the values of the fields
	fmt.Printf("normal: `%s`\n", c.Normal)
	fmt.Printf("snake_case: `%s`\n", c.SnakeCase) // <-- Notice how this is empty
}
