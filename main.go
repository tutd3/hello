package main

import (
    "fmt"
    "net/http"
)

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
    message := "Hello world!" // Corrected the typo here
    fmt.Fprintln(w, message)
    fmt.Println(message) // Log the message to the console
}

func main() {
    http.HandleFunc("/", welcomeHandler)
    fmt.Println("Server is listening on port 3060...")
    if err := http.ListenAndServe(":3060", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
