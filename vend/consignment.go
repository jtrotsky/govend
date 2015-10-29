// Package vend handles interactions with the Vend API.
package vend

import "time"

// ConsignmentPayload contains register data and versioning info.
type ConsignmentPayload struct {
	Data    []Consignment    `json:"data,omitempty"`
	Version map[string]int64 `json:"version,omitempty"`
}

// Consignment is a register object.
type Consignment struct {
	ID              *string    `json:"id,omitempty"`
	OutletID        *string    `json:"outlet_id,omitempty"`
	Name            *string    `json:"name,omitempty"`
	Type            *string    `json:"type,omitempty"`
	Status          *string    `json:"status,omitempty"`
	ConsignmentDate *string    `json:"consignment_date,omitempty"` // NOTE: Using string for ParseVendDT.
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`
}

/*
ENDPOINT:
.vendhq.com/api/2.0/consignments

EXAMPLE PAYLOAD:
{
  "id": "b8ca3a65-0125-11e4-fbb5-9ab19b8222e2",
  "outlet_id": "b8ca3a65-0125-11e4-fbb5-9aab33ef0df8",
  "name": "Order - Tue 13 Jan 2015",
  "due_at": null,
  "type": "SUPPLIER",
  "status": "RECEIVED",
  "supplier_id": null,
  "source_outlet_id": null,
  "consignment_date": "2015-01-12T23:20:33+00:00",
  "received_at": "2015-01-12T23:23:18+00:00",
  "show_inactive": false,
  "supplier_invoice": "",
  "total_count_gain": null,
  "total_cost_gain": null,
  "total_count_loss": null,
  "total_cost_loss": null,
  "created_at": "2015-01-12T23:20:33+00:00",
  "updated_at": "2015-01-12T23:23:18+00:00",
  "deleted_at": null,
  "version": 57967,
  "filters": []
}
*/
