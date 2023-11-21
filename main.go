package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type person struct {
	First string
}

var key = [64]byte{}

func main() {
	for i := 0; i < 64; i++ {
		key[i] = byte(i)
	}
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
		return fmt.Errorf("passwords do not match %w", err)
	}
	return nil
}

func signMessage(msg []byte) ([]byte, error) {
	h := hmac.New(sha512.New, key[:])
	_, err := h.Write(msg)
	if err != nil {
		return nil, fmt.Errorf("error while signing message in sign Message %w", err)
	}
	signature := h.Sum(nil)
	return signature, nil
}

func checkSignature(msg, sig []byte) (bool, error) {
	newSig, err := signMessage(msg)
	if err != nil {
		return false, fmt.Errorf("error in signing message in check Signature %w", err)
	}
	same := hmac.Equal(newSig, sig)
	return same, nil
}
