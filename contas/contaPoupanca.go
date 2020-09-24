package contas

import "../clientes"

// ContaPoupanca de um titular
type ContaPoupanca struct {
	Titular                             clientes.Titular
	NumeroAgencia, NumeroConta, Opercao int
	saldo                               float64
}

// Sacar é o método de saque de uma conta poupança
func (c *ContaPoupanca) Sacar(valorDoSaque float64) bool {
	if valorDoSaque > 0 && valorDoSaque <= c.saldo {
		c.saldo -= valorDoSaque
		return true
	}
	return false
}

// Depositar é o método de depósito de uma conta poupança
func (c *ContaPoupanca) Depositar(valorDoDeposito float64) bool {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return true
	}
	return false
}

// GetSaldo é o método de visualização do saldo de uma conta poupança
func (c *ContaPoupanca) GetSaldo() float64 {
	return c.saldo
}
