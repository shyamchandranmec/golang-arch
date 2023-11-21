package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/Azure/azure-amqp-common-go/v3/uuid"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type person struct {
	First string
}

var key = [64]byte{}
var keys = map[string]Key{}
var currentKid = ""

type UserClaims struct {
	jwt.RegisteredClaims
	SessionId int
}

func (u *UserClaims) Valid() error {
	if !u.VerifyExpiresAt(time.Now(), true) {
		return fmt.Errorf("token has expired")
	}
	if u.SessionId == 0 {
		return fmt.Errorf("invalid sessin id")
	}
	return nil
}

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
	msg := "Shyam Message"
	smsg, _ := signMessage([]byte(msg))
	fmt.Println("Signed message is ", base64.StdEncoding.EncodeToString(smsg))
	same, _ := checkSignature([]byte(msg), smsg)
	fmt.Println("Check signature is ", same)
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
	h := hmac.New(sha512.New, keys[currentKid].key)
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

func CreateToken(c *UserClaims) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, c)
	signedToken, err := t.SignedString(keys[currentKid].key)
	if err != nil {
		return "", fmt.Errorf("error in create token when signing token")
	}
	return signedToken, nil
}

type Key struct {
	key     []byte
	created time.Time
}

func generateNewKey() error {
	newKey := make([]byte, 64)
	_, err := io.ReadFull(rand.Reader, newKey)
	if err != nil {
		return fmt.Errorf("error in generating new key %w", err)
	}
	uid, err := uuid.NewV4()
	if err != nil {
		return fmt.Errorf("error in generating uuid %w", err)
	}
	keys[uid.String()] = Key{
		key:     newKey,
		created: time.Now(),
	}
	currentKid = uid.String()
	return nil
}

func parseToken(token string) (*UserClaims, error) {
	t, err := jwt.ParseWithClaims(token, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS512.Alg() {
			return nil, fmt.Errorf("invalid signing algorith")
		}
		kid, ok := t.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("Invalid key id")
		}
		k, ok := keys[kid]

		if !ok {
			return nil, fmt.Errorf("Invalid key id")
		}
		return k, nil
	})
	if err != nil {
		return nil, fmt.Errorf("error in parsing token %w", err)
	}
	if !t.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	claims := t.Claims.(*UserClaims)
	return claims, nil
}
