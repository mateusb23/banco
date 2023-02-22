package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDaBruna := contas.ContaPoupanca{}
	contaDaBruna.Depositar(100)
	PagarBoleto(&contaDaBruna, 60)

	fmt.Println(contaDaBruna.ObterSaldo())

	contaDoMateus := contas.ContaCorrente{}
	contaDoMateus.Depositar(500)
	PagarBoleto(&contaDoMateus, 380)

	fmt.Println(contaDoMateus.ObterSaldo())
}
