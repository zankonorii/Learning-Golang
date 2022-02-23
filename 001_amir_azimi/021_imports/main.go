package main

import (
	"fmt"
	"github.com/pelletier/go-toml"
	//zanko "github.com/pelletier/go-toml" // zanko is an alias for this imported file
	"log"
	"os"
)

type Config struct {
	//We can create a struct in this struct or any struct
	Login struct {
		User     string
		Password string
		Register struct {
		}
	}
}

func main() {
	file, err := os.Open("config.toml")

	if err != nil {
		log.Fatalf("Error: can't open config file - %s\n", err)
	}

	defer file.Close()

	ctg := &Config{}
	document := toml.NewDecoder(file)

	if err := document.Decode(ctg); err != nil {
		log.Fatalf("Error: can't decode config file - %s\n", err)
	}

	fmt.Printf("%+v", ctg)
}
