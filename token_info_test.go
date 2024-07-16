package gasSDK

import (
	"testing"
)

func TestFetchTokenInfo(t *testing.T) {
	tokenAddress := "EQAxdo7MIQfOWokdf9Dw4yUm4oP43i1m9yzT07cawBVQWiYH"
	tokenInfo, err := FetchTokenInfo(tokenAddress)
	if err != nil {
		t.Fatalf("failed to fetch token info: %v", err)
	}

	if tokenInfo.Name != "Superton " {
		t.Errorf("expected name to be 'Superton ', got %s", tokenInfo.Name)
	}
}
