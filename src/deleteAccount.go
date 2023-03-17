package main

import (
	"context"
	"fmt"
	"os"
)

func DeleteEntry(accNumber string) {
	conn, err := EstablishConnectionDB()
	if err != nil {
		fmt.Println("connection to DB failed. Exiting...")
		os.Exit(1)
	} else {
		fmt.Println("connection to DB: Success")
	}

	query := fmt.Sprintf("DELETE FROM accounts where accnumber = '%s';", accNumber)

	_, err = conn.Exec(context.Background(), query)
	if err != nil {
		fmt.Println("deleting account: ", err.Error())
	} else {
		fmt.Println("deleting account: Success")
	}

	err = conn.Close(context.Background())
	if err != nil {
		fmt.Println("closing connection: ", err.Error())
	} else {
		fmt.Println("closing connection: Success")
	}

	defer conn.Close(context.Background())
}
