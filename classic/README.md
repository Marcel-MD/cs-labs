# Intro to Cryptography. Classical ciphers. Caesar cipher.

### Course: Cryptography & Security

### Author: Marcel Vlasenco

---

## Overview

&ensp;&ensp;&ensp; Caesar cipher has a key which is used to substitute the characters with the next ones, by the order number in a pre-established alphabet. Mathematically it would be expressed as follows:

$em = enc_{k}(x) = x + k (mod \; n),$

$dm = dec_{k}(x) = x + k (mod \; n),$

where:

- em: the encrypted message,
- dm: the decrypted message (i.e. the original one),
- x: input,
- k: key,
- n: size of the alphabet.

&ensp;&ensp;&ensp; Judging by the encryption mechanism one can conclude that this cipher is pretty easy to break. In fact, a brute force attack would have **_O(nm)_** complexity, where **_n_** would be the size of the alphabet and **_m_** the size of the message. This is why there were other variations of this cipher, which are supposed to make the cryptanalysis more complex.

## Objectives

Implement 4 types of the classical ciphers:

- Caesar cipher with one key used for substitution (as explained above)
- Caesar cipher with one key used for substitution, and a permutation of the alphabet
- Vigenere cipher
- Playfair cipher

## Implementation description

- **Caesar cipher**  
  Encryption and decryption functions take a few arguments: `alphabet`, `shift`, `text` and a `seed` for alphabet permutation. The logic is strait-forward we just loop through text's chars and replace them from alphabet with a shift to the right. We can also shuffle the alphabet before we do the encryption.

  ```go
  func Encrypt(alphabet string, shift int, text string) string {
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

  func PermutationEncrypt(alphabet string, seed int64, shift int, text string) string {

    newAlphabet := permute(alphabet, seed)

    return Encrypt(string(newAlphabet), shift, text)
  }
  ```

- **Vigenere cipher**  
  Encryption and decryption functions take a few arguments: `alphabet`, `key` and `text`. We loop through text's chars, then we get the `index` of current char in the alphabet, then we calculate the char from the key that corresponds to that char from the text `k = key[i % len(key)]` then we calculate the `shift` which is the index of this char. Finally current char will be replaced with `alphabet[(index+shift) % len(alphabet)]`.

  ```go
  func Encrypt(alphabet string, key string, text string) string {
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
  ```

- **Playfair cipher**  
  Too much to explain, just go read GeeksForGeeks.

  ```go
  func Encrypt(key string, text string) string {
    text = cleanText(text)
    table := generateTable(key)

    var result strings.Builder
    for i := 0; i < len(text); i += 2 {
      c1 := text[i]
      c2 := text[i+1]

      i1, j1 := search(table, c1)
      i2, j2 := search(table, c2)

      if i1 == i2 {
        result.WriteByte(table[i1][(j1+1)%5])
        result.WriteByte(table[i2][(j2+1)%5])
      } else if j1 == j2 {
        result.WriteByte(table[(i1+1)%5][j1])
        result.WriteByte(table[(i2+1)%5][j2])
      } else {
        result.WriteByte(table[i1][j2])
        result.WriteByte(table[i2][j1])
      }
    }

    return result.String()
  }
  ```

## Run tests

```sh
$ go test ./...
```

## Conclusions

After performing this laboratory work I became familiar with classical ciphers. Back when they were used widely, they might been secure enough, but the advent of computers made them obsolete. Still it's interesting to see where it all started.
