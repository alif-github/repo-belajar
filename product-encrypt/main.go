package main

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	arg := os.Args
	if len(arg) < 1 || len(arg) <= 2 {
		errs := errors.New("error no arguments here or error argument must 2")
		log.Println("error ---> ", errs)
		os.Exit(1)
	}

	cid := arg[0]
	csc := arg[1]

	encryptKey := HashingPassword(cid, csc)
	fmt.Println("Encrypt Key ---> ", encryptKey)
}

func HashingPassword(password string, salt string) string {
	return CheckSumWithSha512([]byte(password + salt))
}

func CheckSumWithSha512(content []byte) string {
	result := sha512.Sum512(content)
	return hex.EncodeToString(result[:])
}
