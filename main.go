package main

import (
	"fmt"
	"log/slog"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dreadster3/definitelynotredis/pkg/engine"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	engine := engine.NewEngine(logger)
	engine.GCInterval = time.Second
	engine.Cache.Logger = logger.With("component", "cache")
	engine.StartGC()

	engine.Cache.SetWithTTL("something", 1, 5*time.Second)
	engine.Cache.Set("somethingelse", 2)

	fmt.Println("Press Ctrl+C to exit")
	<-c
}
