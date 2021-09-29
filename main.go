package main

import (
	"context"
	"fmt"
	"go-blog-jwt-token/api/configs"
	"go-blog-jwt-token/api/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	configs.SetupConfiguration()
}

// @title Go Blog
// @version 1.0
// @description Service Blog
// @termsOfService http://swagger.io/terms/

// @contact.name Muhammad Rais Adlani
// @contact.url https://gitlab.com/mraisadlani
// @contact.email mraisadlani@gmail.com

// @license.name MIT
// @license.url https://gitlab.com/mraisadlani/golang-blog-with-mux-and-jwt-token/-/blob/main/LICENSE

// @securityDefinitions.apikey BearerAuth
// @in Header
// @name Authorization

// @host localhost:9876
// @BasePath /api
func main() {
	r := routes.SetupRoute()

	serve := &http.Server{
		Addr: fmt.Sprintf(":%s", configs.Config.PORT),
		WriteTimeout: configs.Config.WRITETIMEOUT * 10,
		ReadTimeout:  configs.Config.READTIMEOUT * 10,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	go func() {
		err := serve.ListenAndServe()

		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("Connected to port:", configs.Config.PORT)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<- c

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	serve.Shutdown(ctx)
	os.Exit(0)
}
