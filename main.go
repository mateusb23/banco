package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	contaDaBruna := contas.ContaPoupanca{}
	contaDoMateus := contas.ContaCorrente{}

	fmt.Println(contaDaBruna)
	fmt.Println(contaDoMateus)
}
