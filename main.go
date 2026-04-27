package main

import (
	"fmt"
	"gator/internal/config"
	// "os"
)

func main() {
	cfg, ok := config.Read()
	if ok != nil {
		fmt.Printf("main: %v", ok)
	}

	ok = config.SetUser(cfg, "Conor")

	cfg, ok = config.Read()

	fmt.Println(cfg)

}
