package main

import (
	"math/rand"
	"strings"
)

func caesarEncrypt(alphabet string, shift int, text string) string {
	var result strings.Builder
	for _, c := range text {
		if c == ' ' {
			result.WriteByte(' ')
			continue
		}

		index := strings.Index(alphabet, string(c))
		if index == -1 {
			panic("invalid character")
		}

		result.WriteByte(alphabet[(index+shift)%len(alphabet)])
	}
	return result.String()
}

func caesarDecrypt(alphabet string, shift int, text string) string {
	return caesarEncrypt(alphabet, len(alphabet)-shift, text)
}

func caesarPermutationEncrypt(alphabet string, seed int64, shift int, text string) string {

	newAlphabet := permute(alphabet, seed)

	return caesarEncrypt(string(newAlphabet), shift, text)
}

func caesarPermutationDecrypt(alphabet string, seed int64, shift int, text string) string {

	newAlphabet := permute(alphabet, seed)

	return caesarDecrypt(string(newAlphabet), shift, text)
}

func permute(alphabet string, seed int64) string {
	a := []rune(alphabet)

	rand.Seed(seed)
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	return string(a)
}
