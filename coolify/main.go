package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// ヘルパー関数
func duplicateVowel(word []byte, i int) []byte {
	return append(word[:i+1], word[i:]...)
}

func removeVowel(word []byte, i int) []byte {
	return append(word[:i], word[i+1:]...)
}

func randBool() bool {
	return rand.Intn(2) == 0
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := []byte(s.Text())
		if randBool() {
			var vl int = -1
			for i, char := range word {
				switch char {
				case 'a', 'i', 'u', 'e', 'o', 'A', 'I', 'U', 'E', 'O':
					if randBool() {
						vl = i
					}
				}
			}
			if vl >= 0 {
				if randBool() {
					word = duplicateVowel(word, vl)
				} else {
					word = removeVowel(word, vl)
				}
			}
		}
		fmt.Println(string(word))
	}

}
