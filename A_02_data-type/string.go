package data_type

import "fmt"

func String()  {
	//bisa pake double quotes "" atau backtic ``
	var message = `Nama saya "John Wick".
Salam kenal.
Mari belajar "Golang".` //tidak ada yang escape atau semua didalam backtic dianggap string

	fmt.Println(message)

}