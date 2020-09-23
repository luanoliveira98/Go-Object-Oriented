package main

import "fmt"

type CheckingAccount struct {
	holder        string
	numberAgency  int
	numberAccount int
	ballance      float64
}

func main() {
	LuanAccount := CheckingAccount{
		holder:        "Luan",
		numberAgency:  589,
		numberAccount: 123456,
		ballance:      125.5,
	}

	BrunaAccount := CheckingAccount{"Bruna", 222, 111222, 200}

	fmt.Println(LuanAccount)
	fmt.Println(BrunaAccount)
}
