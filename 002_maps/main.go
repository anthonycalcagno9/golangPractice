package main

import "fmt"

func main() {
	fmt.Println("hello world")

	am := map[string]int{
		"Jim":   13,
		"Jam":   15,
		"jambo": 17,
	}

	an := make(map[string]int)
	an["tacocat"] = 30

	//get length of map
	length := len(am)
	fmt.Printf("length of map = %v\n", length)

	fmt.Printf("accessing map an[tacocat] = %v\n", an["tacocat"])
	fmt.Printf("Print whole map = %v\n", am)

	for key, value := range am {
		fmt.Printf("looping through map key = %v\n", key)
		fmt.Printf("looping through map value = %v\n", value)
	}

	//Delete element from map
	delete(am, "Jim")
	fmt.Println(am)

	//comma ok idiom to check if map has a value
	value, ok := am["Jam"]
	if ok {
		fmt.Printf("key did exist, key = %v\n", value)
	} else {
		fmt.Println("key didn't exist")
	}

	//turning above code from comma ok -> comma ok + statement;statement idiom
	if value1, ok := am["Jarbo"]; ok {
		fmt.Printf("key did exist, key = %v\n", value1)
	} else {
		fmt.Println("key didn't exist")
	}

	//PROBLEM STATEMENT: COUNT EVERY WORD IN A BOOK ********************* Seems like this could be important in leet code type questions
	//Loop through the text and use this code
	bookWordsMap := make(map[string]int)
	wordsInBook := []string{"taco", "burrito", "the", "dog", "cat", "taco"}

	for _, value := range wordsInBook {
		bookWordsMap[value]++
	}

	fmt.Printf("map after counting words in book = %v\n", bookWordsMap)

}
