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

// checks if the WishlistListWishlistsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WishlistListWishlistsRequest{}

// WishlistListWishlistsRequest struct for WishlistListWishlistsRequest
type WishlistListWishlistsRequest struct {
	TenantId string `json:"tenantId"`
	PageSize *int64 `json:"pageSize,omitempty"`
	PageToken *string `json:"pageToken,omitempty"`
	Filter *ListWishlistsRequestFilter `json:"filter,omitempty"`
	FilterMask *string `json:"filterMask,omitempty"`
}

type _WishlistListWishlistsRequest WishlistListWishlistsRequest

// NewWishlistListWishlistsRequest instantiates a new WishlistListWishlistsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWishlistListWishlistsRequest(tenantId string) *WishlistListWishlistsRequest {
	this := WishlistListWishlistsRequest{}
	this.TenantId = tenantId
	return &this
}

// NewWishlistListWishlistsRequestWithDefaults instantiates a new WishlistListWishlistsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWishlistListWishlistsRequestWithDefaults() *WishlistListWishlistsRequest {
	this := WishlistListWishlistsRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *WishlistListWishlistsRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *WishlistListWishlistsRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *WishlistListWishlistsRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *WishlistListWishlistsRequest) GetPageSize() int64 {
	if o == nil || IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistListWishlistsRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *WishlistListWishlistsRequest) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *WishlistListWishlistsRequest) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetPageToken returns the PageToken field value if set, zero value otherwise.
func (o *WishlistListWishlistsRequest) GetPageToken() string {
	if o == nil || IsNil(o.PageToken) {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistListWishlistsRequest) GetPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.PageToken) {
		return nil, false
	}
	return o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *WishlistListWishlistsRequest) HasPageToken() bool {
	if o != nil && !IsNil(o.PageToken) {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *WishlistListWishlistsRequest) SetPageToken(v string) {
	o.PageToken = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *WishlistListWishlistsRequest) GetFilter() ListWishlistsRequestFilter {
	if o == nil || IsNil(o.Filter) {
		var ret ListWishlistsRequestFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistListWishlistsRequest) GetFilterOk() (*ListWishlistsRequestFilter, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *WishlistListWishlistsRequest) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given ListWishlistsRequestFilter and assigns it to the Filter field.
func (o *WishlistListWishlistsRequest) SetFilter(v ListWishlistsRequestFilter) {
	o.Filter = &v
}

// GetFilterMask returns the FilterMask field value if set, zero value otherwise.
func (o *WishlistListWishlistsRequest) GetFilterMask() string {
	if o == nil || IsNil(o.FilterMask) {
		var ret string
		return ret
	}
	return *o.FilterMask
}

// GetFilterMaskOk returns a tuple with the FilterMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistListWishlistsRequest) GetFilterMaskOk() (*string, bool) {
	if o == nil || IsNil(o.FilterMask) {
		return nil, false
	}
	return o.FilterMask, true
}

// HasFilterMask returns a boolean if a field has been set.
func (o *WishlistListWishlistsRequest) HasFilterMask() bool {
	if o != nil && !IsNil(o.FilterMask) {
		return true
	}

	return false
}

// SetFilterMask gets a reference to the given string and assigns it to the FilterMask field.
func (o *WishlistListWishlistsRequest) SetFilterMask(v string) {
	o.FilterMask = &v
}

func (o WishlistListWishlistsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WishlistListWishlistsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !IsNil(o.PageToken) {
		toSerialize["pageToken"] = o.PageToken
	}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.FilterMask) {
		toSerialize["filterMask"] = o.FilterMask
	}
	return toSerialize, nil
}

func (o *WishlistListWishlistsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
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

	varWishlistListWishlistsRequest := _WishlistListWishlistsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWishlistListWishlistsRequest)

	if err != nil {
		return err
	}

	*o = WishlistListWishlistsRequest(varWishlistListWishlistsRequest)

	return err
}

type NullableWishlistListWishlistsRequest struct {
	value *WishlistListWishlistsRequest
	isSet bool
}

func (v NullableWishlistListWishlistsRequest) Get() *WishlistListWishlistsRequest {
	return v.value
}

func (v *NullableWishlistListWishlistsRequest) Set(val *WishlistListWishlistsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWishlistListWishlistsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWishlistListWishlistsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWishlistListWishlistsRequest(val *WishlistListWishlistsRequest) *NullableWishlistListWishlistsRequest {
	return &NullableWishlistListWishlistsRequest{value: val, isSet: true}
}

func (v NullableWishlistListWishlistsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWishlistListWishlistsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


