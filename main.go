package main

import (
	"flag"
	"github.com/MKA-Nigeria/mkanmedia-go/config"
	"github.com/MKA-Nigeria/mkanmedia-go/routes"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func init() {

	//environment flag parser
	environment := flag.String("env", "dev", "the environment in which the application should run")
	flag.Parse()

	config.LoadConfig()

	switch *environment {
	case "prod":
		viper.Set("env", "prod")
	case "dev":
		viper.Set("env", "dev")
	}

	log.Printf("%s environment started", viper.Get("env"))
}

func main() {

	env := viper.GetString("env")

	port := viper.GetString(env + ".port")

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
