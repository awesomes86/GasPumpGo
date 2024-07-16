package gasSDK

import (
	"testing"
)

func TestFetchTransactions(t *testing.T) {
	// Test with token address
	tokenAddress := "EQAxdo7MIQfOWokdf9Dw4yUm4oP43i1m9yzT07cawBVQWiYH"
	limit := 100
	transactions, err := FetchTransactions(tokenAddress, limit)
	if err != nil {
		t.Fatalf("failed to fetch transactions with token address: %v", err)
	}

	if len(transactions) == 0 {
		t.Errorf("expected transactions to be non-empty with token address")
	}

	// Test without token address
	transactions, err = FetchTransactions("", limit)
	if err != nil {
		t.Fatalf("failed to fetch transactions without token address: %v", err)
	}

	if len(transactions) == 0 {
		t.Errorf("expected transactions to be non-empty without token address")
	}
}
