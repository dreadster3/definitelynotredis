package engine

import (
	"log/slog"
	"time"
)

type Engine struct {
	Cache      *Cache
	GCInterval time.Duration
	Logger     *slog.Logger
}

func NewEngine(logger *slog.Logger) *Engine {
	return &Engine{
		Cache:      NewCache(),
		GCInterval: time.Minute,
		Logger:     logger,
	}
}

func (e *Engine) StartGC() {
	ticker := time.NewTicker(e.GCInterval)
	go func() {
		logger := e.Logger.With("component", "engine")
		for range ticker.C {
			logger.Debug("Cleaning expired entries")
			e.Cache.removeExpired()
		}
	}()
}
