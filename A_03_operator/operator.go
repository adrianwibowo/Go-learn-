package operator

import "fmt"

func Operator()  {
	
	fmt.Println( "1 == 2:",1 == 2)
	fmt.Println( "1 != 2:",1 != 2)
	fmt.Println( "1 < 2:",1 < 2)
	fmt.Println( "1 > 2:",1 > 2)
	fmt.Println( "1 <= 2:",1 <= 2)
	fmt.Printf( "1 >= 2: %t \n",1 >= 2)

	benar := true 
	salah := false
	fmt.Println("true || false:", benar || salah)
	fmt.Println("true && false:", benar && salah)
	fmt.Println("!true:", !benar)
}