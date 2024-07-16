package gasSDK

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Transaction represents a single transaction in the list.
type Transaction struct {
	Hash            string    `json:"hash"`
	Type            string    `json:"type"`
	Status          string    `json:"status"`
	TokenAddress    string    `json:"token_address"`
	UserInfo        UserInfo  `json:"user_info"`
	TokenInfo       TokenInfo `json:"token_info"`
	CreatedAt       string    `json:"created_at"`
	TonValue        *string   `json:"ton_value"`
	TokenValue      *string   `json:"token_value"`
	TransactionInfo struct {
		FromAddress  string  `json:"from_address"`
		TonAmount    string  `json:"ton_amount"`
		JettonAmount string  `json:"jetton_amount"`
		FeeTonAmount string  `json:"fee_ton_amount"`
		Price        float64 `json:"price"`
	} `json:"transaction_info"`
}

// FetchTransactions fetches the list of transactions from the API.
func FetchTransactions(tokenAddress string, limit int) ([]Transaction, error) {
	var url string
	if tokenAddress != "" {
		url = fmt.Sprintf("https://api.gas111.com/api/v1/transactions/list?limit=%d&token_address=%s", limit, tokenAddress)
	} else {
		url = fmt.Sprintf("https://api.gas111.com/api/v1/transactions/list?limit=%d", limit)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch transactions: %s", resp.Status)
	}

	var transactions []Transaction
	if err := json.NewDecoder(resp.Body).Decode(&transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}
