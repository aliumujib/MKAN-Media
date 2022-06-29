package main

import (
	"log"
	"net/http"

	"github.com/MKA-Nigeria/mkanmedia-go/config"
	"github.com/MKA-Nigeria/mkanmedia-go/routes"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

func init() {

	logrus.SetFormatter(&logrus.TextFormatter{})

	log.Printf("%s environment started", config.Env.ENV)
}

func main() {

	port := config.Env.PORT

	//	APPLY MIDDLEWARES
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})
	log.Printf("server running at %s", port)

	http.ListenAndServe(":"+port, c.Handler(routes.Router()))
}
