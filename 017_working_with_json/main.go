package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {

	p1 := person{
		First: "anthony",
		Last:  "calcagno",
		Age:   29,
	}

	p2 := person{
		First: "rafa",
		Last:  "santos",
		Age:   28,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	//Marshal takes the struct and returns the JSON encoding of it
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println("error marshaling: ", err)
	}

	fmt.Println(string(bs))

	//JSON we got from above
	//[{"First":"anthony","Last":"calcagno","Age":29},{"First":"rafa","Last":"santos","Age":28}]

	s := `[{"First":"anthony","Last":"calcagno","Age":29},{"First":"rafa","Last":"santos","Age":28}]`
	bytes := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bytes)

	var peoples []person

	err = json.Unmarshal(bytes, &peoples)
	if err != nil {
		fmt.Println("error unmarshaling = ", err)
	}

	fmt.Println(peoples)

	for i, v := range peoples {
		fmt.Println(i)
		fmt.Println(v.First)
		fmt.Println(v.Last)
	}
}
