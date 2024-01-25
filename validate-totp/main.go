package main

import (
	"bufio"
	"fmt"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"os"
)

func main() {
	valid, _ := totp.ValidateCustom(passcode, key.Secret(), timeNow, totp.ValidateOpts{
		Period:    120,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA256,
	})
	if valid {
		println("Valid passcode!")
		os.Exit(0)
	} else {
		println("Invalid passcode!")
		os.Exit(1)
	}
}

func promptForPasscode() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Passcode: ")
	text, _ := reader.ReadString('\n')
	return text
}
