package main

import (
	"context"
	"fmt"
	"os"
)

func GetUsersWallets() []string {
	var sqlResponse []string

	conn, err := EstablishConnectionDB()
	if err != nil {
		fmt.Println("connection to DB failed. Exiting...")
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "select accnumber from accounts;")
	if err != nil {
		fmt.Println("extracting list of accounts: ", err.Error())
		os.Exit(1)
	}

	defer rows.Close()

	for rows.Next() {
		var accNumDB string
		rows.Scan(&accNumDB)
		sqlResponse = append(sqlResponse, accNumDB)
	}

	// fmt.Println()
	// fmt.Printf("Accounts:")
	// for i, val := range sqlResponse {
	// 	fmt.Println(i, " - ", val)
	// }

	return sqlResponse
}
