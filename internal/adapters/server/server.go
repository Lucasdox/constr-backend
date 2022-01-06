package server

import (
	"go.uber.org/zap"
	"net/http"
	"time"
)

func StartHttpServer(r http.Handler) {
	l := zap.L()
	s := &http.Server{
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 5 * time.Second,
		Addr:":8080",
		Handler: r,
	}

	l.Info("Serving Http")
	err := s.ListenAndServe()
	if err != nil {
		l.Fatal("Failed to serve http")
	}
}
