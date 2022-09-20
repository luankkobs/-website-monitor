package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoring = 3
const delay = 5

func main() {

	showIntroduction()
	readSites()

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

func returnNameAndAge() (string, int) {
	name := "Luan"
	age := 20
	return name, age
}

func showIntroduction() {
	name := "Luan!"
	version := 2.1
	fmt.Println("Olá, Sr.", name)
	fmt.Println("A versão atual do programa é", version)
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
	fmt.Println("")

	return comandoRead
}

func startMonitoring() {
	fmt.Println("Monitorando..")

	//sites := []string{"https://random-status-code.herokuapp.com/",
	// "https://www.alura.com.br", "https://www.caelum.com.br"}

	sites := readSites()

	for i := 0; i < monitoring; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}

}

func readSites() []string {

	var sites []string

	archive, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(archive)
	return sites
}
