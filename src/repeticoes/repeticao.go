package repeticoes

import "fmt"

/*Show Mostra a estrutura de repetição*/
func Show() {
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for number := 0; number < 10; number++ {
		fmt.Println("loop", number)
	}

}
