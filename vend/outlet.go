// Package vend handles interactions with the Vend API.
package vend

import "time"

// OutletPayload contains outlet data and versioning info.
type OutletPayload struct {
	Data    []Outlet         `json:"data,omitempty"`
	Version map[string]int64 `json:"version,omitempty"`
}

// Outlet is usually a physical store location.
type Outlet struct {
	ID        *string    `json:"id,omitempty"`
	Name      *string    `json:"name,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

/*
ENDPOINT:
.vendhq.com/api/2.0/outlets

EXAMPLE PAYLOAD:
{
  "id": "b8ca3a65-011c-11e4-fbb5-5973b0e19f1a",
  "name": "Jim's Mulch Trolley",
  "default_tax_id": "b8ca3a65-011c-11e4-fbb5-5973b0d04415",
  "currency": "NZD",
  "currency_symbol": "$",
  "display_prices": "inclusive",
  "time_zone": "Pacific/Auckland",
  "deleted_at": null,
  "version": 437275
}
*/
