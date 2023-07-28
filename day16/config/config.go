package config

import (
	"context"
	"os"

	"fmt"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func DatabaseConnection()  {
	

	var err error
	databaseurl := "postgres://postgres:0987@127.0.0.1:5432/myproject_db"

	Conn, _ = pgx.Connect(context.Background(), databaseurl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("you have connected to database")
}