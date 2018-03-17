// passdrill: typing drills for practicing passphrases

package main

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/howeyc/gopass"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/scrypt"
)

const hashFilename = "passdrill.hash"
const derivedKeyLen = 64
const help = "Use -s to save passphrase hash for practice."

var logger = log.New(os.Stderr, "", log.Lshortfile)

func input(msg string) string {
	response := ""
	fmt.Print(msg)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		response = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		logger.Fatalln(err)
	}
	return response
}

func prompt() string {
	fmt.Println("WARNING: the passphrase will be shown so that you can check it!")
	confirmed := ""
	passwd := ""
	for confirmed != "y" {
		passwd = input("Type passphrase to hash (it will be echoed): ")
		if passwd == "" || passwd == "q" {
			fmt.Println("ERROR: the passphrase cannot be empty or 'q'.")
			continue
		}
		fmt.Println("Passphrase to be hashed ->", passwd)
		confirmed = strings.ToLower(input("Confirm (y/n): "))
	}
	return passwd
}

func myPbkdf2(salt, content []byte) []byte {
	const rounds = 200000
	algorithm := sha512.New
	return pbkdf2.Key(content, salt, rounds, derivedKeyLen, algorithm)
}

func myScrypt(salt, content []byte) []byte {
	// The recommended parameters for interactive logins as of 2017 are:
	// N=32768, r=8 and p=1 (https://godoc.org/golang.org/x/crypto/scrypt)
	key, err := scrypt.Key(content, salt, 32768, 8, 1, derivedKeyLen)
	if err != nil {
		logger.Fatalln(err)
	}
	return key
}

func computeHash(keyFunc string, salt []byte, text string) []byte {
	content := []byte(text)
	if keyFunc == "pbkdf2" {
		return myPbkdf2(salt, content)
	} else if keyFunc == "scrypt" {
		return myScrypt(salt, content)
	}
	logger.Fatalln("computeHash: Unknown key function " + keyFunc)
	return nil
}

func buildHash(keyFunc, text string) []byte {
	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	if err != nil {
		logger.Fatalln(err)
	}
	octets := computeHash(keyFunc, salt, text)
	headerStr := keyFunc + ":" + base64.StdEncoding.EncodeToString(salt) +
		":" + base64.StdEncoding.EncodeToString(octets[:])
	return []byte(headerStr)
}

func saveHash(args []string) {
	if len(os.Args) > 2 || os.Args[1] != "-s" {
		fmt.Println("ERROR: invalid argument.", help)
		os.Exit(1)
	}
	// New passdrill hashes use scrypt; pbkdf2 is
	// supported only for reading old hash files
	wrappedHash := buildHash("scrypt", prompt())
	err := ioutil.WriteFile(hashFilename, wrappedHash, 0600)
	if err != nil {
		logger.Fatalln(err)
	}
	fmt.Println("Passphrase hash saved to", hashFilename)
}

func unwrapHash(wrappedHash []byte) (string, []byte, []byte) {
	fields := strings.Split(string(wrappedHash), ":")
	if len(fields) != 3 {
		logger.Fatalln("invalid passphrase hash file:",
			"3 fields expected, found", len(fields))
	}
	keyFunc := fields[0]
	salt, err := base64.StdEncoding.DecodeString(fields[1])
	if err != nil {
		logger.Fatalln(err)
	}
	passwdHash, err := base64.StdEncoding.DecodeString(fields[2])
	if err != nil {
		logger.Fatalln(err)
	}
	return keyFunc, salt, passwdHash
}

func practice() {
	wrappedHash, err := ioutil.ReadFile(hashFilename)
	if os.IsNotExist(err) {
		fmt.Println("ERROR: passphrase hash file not found.", help)
		os.Exit(1)
	}
	if err != nil {
		logger.Fatalln(err)
	}
	keyFunc, salt, passwdHash := unwrapHash(wrappedHash)
	fmt.Println("Type q to end practice.")
	turn := 0
	correct := 0
	for {
		turn++
		fmt.Printf("%d:", turn)
		octets, err := gopass.GetPasswd()
		if err != nil {
			logger.Fatalln(err)
		}
		response := string(octets)
		if response == "" {
			fmt.Println("Type q to quit.")
			turn-- // don't count this response
			continue
		} else if response == "q" {
			turn-- // don't count this response
			break
		}
		answer := "wrong"
		if bytes.Compare(computeHash(keyFunc, salt, response), passwdHash) == 0 {
			correct++
			answer = "OK"
		}
		fmt.Printf("  %s\thits=%d\tmisses=%d\n", answer, correct, turn-correct)
	}
	if turn > 0 {
		pct := float64(correct) / float64(turn) * 100
		fmt.Printf("\n%d exercises. %0.1f%% correct.\n", turn, pct)
	}
}

func main() {
	if len(os.Args) > 1 {
		saveHash(os.Args)
	} else {
		practice()
	}
}
