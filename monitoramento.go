package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoramentos = 5
const delay = 10

func main() {

	showTelaBoasVindas()
	lerSiteDoArquivo()
	montaMenu()

	for {
		executaOpcao(getOpcao())
		montaMenu()
	}

}

func executaOpcao(opcao int) {

	switch opcao {
	case 1:
		inciandoMnitoramento()
	case 2:
		fmt.Println("Exibindo logs")
	case 0:
		fmt.Println("Saindo do programa.")
		os.Exit(0)
	default:
		fmt.Println("Opação desconhecida.")
		os.Exit(-1)
	}
}

func inciandoMnitoramento() {

	sites := lerSiteDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		fmt.Println("Iniciando...")
		fmt.Println("Monitoramento", i+1, ":")
		for _, site := range sites {
			getSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func getSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso")
	} else {
		fmt.Println("Site:", site, "não encontrado ou com problema. Status code:", resp.StatusCode)
	}
}

func lerSiteDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	//arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu erro ao abrir o arquivo:", err)
	}

	info := bufio.NewReader(arquivo)

	for {
		linhaAtual, err := info.ReadString('\n')

		sites = append(sites, strings.TrimSpace(linhaAtual))

		if err == io.EOF {
			break
		}
	}

	return sites
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

func montaMenu() {
	fmt.Println("")
	fmt.Println("1 - Inicar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")
}
