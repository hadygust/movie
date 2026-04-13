package cmd

import "github.com/jackc/pgx"

type application struct {
	config Config
	db     pgx.Conn
}

type config struct {
	
	db
}
