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

// checks if the WishlistBulkCreateSharingResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WishlistBulkCreateSharingResponse{}

// WishlistBulkCreateSharingResponse struct for WishlistBulkCreateSharingResponse
type WishlistBulkCreateSharingResponse struct {
	SharingResponses     []WishlistSharingResponse `json:"sharingResponses,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WishlistBulkCreateSharingResponse WishlistBulkCreateSharingResponse

// NewWishlistBulkCreateSharingResponse instantiates a new WishlistBulkCreateSharingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWishlistBulkCreateSharingResponse() *WishlistBulkCreateSharingResponse {
	this := WishlistBulkCreateSharingResponse{}
	return &this
}

// NewWishlistBulkCreateSharingResponseWithDefaults instantiates a new WishlistBulkCreateSharingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWishlistBulkCreateSharingResponseWithDefaults() *WishlistBulkCreateSharingResponse {
	this := WishlistBulkCreateSharingResponse{}
	return &this
}

// GetSharingResponses returns the SharingResponses field value if set, zero value otherwise.
func (o *WishlistBulkCreateSharingResponse) GetSharingResponses() []WishlistSharingResponse {
	if o == nil || IsNil(o.SharingResponses) {
		var ret []WishlistSharingResponse
		return ret
	}
	return o.SharingResponses
}

// GetSharingResponsesOk returns a tuple with the SharingResponses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistBulkCreateSharingResponse) GetSharingResponsesOk() ([]WishlistSharingResponse, bool) {
	if o == nil || IsNil(o.SharingResponses) {
		return nil, false
	}
	return o.SharingResponses, true
}

// HasSharingResponses returns a boolean if a field has been set.
func (o *WishlistBulkCreateSharingResponse) HasSharingResponses() bool {
	if o != nil && !IsNil(o.SharingResponses) {
		return true
	}

	return false
}

// SetSharingResponses gets a reference to the given []WishlistSharingResponse and assigns it to the SharingResponses field.
func (o *WishlistBulkCreateSharingResponse) SetSharingResponses(v []WishlistSharingResponse) {
	o.SharingResponses = v
}

func (o WishlistBulkCreateSharingResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WishlistBulkCreateSharingResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SharingResponses) {
		toSerialize["sharingResponses"] = o.SharingResponses
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WishlistBulkCreateSharingResponse) UnmarshalJSON(data []byte) (err error) {
	varWishlistBulkCreateSharingResponse := _WishlistBulkCreateSharingResponse{}

	err = json.Unmarshal(data, &varWishlistBulkCreateSharingResponse)

	if err != nil {
		return err
	}

	*o = WishlistBulkCreateSharingResponse(varWishlistBulkCreateSharingResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "sharingResponses")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *WishlistBulkCreateSharingResponse) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *WishlistBulkCreateSharingResponse) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableWishlistBulkCreateSharingResponse struct {
	value *WishlistBulkCreateSharingResponse
	isSet bool
}

func (v NullableWishlistBulkCreateSharingResponse) Get() *WishlistBulkCreateSharingResponse {
	return v.value
}

func (v *NullableWishlistBulkCreateSharingResponse) Set(val *WishlistBulkCreateSharingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWishlistBulkCreateSharingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWishlistBulkCreateSharingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWishlistBulkCreateSharingResponse(val *WishlistBulkCreateSharingResponse) *NullableWishlistBulkCreateSharingResponse {
	return &NullableWishlistBulkCreateSharingResponse{value: val, isSet: true}
}

func (v NullableWishlistBulkCreateSharingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWishlistBulkCreateSharingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
