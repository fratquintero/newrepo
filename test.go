package main

import "fmt"

/* Declarations */

var num int
var name string
var set = 22
const BUILDING int32 = 56

func main() {
	name = "Frat Quintero"
	num = 222
	fmt.Println("Hello World", num, name)
	fmt.Println("Variable set:", set)
	fmt.Println("La constante es:", BUILDING)
}
