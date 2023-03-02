package request

type CreateAccount struct {
	Id       string  `json:"id"`
	Balance  float32 `json:"balance"`
	Currency string  `json:"currency"`
}
