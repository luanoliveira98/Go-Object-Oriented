package main

import (
	"fmt"

	"./contas"
)

func main() {
	contaDaEstefane := contas.ContaCorrente{Titular: "Estefâne", Saldo: 300}
	contaDoLuan := contas.ContaCorrente{Titular: "Luan", Saldo: 100}

	status := contaDaEstefane.Transferir(200, &contaDoLuan)

	fmt.Println(status)
	fmt.Println(contaDaEstefane)
	fmt.Println(contaDoLuan)
}
