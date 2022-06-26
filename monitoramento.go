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

const MONITORAMENTOS = 5
const DELAY = 10

func main() {
	showTelaBoasVindas()
	showMenu()

	for {
		executaOpcao(getOpcao())
		showMenu()
	}
}

func executaOpcao(opcao int) {
	switch opcao {
	case 1:
		inciandoMnitoramento()
	case 2:
		fmt.Println("Exibindo logs...")
		imprimeLog()
	case 0:
		fmt.Println("Saindo do programa.")
		os.Exit(0)
	default:
		fmt.Println("Opação desconhecida.")
		os.Exit(-1)
	}
}

func imprimeLog() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}

func inciandoMnitoramento() {

	sites := lerArquivo("sites.txt")

	for i := 0; i < MONITORAMENTOS; i++ {
		fmt.Println("Iniciando...")
		fmt.Println("Monitoramento", i+1, ":")

		for _, site := range sites {
			getSite(site)
		}

		time.Sleep(DELAY * time.Second)
		fmt.Println("")
	}
}

func getSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu erro:", err)
	}

	var isRunning bool

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso")
		isRunning = true
	} else {
		fmt.Println("Site:", site, "não encontrado ou com problema. Status code:", resp.StatusCode)
		isRunning = false
	}
	registraLog(site, isRunning)

}

func lerArquivo(arq string) []string {
	var sites []string

	arquivo, err := os.Open(arq)
	//arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu erro ao tentar abrir o arquivo:", err)
	}

	dadosDoArquivo := bufio.NewReader(arquivo)

	for {
		linhaAtual, err := dadosDoArquivo.ReadString('\n')

		sites = append(sites, strings.TrimSpace(linhaAtual))

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func showTelaBoasVindas() {
	nome := "Osvaldo"
	versao := 1.1
	fmt.Println("Olá", nome)
	fmt.Println("Versão do programa", versao)
}

func getOpcao() int {
	var op int
	fmt.Scan(&op)

	fmt.Println("A opção escolhida foi:", op)
	fmt.Println("")

	return op
}

func showMenu() {
	fmt.Println("")
	fmt.Println("1 - Inicar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")
}
