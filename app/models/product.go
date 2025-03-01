package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// PriceInfo represents the details in the aggregated price_info JSON.
type PriceInfo struct {
	Provider     string    `json:"provider" db:"provider"`
	CurrentPrice float64   `json:"current_price" db:"current_price"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

type JSONPriceInfo []PriceInfo

func (jp *JSONPriceInfo) Scan(src interface{}) error {
	if src == nil {
		*jp = nil
		return nil
	}
	var data []byte
	switch s := src.(type) {
	case []byte:
		data = s
	case string:
		data = []byte(s)
	default:
		return fmt.Errorf("unsupported type: %T", src)
	}
	return json.Unmarshal(data, jp)
}

func (jp JSONPriceInfo) Value() (driver.Value, error) {
	return json.Marshal(jp)
}

type Product struct {
	ID        int     `json:"id" db:"id"`
	Name      string  `json:"name" db:"name"`
	Brand     string  `json:"brand" db:"brand"`
	Unit      float64 `json:"unit" db:"unit"`
	UnitType  string  `json:"unit_type" db:"unit_type"`
	Image     string  `json:"image" db:"image"`
	PriceInfo string  `json:"price_info" db:"price_info"`
}
