package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/gorilla/mux"
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/joho/godotenv"
    "hello/handlers"
)

var (
    dbPool *pgxpool.Pool
)

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
    message := "Hello world!"
    fmt.Fprintln(w, message)
    fmt.Println(message)
}

func main() {
    // Load environment variables
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Initialize PostgreSQL connection
    dbPool, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatalf("Unable to connect to PostgreSQL: %v", err)
    }
    defer dbPool.Close()

    // Ping the database to check connection
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := dbPool.Ping(ctx); err != nil {
        log.Fatalf("Failed to ping PostgreSQL: %v", err)
    } else {
        log.Println("Successfully connected to PostgreSQL")
    }

    // Setup routes
    router := mux.NewRouter()
    router.HandleFunc("/", welcomeHandler).Methods("GET")
    router.HandleFunc("/items", handlers.GetItems(dbPool)).Methods("GET")
    router.HandleFunc("/item", handlers.CreateItem(dbPool)).Methods("POST")
    router.HandleFunc("/item/{id}", handlers.UpdateItem(dbPool)).Methods("PUT")
    router.HandleFunc("/item/{id}", handlers.DeleteItem(dbPool)).Methods("DELETE")
    router.HandleFunc("/ping", handlers.Ping()).Methods("GET")

    fmt.Println("Server is listening on port 8080...")
    if err := http.ListenAndServe(":8080", router); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
