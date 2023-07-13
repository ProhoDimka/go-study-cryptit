package main

import (
	"fmt"
	"github.com/ProhoDimka/go-study-cryptit/decrypt"
	"github.com/ProhoDimka/go-study-cryptit/encrypt"
)

func main() {
	str := "dimka.pro"
	encryptedStr := encrypt.Nimbud(str)
	fmt.Println("Source value:", str)
	fmt.Println("Encrypted value:", encryptedStr)
	decryptedStr := decrypt.Nimbud(encryptedStr)
	fmt.Println("Decrypted value:", decryptedStr)
}
