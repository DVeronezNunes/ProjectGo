//hello.go
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

const monitoramentos = 3
const delay = 5

func main() {
	registraLog("fake", false)
	exibeIntroducao()

	for {
		exibeMenu()
		leComando()
	}
}

func exibeIntroducao() {

	nome := "Diego"
	versao := 1.1

	fmt.Println("Olá Sr(a).", nome)
	fmt.Println("Estamos na versão: ", versao)
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

}

func leComando() int {

	var comand int

	fmt.Scan(&comand)
	fmt.Println("O valor da variavel comand é :", comand)

	return comand
}

func exibeMenu() {

	comand := leComando()
	switch comand {
	case 1:
		iniciarMonitoramento()

	case 2:
		fmt.Println("Exibindo logs")
		ImprimeLogs()

	case 0:
		fmt.Println("Encerrando programa")
		os.Exit(0)
	default:
		fmt.Println("Comando desconhecido")
		os.Exit(-1)
	}
}

func iniciarMonitoramento() {

	fmt.Println("Monitorando")

	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site: ", i, " site : ", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println(" ")
	}

}

func testaSite(site string) {

	response, error := http.Get(site)

	if error != nil {
		fmt.Println("Ocorreu um erro: ", error)
	}

	if response.StatusCode == 200 {
		fmt.Println("Site: ", site, " foi carregado com sucesso")
		registraLog(site, true)
	} else {
		fmt.Println("Site: ", site, " está com erro: ", response.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, error := os.Open("sites.txt")
	//arquivo, error := ioutil.ReadFile("sites.txt")

	if error != nil {
		fmt.Println("Ocorreu o erro: ", error)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, error := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if error == io.EOF {
			break
		}

	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {
	arquivo, error := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if error != nil {
		fmt.Println("Ocorreu um erro: ", error)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05 - ") + site + "- online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func ImprimeLogs() {

	arquivo, error := ioutil.ReadFile("log.txt")

	if error != nil {
		fmt.Println("Ocorreu um erro: ", error)
	}

	fmt.Println(string(arquivo))

}
