package main

import (
	"bufio"
	"fmt"
	"os"

	"LabW3/Lab1/utils/crack"
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

	target := "6a85dfd77d9cb35770c9dc6728d73d3f"
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

	fmt.Println("Lab 1 - MD5 password cracking")
	fmt.Println("Password found:", found)
	fmt.Println("Attempts:", attempts)
}
