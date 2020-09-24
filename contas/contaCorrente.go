package contas

// ContaCorrente ...
type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

// Sacar ...
func (c *ContaCorrente) Sacar(valorDoSaque float64) bool {
	if valorDoSaque > 0 && valorDoSaque <= c.Saldo {
		c.Saldo -= valorDoSaque
		return true
	}
	return false
}

// Depositar ...
func (c *ContaCorrente) Depositar(valorDoDeposito float64) bool {
	if valorDoDeposito > 0 {
		c.Saldo += valorDoDeposito
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
