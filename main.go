package main

import "fmt"

// ContaCorrente ...
type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

// Sacar ...
func (c *ContaCorrente) Sacar(valorDoSaque float64) bool {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return true
	}
	return false
}

// Depositar ...
func (c *ContaCorrente) Depositar(valorDoDeposito float64) bool {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return true
	}
	return false
}

// Transferir ...
func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if c.Sacar(valorDaTransferencia) {
		contaDestino.Depositar(valorDaTransferencia)
		return true
	}
	return false
}

func main() {
	contaDaEstefane := ContaCorrente{titular: "Estefâne", saldo: 300}
	contaDoLuan := ContaCorrente{titular: "Luan", saldo: 100}

	status := contaDaEstefane.Transferir(500, &contaDoLuan)

	fmt.Println(status)
	fmt.Println(contaDaEstefane)
	fmt.Println(contaDoLuan)
}
