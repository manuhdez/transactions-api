package event

import (
	"encoding/json"
	"log"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/transaction"
)

type DepositCreated struct {
	body []byte
}

type DepositCreatedBody struct {
	Type    string  `json:"type"`
	Account string  `json:"account"`
	Amount  float32 `json:"amount"`
}

func NewDepositCreated(trx transaction.Transaction) DepositCreated {
	body := DepositCreatedBody{string(DepositCreatedType), trx.AccountId, trx.Amount}
	b, _ := json.Marshal(body)
	return DepositCreated{b}
}

func NewDepositCreatedBody(data []byte) (DepositCreatedBody, error) {
	var body DepositCreatedBody
	err := json.Unmarshal(data, &body)
	if err != nil {
		log.Printf("Error parsing created deposit event: %e", err)
		return DepositCreatedBody{}, err
	}

	return body, nil
}

func (d DepositCreated) Type() Type {
	return DepositCreatedType
}

func (d DepositCreated) Body() []byte {
	return d.body
}
