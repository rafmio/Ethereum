package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type ConnectionConfig struct {
	DBMS        string // Database management system
	User        string // User of DBMS
	Password    string // Password to access
	NetLocation string // Net location of database: localhost or IP
	Port        string // :port
	DBName      string // DB name
	Param       string // ?param1=value1&...
}

func (cs *ConnectionConfig) GetConnString() string {
	connString := fmt.Sprintf("%s://%s:%s@%s:%s/%s?%s",
		cs.DBMS,
		cs.User,
		cs.Password,
		cs.NetLocation,
		cs.Port,
		cs.DBName,
		cs.Param,
	)
	return connString
}

// Hardcoded connection config
var ConnConf ConnectionConfig = ConnectionConfig{
	DBMS:        "postgresql",
	User:        "postgres",
	Password:    "qwq121",
	NetLocation: "194.58.102.129",
	Port:        "9432",
	DBName:      "ethcontract",
	Param:       "sslmode=disable",
}

func EstablishConnectionDB() (*pgx.Conn, error) {
	config, err := pgx.ParseConfig(ConnConf.GetConnString())
	if err != nil {
		fmt.Println("parsing config: ", err.Error())
	} else {
		fmt.Println("parsing config: Success")
	}

	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		fmt.Println("connecting to database: ", err.Error())
	} else {
		fmt.Println("connection to database: Success")
	}
	// defer conn.Close(context.Background())

	return conn, err
}
