// Package vend handles interactions with the Vend API.
package vend

import "time"

// ConsignmentProductPayload contains data and versioning info.
type ConsignmentProductPayload struct {
	Data    []ConsignmentProduct `json:"data,omitempty"`
	Version map[string]int64     `json:"version,omitempty"`
}

// ConsignmentProduct ...
type ConsignmentProduct struct {
	ProductID *string `json:"product_id,omitempty"`
	SKU       *string `json:"product_sku,omitempty"`
	Count     *string `json:"count_omitempty"`
	// Name      *string    `json:"name,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

/*
ENDPOINT:
.vendhq.com/api/2.0/consignments/<consignment_id>/products

EXAMPLE PAYLOAD:
{
  "product_id": "b8ca3a65-0125-11e4-fbb5-6fdf26b98b71",
  "product_sku": "200069",
  "count": "50.00000",
  "received": "45.00000",
  "cost": "1.50000",
  "is_included": true,
  "created_at": "2015-01-12T23:21:02+00:00",
  "updated_at": "2015-01-12T23:23:18+00:00",
  "deleted_at": null,
  "version": 1032394
}
*/
