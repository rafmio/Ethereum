package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type AccountEntry struct {
	AccNumber string // 40 symbols
	Password  string // 256 symbols
}

func ConnectDB(entry AccountEntry) {
	fmt.Println("connection to database")
	// assing URL of DB to the variable
	// databaseURL := "postgres://postgres:qwq121@localhost:5433/postgres"
	databaseURL := "postgres://postgres:qwq121@localhost:5432/ethcontract"

	// establish a connection with specified DB
	dbPool, err := pgxpool.Connect(context.Background(), databaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer dbPool.Close()

	// execute insert query
	ExecuteInsertAccountQuery(dbPool, entry)
}

func ExecuteInsertAccountQuery(dbPool *pgxpool.Pool, entry AccountEntry) {
	fmt.Println("starting executing of insert query")
	acc := entry.AccNumber
	psw := entry.Password
	SQLQuery := fmt.Sprintf("INSERT INTO accounts (accnumber, password) VALUES('%s', '%s');", acc, psw)
	rows, err := dbPool.Query(context.Background(), SQLQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error inserting data to DB", err)
	}

	fmt.Println("wrote: ", rows)
}
