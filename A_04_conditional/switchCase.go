package conditional

import "fmt"

func SwitchCase() {

	fmt.Println(`
Regular switch Case
var point = 6

	switch point {
	case 8:
    	fmt.Println("perfect")
	case 7, 6, 5, 4:
    	fmt.Println("awesome")
	default:
    	fmt.Println("not bad")
}
HASILNYA:`)
	var point = 6

	switch point {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	fmt.Println(`
Switch Case gaya If else
var point = 6

	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}
HASILNYA:`)
	var point2 = 6

	switch {
	case point2 == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}
fmt.Println("ada namanya 'fallthrough' untuk tetap melanjutkan case sleanjutkan meskipun case tersebut sudah terpenuhi")
}
