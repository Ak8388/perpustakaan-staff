package main

import (
	"latihangolang/Connection"
	"latihangolang/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := Connection.Connection()

	eng := routes.Routes{
		R:  r,
		Db: db,
	}

	eng.Routers()

	r.Run()
}
