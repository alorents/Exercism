package secret

func Handshake(code uint) []string {
	var handshake []string
	if code&0b1 == 0b1 {
		handshake = append(handshake, "wink")
	}
	if code&0b10 == 0b10 {
		handshake = append(handshake, "double blink")
	}
	if code&0b100 == 0b100 {
		handshake = append(handshake, "close your eyes")
	}
	if code&0b1000 == 0b1000 {
		handshake = append(handshake, "jump")
	}
	if code&0b10000 == 0b10000 {
		for i := 0; i < len(handshake)/2; i++ {
			j := len(handshake) - i - 1
			handshake[i], handshake[j] = handshake[j], handshake[i]
		}
	}

	return handshake
}
