package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/shubhacker/gqlgen-todos/graph"
	"github.com/shubhacker/gqlgen-todos/graph/controller"
	"github.com/shubhacker/gqlgen-todos/graph/generated"
	"github.com/shubhacker/gqlgen-todos/graph/postgres"
	"github.com/shubhacker/gqlgen-todos/graph/auth"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load("dev.env")
	if err != nil {
		log.Println("error loading env", err)
	}

	postgres.InitDbPool()
	pool := postgres.GetPool()
	controller.InitCodes(pool)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	router.Use(auth.AuthMiddleWare(pool))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}