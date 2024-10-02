package main

import "fmt"

func main() {

	introducao()

	fmt.Println("1-Iniciar monitoramento")
	fmt.Println("2-Exibir Logs")
	fmt.Println("0-Encerrar execução")

	var comando int
	fmt.Scan(&comando)

	/*if comando == 1 {
		fmt.Println("Iniciando processo de monitoramento. ")
	} else if comando == 2 {
		fmt.Println("Iniciando processo de Exibição de logs. ")
	} else if comando == 0 {
		fmt.Println("Encerrando execução... ")
	} else {
		fmt.Println("Entrada não identificada. Fim da execução. ")
	}*/

	switch comando {
	case 1:
		fmt.Println("Iniciando processo de monitoramento. ")
	case 2:
		fmt.Println("Iniciando processo de Exibição de logs. ")
	case 0:
		fmt.Println("Encerrando execução... ")
	default:
		fmt.Println("Entrada não identificada. Fim da execução. ")

	}

}
func introducao() {
	nome := "Salazar"
	idade := 28
	versao := 1.23
	fmt.Println("Olá Sr. ", nome, " Sua idade é ", idade)
	fmt.Println("Você está utilizando o Go na versão ", versao)
}
