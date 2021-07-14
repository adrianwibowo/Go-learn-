package hasmap

import "fmt"

func HasMap() {
	var chicken = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}

	for key, val := range chicken {
		fmt.Println(key, "  \t:", val) //templat \t untuk indentasi yang sama pada ":"
	}

	fmt.Println(len(chicken)) // 2
	fmt.Println(chicken)

	delete(chicken, "januari") // delete map

	fmt.Println(len(chicken)) // 1
	fmt.Println(chicken)

	var value, isExist = chicken["februari"] // retuyrn value ke-2 adalah bool yang pertama adalah value

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}



	// kombinasi map dan slice
	var ayam = []map[string]string{
		{"name": "chicken blue",   "gender": "male"},
		{"name": "chicken red",    "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}
	
	for _, chicken := range ayam {
		fmt.Println(chicken["gender"], chicken["name"])
	}
}
