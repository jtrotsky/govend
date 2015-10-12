// Package main handles interactions with the Vend API.
package main

import "time"

// Customer is a customer object.
type Customer struct {
	ID             *string    `json:"id,omitempty"`
	Code           *string    `json:"customer_code,omitempty"`
	FirstName      *string    `json:"first_name,omitempty"`
	LastName       *string    `json:"last_name,omitempty"`
	Email          *string    `json:"email,omitempty"`
	YearToDate     *float64   `json:"year_to_date,omitempty"`
	Balance        *float64   `json:"balance,omitempty"`
	LoyaltyBalance *float64   `json:"loyalty_balance,omitempty"`
	Note           *string    `json:"note,omitempty"`
	Gender         *string    `json:"gender,omitempty"`
	DateOfBirth    *string    `json:"date_of_birth,omitempty"`
	CompanyName    *string    `json:"company_name,omitempty"`
	Phone          *string    `json:"phone,omitempty"`
	Mobile         *string    `json:"mobile,omitempty"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty"`
	Version        *int64     `json:"version,omitempty"`
}

/*
ENDPOINT:
.vendhq.com/api/2.0/customers

EXAMPLE PAYLOAD:
{
  "id": "b8ca3a65-011c-11e4-fbb5-5973b0cb6b3f",
  "customer_code": "WALKIN",
  "first_name": null,
  "last_name": null,
  "email": null,
  "year_to_date": 0,
  "balance": 0,
  "loyalty_balance": 0,
  "note": null,
  "gender": null,
  "date_of_birth": null,
  "company_name": null,
  "phone": null,
  "mobile": null,
  "fax": null,
  "twitter": null,
  "website": null,
  "physical_suburb": null,
  "physical_city": null,
  "physical_postcode": null,
  "physical_state": null,
  "postal_suburb": null,
  "postal_city": null,
  "postal_state": null,
  "customer_group_id": "b8ca3a65-011c-11e4-fbb5-5973b0cb3205",
  "enable_loyalty": false,
  "created_at": "2014-10-21T22:43:35+00:00",
  "updated_at": "2014-11-19T11:42:51+00:00",
  "deleted_at": null,
  "version": 11882051,
  "postal_postcode": null,
  "name": null,
  "physical_address_1": null,
  "physical_address_2": null,
  "physical_country_id": null,
  "postal_address_1": null,
  "postal_address_2": null,
  "postal_country_id": null,
  "custom_field_1": null,
  "custom_field_2": null,
  "custom_field_3": null,
  "custom_field_4": null
}
*/
