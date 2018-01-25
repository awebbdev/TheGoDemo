package main

import (
	"fmt"
)

func main(){
	var cost float64

	cost = 11.99
	fmt.Println("The total cost is", calculateTax(cost))

	length := 12.5
	width :=  6.25
	area, perimeter := calculateAandP(length, width)
	fmt.Println("Area:", area, "Perimeter:", perimeter)
	area1, _ := calculateAandP(length, 5.5)
	fmt.Println("Area:", area1)


}


func calculateTax(cost float64) float64 {
	tax := 0.09
	return cost + (cost * tax)
}

func calculateAandP(length, width float64) (area, perimeter float64){
	area = length * width
	perimeter = (length + width) + 2
	return
}