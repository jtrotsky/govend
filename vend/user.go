// Package vend handles interactions with the Vend API.
package vend

import "time"

// UserPayload contains sales data and versioning info.
type UserPayload struct {
	Data    []User           `json:"data,omitempty"`
	Version map[string]int64 `json:"version,omitempty"`
}

// User is a basic user object.
type User struct {
	ID          *string    `json:"id,omitempty"`
	Username    *string    `json:"username,omitempty"`
	DisplayName *string    `json:"display_name,omitempty"`
	Email       *string    `json:"email,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

/*
ENDPOINT:
.vendhq.com/api/2.0/users

EXAMPLE PAYLOAD:
{
  "id": "b8ca3a65-011c-11e4-fbb5-5973b0ee0c0d",
  "username": "joe@vendhq.com",
  "display_name": "Big J",
  "email": "joe@vendhq.com",
  "restricted_outlet_id": null,
  "account_type": "admin",
  "image_source": "https://honestmulch.vendhq.com/images/placeholder/customer/no-image-white-ss.png",
  "is_primary_user": true,
  "created_at": "2014-10-21T22:43:35+00:00",
  "updated_at": "2015-09-07T20:38:34+00:00",
  "deleted_at": null,
  "version": 238078
}
*/
