package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	showIntroduction()
	for {
		showMenu()

		option := readOption()

		switch option {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este option")
			os.Exit(-1)
		}
	}
}

func startMonitoring() {
	fmt.Println("Monitorando...")
	url := "https://www.youtube.com/"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	} else if response.StatusCode == 200 {
		fmt.Println("Site está online")
	} else {
		fmt.Println("Site está offline")
	}
}

func showIntroduction() {
	name := "Douglas"
	version := 1.1
	fmt.Println("Olá, sr(a).", name)
	fmt.Println("Este programa está na versão", version)
}

func showMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func readOption() int {
	var chosenOption int
	fmt.Scan(&chosenOption)
	fmt.Println("O comando escolhido foi:", chosenOption)

	return chosenOption
}
