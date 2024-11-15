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

// checks if the WishlistListWishlistItemsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WishlistListWishlistItemsRequest{}

// WishlistListWishlistItemsRequest struct for WishlistListWishlistItemsRequest
type WishlistListWishlistItemsRequest struct {
	TenantId string `json:"tenantId"`
	WishlistId string `json:"wishlistId"`
	CustomerGrn *string `json:"customerGrn,omitempty"`
	PageSize *int64 `json:"pageSize,omitempty"`
	PageToken *string `json:"pageToken,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WishlistListWishlistItemsRequest WishlistListWishlistItemsRequest

// NewWishlistListWishlistItemsRequest instantiates a new WishlistListWishlistItemsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWishlistListWishlistItemsRequest(tenantId string, wishlistId string) *WishlistListWishlistItemsRequest {
	this := WishlistListWishlistItemsRequest{}
	this.TenantId = tenantId
	this.WishlistId = wishlistId
	return &this
}

// NewWishlistListWishlistItemsRequestWithDefaults instantiates a new WishlistListWishlistItemsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWishlistListWishlistItemsRequestWithDefaults() *WishlistListWishlistItemsRequest {
	this := WishlistListWishlistItemsRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *WishlistListWishlistItemsRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *WishlistListWishlistItemsRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *WishlistListWishlistItemsRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetWishlistId returns the WishlistId field value
func (o *WishlistListWishlistItemsRequest) GetWishlistId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WishlistId
}

// GetWishlistIdOk returns a tuple with the WishlistId field value
// and a boolean to check if the value has been set.
func (o *WishlistListWishlistItemsRequest) GetWishlistIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WishlistId, true
}

// SetWishlistId sets field value
func (o *WishlistListWishlistItemsRequest) SetWishlistId(v string) {
	o.WishlistId = v
}

// GetCustomerGrn returns the CustomerGrn field value if set, zero value otherwise.
func (o *WishlistListWishlistItemsRequest) GetCustomerGrn() string {
	if o == nil || IsNil(o.CustomerGrn) {
		var ret string
		return ret
	}
	return *o.CustomerGrn
}

// GetCustomerGrnOk returns a tuple with the CustomerGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistListWishlistItemsRequest) GetCustomerGrnOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerGrn) {
		return nil, false
	}
	return o.CustomerGrn, true
}

// HasCustomerGrn returns a boolean if a field has been set.
func (o *WishlistListWishlistItemsRequest) HasCustomerGrn() bool {
	if o != nil && !IsNil(o.CustomerGrn) {
		return true
	}

	return false
}

// SetCustomerGrn gets a reference to the given string and assigns it to the CustomerGrn field.
func (o *WishlistListWishlistItemsRequest) SetCustomerGrn(v string) {
	o.CustomerGrn = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *WishlistListWishlistItemsRequest) GetPageSize() int64 {
	if o == nil || IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistListWishlistItemsRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *WishlistListWishlistItemsRequest) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *WishlistListWishlistItemsRequest) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetPageToken returns the PageToken field value if set, zero value otherwise.
func (o *WishlistListWishlistItemsRequest) GetPageToken() string {
	if o == nil || IsNil(o.PageToken) {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistListWishlistItemsRequest) GetPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.PageToken) {
		return nil, false
	}
	return o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *WishlistListWishlistItemsRequest) HasPageToken() bool {
	if o != nil && !IsNil(o.PageToken) {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *WishlistListWishlistItemsRequest) SetPageToken(v string) {
	o.PageToken = &v
}

func (o WishlistListWishlistItemsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WishlistListWishlistItemsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["wishlistId"] = o.WishlistId
	if !IsNil(o.CustomerGrn) {
		toSerialize["customerGrn"] = o.CustomerGrn
	}
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !IsNil(o.PageToken) {
		toSerialize["pageToken"] = o.PageToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WishlistListWishlistItemsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"wishlistId",
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

	varWishlistListWishlistItemsRequest := _WishlistListWishlistItemsRequest{}

	err = json.Unmarshal(data, &varWishlistListWishlistItemsRequest)

	if err != nil {
		return err
	}

	*o = WishlistListWishlistItemsRequest(varWishlistListWishlistItemsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "wishlistId")
		delete(additionalProperties, "customerGrn")
		delete(additionalProperties, "pageSize")
		delete(additionalProperties, "pageToken")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *WishlistListWishlistItemsRequest) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *WishlistListWishlistItemsRequest) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableWishlistListWishlistItemsRequest struct {
	value *WishlistListWishlistItemsRequest
	isSet bool
}

func (v NullableWishlistListWishlistItemsRequest) Get() *WishlistListWishlistItemsRequest {
	return v.value
}

func (v *NullableWishlistListWishlistItemsRequest) Set(val *WishlistListWishlistItemsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWishlistListWishlistItemsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWishlistListWishlistItemsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWishlistListWishlistItemsRequest(val *WishlistListWishlistItemsRequest) *NullableWishlistListWishlistItemsRequest {
	return &NullableWishlistListWishlistItemsRequest{value: val, isSet: true}
}

func (v NullableWishlistListWishlistItemsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWishlistListWishlistItemsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


