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

// checks if the WishlistGetWishlistBySharedCodeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WishlistGetWishlistBySharedCodeRequest{}

// WishlistGetWishlistBySharedCodeRequest struct for WishlistGetWishlistBySharedCodeRequest
type WishlistGetWishlistBySharedCodeRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	SharedCode *string `json:"sharedCode,omitempty"`
	// If the customer GRN is set on JWT, it will be used as default. Otherwise, it will be used the customer_grn field.
	CustomerGrn *string `json:"customerGrn,omitempty"`
}

// NewWishlistGetWishlistBySharedCodeRequest instantiates a new WishlistGetWishlistBySharedCodeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWishlistGetWishlistBySharedCodeRequest() *WishlistGetWishlistBySharedCodeRequest {
	this := WishlistGetWishlistBySharedCodeRequest{}
	return &this
}

// NewWishlistGetWishlistBySharedCodeRequestWithDefaults instantiates a new WishlistGetWishlistBySharedCodeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWishlistGetWishlistBySharedCodeRequestWithDefaults() *WishlistGetWishlistBySharedCodeRequest {
	this := WishlistGetWishlistBySharedCodeRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *WishlistGetWishlistBySharedCodeRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistGetWishlistBySharedCodeRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *WishlistGetWishlistBySharedCodeRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *WishlistGetWishlistBySharedCodeRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetSharedCode returns the SharedCode field value if set, zero value otherwise.
func (o *WishlistGetWishlistBySharedCodeRequest) GetSharedCode() string {
	if o == nil || IsNil(o.SharedCode) {
		var ret string
		return ret
	}
	return *o.SharedCode
}

// GetSharedCodeOk returns a tuple with the SharedCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistGetWishlistBySharedCodeRequest) GetSharedCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SharedCode) {
		return nil, false
	}
	return o.SharedCode, true
}

// HasSharedCode returns a boolean if a field has been set.
func (o *WishlistGetWishlistBySharedCodeRequest) HasSharedCode() bool {
	if o != nil && !IsNil(o.SharedCode) {
		return true
	}

	return false
}

// SetSharedCode gets a reference to the given string and assigns it to the SharedCode field.
func (o *WishlistGetWishlistBySharedCodeRequest) SetSharedCode(v string) {
	o.SharedCode = &v
}

// GetCustomerGrn returns the CustomerGrn field value if set, zero value otherwise.
func (o *WishlistGetWishlistBySharedCodeRequest) GetCustomerGrn() string {
	if o == nil || IsNil(o.CustomerGrn) {
		var ret string
		return ret
	}
	return *o.CustomerGrn
}

// GetCustomerGrnOk returns a tuple with the CustomerGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistGetWishlistBySharedCodeRequest) GetCustomerGrnOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerGrn) {
		return nil, false
	}
	return o.CustomerGrn, true
}

// HasCustomerGrn returns a boolean if a field has been set.
func (o *WishlistGetWishlistBySharedCodeRequest) HasCustomerGrn() bool {
	if o != nil && !IsNil(o.CustomerGrn) {
		return true
	}

	return false
}

// SetCustomerGrn gets a reference to the given string and assigns it to the CustomerGrn field.
func (o *WishlistGetWishlistBySharedCodeRequest) SetCustomerGrn(v string) {
	o.CustomerGrn = &v
}

func (o WishlistGetWishlistBySharedCodeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WishlistGetWishlistBySharedCodeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.SharedCode) {
		toSerialize["sharedCode"] = o.SharedCode
	}
	if !IsNil(o.CustomerGrn) {
		toSerialize["customerGrn"] = o.CustomerGrn
	}
	return toSerialize, nil
}

type NullableWishlistGetWishlistBySharedCodeRequest struct {
	value *WishlistGetWishlistBySharedCodeRequest
	isSet bool
}

func (v NullableWishlistGetWishlistBySharedCodeRequest) Get() *WishlistGetWishlistBySharedCodeRequest {
	return v.value
}

func (v *NullableWishlistGetWishlistBySharedCodeRequest) Set(val *WishlistGetWishlistBySharedCodeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWishlistGetWishlistBySharedCodeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWishlistGetWishlistBySharedCodeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWishlistGetWishlistBySharedCodeRequest(val *WishlistGetWishlistBySharedCodeRequest) *NullableWishlistGetWishlistBySharedCodeRequest {
	return &NullableWishlistGetWishlistBySharedCodeRequest{value: val, isSet: true}
}

func (v NullableWishlistGetWishlistBySharedCodeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWishlistGetWishlistBySharedCodeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


