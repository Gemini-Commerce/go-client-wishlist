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

// checks if the WishlistMergeWishlistsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WishlistMergeWishlistsRequest{}

// WishlistMergeWishlistsRequest struct for WishlistMergeWishlistsRequest
type WishlistMergeWishlistsRequest struct {
	TenantId string `json:"tenantId"`
	StartingWishlistId string `json:"startingWishlistId"`
	// If the customer GRN is set on JWT, it will be used as default. Otherwise, it will be used the customer_grn field.
	CustomerGrn string `json:"customerGrn"`
	// The wishlist to merge into the starting wishlist.
	TargetWishlistId *string `json:"targetWishlistId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WishlistMergeWishlistsRequest WishlistMergeWishlistsRequest

// NewWishlistMergeWishlistsRequest instantiates a new WishlistMergeWishlistsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWishlistMergeWishlistsRequest(tenantId string, startingWishlistId string, customerGrn string) *WishlistMergeWishlistsRequest {
	this := WishlistMergeWishlistsRequest{}
	this.TenantId = tenantId
	this.StartingWishlistId = startingWishlistId
	this.CustomerGrn = customerGrn
	return &this
}

// NewWishlistMergeWishlistsRequestWithDefaults instantiates a new WishlistMergeWishlistsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWishlistMergeWishlistsRequestWithDefaults() *WishlistMergeWishlistsRequest {
	this := WishlistMergeWishlistsRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *WishlistMergeWishlistsRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *WishlistMergeWishlistsRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *WishlistMergeWishlistsRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetStartingWishlistId returns the StartingWishlistId field value
func (o *WishlistMergeWishlistsRequest) GetStartingWishlistId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartingWishlistId
}

// GetStartingWishlistIdOk returns a tuple with the StartingWishlistId field value
// and a boolean to check if the value has been set.
func (o *WishlistMergeWishlistsRequest) GetStartingWishlistIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartingWishlistId, true
}

// SetStartingWishlistId sets field value
func (o *WishlistMergeWishlistsRequest) SetStartingWishlistId(v string) {
	o.StartingWishlistId = v
}

// GetCustomerGrn returns the CustomerGrn field value
func (o *WishlistMergeWishlistsRequest) GetCustomerGrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerGrn
}

// GetCustomerGrnOk returns a tuple with the CustomerGrn field value
// and a boolean to check if the value has been set.
func (o *WishlistMergeWishlistsRequest) GetCustomerGrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerGrn, true
}

// SetCustomerGrn sets field value
func (o *WishlistMergeWishlistsRequest) SetCustomerGrn(v string) {
	o.CustomerGrn = v
}

// GetTargetWishlistId returns the TargetWishlistId field value if set, zero value otherwise.
func (o *WishlistMergeWishlistsRequest) GetTargetWishlistId() string {
	if o == nil || IsNil(o.TargetWishlistId) {
		var ret string
		return ret
	}
	return *o.TargetWishlistId
}

// GetTargetWishlistIdOk returns a tuple with the TargetWishlistId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistMergeWishlistsRequest) GetTargetWishlistIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetWishlistId) {
		return nil, false
	}
	return o.TargetWishlistId, true
}

// &#39;Has&#39;TargetWishlistId returns a boolean if a field has been set.
func (o *WishlistMergeWishlistsRequest) &#39;Has&#39;TargetWishlistId() bool {
	if o != nil && !IsNil(o.TargetWishlistId) {
		return true
	}

	return false
}

// SetTargetWishlistId gets a reference to the given string and assigns it to the TargetWishlistId field.
func (o *WishlistMergeWishlistsRequest) SetTargetWishlistId(v string) {
	o.TargetWishlistId = &v
}

func (o WishlistMergeWishlistsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WishlistMergeWishlistsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["startingWishlistId"] = o.StartingWishlistId
	toSerialize["customerGrn"] = o.CustomerGrn
	if !IsNil(o.TargetWishlistId) {
		toSerialize["targetWishlistId"] = o.TargetWishlistId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WishlistMergeWishlistsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"startingWishlistId",
		"customerGrn",
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

	varWishlistMergeWishlistsRequest := _WishlistMergeWishlistsRequest{}

	err = json.Unmarshal(data, &varWishlistMergeWishlistsRequest)

	if err != nil {
		return err
	}

	*o = WishlistMergeWishlistsRequest(varWishlistMergeWishlistsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "startingWishlistId")
		delete(additionalProperties, "customerGrn")
		delete(additionalProperties, "targetWishlistId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *WishlistMergeWishlistsRequest) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *WishlistMergeWishlistsRequest) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableWishlistMergeWishlistsRequest struct {
	value *WishlistMergeWishlistsRequest
	isSet bool
}

func (v NullableWishlistMergeWishlistsRequest) Get() *WishlistMergeWishlistsRequest {
	return v.value
}

func (v *NullableWishlistMergeWishlistsRequest) Set(val *WishlistMergeWishlistsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWishlistMergeWishlistsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWishlistMergeWishlistsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWishlistMergeWishlistsRequest(val *WishlistMergeWishlistsRequest) *NullableWishlistMergeWishlistsRequest {
	return &NullableWishlistMergeWishlistsRequest{value: val, isSet: true}
}

func (v NullableWishlistMergeWishlistsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWishlistMergeWishlistsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


