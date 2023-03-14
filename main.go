package main

import (
	"fmt"
)

func main() {
	byWord("selamat malam")
	byMap("selamat malam")
}

func byWord(teks string) string {
	for i := 0; i < (len(teks)); i++ {
		fmt.Println(string(teks[i]))
	}
	return teks

}

func byMap(nextText string) {
	counts := make(map[rune]int)

	for _, char := range nextText {
		counts[char]++

	}
	fmt.Print("map[")
	for char, count := range counts {
		fmt.Printf("%c:%d ", char, count) // menampilkan key dan value pada map
	}
	fmt.Print("]")
	fmt.Println()

	fmt.Println(counts) // jika saya print menggunakan ini maka seluruh format print nya menggunakan integer
	// begitu juga sebaliknya jika saya print menggunakan %c maka setiap hasil print an menampilkan karakter

}
