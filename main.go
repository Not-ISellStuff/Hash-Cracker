package main

// Simple script for cracking hashes
// Haven't codded in go in a while so I'm just practicing

// MD5 = 1
// SHA1 = 2
// SHA256 = 3

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

// Functions

func hash(str string, meth string) string {
	if meth == "1" {
		h := md5.New()
		h.Write([]byte(str))
		ha := h.Sum(nil)
		hashed := hex.EncodeToString(ha)
		return hashed

	} else if meth == "2" {
		h := sha1.New()
		h.Write([]byte(str))
		ha := h.Sum(nil)
		hashed := hex.EncodeToString(ha)
		return hashed

	} else if meth == "3" {
		h := sha256.New()
		h.Write([]byte(str))
		ha := h.Sum(nil)
		hashed := hex.EncodeToString(ha)
		return hashed
	}

	return ""
}

func crack(hash_ string, meth string) string {
	//     	ERRORS

	// 1 = Error Reading File
	// 2 = Failed To Crack Hash
	// 3 = Algorithm Not Supported

	f, e := os.Open("wordlist.txt")
	if e != nil {
		return "e:1"
	}
	s := bufio.NewScanner(f)

	if meth == "1" {
		for s.Scan() {
			ln := s.Text()
			hashed := hash(ln, "1")

			if hashed == hash_ {
				return "t:" + ln
			}
		}
		return "e:2"
	} else if meth == "2" {
		for s.Scan() {
			ln := s.Text()
			hashed := hash(ln, "2")

			if hashed == hash_ {
				return "t:" + ln
			}
		}
		return "e:2"
	} else if meth == "3" {
		for s.Scan() {
			ln := s.Text()
			hashed := hash(ln, "3")

			if hashed == hash_ {
				return "t:" + ln
			}
		}
		return "e:2"
	}

	return "e:3"
}

//-----------------------------------------------------------------//

func main() {
	var op string
	var hash_ string

	menu := "Made By --> [ISellStuff]\nhttps://github.com/Not-ISellStuff\n\nSelect Algorithm:\n[1] MD5\n[2] SHA1\n[3] SHA256\n"
	fmt.Println(menu)
	fmt.Println("")
	fmt.Scan(&op)

	fmt.Println("\nNow Enter The Hash: ")
	fmt.Scan(&hash_)

	fmt.Println("\n[+] Attempting To Crack Hash")
	result := crack(hash_, op)
	mainr := strings.Split(result, ":")

	if mainr[0] == "e" {
		var exit string
		var error string
		err := mainr[1]

		if err == "1" {
			error = "Error Reading File"
		} else if err == "2" {
			error = "Failed To Crack Hash"
		} else {
			error = "Algorithm Not Supported"
		}

		fmt.Println("\n[!] Error: " + error + "\nPress Any Key Then Enter To Close > ")
		fmt.Scan(&exit)

		return
	}

	var ext string
	var cracked = string(result[1])
	fmt.Println("[+] Cracked Hash | " + cracked + "\n" + "\nPress Any Key Then Enter To Close > ")
	fmt.Scan(&ext)
}
