package main

import (
	"bufio"
	"fmt"
	"os"

	"LabW3/Lab2/utils/crack"
)

func main() {
	wordlist, err := os.Open("../wordlists/nord_vpn.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wordlist.Close()

	verbose, err := os.Create("verbose.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer verbose.Close()

	target := "aa1c7d931cf140bb35a5a16adeb83a551649c3b9"
	scanner := bufio.NewScanner(wordlist)
	attempts := 0
	found := ""

	for scanner.Scan() {
		password := scanner.Text()
		hash := crack.Hash(password)
		attempts++
		fmt.Fprintf(verbose, "Attempt %d: %s -> %s\n", attempts, password, hash)

		if hash == target {
			found = password
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Lab 2 - SHA-1 password cracking")
	fmt.Println("Password found:", found)
	fmt.Println("Attempts:", attempts)
}
