package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {

	//versionName := "9.9.9"
	//versionCode := 0
	//
	//fmt.Println(math.Pow10(3))
	//
	//vns := strings.Split(versionName, ".")
	//for i, vn := range vns {
	//	v, _ := com.StrTo(vn).Int()
	//	x := int(math.Pow10((len(vns) - i - 1) * 2))
	//	versionCode += v * x
	//	//com.ex
	//}
	//fmt.Println(versionCode)
	//a := "123456.aar"
	//fmt.Print(a[:len(a)-4])

	//fmt.Println(base64.URLEncoding.EncodeToString([]byte("123123131321231")))

	//plaintext := base64.StdEncoding.EncodeToString([]byte("123"))

	plaintext := []byte("aabbccddeeff")

	key := []byte("b573f0b036f194a88f8151bdad14c02f")
	//if _, err := io.ReadFull(rand.Reader, key); err != nil {
	//	panic(err.Error())
	//}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	nonce, _ := hex.DecodeString("b51074cdfdb2fd1a")
	//if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
	//	panic(err.Error())
	//}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, []byte(plaintext), []byte{})

	fmt.Println(base64.StdEncoding.EncodeToString(ciphertext))
	//fmt.Printf("Key:\t\t%x", key)
	//fmt.Printf("\nCiphertext:\t%x", ciphertext)

}
