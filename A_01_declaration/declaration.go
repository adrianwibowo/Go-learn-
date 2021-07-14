package declaration

import "fmt"

func Declaration() {
	fmt.Println("---------------Declaration-------------")

	//komentar single line
	/*
		Multiple
	*/

	var name string = "john"
	fmt.Println("Ini Print biasa")

	//printF dimana struktur nya sudah didefinisikan diawal
	fmt.Printf("halo %s !\n", name) // harus dikasih literal newline "\n" agar ada baris baru

	//variable tanpa vtype data

	lastname := "wick"    //yang pertama dengan "inference" yaitu tanda := dan ini hanya bisa dibuat di blok fungsi
	var newName = "wicky" //yang kedua tanpa type
	lastname = "re-assign value"
	fmt.Println(lastname, newName)

	// declare multi variable

	var satu, dua, tiga int
	satu, dua, tiga = 1, 2, 3 // bisa langsung assign ketidanya sekaligus
	println(satu, dua, tiga)

	empat, lima, enam := 4, 5.0, "enam" //for more handy declaration
	println(empat, lima, enam)

	//variable _ underscore
	_ = "tempat menyimpan value yang tidak berguna"
	/*
		Underscore (_) adalah reserved variable yang bisa dimanfaatkan untuk menampung nilai yang tidak dipakai.
		Bisa dibilang variabel ini merupakan keranjang sampah.

		Biasanya variabel underscore sering dimanfaatkan untuk menampung nilai balik fungsi yang tidak digunakan.

		Perlu diketahui, bahwa isi variabel underscore tidak dapat ditampilkan.
		Data yang sudah masuk variabel tersebut akan hilang.
		Ibaratkan variabel underscore seperti blackhole, objek apapun yang masuk kedalamnya,
		akan terjebak selamanya di-dalam singularity dan tidak akan bisa keluar
	*/




	//deklarasi variable menggunakan new
	menggunakanNew := new(string) //Perlu diketahui, deklarasi menggunakan := hanya bisa dilakukan di dalam blok fungsi.

	fmt.Println(menggunakanNew, "<-- total kapasitas") // ini akan menghasilkan kapasitas var yg dipakai
	fmt.Println(*menggunakanNew, "<-- kosong")         // harus pakai * "dereference" jika mau ngeprint kosongan value

	// Variable constan adalah variable yang tidak akan berubah
	const phi float32 = 3.14
	fmt.Print(phi + 2) // method print akan memprint arguen didalamnya namun tidak akan mebuat baris baru

	fmt.Println("\n" + `---------------------------------------
	
	`)

}
