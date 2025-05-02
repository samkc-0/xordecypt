package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

const ascii_lower = "abcdefghijklmnopqrstuvwxyz"
const ascii_upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const ascii_digits = "0123456789"
const keyspace = ascii_lower + ascii_upper + ascii_digits

func xorDecrypt(cipher, key []byte) []byte {
	out := make([]byte, len(cipher))
	for i := range cipher {
		out[i] = cipher[i] ^ key[i%len(key)]
	}
	return out
}

func recoverPartialKey(cipher, known []byte) []byte {
	key := make([]byte, len(known))
	for i := range known {
		key[i] = cipher[i] ^ known[i]
	}
	return key
}

func bruteForce(cipher, knownKey []byte, keyLen int) {
	pad := keyLen - len(knownKey)
	if pad != 1 {
		fmt.Println("only 1 unknown byte supported atm")
		os.Exit(1)
	}

	for i := 0; i < len(keyspace); i++ {
		c := keyspace[i]
		key := append(knownKey, c)
		pt := xorDecrypt(cipher, key)
		if isPrintable(pt) {
			fmt.Printf("key: %q => %s\n", string(key), string(pt))
		}
	}
}

func isPrintable(b []byte) bool {
	for _, c := range b {
		if c < 32 || c > 126 {
			return false
		}
	}
	return true
}
func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `xordecrypt: xor decryption with partial known plaintext

Usage:
  xordecrypt -c CIPHERTEXT -p KNOWN_PLAINTEXT_PREFIX -l KEYLEN 

Options:
`)
		flag.PrintDefaults()
	}

	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(1)
	}

	c := flag.String("c", "", "hex-encoded input file")
	prefix := flag.String("p", "", "known plaintext prefix")
	keylen := flag.Int("l", 0, "total key length")
	flag.Parse()

	if *c == "" {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, "failed to read from stdin")
			os.Exit(1)
		}
		*c = strings.TrimSpace(string(data))
	}
	if *prefix == "" {
		fmt.Fprintln(os.Stderr, "missing required -p (plaintext prefix) argument")
		flag.Usage()
		os.Exit(1)
	}
	if *keylen == 0 {
		fmt.Fprintln(os.Stderr, "missing or invalid -l (keylen) argument")
		flag.Usage()
		os.Exit(1)
	}
	if *keylen < len(*prefix) {
		fmt.Fprintf(os.Stderr, "keylen (%d) cannot be shorter than prefix length (%d)\n", *keylen, len(*prefix))
		os.Exit(1)
	}

	cipher, err := hex.DecodeString(*c)
	if err != nil {
		panic(err)
	}

	known := []byte(*prefix)
	partial := recoverPartialKey(cipher, known)
	bruteForce(cipher, partial, *keylen)
}
