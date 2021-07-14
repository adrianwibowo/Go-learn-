package data_type

import "fmt"

func Number() {

	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	//Template %d pada fmt.Printf() digunakan untuk memformat data numerik non-desimal.
}
