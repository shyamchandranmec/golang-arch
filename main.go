package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.HandleFunc("/encode2", foo2)
	http.HandleFunc("/decode2", bar2)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "Jenny",
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println("Got an error encoding p1 ", err)
	}

}

/*
curl -XPOST -H "Content-type: application/json" -d '{"First":"Shyam"}' 'localhost:8080/decode'
*/

func bar(w http.ResponseWriter, r *http.Request) {
	var p1 person
	err := json.NewDecoder(r.Body).Decode(&p1)
	if err != nil {
		log.Println("Got an error Decoding to  p1 ", err)
	}
	log.Println("P1 is ", p1)
}

func foo2(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "Jenny",
	}
	p2 := person{
		First: "James",
	}

	xp := []person{p1, p2}
	if err := json.NewEncoder(w).Encode(xp); err != nil {
		log.Println("Unable to encode xp ", err)
	}
}

/*
curl -XPOST -H "Content-type: application/json" -d '[{"First":"Jenny"},{"First":"James"}]' 'localhost:8080/decode2'
*/
func bar2(w http.ResponseWriter, r *http.Request) {
	xp := []person{}
	if err := json.NewDecoder(r.Body).Decode(&xp); err != nil {
		log.Println("Unable to decode body to xp ", err)
	}
	log.Println("Decoded body is ", xp)
}
