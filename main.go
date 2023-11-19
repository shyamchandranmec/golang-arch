package main

import (
	"encoding/base64"
	"fmt"
)

type person struct {
	First string
}

func main() {

	unp := "username:password"
	encValue := base64.StdEncoding.EncodeToString([]byte(unp))
	fmt.Println("Base 64 encode value is ", encValue)

}
