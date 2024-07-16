package gasSDK

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// User represents the JSON structure of the user information.
type User struct {
	TelegramID       int64  `json:"telegram_id"`
	Name             string `json:"name"`
	ImageURL         string `json:"image_url"`
	TelegramUsername string `json:"telegram_username"`
	WalletAddress    string `json:"wallet_address"`
}

// FetchUserInfo fetches the user information from the API.
func FetchUserInfo(telegramID int64) (*User, error) {
	url := fmt.Sprintf("https://api.gas111.com/api/v1/users/info?telegram_id=%d", telegramID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch user info: %s", resp.Status)
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
