package main

import (
	"fmt"
)

func main() {

	var str string
	fmt.Println("Input str:")
	fmt.Scanf("%s", &str)
	length := len(str)
	for i := 0; i <= length/2; i++ {
		if str[i] == str[length-i-1] {
			fmt.Println(str)
			break
		} else {
			fmt.Println(false)
			break
		}
	}

}
