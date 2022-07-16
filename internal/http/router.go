package http

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sizijay/iban-validator/internal/server"
	"net/http"
	"time"
)

var httpServer *http.Server

func InitRouter(ctx context.Context) {
	router := mux.NewRouter()
	port := 8081

	// ping
	router.Handle("/ping", methodControl(http.MethodGet, server.Ping()))

	// iban validator
	router.Handle("/validate/iban", methodControl(http.MethodPost, server.ValidateIBAN()))

	StartServer(ctx, port, router)
}

func StartServer(ctx context.Context, port int, r http.Handler) {
	running := make(chan interface{}, 1)

	httpServer := &http.Server{
		Addr:         fmt.Sprintf(`:%d`, port),
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		IdleTimeout:  time.Second * 5,
		Handler:      r,
	}

	go func(ctx context.Context) {
		err := httpServer.ListenAndServe()
		if err != nil {
			fmt.Printf(`Cannot start web server: %v`, err)
		}
		running <- `done`
	}(ctx)

	go func(ctx context.Context) {
		fmt.Println(http.ListenAndServe(":6069", nil))
		<- running
	}(ctx)

	fmt.Println(fmt.Sprintf(`HTTP router started on port [%d]`, port))
}

func StopServer(ctx context.Context) {
	if err := httpServer.Shutdown(context.Background()); err != nil {
		fmt.Printf(`Failed to gracefully shutdown server`)
	}

	fmt.Println(fmt.Printf(`Success! Gracefully shutting down server`))
}

func methodControl(method string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method {
			h.ServeHTTP(w, r)
		} else {
			http.Error(w, `Method not allowed`, http.StatusMethodNotAllowed)
		}
	})
}