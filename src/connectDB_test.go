// https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package
package main

import (
	"fmt"
	"testing"
)

func TestEstablishConnectionDB(t *testing.T) {
	conn, err := EstablishConnectionDB()
	if err != nil {
		t.Errorf("Failed connection to database")
	} else {
		fmt.Println("Connection: success")
		fmt.Println("Connection preferences:")
		fmt.Println(conn)
		fmt.Println()
	}
}

func BenchmarkEstablishConnectionDB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EstablishConnectionDB()
	}
}

// go test
// go test -v
// go test -coverprofile=coverage.out
// go tool cover -html=coverage.out
// go test -bench=.
// go test -bench=EstablishConnectionDB
