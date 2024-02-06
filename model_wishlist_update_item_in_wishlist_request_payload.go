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

// checks if the WishlistUpdateItemInWishlistRequestPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WishlistUpdateItemInWishlistRequestPayload{}

// WishlistUpdateItemInWishlistRequestPayload struct for WishlistUpdateItemInWishlistRequestPayload
type WishlistUpdateItemInWishlistRequestPayload struct {
	PreferredQuantity *string `json:"preferredQuantity,omitempty"`
	Description *WishlistLocalizedText `json:"description,omitempty"`
	CustomerGrn *string `json:"customerGrn,omitempty"`
}

// NewWishlistUpdateItemInWishlistRequestPayload instantiates a new WishlistUpdateItemInWishlistRequestPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWishlistUpdateItemInWishlistRequestPayload() *WishlistUpdateItemInWishlistRequestPayload {
	this := WishlistUpdateItemInWishlistRequestPayload{}
	return &this
}

// NewWishlistUpdateItemInWishlistRequestPayloadWithDefaults instantiates a new WishlistUpdateItemInWishlistRequestPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWishlistUpdateItemInWishlistRequestPayloadWithDefaults() *WishlistUpdateItemInWishlistRequestPayload {
	this := WishlistUpdateItemInWishlistRequestPayload{}
	return &this
}

// GetPreferredQuantity returns the PreferredQuantity field value if set, zero value otherwise.
func (o *WishlistUpdateItemInWishlistRequestPayload) GetPreferredQuantity() string {
	if o == nil || IsNil(o.PreferredQuantity) {
		var ret string
		return ret
	}
	return *o.PreferredQuantity
}

// GetPreferredQuantityOk returns a tuple with the PreferredQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistUpdateItemInWishlistRequestPayload) GetPreferredQuantityOk() (*string, bool) {
	if o == nil || IsNil(o.PreferredQuantity) {
		return nil, false
	}
	return o.PreferredQuantity, true
}

// HasPreferredQuantity returns a boolean if a field has been set.
func (o *WishlistUpdateItemInWishlistRequestPayload) HasPreferredQuantity() bool {
	if o != nil && !IsNil(o.PreferredQuantity) {
		return true
	}

	return false
}

// SetPreferredQuantity gets a reference to the given string and assigns it to the PreferredQuantity field.
func (o *WishlistUpdateItemInWishlistRequestPayload) SetPreferredQuantity(v string) {
	o.PreferredQuantity = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WishlistUpdateItemInWishlistRequestPayload) GetDescription() WishlistLocalizedText {
	if o == nil || IsNil(o.Description) {
		var ret WishlistLocalizedText
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistUpdateItemInWishlistRequestPayload) GetDescriptionOk() (*WishlistLocalizedText, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WishlistUpdateItemInWishlistRequestPayload) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given WishlistLocalizedText and assigns it to the Description field.
func (o *WishlistUpdateItemInWishlistRequestPayload) SetDescription(v WishlistLocalizedText) {
	o.Description = &v
}

// GetCustomerGrn returns the CustomerGrn field value if set, zero value otherwise.
func (o *WishlistUpdateItemInWishlistRequestPayload) GetCustomerGrn() string {
	if o == nil || IsNil(o.CustomerGrn) {
		var ret string
		return ret
	}
	return *o.CustomerGrn
}

// GetCustomerGrnOk returns a tuple with the CustomerGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistUpdateItemInWishlistRequestPayload) GetCustomerGrnOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerGrn) {
		return nil, false
	}
	return o.CustomerGrn, true
}

// HasCustomerGrn returns a boolean if a field has been set.
func (o *WishlistUpdateItemInWishlistRequestPayload) HasCustomerGrn() bool {
	if o != nil && !IsNil(o.CustomerGrn) {
		return true
	}

	return false
}

// SetCustomerGrn gets a reference to the given string and assigns it to the CustomerGrn field.
func (o *WishlistUpdateItemInWishlistRequestPayload) SetCustomerGrn(v string) {
	o.CustomerGrn = &v
}

func (o WishlistUpdateItemInWishlistRequestPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WishlistUpdateItemInWishlistRequestPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PreferredQuantity) {
		toSerialize["preferredQuantity"] = o.PreferredQuantity
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.CustomerGrn) {
		toSerialize["customerGrn"] = o.CustomerGrn
	}
	return toSerialize, nil
}

type NullableWishlistUpdateItemInWishlistRequestPayload struct {
	value *WishlistUpdateItemInWishlistRequestPayload
	isSet bool
}

func (v NullableWishlistUpdateItemInWishlistRequestPayload) Get() *WishlistUpdateItemInWishlistRequestPayload {
	return v.value
}

func (v *NullableWishlistUpdateItemInWishlistRequestPayload) Set(val *WishlistUpdateItemInWishlistRequestPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableWishlistUpdateItemInWishlistRequestPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableWishlistUpdateItemInWishlistRequestPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWishlistUpdateItemInWishlistRequestPayload(val *WishlistUpdateItemInWishlistRequestPayload) *NullableWishlistUpdateItemInWishlistRequestPayload {
	return &NullableWishlistUpdateItemInWishlistRequestPayload{value: val, isSet: true}
}

func (v NullableWishlistUpdateItemInWishlistRequestPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWishlistUpdateItemInWishlistRequestPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

