package main

import (
	api "github.com/nekoze1210/echo-api-w3w/interfaces/api"
	"github.com/nekoze1210/echo-api-w3w/interfaces/api/server/middleware"
)

func main() {
	middleware.ExportDotEnv()
	api.RunServer()
}
