package gasSDK

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// TokenInfo represents the JSON structure of the token information.
type TokenInfo struct {
	Name                      string   `json:"name"`
	Ticker                    string   `json:"ticker"`
	TokenAddress              string   `json:"token_address"`
	ImageURL                  string   `json:"image_url"`
	UserInfo                  UserInfo `json:"user_info"`
	ContractVersion           int      `json:"contract_version"`
	Status                    string   `json:"status"`
	StickerPackURL            string   `json:"sticker_pack_url"`
	IsFull                    *bool    `json:"is_full"`
	MarketCap                 string   `json:"market_cap"`
	PoolAddress               *string  `json:"pool_address"`
	UnwrappedJettonMasterAddr *string  `json:"unwrapped_jetton_master_address"`
	DedustWhitelistReqSent    *bool    `json:"dedust_whitelist_request_sent"`
	TgChannelLink             string   `json:"tg_channel_link"`
	TgChatLink                string   `json:"tg_chat_link"`
	TwitterLink               string   `json:"twitter_link"`
	WebsiteLink               string   `json:"website_link"`
	Description               string   `json:"description"`
	UnwrapDeadlineAt          *string  `json:"unwrap_deadline_at"`
}

// UserInfo represents the user information structure in the JSON response.
type UserInfo struct {
	TelegramID       int    `json:"telegram_id"`
	Name             string `json:"name"`
	ImageURL         string `json:"image_url"`
	TelegramUsername string `json:"telegram_username"`
	WalletAddress    string `json:"wallet_address"`
}

// FetchTokenInfo fetches the token information from the API.
func FetchTokenInfo(tokenAddress string) (*TokenInfo, error) {
	url := fmt.Sprintf("https://api.gas111.com/api/v1/tokens/info?token_address=%s", tokenAddress)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch token info: %s", resp.Status)
	}

	var tokenInfo TokenInfo
	if err := json.NewDecoder(resp.Body).Decode(&tokenInfo); err != nil {
		return nil, err
	}

	return &tokenInfo, nil
}
