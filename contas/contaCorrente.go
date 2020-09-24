package contas

import "../clientes"

// ContaCorrente de um titular
type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

// Sacar é o método de saque de uma conta corrente
func (c *ContaCorrente) Sacar(valorDoSaque float64) bool {
	if valorDoSaque > 0 && valorDoSaque <= c.Saldo {
		c.Saldo -= valorDoSaque
		return true
	}
	return false
}

// Depositar é o método de depósito de uma conta corrente
func (c *ContaCorrente) Depositar(valorDoDeposito float64) bool {
	if valorDoDeposito > 0 {
		c.Saldo += valorDoDeposito
		return true
	}
	return false
}

// Transferir é o método de transferência de uma conta corrente
func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if c.Sacar(valorDaTransferencia) {
		contaDestino.Depositar(valorDaTransferencia)
		return true
	}
	return false
}
