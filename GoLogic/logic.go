package main

import "fmt"

func main(){

	//ifelse
	num := 15
	if num % 2 == 0{
		fmt.Println("Even")
	}else{
		fmt.Println("odd")
	}

	if num < 10 {
		fmt.Println("Less then ten.")
	}else if num >= 10 && num <= 20{
		fmt.Println("It's between ten and twenty")
	}else{
		fmt.Println("Greater than 20")
	}

	//Loops
	for i := 0; i <= 10; i++{
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	for i := 0; i <= 10; i++{
		if i > 5 {
			break
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	for i := 0; i <= 10; i++{
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	x := 0
	for x <= 10{
		fmt.Printf("%d ", x)
		x += 3
	}
	fmt.Println()

	//Switches
	choice := 2
	switch choice {
	case 1:
		fmt.Println("Sandwich")
	case 2:
		fmt.Println("Salad")
	case 3:
		fmt.Println("Pizza")
	case 4:
		fmt.Println("Hamburger")
	default:
		fmt.Println("You will eat what I serve you.")
	}

	grade := "B"
	switch grade {
	case "A", "B", "C":
		fmt.Println("PASS")
	case "D":
		fmt.Println("NEEDS IMPROVEMENT")
	case "F":
		fmt.Println("FAIL")
	}

	num = 5
	switch {
	case num < 5 :
		fmt.Printf("%d is less than Five\n", num)
		fallthrough
	case num < 10:
		fmt.Printf("%d is less than Ten\n", num)
		fallthrough
	case num < 20 :
		fmt.Printf("%d is less than Twenty\n", num)
	}
}
