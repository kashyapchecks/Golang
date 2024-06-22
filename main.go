package main

import "fmt"

type Bill struct {
	Name string
	Price float64
	Quantity string
}




func main() {
	var a Bill =  Bill{Name: "Dosa" , Price:140, Quantity: "Two"}
	fmt.Println(a.Price)
}