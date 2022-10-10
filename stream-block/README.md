# Topic: Symmetric Ciphers. Stream Ciphers. Block Ciphers.

### Course: Cryptography & Security

### Author: Marcel Vlasenco

---

## Overview

&ensp;&ensp;&ensp; Symmetric Cryptography deals with the encryption of plain text when having only one encryption key which needs to remain private. Based on the way the plain text is processed/encrypted there are 2 types of ciphers:

- Stream ciphers:
  - The encryption is done one byte at a time.
  - Stream ciphers use confusion to hide the plain text.
  - Make use of substitution techniques to modify the plain text.
  - The implementation is fairly complex.
  - The execution is fast.
- Block ciphers:
  - The encryption is done one block of plain text at a time.
  - Block ciphers use confusion and diffusion to hide the plain text.
  - Make use of transposition techniques to modify the plain text.
  - The implementation is simpler relative to the stream ciphers.
  - The execution is slow compared to the stream ciphers.

## Objectives

1. Get familiar with the symmetric cryptography, stream and block ciphers.

2. Implement an example of a stream cipher RC4.

3. Implement an example of a block cipher DES.

## Implementation description

## Run tests

```sh
$ go test ./...
```

## Conclusions
