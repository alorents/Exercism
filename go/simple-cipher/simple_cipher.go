package cipher

import (
	"regexp"
	"strings"
	"unicode"
)

// VigenereCipher contains the key for a Vigenere Cipher 
// encode bool controls the direction(encode/decode) of the shift when applying key
type VigenereCipher struct {
	key string
	encode bool
}

// NewCaesar creates a specific implementation of a Shift Cipher
func NewCaesar() Cipher {
	return NewShift(3)
}

// NewShift creates a Shift Cipher from a provide shift distance by implementing a simplified Vigenere Cipher
func NewShift(distance int) Cipher {
	if distance == 0 || distance <= -26 || distance >= 26 {
		return nil
	}
	return VigenereCipher{string('a' + distance), true}
}

// NewVigenere crates a Vigenere Cipher from a provided key
func NewVigenere(key string) Cipher {
	isValidKey := false
	for _, r := range key {
		if r == 'a' {
			continue
		}
		if !unicode.IsLower(r) {
			return nil
		}
		isValidKey = true
	}
	if !isValidKey {
		return nil
	}

	return VigenereCipher{key, true}
}

// NewDecodeVigenere creates a decoding Vigenere Cipher by applying the provided key backwards.
// Reverse compatible with Shit and Ceaser Ciphers
func NewDecodeVigenere(key string) Cipher {
	return VigenereCipher{key, false}
}

// Encode applies the key in the direction specified by VigenereCipher
func (cipher VigenereCipher) Encode(plaintText string) string {
	normailizedText := regexp.MustCompile(`[^a-z]`).ReplaceAllString(strings.ToLower(plaintText), "")
	encoded := ""
	for i, r := range normailizedText {
		if r < 'a' || r > 'z' {
			continue
		}
		distance := int32(cipher.key[i%len(cipher.key)]) - 'a'
		if !cipher.encode {
			distance *= -1
		}
		if distance < 0 {
			distance += 26
		}
		encodedRune := r + distance
		if encodedRune > 'z' {
			encodedRune -= 26
		}
		encoded += string(encodedRune)
	}

	return encoded
}

// Decode decodes the encoded text by applying the specified cipher key in reverse
func (cipher VigenereCipher) Decode(encodedText string) string {
	decode := NewDecodeVigenere(cipher.key)

	return decode.Encode(encodedText)
}
