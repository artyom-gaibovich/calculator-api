package main

import (
	"calculator-api/internal/config"
	"fmt"
)

func main() {
	cfg, _ := config.InitConfig()

	fmt.Printf(cfg.AppName)
	fmt.Println("hell")

}
