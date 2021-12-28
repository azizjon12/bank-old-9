package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleTotal() {
	fmt.Println(Total([]types.Card{
		{
			Balance: 10_000_00,
			Active:  true,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: 10_000_00,
			Active:  true,
		},
		{
			Balance: 20_000_00,
			Active:  true,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: 10_000_00,
			Active:  false,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: -10_000_00,
			Active:  true,
		},
	}))

	// Output:
	// 1000000
	// 3000000
	// 0
	// 0

}

func ExamplePaymentSources() {
	cards := []types.Card{
		{
			PAN:     "5058 0001 0000 8888",
			Balance: 100_00,
			Active:  true,
		},
		{
			PAN:     "5058 0002 0000 8888",
			Balance: 200_00,
			Active:  false,
		},
		{
			PAN:     "5058 0003 0000 8888",
			Balance: 0,
			Active:  true,
		},
		{
			PAN:     "5058 0004 0000 8888",
			Balance: 1_00,
			Active:  true,
		},
	}

	paymentSources := PaymentSources(cards)
	for _, source := range paymentSources {
		fmt.Println(source.Number)
	}

	// Output:
	// 5058 0001 0000 8888
	// 5058 0004 0000 8888
}
