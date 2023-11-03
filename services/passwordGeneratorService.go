package services

import (
	"fmt"
	"math/rand"
)

func PasswordGenerator(length int) {

	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"

	password := make([]byte, length)

	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(charset))

		password[i] = charset[randomIndex]

	}

	fmt.Println(string(password))

}
