package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("sample-file.txt")
	if err != nil {
		errx := fmt.Errorf("error in reading file %w", err)
		log.Fatal(errx)
	}
	defer file.Close()
	data := make([]byte, 100)
	count, err := file.Read(data)
	fmt.Printf("Read %d bytes:%q", count, data[:count])
	fmt.Println()
	h := sha256.New()
	h.Write(data)
	fmt.Printf("Here is the type Before sum %T", h)
	fmt.Println()
	xb := h.Sum(nil)
	fmt.Printf("Here is the type after sum %T", xb)
	fmt.Println()
	fmt.Printf("xb is %x", xb)
	fmt.Println()

	xb = h.Sum(nil)
	fmt.Printf("2 Here is the type after sum %T", xb)
	fmt.Println()
	fmt.Printf("2 xb is %x", xb)
	fmt.Println()

	xb = h.Sum(xb)
	fmt.Printf("3 Here is the type after sum %T", xb)
	fmt.Println()
	fmt.Printf("3 xb is %x", xb)
	fmt.Println()
}
