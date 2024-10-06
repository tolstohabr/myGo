package models

import (
	"encoding/json"
)

type Banner struct {
	ID        int             `json:"id" db:"id"`
	JSONData  json.RawMessage `json:"json_data" db:"json_data"`
	FeatureID int             `json:"feature_id" db:"feature_id"`
	IsActive  bool            `json:"is_active" db:"is_active"`
}
