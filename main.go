package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	}
	return "Saldo insuficiente!"
}

func main() {
	contaDaEstefane := ContaCorrente{}
	contaDaEstefane.titular = "Estefâne"
	contaDaEstefane.saldo = 500

	fmt.Println(contaDaEstefane.saldo)

	fmt.Println(contaDaEstefane.Sacar(-100))
	fmt.Println(contaDaEstefane.saldo)
}
