package main

import "fmt"

type CheckingAccount struct {
	holder        string
	numberAgency  int
	numberAccount int
	ballance      float64
}

func main() {
	fmt.Println(CheckingAccount{})
}
