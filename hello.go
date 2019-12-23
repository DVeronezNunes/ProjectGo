//hello.go
package main

import "fmt"
import "os"
import "net/http"

func main() {
	
	exibeIntroducao()

	for{
		exibeMenu()
		leComando()
	}
}

func exibeIntroducao(){

	nome   := "Diego"
	versao := 1.1

	fmt.Println("Olá sr(a).", nome)	
	fmt.Println("Estamos na versão: ", versao)
	fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")

}

func leComando() int {
	
	var comand int

	fmt.Scan(&comand)	
	fmt.Println("O valor da variavel comand é :" , comand)
	
	return comand
}

func exibeMenu(){
	switch leComando() {
	case 1:
		iniciarMonitoramento()	
	
	case 2:
		fmt.Println("Exibindo logs")
		
	case 0:
		fmt.Println("Encerrando programa")
		os.Exit(0)
	default:
		fmt.Println("Comando desconhecido")
		os.Exit(-1)
	}
}

func iniciarMonitoramento(){

	fmt.Println("Monitorando")
	site := "https://www.alura.com.br/"
	response,_ := http.Get(site)
	//fmt.Println(response)
	
	if response.StatusCode == 200{
		fmt.Println("Site: ", site, " foi carregado com sucesso")
	}else{
		fmt.Println("Site: ", site, " está com erro: ", response.StatusCode)

	}

}

