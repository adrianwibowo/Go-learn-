package iteration

import "fmt"

func ForLoop() {
	fmt.Println("Loop biasa")
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	fmt.Println("Loop sederhana")
	var i = 0

	for i < 5 {
		fmt.Println("Angka", i)
		i++
	}

	fmt.Println("Loop dengan break dan continue")
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

	fmt.Println("Loop nested")
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}

	fmt.Println("Loop dengan label")
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}

	
}
