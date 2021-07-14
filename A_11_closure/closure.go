package main

import "fmt"

func main() {
    var getMinMax = func(n []int) (int, int) {
        var min, max int
        for i, e := range n {
            switch {
            case i == 0:
                max, min = e, e
            case e > max:
                max = e
            case e < min:
                min = e
            }
        }
        return min, max
    }

    var numbers = []int{2, 3, 4, 3, 4, 2, 3}
    var min, max = getMinMax(numbers)
    fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max) //template %v untuk menampilkan value smeua jenis type data


	// Imediately-invoked function expression
	var nomors = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

    var newNumbers = func(min int) []int {
        var r []int
        for _, e := range nomors {
            if e < min {
                continue
            }
            r = append(r, e)
        }
        return r
    }(3) // parameter langsung diisi disini stelah deklarasi

    fmt.Println("original number :", nomors)
    fmt.Println("filtered number :", newNumbers)

	var howMany, getNumbers = findMax(numbers, max)
	var theNumbers = getNumbers()

    fmt.Println("numbers\t:", numbers)
    fmt.Printf("find \t: %d\n\n", max)

    fmt.Println("found \t:", howMany)    // 9
    fmt.Println("value \t:", theNumbers) 
	
}

func findMax(numbers []int, max int) (int, func() []int) {
    var res []int
    for _, e := range numbers {
        if e <= max {
            res = append(res, e)
        }
    }
	// return function anonymous tapi juga bisa di beri label
    return len(res), func() []int { 
        return res
    }
}