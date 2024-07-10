package main

import (
	"awesomeProject/internal/config"
	"awesomeProject/internal/db"
	"awesomeProject/internal/logger"
	"awesomeProject/internal/routes"
	"awesomeProject/internal/web"
	"context"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/julienschmidt/httprouter"
)

func main() {
	defer logger.Sync()
	defer db.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer handlePanic(ctx, cancel)

	if err := config.LoadEnv(); err != nil {
		panic(err)
	}

	config.SetupLogger(config.Env.Environment)

	config.SetupDBConnection(context.Background())

	router := httprouter.New()
	routes.Init(router)

	log.Fatal(http.ListenAndServe(config.Env.Port, router))
}

func handlePanic(ctx context.Context, cancel context.CancelFunc) {
	if recvr := recover(); recvr != nil {
		errorMessage := fmt.Sprintf("%v", recvr)
		err := web.ErrInternalServerError(errorMessage)
		logger.E(ctx, err, "panic",
			logger.Field("status", err.HTTPStatusCode()),
			logger.Field("stack", string(debug.Stack())),
		)
	}

	cancel()
}
