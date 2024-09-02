package main

import "fmt"

func main(){

	for i:= 1; i <= 100; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("pin pan")
		}else if i % 5 == 0 {
			fmt.Println("Pan")
		}else if i % 3 == 0{
			fmt.Println("pin")
		}else{
			fmt.Println(i)

		}
		
	}
	
}
//  contar de um a cem 
// verificar os divisives por 3 e 5 