package http

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var httpServer *http.Server

func InitRouter(ctx context.Context) {
	router := mux.NewRouter()
	port := 8081

	//route
	//router.Handle("/ping", server.Ping()).Methods(http.MethodGet)

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

	fmt.Printf(`HTTP router started on port [%d]`, port)

	<-running
}

func StopServer(ctx context.Context) {
	if err := httpServer.Shutdown(ctx); err != nil {
		fmt.Printf(`Failed to gracefully shutdown server`)
	}

	fmt.Printf(`Success! Gracefully shutting down server`)
}