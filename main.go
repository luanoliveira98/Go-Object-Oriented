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

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso!", c.saldo
	}
	return "Valor do deposito menor que zero!", c.saldo
}

func main() {
	contaDaEstefane := ContaCorrente{}
	contaDaEstefane.titular = "Estefâne"
	contaDaEstefane.saldo = 500

	fmt.Println(contaDaEstefane.saldo)
	status, valor := contaDaEstefane.Depositar(1000)
	fmt.Println(status, valor)
}
