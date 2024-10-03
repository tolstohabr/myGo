package models

import (
	"encoding/json"
)

// Banner представляет баннер.
type Banner struct {
	ID        int             `json:"id" db:"id"`                 // ID баннера
	JSONData  json.RawMessage `json:"json_data" db:"json_data"`   // Данные в формате JSON
	FeatureID int             `json:"feature_id" db:"feature_id"` // ID фичи, к которой относится баннер
	IsActive  bool            `json:"is_active" db:"is_active"`   // Статус активации баннера
}

// Tag представляет тег.
type Tag struct {
	ID   int    `json:"id" db:"id"`     // ID тега
	Name string `json:"name" db:"name"` // Имя тега
}

// Feature представляет фичу.
type Feature struct {
	ID   int    `json:"id" db:"id"`     // ID фичи
	Name string `json:"name" db:"name"` // Имя фичи
}
