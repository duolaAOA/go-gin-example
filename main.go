package main

import (
	"fmt"
	"log"

	"go-gin-example/pkg/settings"
	"go-gin-example/routers"
)

func main() {
	router := routers.InitRouter()
	addr := fmt.Sprintf("0.0.0.0:%d", settings.HTTPPort)
	log.Println("Ready to run http service on", addr)
	if err := router.Run(addr); err != nil {
		log.Fatal(err)
	}
}
