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

// checks if the WishlistSharingRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WishlistSharingRequest{}

// WishlistSharingRequest struct for WishlistSharingRequest
type WishlistSharingRequest struct {
	WishlistId            *string             `json:"wishlistId,omitempty"`
	Permission            *WishlistPermission `json:"permission,omitempty"`
	CustomerGrn           *string             `json:"customerGrn,omitempty"`
	CustomerAggregationId *string             `json:"customerAggregationId,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _WishlistSharingRequest WishlistSharingRequest

// NewWishlistSharingRequest instantiates a new WishlistSharingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWishlistSharingRequest() *WishlistSharingRequest {
	this := WishlistSharingRequest{}
	var permission WishlistPermission = WISHLISTPERMISSION_UNKNOWN_PERMISSION
	this.Permission = &permission
	return &this
}

// NewWishlistSharingRequestWithDefaults instantiates a new WishlistSharingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWishlistSharingRequestWithDefaults() *WishlistSharingRequest {
	this := WishlistSharingRequest{}
	var permission WishlistPermission = WISHLISTPERMISSION_UNKNOWN_PERMISSION
	this.Permission = &permission
	return &this
}

// GetWishlistId returns the WishlistId field value if set, zero value otherwise.
func (o *WishlistSharingRequest) GetWishlistId() string {
	if o == nil || IsNil(o.WishlistId) {
		var ret string
		return ret
	}
	return *o.WishlistId
}

// GetWishlistIdOk returns a tuple with the WishlistId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistSharingRequest) GetWishlistIdOk() (*string, bool) {
	if o == nil || IsNil(o.WishlistId) {
		return nil, false
	}
	return o.WishlistId, true
}

// HasWishlistId returns a boolean if a field has been set.
func (o *WishlistSharingRequest) HasWishlistId() bool {
	if o != nil && !IsNil(o.WishlistId) {
		return true
	}

	return false
}

// SetWishlistId gets a reference to the given string and assigns it to the WishlistId field.
func (o *WishlistSharingRequest) SetWishlistId(v string) {
	o.WishlistId = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *WishlistSharingRequest) GetPermission() WishlistPermission {
	if o == nil || IsNil(o.Permission) {
		var ret WishlistPermission
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistSharingRequest) GetPermissionOk() (*WishlistPermission, bool) {
	if o == nil || IsNil(o.Permission) {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *WishlistSharingRequest) HasPermission() bool {
	if o != nil && !IsNil(o.Permission) {
		return true
	}

	return false
}

// SetPermission gets a reference to the given WishlistPermission and assigns it to the Permission field.
func (o *WishlistSharingRequest) SetPermission(v WishlistPermission) {
	o.Permission = &v
}

// GetCustomerGrn returns the CustomerGrn field value if set, zero value otherwise.
func (o *WishlistSharingRequest) GetCustomerGrn() string {
	if o == nil || IsNil(o.CustomerGrn) {
		var ret string
		return ret
	}
	return *o.CustomerGrn
}

// GetCustomerGrnOk returns a tuple with the CustomerGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistSharingRequest) GetCustomerGrnOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerGrn) {
		return nil, false
	}
	return o.CustomerGrn, true
}

// HasCustomerGrn returns a boolean if a field has been set.
func (o *WishlistSharingRequest) HasCustomerGrn() bool {
	if o != nil && !IsNil(o.CustomerGrn) {
		return true
	}

	return false
}

// SetCustomerGrn gets a reference to the given string and assigns it to the CustomerGrn field.
func (o *WishlistSharingRequest) SetCustomerGrn(v string) {
	o.CustomerGrn = &v
}

// GetCustomerAggregationId returns the CustomerAggregationId field value if set, zero value otherwise.
func (o *WishlistSharingRequest) GetCustomerAggregationId() string {
	if o == nil || IsNil(o.CustomerAggregationId) {
		var ret string
		return ret
	}
	return *o.CustomerAggregationId
}

// GetCustomerAggregationIdOk returns a tuple with the CustomerAggregationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistSharingRequest) GetCustomerAggregationIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerAggregationId) {
		return nil, false
	}
	return o.CustomerAggregationId, true
}

// HasCustomerAggregationId returns a boolean if a field has been set.
func (o *WishlistSharingRequest) HasCustomerAggregationId() bool {
	if o != nil && !IsNil(o.CustomerAggregationId) {
		return true
	}

	return false
}

// SetCustomerAggregationId gets a reference to the given string and assigns it to the CustomerAggregationId field.
func (o *WishlistSharingRequest) SetCustomerAggregationId(v string) {
	o.CustomerAggregationId = &v
}

func (o WishlistSharingRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WishlistSharingRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WishlistId) {
		toSerialize["wishlistId"] = o.WishlistId
	}
	if !IsNil(o.Permission) {
		toSerialize["permission"] = o.Permission
	}
	if !IsNil(o.CustomerGrn) {
		toSerialize["customerGrn"] = o.CustomerGrn
	}
	if !IsNil(o.CustomerAggregationId) {
		toSerialize["customerAggregationId"] = o.CustomerAggregationId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WishlistSharingRequest) UnmarshalJSON(data []byte) (err error) {
	varWishlistSharingRequest := _WishlistSharingRequest{}

	err = json.Unmarshal(data, &varWishlistSharingRequest)

	if err != nil {
		return err
	}

	*o = WishlistSharingRequest(varWishlistSharingRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "wishlistId")
		delete(additionalProperties, "permission")
		delete(additionalProperties, "customerGrn")
		delete(additionalProperties, "customerAggregationId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *WishlistSharingRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *WishlistSharingRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableWishlistSharingRequest struct {
	value *WishlistSharingRequest
	isSet bool
}

func (v NullableWishlistSharingRequest) Get() *WishlistSharingRequest {
	return v.value
}

func (v *NullableWishlistSharingRequest) Set(val *WishlistSharingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWishlistSharingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWishlistSharingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWishlistSharingRequest(val *WishlistSharingRequest) *NullableWishlistSharingRequest {
	return &NullableWishlistSharingRequest{value: val, isSet: true}
}

func (v NullableWishlistSharingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWishlistSharingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
