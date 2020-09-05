package main

import (
	"Go_Go_Go/graph"
	"Go_Go_Go/graph/generated"
	"Go_Go_Go/postgre"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-pg/pg/v10"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

const defaultPort = "7000"

func main() {
	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200", "https://eltube-263a1.web.app", "http://localhost:7000"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	db := postgre.New(&pg.Options{
		Addr:     ":5432",
		User:     "gyhkptjvuefhql",
		Password: "e40ac6782e2c71a5a8b379de4d863f4d08dc3e824d7963d15d3713ec8ae60bac",
		Database: "d5d8tj6dob63ih",
	})

	option, err := pg.ParseURL("postgres://gyhkptjvuefhql:e40ac6782e2c71a5a8b379de4d863f4d08dc3e824d7963d15d3713ec8ae60bac@ec2-184-73-249-9.compute-1.amazonaws.com:5432/dadrcqhmc07ii")
	if err != nil {
		panic(err);
	}

	db = pg.Connect(option);

	db.AddQueryHook(postgre.DBLogger{})

	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
