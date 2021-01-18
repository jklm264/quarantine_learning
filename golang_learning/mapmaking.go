package main

import "fmt"

//src: https://yourbasic.org/golang/maps-explained/
func main(){
	fmt.Println("We are going to be working with maps now.")
	// map[KeyType]ValueType
	//var m map[string]int //string-int pairs
	map_float := make(map[string]float64, 100) // Prealloccates room for 100 entries
	map_literal := map[string]float64{
		"lucky": 613.69,
		"hand": 5,
		"pi": 3.1416,
	}
	fmt.Println(len(map_float))
	fmt.Println(len(map_literal))

	for key, value := range map_literal {
    		fmt.Println("Key:", key, "Value:", value)
	}

	delete(map_literal,"hand")
	fmt.Println(len(map_literal))
	fmt.Print(map_literal["lucky"])
	fmt.Println("------------------------------------")
	
	var account_db = make(map[string]Account, 10)
	account_db["hiya"] = makeme()
	fmt.Println(account_db["hiya"].username)
	fmt.Println(account_db["hiya"].bb.fivers + 5)

}

type Account struct {
	username      string
	accountnumber int32
	bb 			  Bumper
}
type Bumper struct {
	fivers int
}

func makeme() Account {
	return Account{
		username: "james",
		accountnumber: 4443,
		bb: cbump(),
	}
}

func cbump() Bumper {
	return Bumper{
		fivers: 5,
	}
}