package main

import (
	"encoding/base64"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type person struct {
	First string
}

func main() {

	unp := "username:password"
	encValue := base64.StdEncoding.EncodeToString([]byte(unp))
	fmt.Println("Base 64 encode value is ", encValue)
	password := "1234455"
	bh, err := hashPassword(password)
	if err != nil {
		panic(err)
	}
	fmt.Println("hash password is ", string(bh))
	if err = comparePassword(password, bh); err != nil {
		log.Fatal("Not logged in ")
	}

	log.Println("Logged in")
}

func hashPassword(password string) ([]byte, error) {
	bh, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error while generation bcrypt password %w", err)
	}
	return bh, nil
}

func comparePassword(password string, hp []byte) error {
	if err := bcrypt.CompareHashAndPassword(hp, []byte(password)); err != nil {
		return fmt.Errorf("Passwords do not match %w", err)
	}
	return nil
}
