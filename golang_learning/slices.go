package main

import "fmt"

var sli = [5]string{"Sasuke", "Sakrua", "Shino", "Naruto", "Gara"}

func main() {
	//make slice

	fmt.Println(sli)
	fmt.Println(sli[2])
	changename()
	fmt.Println(sli)

}

func changename() {
	fmt.Println("> Will replace with ino")
	sli[2] = "Ino"
}
