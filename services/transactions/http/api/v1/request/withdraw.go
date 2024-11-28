package request

type Withdraw struct {
	Account  string  `json:"account" validate:"required"`
	Amount   float32 `json:"amount" validate:"required"`
	Currency string  `json:"currency" validate:"required,iso4217"`
}
