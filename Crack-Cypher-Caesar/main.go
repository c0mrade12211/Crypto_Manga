package main

import (
	"fmt"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func CrackCypher(plain_text string) {
	for shift := 1; shift <= 25; shift++ {
		cipher_text := ""
		for i := 0; i < len(plain_text); i++ {
			if plain_text[i] == 32 {
				cipher_text += " "
				continue
			}
			index := (strings.Index(alphabet, string(plain_text[i])) + shift) % len(alphabet)
			cipher_text += string(alphabet[index])
		}
		fmt.Printf("Shift %d: %s\n", shift, cipher_text)
	}
}

func main() {
	var plain_text string

	fmt.Println("[+] Input Cypher Text")
	fmt.Scan(&plain_text)
	CrackCypher(plain_text)
}
