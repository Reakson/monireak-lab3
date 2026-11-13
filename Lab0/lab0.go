package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"golang.org/x/crypto/sha3"
	"fmt"
)

// Function for MD5 hashing
func hashMD5(input string) string {

	data := []byte(input)
	hash := md5.Sum(data)

	return hex.EncodeToString(hash[:])
}

// Function for SHA1 hashing
func hashSHA1(input string) string {

	data := []byte(input)
	hash := sha1.Sum(data)

	return hex.EncodeToString(hash[:])
}

// Function for SHA256 hashing
func hashSHA256(input string) string {
	data := []byte(input)
	hash := sha256.Sum256(data)

	return hex.EncodeToString(hash[:])
}

// Function for 512 hashing
func hashSHA512(input string) string {
	data := []byte(input)
	hash := sha512.Sum512(data)

	return hex.EncodeToString(hash[:])
}

// Function for SHA3 hashing 
func hashSHA3(input string) string {
	data := []byte(input)
	h := sha3.New256()
	h.Write(data)
	hash := h.Sum(nil)

	return hex.EncodeToString(hash)
}
func main() {
	var value1, value2 string

	fmt.Println("=== Name + Hashing Program ===")

	fmt.Print("Please input value 1: ")
	fmt.Scan(&value1)

	fmt.Print("Please input value 2: ")
	fmt.Scan(&value2)

	fmt.Printf("Value 1 is %s, Value 2 is %s", value1, value2)
	fmt.Println()

	md5Value1 := hashMD5(value1)
	md5Value2 := hashMD5(value2)

	if md5Value1 == md5Value2 {
		fmt.Printf("Hash (MD5):\n Output A = %s\n Output B = %s\n => Match!", md5Value1, md5Value2)
	} else {
		fmt.Printf("Hash (MD5):\n Output A = %s\n Output B = %s\n => Not Match!", md5Value1, md5Value2)
	}

	fmt.Println()
	
	sha1Value1 := hashSHA1(value1)
	sha1Value2 := hashSHA1(value2)


	if sha1Value1 == sha1Value2 {
		fmt.Printf("Hash (SHA1):\n Output A = %s\n Output B = %s\n => Match!", sha1Value1, sha1Value2)
	} else {
		fmt.Printf("Hash (SHA1):\n Output A = %s\n Output B = %s\n => Not Match!", sha1Value1, sha1Value2)
	}

	fmt.Println()
	
	sha256Value1 := hashSHA256(value1)
	sha256Value2 := hashSHA256(value2)

	if sha256Value1 == sha256Value2 {
		fmt.Printf("Hash (SHA256):\n Output A = %s\n Output B = %s\n => Match!", sha256Value1, sha256Value2)
	} else {
		fmt.Printf("Hash (SHA256):\n Output A = %s\n Output B = %s\n => Not Match!", sha256Value1, sha256Value2)
	}

	fmt.Println()
	
	sha512Value1 := hashSHA512(value1)
	sha512Value2 := hashSHA512(value2)

	if sha512Value1 == sha512Value2 {
		fmt.Printf("Hash (SHA512):\n Output A = %s\n Output B = %s\n => Match!", sha512Value1, sha512Value2)
	} else {
		fmt.Printf("Hash (SHA512):\n Output A = %s\n Output B = %s\n => Not Match!", sha512Value1, sha512Value2)
	}

	fmt.Println()
	
	sha3Value1 := hashSHA3(value1)
	sha3Value2 := hashSHA3(value2)

	if sha3Value1 == sha3Value2 {
		fmt.Printf("Hash (SHA3-256):\n Output A = %s\n Output B = %s\n => Match!", sha3Value1, sha3Value2)
	} else {
		fmt.Printf("Hash (SHA3-256):\n Output A = %s\n Output B = %s\n => Not Match!", sha3Value1, sha3Value2)
	}
}
