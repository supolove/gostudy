package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 4 {
		fmt.Printf("error %d arg", len(os.Args))
		os.Exit(0)
	}

	//for key, val := range os.Args {
	//	log.Println(key, ":", val)
	//}

	//fmt.Println(os.Args[0])
	//fmt.Println(os.Args[1])
	//fmt.Println(os.Args[2])
	//fmt.Println(os.Args[3])

	//nonce := "37b8e8a308c354048d245f6d"
	//key := "AES256Key-32Characters1234567890"
	//plainText := "172.10.99.88"

	nonce := os.Args[1]
	key := os.Args[2]
	plainText := os.Args[3]
	cipherText := ExampleNewGCM_encrypt(plainText, key, nonce)
	//newPlain := ExampleNewGCM_decrypt(cipherText, key, nonce)

	//fmt.Println("cipher:", cipherText)
	//fmt.Println("new plain:", newPlain)
	fmt.Print(cipherText)
	os.Exit(0)
}

func ExampleNewGCM_encrypt(src, k, n string) string {
	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	key := []byte(k)
	plaintext := []byte(src)

	block, err := aes.NewCipher(key)
	if err != nil {
		return ""
	}

	nonce, _ := hex.DecodeString(n)

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return ""
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	return base64.StdEncoding.EncodeToString(ciphertext)
}

func ExampleNewGCM_decrypt(src, k, n string) string {
	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	key := []byte(k)
	ciphertext, _ := hex.DecodeString(src)

	nonce, _ := hex.DecodeString(n)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return string(plaintext)
}
