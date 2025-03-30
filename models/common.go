package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// StringSlice handles []string as JSONB in PostgreSQL
type StringSlice []string

// Convert Go []string to JSON before saving to DB
func (s StringSlice) Value() (driver.Value, error) {
	return json.Marshal(s)
}

// Convert JSON from DB back into Go []string
func (s *StringSlice) Scan(value interface{}) error {
	if value == nil {
		*s = []string{}
		return nil
	}

	str, ok := value.([]byte) // PostgreSQL stores JSON as byte array
	if !ok {
		return errors.New("invalid type for StringSlice")
	}

	return json.Unmarshal(str, s)
}
