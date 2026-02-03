package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    switch card {
	case "ace": return 11
	case "two": return 2
	case "three": return 3
	case "four": return 4
	case "five": return 5
	case "six": return 6
	case "seven": return 7
	case "eight": return 8
	case "nine": return 9
	case "ten", "jack", "queen", "king": return 10
	default: return 0
	}
	panic("Please implement the ParseCard function")
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    sum := ParseCard(card1) + ParseCard(card2)
	dealerValue := ParseCard(dealerCard)

	switch {
	// Trường hợp đặc biệt: Hai quân Át (Pair of Aces)
	case card1 == "ace" && card2 == "ace":
		return "P" // Split

	// Trường hợp Blackjack (21 điểm)
	case sum == 21:
		if dealerValue < 10 {
			return "W" // Win automatically
		}
		return "S" // Stand và đợi xem dealer có Blackjack không

	// Điểm từ 17 đến 20
	case sum >= 17:
		return "S" // Stand

	// Điểm từ 12 đến 16
	case sum >= 12:
		if dealerValue >= 7 {
			return "H" // Hit nếu dealer có bài mạnh
		}
		return "S" // Stand nếu dealer có bài yếu

	// Điểm từ 11 trở xuống
	default:
		return "H" // Hit để lấy thêm bài
	}
	panic("Please implement the FirstTurn function")
}
