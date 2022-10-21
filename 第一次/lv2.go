package main

import (
	"fmt"
)

func main() {
	var a, b float64
	var cmd byte
	fmt.Println("Input a,cmd,b:")
	fmt.Scanf("%f%c%f", &a, &cmd, &b)
	switch cmd {
	case '+':
		{
			fmt.Println("%f+%f=%f", a, b, a+b)
			break
		}
	case '-':
		{
			fmt.Println("%f-%f=%f", a, b, a-b)
			break
		}
	case '*':
		{
			fmt.Println("%f*%f=%f", a, b, a*b)
			break
		}
	case '/':
		{
			if 0 == b {
				fmt.Println("Division by zero!")
			} else {
				fmt.Println("%f/%f=%f", a, b, a/b)
			}
			break
		}
	default:
		fmt.Println("Invalid operator!")
	}
}
