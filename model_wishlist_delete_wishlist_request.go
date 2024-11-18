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

// checks if the WishlistDeleteWishlistRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WishlistDeleteWishlistRequest{}

// WishlistDeleteWishlistRequest struct for WishlistDeleteWishlistRequest
type WishlistDeleteWishlistRequest struct {
	TenantId             string `json:"tenantId"`
	Id                   string `json:"id"`
	AdditionalProperties map[string]interface{}
}

type _WishlistDeleteWishlistRequest WishlistDeleteWishlistRequest

// NewWishlistDeleteWishlistRequest instantiates a new WishlistDeleteWishlistRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWishlistDeleteWishlistRequest(tenantId string, id string) *WishlistDeleteWishlistRequest {
	this := WishlistDeleteWishlistRequest{}
	this.TenantId = tenantId
	this.Id = id
	return &this
}

// NewWishlistDeleteWishlistRequestWithDefaults instantiates a new WishlistDeleteWishlistRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWishlistDeleteWishlistRequestWithDefaults() *WishlistDeleteWishlistRequest {
	this := WishlistDeleteWishlistRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *WishlistDeleteWishlistRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *WishlistDeleteWishlistRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *WishlistDeleteWishlistRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetId returns the Id field value
func (o *WishlistDeleteWishlistRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WishlistDeleteWishlistRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WishlistDeleteWishlistRequest) SetId(v string) {
	o.Id = v
}

func (o WishlistDeleteWishlistRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WishlistDeleteWishlistRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["id"] = o.Id

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WishlistDeleteWishlistRequest) UnmarshalJSON(data []byte) (err error) {
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
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varWishlistDeleteWishlistRequest := _WishlistDeleteWishlistRequest{}

	err = json.Unmarshal(data, &varWishlistDeleteWishlistRequest)

	if err != nil {
		return err
	}

	*o = WishlistDeleteWishlistRequest(varWishlistDeleteWishlistRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *WishlistDeleteWishlistRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *WishlistDeleteWishlistRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableWishlistDeleteWishlistRequest struct {
	value *WishlistDeleteWishlistRequest
	isSet bool
}

func (v NullableWishlistDeleteWishlistRequest) Get() *WishlistDeleteWishlistRequest {
	return v.value
}

func (v *NullableWishlistDeleteWishlistRequest) Set(val *WishlistDeleteWishlistRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWishlistDeleteWishlistRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWishlistDeleteWishlistRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWishlistDeleteWishlistRequest(val *WishlistDeleteWishlistRequest) *NullableWishlistDeleteWishlistRequest {
	return &NullableWishlistDeleteWishlistRequest{value: val, isSet: true}
}

func (v NullableWishlistDeleteWishlistRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWishlistDeleteWishlistRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
