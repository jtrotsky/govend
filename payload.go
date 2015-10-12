// Package main handles interactions with the Vend API.
package main

import "encoding/json"

// Payload contains resource data and versioning info.
// This is the default format returned by 2.0 endpoints.
type Payload struct {
	Data    json.RawMessage  `json:"data,omitempty"`
	Version map[string]int64 `json:"version,omitempty"`
}
