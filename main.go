package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
}

func main() {
	p1 := person{
		First: "Jenny",
	}
	p2 := person{
		First: "James",
	}

	xp := []person{p1, p2}
	bs, err := json.Marshal(xp)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Pring JSON", string(bs))

	xp2 := []person{}
	err = json.Unmarshal(bs, &xp2)

	if err != nil {
		log.Panic("Error on Unmarshall ", err)
	}
	fmt.Println("Back to go data structure ", xp2)
}
