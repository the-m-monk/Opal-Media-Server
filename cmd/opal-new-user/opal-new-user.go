package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"unicode"

	"golang.org/x/crypto/argon2"
	"golang.org/x/term"
)

func checkForInvalidChars(userName string) {
	for _, char := range userName {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) && char != '_' {
			fmt.Println("Error: invalid character in username, username can only contain letters, numbers, and underscores.")
			os.Exit(1)
		}
	}
}

func generateSalt() []byte {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		fmt.Printf("Salting error: %s", err.Error())
		os.Exit(1)
	}
	return salt
}

func generateMCF(passStr string, salt []byte) string {
	//TODO: make this configurable to allow for less/more secure hashes depending if the server is local or public facing
	timeCost := uint32(4)
	mem := uint32(256 * 1024)
	threads := uint8(1)
	keyLen := uint32(32)

	hash := argon2.IDKey([]byte(passStr), salt, timeCost, mem, threads, keyLen)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	//hash algo id:version number:options:salt:pasword hash
	mcf := fmt.Sprintf("argon2id:19:m=%d,t=%d,p=%d:%s:%s", mem, timeCost, threads, b64Salt, b64Hash)
	return mcf
}

func main() {
	fmt.Print("Username: ")
	var userName string
	fmt.Scanln(&userName)
	//TODO/BUG: if username contains spaces, it doesnt set off the username check but instead everything after the space is used for the password
	checkForInvalidChars(userName)

	//TODO: abitrary limit
	if len(userName) > 16 {
		fmt.Println("Error: username cannot be longer than 16 characters in length")
		os.Exit(1)
	}

	fmt.Print("Password: ")
	pass, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	fmt.Println()

	fmt.Print("Confirm password: ")
	secondPass, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	fmt.Println()

	passStr := strings.TrimSpace(string(pass))
	secondPassStr := strings.TrimSpace(string(secondPass))

	if passStr != secondPassStr {
		fmt.Println("Error: passwords do not match.")
		os.Exit(1)
	}

	salt := generateSalt()
	mcf := generateMCF(passStr, salt)

	file, err := os.OpenFile("./db/users", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Println("Error opening user db:", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(fmt.Sprintf("%s:%s\n", userName, mcf)); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Successfully added user")
}
