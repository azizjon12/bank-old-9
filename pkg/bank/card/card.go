package card

import "bank/pkg/bank/types"

// Total calculates the total amount of money in all Cards
// Negative amounts and blocked cards will be ignored
func Total(cards []types.Card) types.Money {
	sum := types.Money(0)
	for _, card := range cards {
		if !card.Active {
			continue
		}

		if card.Balance <= 0 {
			continue
		}

		sum += card.Balance
	}
	return sum
}
