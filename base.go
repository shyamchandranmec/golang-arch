package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	msg := "This is totally fun. /Thanks for helping me to learn this"
	encoded := encode(msg)
	fmt.Println(encoded)

	// x := make([]byte, base64.URLEncoding.EncodedLen(len(msg)))
	// base64.StdEncoding.Encode(x, []byte(msg))
	// fmt.Print(string(x))
	decoded, err := decode(encoded)
	if err != nil {
		fmt.Println("Unable to decode ", encoded)
	} else {
		fmt.Println("Decode string is ", decoded)
	}

}

func encode(msg string) string {
	return base64.URLEncoding.EncodeToString([]byte(msg))
}

func decode(s string) (string, error) {
	x, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return "", fmt.Errorf("Unable to decode base 64 string %w", err)
	}
	return string(x), nil
}
