package models

type Currency struct {
	ID           int     `json:"id" db:"id"`
	BuyRate      float64 `json:"buyRate" db:"buy_rate"`
	SellRate     float64 `json:"sellRate" db:"sell_rate"`
	ProviderName string  `json:"providerName" db:"name"`
	ProviderLogo string  `json:"providerLogo" db:"logo_url"`
	Code         string  `json:"code" db:"code"`
}
