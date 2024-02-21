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

// checks if the WishlistUpdateWishlistRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WishlistUpdateWishlistRequest{}

// WishlistUpdateWishlistRequest struct for WishlistUpdateWishlistRequest
type WishlistUpdateWishlistRequest struct {
	TenantId string `json:"tenantId"`
	Id string `json:"id"`
	Payload WishlistUpdateWishlistRequestPayload `json:"payload"`
	PayloadMask string `json:"payloadMask"`
}

type _WishlistUpdateWishlistRequest WishlistUpdateWishlistRequest

// NewWishlistUpdateWishlistRequest instantiates a new WishlistUpdateWishlistRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWishlistUpdateWishlistRequest(tenantId string, id string, payload WishlistUpdateWishlistRequestPayload, payloadMask string) *WishlistUpdateWishlistRequest {
	this := WishlistUpdateWishlistRequest{}
	this.TenantId = tenantId
	this.Id = id
	this.Payload = payload
	this.PayloadMask = payloadMask
	return &this
}

// NewWishlistUpdateWishlistRequestWithDefaults instantiates a new WishlistUpdateWishlistRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWishlistUpdateWishlistRequestWithDefaults() *WishlistUpdateWishlistRequest {
	this := WishlistUpdateWishlistRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *WishlistUpdateWishlistRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *WishlistUpdateWishlistRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *WishlistUpdateWishlistRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetId returns the Id field value
func (o *WishlistUpdateWishlistRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WishlistUpdateWishlistRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WishlistUpdateWishlistRequest) SetId(v string) {
	o.Id = v
}

// GetPayload returns the Payload field value
func (o *WishlistUpdateWishlistRequest) GetPayload() WishlistUpdateWishlistRequestPayload {
	if o == nil {
		var ret WishlistUpdateWishlistRequestPayload
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *WishlistUpdateWishlistRequest) GetPayloadOk() (*WishlistUpdateWishlistRequestPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *WishlistUpdateWishlistRequest) SetPayload(v WishlistUpdateWishlistRequestPayload) {
	o.Payload = v
}

// GetPayloadMask returns the PayloadMask field value
func (o *WishlistUpdateWishlistRequest) GetPayloadMask() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayloadMask
}

// GetPayloadMaskOk returns a tuple with the PayloadMask field value
// and a boolean to check if the value has been set.
func (o *WishlistUpdateWishlistRequest) GetPayloadMaskOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayloadMask, true
}

// SetPayloadMask sets field value
func (o *WishlistUpdateWishlistRequest) SetPayloadMask(v string) {
	o.PayloadMask = v
}

func (o WishlistUpdateWishlistRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WishlistUpdateWishlistRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["id"] = o.Id
	toSerialize["payload"] = o.Payload
	toSerialize["payloadMask"] = o.PayloadMask
	return toSerialize, nil
}

func (o *WishlistUpdateWishlistRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"id",
		"payload",
		"payloadMask",
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

	varWishlistUpdateWishlistRequest := _WishlistUpdateWishlistRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWishlistUpdateWishlistRequest)

	if err != nil {
		return err
	}

	*o = WishlistUpdateWishlistRequest(varWishlistUpdateWishlistRequest)

	return err
}

type NullableWishlistUpdateWishlistRequest struct {
	value *WishlistUpdateWishlistRequest
	isSet bool
}

func (v NullableWishlistUpdateWishlistRequest) Get() *WishlistUpdateWishlistRequest {
	return v.value
}

func (v *NullableWishlistUpdateWishlistRequest) Set(val *WishlistUpdateWishlistRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWishlistUpdateWishlistRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWishlistUpdateWishlistRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWishlistUpdateWishlistRequest(val *WishlistUpdateWishlistRequest) *NullableWishlistUpdateWishlistRequest {
	return &NullableWishlistUpdateWishlistRequest{value: val, isSet: true}
}

func (v NullableWishlistUpdateWishlistRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWishlistUpdateWishlistRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


