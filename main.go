package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	pgun := os.Getenv("PG_USERNAME")
	pgpw := os.Getenv("PG_PASSWORD")
	pghost := os.Getenv("PG_HOST")
	pgdb := os.Getenv("PG_DATABASE")
	pgtbl := os.Getenv("PG_TABLE")

	pgstr := makePGConnectionStr(pgun, pgpw, pghost, pgdb)
	conn, err := pgx.Connect(context.Background(), pgstr)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close(context.Background())

	inet := net{}
	inet.GetPublicIP()

	tx, err := conn.Begin(context.Background())
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(), "update " + pgtbl + " set public = $1 where id = $2", inet.PublicIP, 1)
	if err != nil {
		panic(err.Error())
	}

	err = tx.Commit(context.Background())
	if err != nil {
		panic(err.Error())
	}

	conn.Close(context.Background())
	fmt.Println("Success!")
	os.Exit(0)
}

func makePGConnectionStr(un, pw, host, db string) string {
	return "postgres://" + un + ":" + pw + "@" + host + "/" + db
}
