package main

import (
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gorilla/websocket"
	"github.com/shubhacker/gqlgen-todos/graph/auth"
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
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				return r.Host == "*"
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})
	router.Use(auth.AuthMiddleware(pool))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	//router.Post("/authenticate",mapper.AuthenticateUserRest())
	http.Handle("/", http.FileServer(http.Dir("/tmp")))
	router.Handle("/query", srv)
	router.Handle("/v2/query/",srv)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}