package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	var sites [4]string
	sites[0] = "https://random-status-code.herokuapp.com/"
	sites[1] = "https://www.alura.com.br"
	sites[2] = "https://www.caelum.com.br"

	fmt.Println(sites)
	showIntroduction()

	for {
		showMenu()

		comando := readCommand()

		switch comando {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Exibindo logs..")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando")
			os.Exit(-1)
		}
	}
}

func devolveNomeEIdade() (string, int) {
	nome := "Luan"
	idade := 20
	return nome, idade
}

func showIntroduction() {
	nome := "Luan!"
	versao := 2.1
	fmt.Println("Olá, Sr.", nome)
	fmt.Println("A versão atual do programa é", versao)
}

func showMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Saindo do programa..")
}

func readCommand() int {
	var comandoRead int
	fmt.Scan(&comandoRead)
	fmt.Println("O comando escolhido foi", comandoRead)

	return comandoRead
}

func startMonitoring() {
	fmt.Println("Monitorando..")
	var sites [4]string
	sites[0] = "https://random-status-code.herokuapp.com/"
	sites[1] = "https://www.alura.com.br"
	sites[2] = "https://www.caelum.com.br"

	fmt.Println(sites)

	site := "https://random-status-code.herokuapp.com/"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}
