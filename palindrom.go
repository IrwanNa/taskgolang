package main

import (
	"fmt"
	"strings"
)

// fungsi
func isPalindrome(s string) bool {
	// Konversi ke huruf
	s = strings.ToLower(s)
	// Menghilangkan spasi
	s = strings.ReplaceAll(s, " ", "")
	// Inisialisasi dua indeks
	i, j := 0, len(s)-1
	// Looping untuk membandingkan karakter dari depan dan belakang
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	
	return true
}

func main() {
	
	var input string
	fmt.Print("Masukkan kata atau kalimat: ")
	fmt.Scanln(&input)

	if isPalindrome(input) {
		fmt.Println("Itu Palindrom")
	} else {
		fmt.Println("Bukan Palindrom")
	}
}
