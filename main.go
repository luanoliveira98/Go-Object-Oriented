package main

import (
	"fmt"

	"./contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) bool
}

func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 60)

	fmt.Println(contaDoDenis.GetSaldo())

	contaDaLuisa := contas.ContaPoupanca{}
	contaDaLuisa.Depositar(500)
	PagarBoleto(&contaDaLuisa, -400)

	fmt.Println(contaDaLuisa.GetSaldo())
}
