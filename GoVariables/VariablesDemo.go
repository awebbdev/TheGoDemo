package main

import (
	"fmt"
	"unsafe"
)

func main(){
	var age int //variable declaration
	var initage int = 35 // initialized variable
	var inferedage = 30 //variable type inference
	shorthandage := 62 //shorthand variable declaration

	//age = "56"  //Strict typing cannot change from int to string
	//inferedage = false

	age = 15  //variable assignment
	fmt.Println("My age is", age)
	fmt.Println("Their age is", initage)
	fmt.Println("Infered age is", inferedage)
	fmt.Println("Shorthand age is: ", shorthandage)

	var watts, kg int = 350, 85 //multiple variable declaration

	fmt.Println("Watts is", watts, "kilograms is", kg)

	var (
		statement = "My age is"
		age1 = 40
	)
	fmt.Println(statement, age1)

	x, y := 10, 20
	fmt.Println("x is", x, " y is", y)

	//x, y := 30, 40 //error no new variables

	//pointers
	a, b := 42, 2701

	p := &a         // point to a
	fmt.Println(*p) // read a through the pointer
	*p = 21         // set a through the pointer
	fmt.Println(a)  // see the new value of a

	p = &b         // point to b
	*p = *p / 37   // divide b through the pointer
	fmt.Println(b) // see the new value of b

	//Numeric Types
	//int, int32, uint32, float32, complex64
	var alpha int16 = -30000
	var beta int = 100

	fmt.Printf("Alpha's value: %d, Type; %T, Size: %d\n", alpha, alpha, unsafe.Sizeof(alpha))
	fmt.Printf("Beta's value: %d, Type; %T, Size: %d\n", beta, beta, unsafe.Sizeof(beta))

	//Type Conversion

	i := 15.45  //float64
	j := 10 //int
	//sum := i + j //error, mismatched types
	sum := i + float64(j)
	fmt.Println("The sum is", sum)
	intsum := int(i) + j
	fmt.Println("The intSum is", intsum)

	//Constants
	const gamma = 90
	//gamma = 100 cannot reassign
	//Known at compile
	

}
