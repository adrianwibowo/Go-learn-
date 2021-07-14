package conditional

import "fmt"

func Conditional() {
	var point = 8

	fmt.Println(`var point = 8`)
	fmt.Println(`
Regular If Else
	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	} 
HASILNYA:`)

	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}

	
	
	
	
	var point2 = 8840.0
	fmt.Println(`var point2 = 8840.0`)

	fmt.Println(`
	
If Else dengan operasi dan variable didalam kondisinya:
	if percent := point2 / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}
HASILNYA:`)
	if percent := point2 / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}
}
