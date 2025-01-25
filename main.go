package main

import (
	"fmt"
	"log"

	"github.com/jzaager/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Current config: %+v\n", cfg)

	err = cfg.SetUser("justin")
	if err != nil {
		log.Fatal(err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Current config: %+v\n", cfg)
}
