package main

import "fmt"

/*
The VigenereEncrypt function takes two strings as arguments: PlainText (the original text to be encrypted) and key (the key to be encrypted).
In a loop over the symbolic source text `plainText`, each character offset has a value obtained from the key `key`.
The shift is carried out by entering character values ​​with the position of the key characters in the alphabet.
If the character source text is a letter in preserved case, then it is encrypted using an offset to remain within the Latin alphabet (a-z).
*/
func VigenereEncrypt(plainText string, key string) string {
	cipherText := ""
	keyIndex := 0

	for i := 0; i < len(plainText); i++ {
		char := plainText[i]
		shift := getShift(key[keyIndex])

		if char >= 'a' && char <= 'z' {
			cipherChar := 'a' + (int(char-'a')+shift)%26
			cipherText += string(cipherChar)
		} else {
			cipherText += string(char)
		}

		keyIndex = (keyIndex + 1) % len(key)
	}

	return cipherText
}

func getShift(char byte) int {
	return int(char - '0')
}

func main() {
	plainText := "momotaro"
	key := "2531"
	/*
			key
		1 letter shift 2
		2 letter shift 5
		3 letter shift 3
		4 letter shift 1
	*/
	cipherText := VigenereEncrypt(plainText, key)
	fmt.Println(cipherText)
}
