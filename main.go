package main

import "fmt"
import "github.com/yanvic/projeto-2-Go/contas"
import "github.com/yanvic/projeto-2-Go/clientes"


func main() {
	//ContaYan := ContaConrrente{titular: "yan", numeroAgencia: 250, numeroConta: 12456, saldo: 255}
	contaYan := contas.ContaConrrente{Titular: "yan", Saldo: 300}
	contaVitor := contas.ContaConrrente{Titular: "vitor", Saldo: 200}

	status := contaYan.Tranferir(200, &contaVitor)

	fmt.Println(status)
	fmt.Println(contaYan)
	fmt.Println(contaVitor)
	//contaYan.titular = "yan"
	//contaYan.saldo = 100

	//fmt.Println(contaYan.saldo)

	//fmt.Println(contaYan.Saque(50))
	//fmt.Println(contaYan.saldo)

	//status, valor := contaYan.Depositar(500)
	//fmt.Println(status, valor)

}
