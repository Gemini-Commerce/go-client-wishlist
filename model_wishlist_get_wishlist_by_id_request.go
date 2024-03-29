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
	"bytes"
	"fmt"
)

// checks if the WishlistGetWishlistByIdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WishlistGetWishlistByIdRequest{}

// WishlistGetWishlistByIdRequest struct for WishlistGetWishlistByIdRequest
type WishlistGetWishlistByIdRequest struct {
	TenantId string `json:"tenantId"`
	Id string `json:"id"`
}

type _WishlistGetWishlistByIdRequest WishlistGetWishlistByIdRequest

// NewWishlistGetWishlistByIdRequest instantiates a new WishlistGetWishlistByIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWishlistGetWishlistByIdRequest(tenantId string, id string) *WishlistGetWishlistByIdRequest {
	this := WishlistGetWishlistByIdRequest{}
	this.TenantId = tenantId
	this.Id = id
	return &this
}

// NewWishlistGetWishlistByIdRequestWithDefaults instantiates a new WishlistGetWishlistByIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWishlistGetWishlistByIdRequestWithDefaults() *WishlistGetWishlistByIdRequest {
	this := WishlistGetWishlistByIdRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *WishlistGetWishlistByIdRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *WishlistGetWishlistByIdRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *WishlistGetWishlistByIdRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetId returns the Id field value
func (o *WishlistGetWishlistByIdRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WishlistGetWishlistByIdRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WishlistGetWishlistByIdRequest) SetId(v string) {
	o.Id = v
}

func (o WishlistGetWishlistByIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WishlistGetWishlistByIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *WishlistGetWishlistByIdRequest) UnmarshalJSON(data []byte) (err error) {
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

	varWishlistGetWishlistByIdRequest := _WishlistGetWishlistByIdRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWishlistGetWishlistByIdRequest)

	if err != nil {
		return err
	}

	*o = WishlistGetWishlistByIdRequest(varWishlistGetWishlistByIdRequest)

	return err
}

type NullableWishlistGetWishlistByIdRequest struct {
	value *WishlistGetWishlistByIdRequest
	isSet bool
}

func (v NullableWishlistGetWishlistByIdRequest) Get() *WishlistGetWishlistByIdRequest {
	return v.value
}

func (v *NullableWishlistGetWishlistByIdRequest) Set(val *WishlistGetWishlistByIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWishlistGetWishlistByIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWishlistGetWishlistByIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWishlistGetWishlistByIdRequest(val *WishlistGetWishlistByIdRequest) *NullableWishlistGetWishlistByIdRequest {
	return &NullableWishlistGetWishlistByIdRequest{value: val, isSet: true}
}

func (v NullableWishlistGetWishlistByIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWishlistGetWishlistByIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


