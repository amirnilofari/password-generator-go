package main

import (
	"fmt"
	"math/rand"
	"time"
)

// password = string + number + digit

//A aB bC cD dE eF fG gH hI iJ jK kL lM mN nO oP pQ qR rS sT tU uV vW wX xY yZ z

var letter = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$0123456789")

func main() {
	fmt.Println("Your password: ", randGeneratorPassword(16))
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randGeneratorPassword(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
