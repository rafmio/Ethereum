package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type AccountEntry struct {
	AccNumber string
	Password  string
}

func InsertAccount(conn *pgx.Conn, entry AccountEntry) error {
	fmt.Println("Adding the entry to database...")

	acc := entry.AccNumber
	psw := entry.Password

	query := fmt.Sprintf("INSERT INTO accounts (accnumber, password) VALUES('%s', '%s');",
		acc,
		psw,
	)

	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		fmt.Println("adding entry: ", err.Error())
	} else {
		fmt.Println("adding entry: Success")
	}

	err = conn.Close(context.Background())
	if err != nil {
		fmt.Println("closing connection: ", err.Error())
	} else {
		fmt.Println("closing connection: Success")
	}

	defer conn.Close(context.Background())
	
	return err
}
