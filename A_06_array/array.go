package array

import "fmt"

func Array() {
	var names [4]string //4 disini adalah length
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"
	// names[4] akan eror karena tidak sesuai length

	fmt.Println(names[0], names[1], names[2], names[3])

	// inisiasi diawal
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	fmt.Println("Jumlah element \t\t", len(fruits)) //len untuk memunculkan panjang array
	fmt.Println("Isi semua element \t", fruits)


	// deklarasi array tanpa length
	var numbers = [...]int{2, 3, 2, 4, 3}

	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	//array multi dimensi
	var numbers1 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	//Iterasi dalam array
	var buah = [4]string{"apple", "grape", "banana", "melon"}

	for i := 0; i < len(buah); i++ {
		fmt.Printf("elemen %d : %s\n", i, buah[i])
	}

	//Iterasi array dengan for-range
	/*
		Array fruits diambil elemen-nya secara berurutan. Nilai tiap elemen ditampung variabel oleh fruit (tanpa huruf s),
		sedangkan indeks nya ditampung variabel i.
	*/
	var buah2 = [4]string{"apple", "grape", "banana", "melon"}

	for i, fruit := range buah2 {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}

	//menggunakan _
	var fruits2 = [4]string{"apple", "grape", "banana", "melon"}

	for _, fruit := range fruits2 {
		fmt.Printf("nama buah : %s\n", fruit)
	}

	//Jika yang dibutuhkan hanya indeks elemen-nya saja, bisa gunakan 1 buah variabel setelah keyword for.
	// for i, _ := range fruits {
	// }
	// atau
	// for i := range fruits {
	// }

}
