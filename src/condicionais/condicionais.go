package condicionais

import "fmt"

/*Show Mostra o resultado do número MAIOR QUE*/
func Show() {
	numero := 20
	n := 10

	if n := 5; n > 0 {
		fmt.Println(n)
		fmt.Println("MAIOR QUE 0")
	} else {
		fmt.Println("MENOR QUE 0")
	}

	fmt.Println(n)

	switch numero {
	case 2:
		fmt.Println("Número 2")
	case 5:
		fmt.Println("Número 5")
	default:
		fmt.Println("PADRÂO")
	}
}
