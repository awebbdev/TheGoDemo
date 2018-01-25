package main

import (
	"github.com/thebigbadox/TheGoDemo/GoPackages/GoPower"
	"fmt"
)

func main(){
	var watts, kg = 285.5, 83.5

	fmt.Println("FTP is",GoPower.CalcFTP(watts))
	fmt.Println("Watts per kilo is", GoPower.WattsperKG(watts, kg))
}
