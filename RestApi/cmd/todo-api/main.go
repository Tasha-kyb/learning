package main

import (
    "log"
    "net/http"
    "os"

    "RestApi/internal/http/middleware"
    "RestApi/internal/http/handlers"
    myhttp "RestApi/internal/http"
    "RestApi/internal/service"
    "RestApi/internal/storage/mem"
)

func main() {
    repo := mem.NewListRepo()
    listService := service.NewListService(repo)
    listHandler := handlers.NewListHandler(listService)

    server := myhttp.NewHTTPServer(listHandler)

    handler := middleware.RequestID(server)
    handler = middleware.Logging(handler)
    
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    
    log.Printf("Server starting on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, handler))
}