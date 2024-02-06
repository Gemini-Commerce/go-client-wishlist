/*
wishlist/service.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package wishlist

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the WishlistRemoveItemFromWishlistRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WishlistRemoveItemFromWishlistRequest{}

// WishlistRemoveItemFromWishlistRequest struct for WishlistRemoveItemFromWishlistRequest
type WishlistRemoveItemFromWishlistRequest struct {
	TenantId string `json:"tenantId"`
	Id string `json:"id"`
}

type _WishlistRemoveItemFromWishlistRequest WishlistRemoveItemFromWishlistRequest

// NewWishlistRemoveItemFromWishlistRequest instantiates a new WishlistRemoveItemFromWishlistRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWishlistRemoveItemFromWishlistRequest(tenantId string, id string) *WishlistRemoveItemFromWishlistRequest {
	this := WishlistRemoveItemFromWishlistRequest{}
	this.TenantId = tenantId
	this.Id = id
	return &this
}

// NewWishlistRemoveItemFromWishlistRequestWithDefaults instantiates a new WishlistRemoveItemFromWishlistRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWishlistRemoveItemFromWishlistRequestWithDefaults() *WishlistRemoveItemFromWishlistRequest {
	this := WishlistRemoveItemFromWishlistRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *WishlistRemoveItemFromWishlistRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *WishlistRemoveItemFromWishlistRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *WishlistRemoveItemFromWishlistRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetId returns the Id field value
func (o *WishlistRemoveItemFromWishlistRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WishlistRemoveItemFromWishlistRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WishlistRemoveItemFromWishlistRequest) SetId(v string) {
	o.Id = v
}

func (o WishlistRemoveItemFromWishlistRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WishlistRemoveItemFromWishlistRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *WishlistRemoveItemFromWishlistRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varWishlistRemoveItemFromWishlistRequest := _WishlistRemoveItemFromWishlistRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWishlistRemoveItemFromWishlistRequest)

	if err != nil {
		return err
	}

	*o = WishlistRemoveItemFromWishlistRequest(varWishlistRemoveItemFromWishlistRequest)

	return err
}

type NullableWishlistRemoveItemFromWishlistRequest struct {
	value *WishlistRemoveItemFromWishlistRequest
	isSet bool
}

func (v NullableWishlistRemoveItemFromWishlistRequest) Get() *WishlistRemoveItemFromWishlistRequest {
	return v.value
}

func (v *NullableWishlistRemoveItemFromWishlistRequest) Set(val *WishlistRemoveItemFromWishlistRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWishlistRemoveItemFromWishlistRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWishlistRemoveItemFromWishlistRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWishlistRemoveItemFromWishlistRequest(val *WishlistRemoveItemFromWishlistRequest) *NullableWishlistRemoveItemFromWishlistRequest {
	return &NullableWishlistRemoveItemFromWishlistRequest{value: val, isSet: true}
}

func (v NullableWishlistRemoveItemFromWishlistRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWishlistRemoveItemFromWishlistRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

