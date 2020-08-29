package main

import (
	"Go_Go_Go/graph"
	"Go_Go_Go/graph/generated"
	"Go_Go_Go/postgre"
	"github.com/go-chi/chi"
	"github.com/go-pg/pg/v10"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
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
		User:     "dxnovgghvhqzgd",
		Password: "6b52c9b13f3d63b5fa48db501e21f8a1b054c1cea225d221c1d9f081567c6ad5",
		Database: "d5d8tj6dob63ih",
	})

	option, err := pg.ParseURL("postgres://dxnovgghvhqzgd:6b52c9b13f3d63b5fa48db501e21f8a1b054c1cea225d221c1d9f081567c6ad5@ec2-54-91-178-234.compute-1.amazonaws.com:5432/d5d8tj6dob63ih")
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
