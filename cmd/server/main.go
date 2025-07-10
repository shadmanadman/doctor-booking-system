package main

import (
    "doctor-booking-system/internal/graphql/graph"
    "doctor-booking-system/internal/db"
    "doctor-booking-system/internal/middleware"
    "github.com/go-chi/chi/v5"
    "github.com/joho/godotenv"
    "log"
    "net/http"
    "os"

    "github.com/99designs/gqlgen/graphql/handler"
    "github.com/99designs/gqlgen/graphql/playground"
)

func main(){
	// Load .env (only for local development)
    if err := godotenv.Load(); err != nil {
        log.Println("⚠️  No .env file found, using system environment variables")
    }

    db.Connect()

    r := chi.NewRouter()
    r.Use(middleware.JWTAuthMiddleware)

    srv := handler.NewDefaultServer(generated.NewExecutableSchema(
        generated.Config{Resolvers: &graph.Resolver{DB: db.DB}},
    ))

    r.Handle("/", playground.Handler("GraphQL Playground", "/query"))
    r.Handle("/query", srv)


      port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("🚀 Server running at http://localhost:%s/", port)
    log.Fatal(http.ListenAndServe(":"+port, r))
}