package main

import (
	"fmt"
	"testing"
)

func TestGenerateAcc(t *testing.T) {
	got := generateAcc()
	want1, want2 := 10, 30

	if len(got.AccNumber) != want1 || len(got.Password) != want2 {
		t.Errorf("len of got.AccNumber %q, wanted %q", len(got.AccNumber), want1)
		t.Errorf("len of got.Password %q, wanted %q", len(got.Password), want2)
	} else {
		fmt.Println("len: ", len(got.AccNumber), " ", len(got.Password))
	}
}

func BenchmarkGenerateAcc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateAcc()
	}
}

// go test
// go test -v
// go test -coverprofile=coverage.out
// go tool cover -html=coverage.out
// go test -bench=.
// go test -bench=generateAcc
