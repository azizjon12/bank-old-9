package types

// Money represents amount of money in minimum units e.g. dirams, cents, pennies
type Money int64

// Currency represents the currency code, e.g. "TJS", "GBP", "USD"
type Currency string

// Currency codes
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// PAN (Payment Card Number)
type PAN string

// Card represents information regarding the payment card
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money // Money from above
	Currency   Currency
	Color      string
	Name       string
	Active     bool
	MinBalance Money
}

// Payment represents information about the payments
type Payment struct {
	ID     int
	Amount Money
}

type PaymentSource struct {
	Type    string
	Number  string
	Balance Money
}
