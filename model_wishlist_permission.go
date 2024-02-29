/*
Wishlist Service

API for managing wishlists

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package wishlist

import (
	"encoding/json"
	"fmt"
)

// WishlistPermission the model 'WishlistPermission'
type WishlistPermission string

// List of wishlistPermission
const (
	UNKNOWN_PERMISSION WishlistPermission = "UNKNOWN_PERMISSION"
	VIEW_PERMISSION WishlistPermission = "VIEW_PERMISSION"
	EDIT_PERMISSION WishlistPermission = "EDIT_PERMISSION"
)

// All allowed values of WishlistPermission enum
var AllowedWishlistPermissionEnumValues = []WishlistPermission{
	"UNKNOWN_PERMISSION",
	"VIEW_PERMISSION",
	"EDIT_PERMISSION",
}

func (v *WishlistPermission) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WishlistPermission(value)
	for _, existing := range AllowedWishlistPermissionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WishlistPermission", value)
}

// NewWishlistPermissionFromValue returns a pointer to a valid WishlistPermission
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWishlistPermissionFromValue(v string) (*WishlistPermission, error) {
	ev := WishlistPermission(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WishlistPermission: valid values are %v", v, AllowedWishlistPermissionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WishlistPermission) IsValid() bool {
	for _, existing := range AllowedWishlistPermissionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to wishlistPermission value
func (v WishlistPermission) Ptr() *WishlistPermission {
	return &v
}

type NullableWishlistPermission struct {
	value *WishlistPermission
	isSet bool
}

func (v NullableWishlistPermission) Get() *WishlistPermission {
	return v.value
}

func (v *NullableWishlistPermission) Set(val *WishlistPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableWishlistPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableWishlistPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWishlistPermission(val *WishlistPermission) *NullableWishlistPermission {
	return &NullableWishlistPermission{value: val, isSet: true}
}

func (v NullableWishlistPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWishlistPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

