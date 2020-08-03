package main

import (
	"fmt"
	"net/http"

	"go-gin-example/pkg/settings"
	"go-gin-example/routers"
)

func main() {
	router := routers.InitRouter()
	
	s := &http.Server{
		Addr: fmt.Sprintf("%d", settings.HTTPPort),
		Handler: router,
		ReadTimeout: settings.ReadTimeout,
		WriteTimeout: settings.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	_ = s.ListenAndServe()
}
