package main

import (
	"github.com/hadygust/movie/internal/env"
	"github.com/hadygust/movie/internal/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	cfg := config{
		address: ":8080",
		db: dbConfig{
			dsn: env.GetString("DSN", "host=localhost user=postgres password=postgres dbname=movie sslmode=disable"),
		},
	}

	conn, err := gorm.Open(postgres.Open(cfg.db.dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	app := application{
		config: cfg,
		db:     *conn,
	}

	app.db.AutoMigrate(&model.User{})
	app.db.AutoMigrate(&model.Movie{})

	app.run(app.mount())
}
