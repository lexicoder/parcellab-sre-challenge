package main

import (
	"fmt"
	"log"
	"net/http"
	"parcellab-sre-challenge/internal/pkg/config"
	"parcellab-sre-challenge/internal/pkg/handlers"
	"parcellab-sre-challenge/internal/pkg/health"
)

func main() {
	cfg := config.NewConfig()

	r := http.NewServeMux()

	r.HandleFunc("GET /", handlers.SalutationHandler(cfg.Salutation))
	r.HandleFunc("GET /health", health.HealthGet())

	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.ListenPort),
		Handler:      r,
		ReadTimeout:  cfg.ServerReadTimeout,
		WriteTimeout: cfg.ServerWriteTimeout,
		IdleTimeout:  cfg.ServerIdleTimeout,
	}

	fmt.Printf("Server is running on port %s\n", cfg.ListenPort)
	log.Fatal(s.ListenAndServe())
}
