package server

import (
	"net/http"
	"time"
)

func StartHttpServer(r http.Handler) {
	s := &http.Server{
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 5 * time.Second,
		Addr:":8080",
		Handler: r,
	}

	s.ListenAndServe()
}
