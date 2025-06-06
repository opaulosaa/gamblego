package main

import "fmt"

func GetName() string {
	name := ""

	fmt.Println("Bem vindo ao Cassino do Paulosa...")
	fmt.Printf("Digite seu nome: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return ""
	}
	fmt.Printf("Aposte com cuidado, %s\n", name)
	return name
}

func GetBet(balance uint) uint {
	var bet uint
	for true {
		fmt.Printf("Digite o valor para apostar, ou digite 0 para sair (carteira = R$%d): ", balance)
		fmt.Scan(&bet)

		if bet > balance {
			fmt.Println("Aposta inválida, dinheiro insuficiente")
		} else {
			break
		}
	}
	return bet
}
