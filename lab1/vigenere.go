package main

import "strings"

func vigenereEncrypt(alphabet string, key string, text string) string {
	var result strings.Builder
	for i, c := range text {
		if c == ' ' {
			result.WriteByte(' ')
			continue
		}

		index := strings.Index(alphabet, string(c))
		if index == -1 {
			panic("invalid character")
		}

		k := key[i%len(key)]
		shift := strings.Index(alphabet, string(k))
		if shift == -1 {
			panic("invalid character")
		}

		result.WriteByte(alphabet[(index+shift)%len(alphabet)])
	}
	return result.String()
}

func vigenereDecrypt(alphabet string, key string, text string) string {
	var result strings.Builder
	for i, c := range text {
		if c == ' ' {
			result.WriteByte(' ')
			continue
		}

		index := strings.Index(alphabet, string(c))
		if index == -1 {
			panic("invalid character")
		}

		k := key[i%len(key)]
		shift := strings.Index(alphabet, string(k))
		if shift == -1 {
			panic("invalid character")
		}

		result.WriteByte(alphabet[(index-shift+len(alphabet))%len(alphabet)])
	}
	return result.String()
}
