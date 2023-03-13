package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

func UpdatePassword(conn *pgx.Conn, entry AccountEntry) {
	acc := entry.AccNumber
	psw := entry.Password

	query := fmt.Sprintf("UPDATE accounts SET password = '%s' where accnumber = '%s';",
		psw,
		acc,
	)

	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		fmt.Println("updating entry: ", err.Error())
	} else {
		fmt.Println("updating entry: Success")
	}

	err = conn.Close(context.Background())
	if err != nil {
		fmt.Println("closing connection: ", err.Error())
	} else {
		fmt.Println("closing connection: Success")
	}

	defer conn.Close(context.Background())
	fmt.Println("Database connection closed")
}
