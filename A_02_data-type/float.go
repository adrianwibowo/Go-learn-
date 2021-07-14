package data_type

import "fmt"

func Float() {
	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber) 
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	/*
	Template %f digunakan untuk memformat data numerik desimal menjadi string. 
	Digit desimal yang akan dihasilkan adalah 6 digit. Pada contoh di atas, hasil format variabel decimalNumber adalah 
	2.620000.
	Jumlah digit yang muncul bisa dikontrol menggunakan %.nf, tinggal ganti n dengan angka yang diinginkan. Contoh: %.3f 
	maka akan menghasilkan 3 digit desimal, %.10f maka akan menghasilkan 10 digit desimal.
	*/
}
