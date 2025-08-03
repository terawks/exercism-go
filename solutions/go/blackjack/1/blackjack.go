package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch {
	case card == "ace":
		return 11
	case card == "two":
		return 2
	case card == "three":
		return 3
	case card == "four":
		return 4
	case card == "five":
		return 5
	case card == "six":
		return 6
	case card == "seven":
		return 7
	case card == "eight":
		return 8
	case card == "nine":
		return 9
	case card == "ten":
		return 10
	case card == "jack":
		return 10
	case card == "queen":
		return 10
	case card == "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	val1, val2, dval := ParseCard(card1), ParseCard(card2), ParseCard(dealerCard)
	if val1 == 11 && val2 == 11 {
		return "P"
	}
	pairsum := val1 + val2
	if pairsum == 21 {
		if dval < 10 {
			return "W"
		}
		return "S"
	}
	if 17 <= pairsum {
		return "S"
	}
	if 12 <= pairsum {
		if 7 <= dval {
			return "H"
		}
		return "S"
	}
	return "H"
}
