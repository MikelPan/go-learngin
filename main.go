package main

import (
	"fmt"
	"github.com/MikelPan/go-learning/pkg/setting"
	"github.com/MikelPan/go-learning/routers"
	"net/http"
)

// @title Golang Gin API
// @version 1.0
// @description An devops of gin
// @termsOfService https://github.com/MikelPan/go-learning
// @host 127.0.0.1:8000
// @BasePath /api/v1

// @contact.url http://www.swagger.io/support

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}

