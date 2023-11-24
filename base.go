// package main

// import (
// 	"bytes"
// 	"crypto/aes"
// 	"crypto/cipher"
// 	"encoding/base64"
// 	"fmt"
// 

// 	"golang.org/x/crypto/bcrypt"
// )

// func main() {
// 	msg := "This is totally fun. Thanks for helping me to learn this"
// 	key, err := getByteSlice("ilovedogs", 16)
// 	if err != nil {
// 		fmt.Errorf("errorn in creating key %w", err)
// 	}
// 	fmt.Println("Key is ", base64.StdEncoding.EncodeToString(key))
// 	result, iv, err := enDecode(key, nil, msg)
// 	if err != nil {
// 		fmt.Errorf("errorn in endecode msg %w", err)
// 	}
// 	fmt.Println("Before b64 Result is ", string(result))

// 	eResult := base64.StdEncoding.EncodeToString(result)
// 	fmt.Println("After b64 eResult is ", string(eResult))
// 	fmt.Println()
// 	fmt.Println()

// 	dr, err := base64.StdEncoding.DecodeString(eResult)
// 	if err != nil {
// 		fmt.Errorf("errorn in decode dr %w", err)
// 	}

// 	dResult, _, err := enDecode(key, iv, string(dr))
// 	if err != nil {
// 		fmt.Errorf("errorn in endecode  r2 %w", err)
// 	}
// 	fmt.Println("Decoded result is ", string(dResult))

// }

// func getByteSlice(s string, n int) ([]byte, error) {
// 	password := s
// 	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
// 	if err != nil {
// 		return nil, fmt.Errorf("errorn in generating password %w", err)
// 	}
// 	key := b[:n]
// 	return key, nil
// }

// func enDecode(key []byte, iv []byte, input string) ([]byte, []byte, error) {
// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 		return nil, nil, fmt.Errorf("Unable to create new cipher %w", err)
// 	}
// 	if iv == nil {
// 		iv, err = getByteSlice("ivsaltstring", 16)
// 	}

// 	//iv := make([]byte, getByteSlice("ivsaltstring", 16))
// 	fmt.Println("IV is ", base64.StdEncoding.EncodeToString(iv))
// 	if err != nil {
// 		return nil, nil, fmt.Errorf("Unable to create salt %w", err)
// 	}

// 	s := cipher.NewCTR(block, iv)
// 	buff := &bytes.Buffer{}
// 	sw := cipher.StreamWriter{
// 		S: s,
// 		W: buff,
// 	}
// 	_, err = sw.Write([]byte(input))
// 	if err != nil {
// 		return nil, nil, fmt.Errorf("Unable to SW write input %w", err)
// 	}
// 	output := buff.Bytes()
// 	return output, iv, nil

// }

// func encrypt(key []byte, msg string)
