package main

import (
	"fmt" 
	"basic-golang/A_01_declaration"
	"basic-golang/A_02_data-type"
	"basic-golang/A_03_operator"
	"basic-golang/A_04_conditional"
	"basic-golang/A_05_iteration"
	"basic-golang/A_06_array"
	"basic-golang/A_07_map"
)


func main() {
	
	declaration.Declaration()

	//TYPE DATA
	fmt.Println("----------------Data Type--------------")
	data_type.Number()
	data_type.Float()
	data_type.Boolean()
	data_type.String()
	fmt.Println("\n" +`---------------------------------------
	
	`)

	//OPERATOR
	fmt.Println("----------------Operator--------------")
	operator.Operator()
	fmt.Println("\n" +`---------------------------------------
	
	`)
	

	//CONDITIONAL
	fmt.Println("----------------Conditional--------------")
	conditional.Conditional()
	conditional.SwitchCase()
	fmt.Println("\n" +`---------------------------------------
	
	`)

	//iteration
	fmt.Println("----------------Iteration--------------")
	iteration.ForLoop()
	fmt.Println("\n" +`---------------------------------------
	
	`)

	fmt.Println("----------------Array--------------")
	array.Array()
	fmt.Println("\n" +`---------------------------------------
	
	`)


	fmt.Println("----------------HasMap--------------")
	hasmap.HasMap()
	fmt.Println("\n" +`---------------------------------------
	
	`)


}


