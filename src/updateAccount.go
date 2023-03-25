package main

import (
	"context"
	"fmt"
	"os"
)

func UpdatePassword(entry AccountEntry) {
	conn, err := EstablishConnectionDB()
	if err != nil {
		fmt.Println("connection to DB failed. Exiting...")
		os.Exit(1)
	}
	acc := entry.AccNumber
	psw := entry.Password

	query := fmt.Sprintf("UPDATE accounts SET password = '%s' where accnumber = '%s';",
		psw,
		acc,
	)

	_, err = conn.Exec(context.Background(), query)
	if err != nil {
		fmt.Println("updating data: ", err.Error())
	} else {
		fmt.Println("updating data: Success")
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