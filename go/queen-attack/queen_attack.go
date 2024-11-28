package queenattack

import "fmt"

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
	if len(whitePosition) != 2 || len(blackPosition) != 2 || whitePosition == blackPosition {
		return false, fmt.Errorf("invalid positions %s, %s", whitePosition, blackPosition)
	}

	whiteFile := int(whitePosition[0])
	whiteRank := int(whitePosition[1])
	blackFile := int(blackPosition[0])
	blackRank := int(blackPosition[1])

	if whiteFile < 'a' || whiteFile > 'h' || blackFile < 'a' || blackFile > 'h' ||
		whiteRank < '1' || whiteRank > '8' || blackRank < '1' || blackRank > '8' {
		return false, fmt.Errorf("invalid positions %s, %s", whitePosition, blackPosition)
	}

	rise, run := whiteFile-blackFile, whiteRank-blackRank
	if rise == 0 || run == 0 || rise == run || rise == -run {
		return true, nil
	}
	return false, nil
}
