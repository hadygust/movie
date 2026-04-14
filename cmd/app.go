package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hadygust/movie/internal/handler"
	"github.com/hadygust/movie/internal/middleware"
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

	authMiddleware := middleware.NewAuthMiddleware(userService)

	users := r.Group("/user")
	users.GET("/all", authMiddleware.RequireAuth, userHandler.AllUser)
	users.POST("/register", authMiddleware.RequireNonUser, userHandler.Register)
	users.POST("/login", authMiddleware.RequireNonUser, userHandler.Login)
	users.GET("/logged-in", authMiddleware.RequireAuth, userHandler.LoggedIn)
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
