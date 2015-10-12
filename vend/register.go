// Package vend handles interactions with the Vend API.
package vend

import "time"

// RegisterPayload contains register data and versioning info.
type RegisterPayload struct {
	Data    []Register       `json:"data,omitempty"`
	Version map[string]int64 `json:"version,omitempty"`
}

// Register is a register object.
type Register struct {
	ID        *string    `json:"id,omitempty"`
	Name      *string    `json:"name,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

/*
ENDPOINT:
.vendhq.com/api/2.0/registers

EXAMPLE PAYLOAD:
{
  "id": "b8ca3a65-011c-11e4-fbb5-5973b0e372df",
  "name": "Main Register",
  "outlet_id": "b8ca3a65-011c-11e4-fbb5-5973b0e19f1a",
  "ask_for_note_on_save": 1,
  "print_note_on_receipt": false,
  "ask_for_user_on_sale": false,
  "show_discounts_on_receipts": true,
  "print_receipt": true,
  "email_receipt": true,
  "invoice_prefix": "",
  "invoice_suffix": "",
  "invoice_sequence": 99,
  "button_layout_id": "f804bca9-88c4-11e4-9bb5-b8ca3a65011c",
  "is_open": true,
  "deleted_at": null,
  "register_open_time": "2015-08-30T02:27:33+00:00",
  "register_close_time": null,
  "register_open_sequence_id": "e52b2846-e925-11e5-f98b-4ebeac20e500",
  "cash_managed_payment_type_id": null,
  "version": 3143163
}
*/
