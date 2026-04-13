package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hadygust/movie/internal/handler"
	"github.com/hadygust/movie/internal/repository"
	"github.com/hadygust/movie/internal/service"
	"gorm.io/gorm"
)

// mount
func (app *application) mount() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.Data(http.StatusOK, "application/json", []byte("Hello World"))
	})

	userRepo := repository.NewUserRepository(&app.db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	users := r.Group("/user")
	users.GET("/", userHandler.AllUser)
	users.POST("/register", userHandler.Register)
	users.POST("/login", userHandler.Login)
	return r
}

// run
func (app *application) run(h *gin.Engine) {

	h.Run(app.config.address)
}

type application struct {
	config config
	db     gorm.DB
}

type config struct {
	address string
	db      dbConfig
}

type dbConfig struct {
	dsn string
}
