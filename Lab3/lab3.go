package main

import (
	"bufio"
	"fmt"
	"os"

	"LabW3/Lab3/utils/crack"
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

	target := "485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38"
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

	fmt.Println("Lab 3 - SHA-512 password cracking")
	fmt.Println("Password found:", found)
	fmt.Println("Attempts:", attempts)
}
