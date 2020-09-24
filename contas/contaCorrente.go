package contas

import "../clientes"

// ContaCorrente de um titular
type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

// Sacar é o método de saque de uma conta corrente
func (c *ContaCorrente) Sacar(valorDoSaque float64) bool {
	if valorDoSaque > 0 && valorDoSaque <= c.saldo {
		c.saldo -= valorDoSaque
		return true
	}
	return false
}

// Depositar é o método de depósito de uma conta corrente
func (c *ContaCorrente) Depositar(valorDoDeposito float64) bool {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
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

// GetSaldo é o método de visualização do saldo de uma conta corrente
func (c *ContaCorrente) GetSaldo() float64 {
	return c.saldo
}
