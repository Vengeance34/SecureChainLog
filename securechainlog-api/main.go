package main

import (
	"fmt"
	"log"
	"net/http"

	"securechainlog-api/db"
	"securechainlog-api/handlers"

	"github.com/gorilla/mux"
)
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("🔍 %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer database.Close()

	router := mux.NewRouter()
	router.Use(loggingMiddleware)

	// 🔹 Root test route
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "🚀 SecureChainLog API is running!")
	}).Methods("GET")

	// 🔹 API routes
	router.HandleFunc("/logs", handlers.GetAssetLogs(database)).Methods("GET")
	router.HandleFunc("/logs", handlers.CreateAssetLog(database)).Methods("POST")


	log.Println("✅ Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
