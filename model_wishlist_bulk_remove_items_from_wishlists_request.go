/*
wishlist/service.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package wishlist

import (
	"encoding/json"
)

// checks if the WishlistBulkRemoveItemsFromWishlistsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WishlistBulkRemoveItemsFromWishlistsRequest{}

// WishlistBulkRemoveItemsFromWishlistsRequest struct for WishlistBulkRemoveItemsFromWishlistsRequest
type WishlistBulkRemoveItemsFromWishlistsRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	// Item GRNs to remove from wishlists. Max 500 items per request.
	ItemGrns []string `json:"itemGrns,omitempty"`
	// Wishlist IDs to remove items from. Max 500 wishlists per request. If not provided, items will be removed from all wishlists.
	WishlistIds []string `json:"wishlistIds,omitempty"`
}

// NewWishlistBulkRemoveItemsFromWishlistsRequest instantiates a new WishlistBulkRemoveItemsFromWishlistsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWishlistBulkRemoveItemsFromWishlistsRequest() *WishlistBulkRemoveItemsFromWishlistsRequest {
	this := WishlistBulkRemoveItemsFromWishlistsRequest{}
	return &this
}

// NewWishlistBulkRemoveItemsFromWishlistsRequestWithDefaults instantiates a new WishlistBulkRemoveItemsFromWishlistsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWishlistBulkRemoveItemsFromWishlistsRequestWithDefaults() *WishlistBulkRemoveItemsFromWishlistsRequest {
	this := WishlistBulkRemoveItemsFromWishlistsRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *WishlistBulkRemoveItemsFromWishlistsRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistBulkRemoveItemsFromWishlistsRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *WishlistBulkRemoveItemsFromWishlistsRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *WishlistBulkRemoveItemsFromWishlistsRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetItemGrns returns the ItemGrns field value if set, zero value otherwise.
func (o *WishlistBulkRemoveItemsFromWishlistsRequest) GetItemGrns() []string {
	if o == nil || IsNil(o.ItemGrns) {
		var ret []string
		return ret
	}
	return o.ItemGrns
}

// GetItemGrnsOk returns a tuple with the ItemGrns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistBulkRemoveItemsFromWishlistsRequest) GetItemGrnsOk() ([]string, bool) {
	if o == nil || IsNil(o.ItemGrns) {
		return nil, false
	}
	return o.ItemGrns, true
}

// HasItemGrns returns a boolean if a field has been set.
func (o *WishlistBulkRemoveItemsFromWishlistsRequest) HasItemGrns() bool {
	if o != nil && !IsNil(o.ItemGrns) {
		return true
	}

	return false
}

// SetItemGrns gets a reference to the given []string and assigns it to the ItemGrns field.
func (o *WishlistBulkRemoveItemsFromWishlistsRequest) SetItemGrns(v []string) {
	o.ItemGrns = v
}

// GetWishlistIds returns the WishlistIds field value if set, zero value otherwise.
func (o *WishlistBulkRemoveItemsFromWishlistsRequest) GetWishlistIds() []string {
	if o == nil || IsNil(o.WishlistIds) {
		var ret []string
		return ret
	}
	return o.WishlistIds
}

// GetWishlistIdsOk returns a tuple with the WishlistIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistBulkRemoveItemsFromWishlistsRequest) GetWishlistIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.WishlistIds) {
		return nil, false
	}
	return o.WishlistIds, true
}

// HasWishlistIds returns a boolean if a field has been set.
func (o *WishlistBulkRemoveItemsFromWishlistsRequest) HasWishlistIds() bool {
	if o != nil && !IsNil(o.WishlistIds) {
		return true
	}

	return false
}

// SetWishlistIds gets a reference to the given []string and assigns it to the WishlistIds field.
func (o *WishlistBulkRemoveItemsFromWishlistsRequest) SetWishlistIds(v []string) {
	o.WishlistIds = v
}

func (o WishlistBulkRemoveItemsFromWishlistsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WishlistBulkRemoveItemsFromWishlistsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.ItemGrns) {
		toSerialize["itemGrns"] = o.ItemGrns
	}
	if !IsNil(o.WishlistIds) {
		toSerialize["wishlistIds"] = o.WishlistIds
	}
	return toSerialize, nil
}

type NullableWishlistBulkRemoveItemsFromWishlistsRequest struct {
	value *WishlistBulkRemoveItemsFromWishlistsRequest
	isSet bool
}

func (v NullableWishlistBulkRemoveItemsFromWishlistsRequest) Get() *WishlistBulkRemoveItemsFromWishlistsRequest {
	return v.value
}

func (v *NullableWishlistBulkRemoveItemsFromWishlistsRequest) Set(val *WishlistBulkRemoveItemsFromWishlistsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWishlistBulkRemoveItemsFromWishlistsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWishlistBulkRemoveItemsFromWishlistsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWishlistBulkRemoveItemsFromWishlistsRequest(val *WishlistBulkRemoveItemsFromWishlistsRequest) *NullableWishlistBulkRemoveItemsFromWishlistsRequest {
	return &NullableWishlistBulkRemoveItemsFromWishlistsRequest{value: val, isSet: true}
}

func (v NullableWishlistBulkRemoveItemsFromWishlistsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWishlistBulkRemoveItemsFromWishlistsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


