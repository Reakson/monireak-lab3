package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
)

const (
	password   = "meowmeow"
	flagText   = "cryptoCTF{meowmeow}"
	leakedHash = "b6ccb4ece5454dcae51778b3e239ebc2"
)

func md5Hash(value string) string {
	digest := md5.Sum([]byte(value))
	return hex.EncodeToString(digest[:])
}

func main() {
	encodedSearch := "%62%36%63%63%62%34%65%63%65%35%34%35%34%64%63%61%65%35%31%37%37%38%62%33%65%32%33%39%65%62%63%32"
	decodedSearch, err := url.QueryUnescape(encodedSearch)
	if err != nil {
		fmt.Println("failed to decode leaked search: ", err)
		return
	}

	fmt.Println("Lab 4 - Mini-CTF (Cat Leak)")
	fmt.Println("Decoded leaked hash:", decodedSearch)
	fmt.Println("MD5(meowmeow):", md5Hash(password))
	fmt.Println("Hash match:", decodedSearch == leakedHash && decodedSearch == md5Hash(password))
	fmt.Println("Flag:", flagText)
}
