package internal

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/sizijay/iban-validator/internal/http"
	"os"
	"os/signal"
	"syscall"
)

func Init() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	tCtx := context.WithValue(context.Background(), "uuid", uuid.New())
	ctx, cancel := context.WithCancel(tCtx)
	defer cancel()

	http.InitRouter(ctx)

	select {
	case <-sigs:
		fmt.Println("Shutting down server", "OS interrupt")
		http.StopServer(ctx)
	}
}