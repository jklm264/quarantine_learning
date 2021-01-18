package main

import "fmt"

var mymap = make(map[string]*Bumper, 10)

func main() {
	fmt.Println("We are going to be working with maps now.")
	filler := fillbump()
	mymap["sister"] = &filler // src: https://stackoverflow.com/questions/42716852/how-to-update-map-values-in-go
	fmt.Println(mymap)
	var six int = 6
	mymap["sister"].fivers = six // src: https://suraj.pro/post/golang_workaround/
	fmt.Println(mymap["sister"].fivers)
}

type Bumper struct {
	fivers int
	rex    string
}

func fillbump() Bumper {
	return Bumper{
		fivers: 5,
		rex:    "Rex",
	}
}
