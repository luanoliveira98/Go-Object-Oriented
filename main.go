package main

import (
	"fmt"

	"./clientes"
	"./contas"
)

func main() {
	clienteBruno := clientes.Titular{"Bruno", "123.111.123.12", "Desenvolvedor GO"}
	contaDoBruno := contas.ContaCorrente{clienteBruno, 123, 123456, 100}

	fmt.Println(contaDoBruno)
}
