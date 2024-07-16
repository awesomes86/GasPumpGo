package gasSDK

import (
	"testing"
)

func TestFetchUserInfo(t *testing.T) {
	telegramID := int64(6917514253)
	user, err := FetchUserInfo(telegramID)
	if err != nil {
		t.Fatalf("failed to fetch user info: %v", err)
	}

	if len(user.Name) == 0 {
		t.Errorf("expected name, got %s", user.Name)
	}
}
