package request

type Withdraw struct {
	Account  string  `json:"account" binding:"required"`
	Amount   float32 `json:"amount" binding:"required"`
	Currency string  `json:"currency" binding:"required"`
}
