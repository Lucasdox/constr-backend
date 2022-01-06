package main

import (
	"github.com/Lucasdox/constr-backend/internal/adapters/http/router"
	"github.com/Lucasdox/constr-backend/internal/adapters/server"
)

func main() {
	r := router.Router()

	server.StartHttpServer(r)
}
