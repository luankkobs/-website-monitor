package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoring = 3
const delay = 5

func main() {

	showIntroduction()

	for {
		showMenu()

		comando := readCommand()

		switch comando {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Exibindo logs..")
			showLogs()
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
		logRegister(site, true)

	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		logRegister(site, false)
	}

}

func readSites() []string {

	var sites []string

	archive, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(archive)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	archive.Close()

	return sites
}

func logRegister(site string, status bool) {
	archive, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Erro:", err)
	}

	archive.WriteString(time.Now().Format("02/01/2006 15:04:05") +
		" - " + site + "- Online: " + strconv.FormatBool(status) + "\n")

	archive.Close()
}

func showLogs() {
	archive, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(archive))
}
