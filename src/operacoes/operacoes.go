package operacoes

import "fmt"

/*Show Exibe operações matemáticas básica */
func Show() {
	numero1, numero2 := 10, 8
	fmt.Println("Multiplicação:\n", numero1*numero2)
	fmt.Println("Divisão:\n", numero1/numero2)
	fmt.Println("Subtração\n", numero1-numero2)
}
