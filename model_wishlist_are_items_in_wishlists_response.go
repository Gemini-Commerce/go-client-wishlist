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
)

// checks if the WishlistAreItemsInWishlistsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WishlistAreItemsInWishlistsResponse{}

// WishlistAreItemsInWishlistsResponse struct for WishlistAreItemsInWishlistsResponse
type WishlistAreItemsInWishlistsResponse struct {
	ItemGrnMap *map[string]WishlistAreItemsInWishlistsResponsePayload `json:"itemGrnMap,omitempty"`
}

// NewWishlistAreItemsInWishlistsResponse instantiates a new WishlistAreItemsInWishlistsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWishlistAreItemsInWishlistsResponse() *WishlistAreItemsInWishlistsResponse {
	this := WishlistAreItemsInWishlistsResponse{}
	return &this
}

// NewWishlistAreItemsInWishlistsResponseWithDefaults instantiates a new WishlistAreItemsInWishlistsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWishlistAreItemsInWishlistsResponseWithDefaults() *WishlistAreItemsInWishlistsResponse {
	this := WishlistAreItemsInWishlistsResponse{}
	return &this
}

// GetItemGrnMap returns the ItemGrnMap field value if set, zero value otherwise.
func (o *WishlistAreItemsInWishlistsResponse) GetItemGrnMap() map[string]WishlistAreItemsInWishlistsResponsePayload {
	if o == nil || IsNil(o.ItemGrnMap) {
		var ret map[string]WishlistAreItemsInWishlistsResponsePayload
		return ret
	}
	return *o.ItemGrnMap
}

// GetItemGrnMapOk returns a tuple with the ItemGrnMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistAreItemsInWishlistsResponse) GetItemGrnMapOk() (*map[string]WishlistAreItemsInWishlistsResponsePayload, bool) {
	if o == nil || IsNil(o.ItemGrnMap) {
		return nil, false
	}
	return o.ItemGrnMap, true
}

// HasItemGrnMap returns a boolean if a field has been set.
func (o *WishlistAreItemsInWishlistsResponse) HasItemGrnMap() bool {
	if o != nil && !IsNil(o.ItemGrnMap) {
		return true
	}

	return false
}

// SetItemGrnMap gets a reference to the given map[string]WishlistAreItemsInWishlistsResponsePayload and assigns it to the ItemGrnMap field.
func (o *WishlistAreItemsInWishlistsResponse) SetItemGrnMap(v map[string]WishlistAreItemsInWishlistsResponsePayload) {
	o.ItemGrnMap = &v
}

func (o WishlistAreItemsInWishlistsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WishlistAreItemsInWishlistsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ItemGrnMap) {
		toSerialize["itemGrnMap"] = o.ItemGrnMap
	}
	return toSerialize, nil
}

type NullableWishlistAreItemsInWishlistsResponse struct {
	value *WishlistAreItemsInWishlistsResponse
	isSet bool
}

func (v NullableWishlistAreItemsInWishlistsResponse) Get() *WishlistAreItemsInWishlistsResponse {
	return v.value
}

func (v *NullableWishlistAreItemsInWishlistsResponse) Set(val *WishlistAreItemsInWishlistsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWishlistAreItemsInWishlistsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWishlistAreItemsInWishlistsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWishlistAreItemsInWishlistsResponse(val *WishlistAreItemsInWishlistsResponse) *NullableWishlistAreItemsInWishlistsResponse {
	return &NullableWishlistAreItemsInWishlistsResponse{value: val, isSet: true}
}

func (v NullableWishlistAreItemsInWishlistsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWishlistAreItemsInWishlistsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


