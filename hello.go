//hello.go
package main

//imports used on application
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

//Constants to easily set up the project

//Number of times that the monitoring runs
const monitoramentos = 3

//Time between the execution of the monitoring in seconds
const delay = 5

//Name of the file log
const logFilename = "logfile.txt"

//Main function of the program
func main() {

	exibeIntroducao()

	for {
		//Display the 'menu' of the program
		exibeMenu()
		//Read the typed option
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
		imprimeLogs()

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

	//Read the file with the sites and put the values on a slice(type of array)
	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		//For using range to get all the values of slice(array)
		for i, site := range sites {
			fmt.Println("Testando site: ", i, " site : ", site)
			testaSite(site)
		}
		//Pause between the executions of the monitoring loop
		time.Sleep(delay * time.Second)
		//Just add one empty line to easily read the outputs
		fmt.Println(" ")
	}

}

func testaSite(site string) {

	response, error := http.Get(site)

	//Displays an error if it occurs
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

	//Displays an error if it occurs
	if error != nil {
		fmt.Println("Ocorreu o erro: ", error)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		//Read each line considering the end of the line the binary '\n'
		linha, error := leitor.ReadString('\n')
		//Remove all blanck spaces and '\n'
		linha = strings.TrimSpace(linha)

		//Add each site on the file to a new position on slice(array)
		sites = append(sites, linha)

		//Breaks the loop when the file ends
		if error == io.EOF {
			break
		}

	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {
	//Open the log file if exists or create a new file to put infos line by line
	arquivo, error := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	//Displays an error if it occurs
	if error != nil {
		fmt.Println("Ocorreu um erro: ", error)
	}

	//Add infos to log file with current date site name and status of the site
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05 - ") + site + "- online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func imprimeLogs() {
	//Read the content of log file
	arquivo, error := ioutil.ReadFile(logFilename)

	//Displays an error if it occurs
	if error != nil {
		fmt.Println("Ocorreu um erro: ", error)
	}

	fmt.Println(string(arquivo))

}
