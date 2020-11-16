package cipher

import (
	"strings"
)

type ShiftCipher int32
type VigenereCipher string

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance == 0 || distance <= -26 || distance >= 26 {
		return nil
	}
	return ShiftCipher(distance)
}

func NewVigenere(key string) Cipher {
	var cipher Cipher
	return cipher
}

func (cipher ShiftCipher) Encode(plaintText string) string {
	distance := int32(cipher)
	if distance < 0 {
		distance += 26
	}
	plaintText = strings.ToLower(plaintText)
	encoded := ""
	for _, r := range plaintText {
		if r < 'a' || r > 'z' {
			continue
		}
		encodedRune := r + distance
		if encodedRune > 'z' {
			encodedRune -= 26
		}
		encoded += string(encodedRune)
	}

	return encoded
}

func (cipher ShiftCipher) Decode(plaintText string) string {
	decode := NewShift(-int(cipher))

	return decode.Encode(plaintText)
}
