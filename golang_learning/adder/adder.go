package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// fmt.Printf("\t%T\n", os.Args[1])
	one, _ := strconv.Atoi(os.Args[1])
	two, _ := strconv.Atoi(os.Args[2])
	//fmt.Printf("%T\t%T\n", one,os.Args[2])
	//fmt.Println(fmt.Sprintf("One: %v \tTwo: %v", os.Args[1], os.Args[2]))
	//fmt.Println(fmt.Sprintf("One: %v \tTwo: %v", one, two))
	//fmt.Println(fmt.Sprintf("Will be adding %v and %v", one, two))
	// fmt.Println(addme(one, two))
	fmt.Printf("%v + %v = %v", one, two, addme(one, two))
}

func addme(a int, b int) int {
	return a + b
}
