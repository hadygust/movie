package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
)

// mount
func (app *application) mount() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.Data(http.StatusOK, "application/json", []byte("Hello World"))
	})

	return r
}

// run
func (app *application) run(h *gin.Engine) {
	h.Run(app.config.address)
}

type application struct {
	config config
	db     pgx.Conn
}

type config struct {
	address string
	db      dbConfig
}

type dbConfig struct {
	dsn string
}
