package main

import (
	"github.com/hadygust/movie/internal/env"
	"github.com/jackc/pgx"
)

func main() {

	cfg := config{
		address: ":8080",
		db: dbConfig{
			dsn: env.GetString("dsn", ""),
		},
	}

	app := application{
		config: cfg,
		db:     pgx.Conn{},
	}

	app.run(app.mount())
}
