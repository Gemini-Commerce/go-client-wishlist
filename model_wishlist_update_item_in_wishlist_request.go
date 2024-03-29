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

// checks if the WishlistUpdateItemInWishlistRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WishlistUpdateItemInWishlistRequest{}

// WishlistUpdateItemInWishlistRequest struct for WishlistUpdateItemInWishlistRequest
type WishlistUpdateItemInWishlistRequest struct {
	TenantId string `json:"tenantId"`
	Id string `json:"id"`
	Payload WishlistUpdateItemInWishlistRequestPayload `json:"payload"`
	PayloadMask string `json:"payloadMask"`
}

type _WishlistUpdateItemInWishlistRequest WishlistUpdateItemInWishlistRequest

// NewWishlistUpdateItemInWishlistRequest instantiates a new WishlistUpdateItemInWishlistRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWishlistUpdateItemInWishlistRequest(tenantId string, id string, payload WishlistUpdateItemInWishlistRequestPayload, payloadMask string) *WishlistUpdateItemInWishlistRequest {
	this := WishlistUpdateItemInWishlistRequest{}
	this.TenantId = tenantId
	this.Id = id
	this.Payload = payload
	this.PayloadMask = payloadMask
	return &this
}

// NewWishlistUpdateItemInWishlistRequestWithDefaults instantiates a new WishlistUpdateItemInWishlistRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWishlistUpdateItemInWishlistRequestWithDefaults() *WishlistUpdateItemInWishlistRequest {
	this := WishlistUpdateItemInWishlistRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *WishlistUpdateItemInWishlistRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *WishlistUpdateItemInWishlistRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *WishlistUpdateItemInWishlistRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetId returns the Id field value
func (o *WishlistUpdateItemInWishlistRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WishlistUpdateItemInWishlistRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WishlistUpdateItemInWishlistRequest) SetId(v string) {
	o.Id = v
}

// GetPayload returns the Payload field value
func (o *WishlistUpdateItemInWishlistRequest) GetPayload() WishlistUpdateItemInWishlistRequestPayload {
	if o == nil {
		var ret WishlistUpdateItemInWishlistRequestPayload
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *WishlistUpdateItemInWishlistRequest) GetPayloadOk() (*WishlistUpdateItemInWishlistRequestPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *WishlistUpdateItemInWishlistRequest) SetPayload(v WishlistUpdateItemInWishlistRequestPayload) {
	o.Payload = v
}

// GetPayloadMask returns the PayloadMask field value
func (o *WishlistUpdateItemInWishlistRequest) GetPayloadMask() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayloadMask
}

// GetPayloadMaskOk returns a tuple with the PayloadMask field value
// and a boolean to check if the value has been set.
func (o *WishlistUpdateItemInWishlistRequest) GetPayloadMaskOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayloadMask, true
}

// SetPayloadMask sets field value
func (o *WishlistUpdateItemInWishlistRequest) SetPayloadMask(v string) {
	o.PayloadMask = v
}

func (o WishlistUpdateItemInWishlistRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WishlistUpdateItemInWishlistRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["id"] = o.Id
	toSerialize["payload"] = o.Payload
	toSerialize["payloadMask"] = o.PayloadMask
	return toSerialize, nil
}

func (o *WishlistUpdateItemInWishlistRequest) UnmarshalJSON(data []byte) (err error) {
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

	varWishlistUpdateItemInWishlistRequest := _WishlistUpdateItemInWishlistRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWishlistUpdateItemInWishlistRequest)

	if err != nil {
		return err
	}

	*o = WishlistUpdateItemInWishlistRequest(varWishlistUpdateItemInWishlistRequest)

	return err
}

type NullableWishlistUpdateItemInWishlistRequest struct {
	value *WishlistUpdateItemInWishlistRequest
	isSet bool
}

func (v NullableWishlistUpdateItemInWishlistRequest) Get() *WishlistUpdateItemInWishlistRequest {
	return v.value
}

func (v *NullableWishlistUpdateItemInWishlistRequest) Set(val *WishlistUpdateItemInWishlistRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWishlistUpdateItemInWishlistRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWishlistUpdateItemInWishlistRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWishlistUpdateItemInWishlistRequest(val *WishlistUpdateItemInWishlistRequest) *NullableWishlistUpdateItemInWishlistRequest {
	return &NullableWishlistUpdateItemInWishlistRequest{value: val, isSet: true}
}

func (v NullableWishlistUpdateItemInWishlistRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWishlistUpdateItemInWishlistRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


