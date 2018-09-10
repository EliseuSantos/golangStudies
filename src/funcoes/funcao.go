package funcoes

import "fmt"

/*Show Mostra a soma de 2 número utilizando uma função*/
func Show() {
	data, msg := soma(10, 20)
	if msg != "" {
		fmt.Println(msg)
	}

	fmt.Println("O Resultado é:", data)
}

func soma(n1, n2 int) (int, string) {
	fmt.Println("Iniciando chamada da função soma")

	data := n1 + n2
	if data > 10 {
		return data, "Valor é maior que 10"
	}

	return data, ""
}
