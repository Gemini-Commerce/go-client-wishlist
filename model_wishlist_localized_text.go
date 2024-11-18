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

// checks if the WishlistLocalizedText type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WishlistLocalizedText{}

// WishlistLocalizedText struct for WishlistLocalizedText
type WishlistLocalizedText struct {
	Value                *map[string]string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WishlistLocalizedText WishlistLocalizedText

// NewWishlistLocalizedText instantiates a new WishlistLocalizedText object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWishlistLocalizedText() *WishlistLocalizedText {
	this := WishlistLocalizedText{}
	return &this
}

// NewWishlistLocalizedTextWithDefaults instantiates a new WishlistLocalizedText object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWishlistLocalizedTextWithDefaults() *WishlistLocalizedText {
	this := WishlistLocalizedText{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *WishlistLocalizedText) GetValue() map[string]string {
	if o == nil || IsNil(o.Value) {
		var ret map[string]string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistLocalizedText) GetValueOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *WishlistLocalizedText) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]string and assigns it to the Value field.
func (o *WishlistLocalizedText) SetValue(v map[string]string) {
	o.Value = &v
}

func (o WishlistLocalizedText) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WishlistLocalizedText) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WishlistLocalizedText) UnmarshalJSON(data []byte) (err error) {
	varWishlistLocalizedText := _WishlistLocalizedText{}

	err = json.Unmarshal(data, &varWishlistLocalizedText)

	if err != nil {
		return err
	}

	*o = WishlistLocalizedText(varWishlistLocalizedText)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWishlistLocalizedText struct {
	value *WishlistLocalizedText
	isSet bool
}

func (v NullableWishlistLocalizedText) Get() *WishlistLocalizedText {
	return v.value
}

func (v *NullableWishlistLocalizedText) Set(val *WishlistLocalizedText) {
	v.value = val
	v.isSet = true
}

func (v NullableWishlistLocalizedText) IsSet() bool {
	return v.isSet
}

func (v *NullableWishlistLocalizedText) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWishlistLocalizedText(val *WishlistLocalizedText) *NullableWishlistLocalizedText {
	return &NullableWishlistLocalizedText{value: val, isSet: true}
}

func (v NullableWishlistLocalizedText) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWishlistLocalizedText) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
