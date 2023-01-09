package src

import "fmt"

func PrintSegitiga(number int) {

	for i := 0; i < number; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
			
		}
		for k := 0; k < 2*(number-i)-1; k++ {
			fmt.Print(k+1)
		}
		fmt.Println()
	}

}