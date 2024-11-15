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

// checks if the WishlistAreItemsInWishlistsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WishlistAreItemsInWishlistsRequest{}

// WishlistAreItemsInWishlistsRequest struct for WishlistAreItemsInWishlistsRequest
type WishlistAreItemsInWishlistsRequest struct {
	TenantId string `json:"tenantId"`
	WishlistId *string `json:"wishlistId,omitempty"`
	CustomerGrn *string `json:"customerGrn,omitempty"`
	ItemGrns []string `json:"itemGrns"`
	AdditionalProperties map[string]interface{}
}

type _WishlistAreItemsInWishlistsRequest WishlistAreItemsInWishlistsRequest

// NewWishlistAreItemsInWishlistsRequest instantiates a new WishlistAreItemsInWishlistsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWishlistAreItemsInWishlistsRequest(tenantId string, itemGrns []string) *WishlistAreItemsInWishlistsRequest {
	this := WishlistAreItemsInWishlistsRequest{}
	this.TenantId = tenantId
	this.ItemGrns = itemGrns
	return &this
}

// NewWishlistAreItemsInWishlistsRequestWithDefaults instantiates a new WishlistAreItemsInWishlistsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWishlistAreItemsInWishlistsRequestWithDefaults() *WishlistAreItemsInWishlistsRequest {
	this := WishlistAreItemsInWishlistsRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *WishlistAreItemsInWishlistsRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *WishlistAreItemsInWishlistsRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *WishlistAreItemsInWishlistsRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetWishlistId returns the WishlistId field value if set, zero value otherwise.
func (o *WishlistAreItemsInWishlistsRequest) GetWishlistId() string {
	if o == nil || IsNil(o.WishlistId) {
		var ret string
		return ret
	}
	return *o.WishlistId
}

// GetWishlistIdOk returns a tuple with the WishlistId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistAreItemsInWishlistsRequest) GetWishlistIdOk() (*string, bool) {
	if o == nil || IsNil(o.WishlistId) {
		return nil, false
	}
	return o.WishlistId, true
}

// HasWishlistId returns a boolean if a field has been set.
func (o *WishlistAreItemsInWishlistsRequest) HasWishlistId() bool {
	if o != nil && !IsNil(o.WishlistId) {
		return true
	}

	return false
}

// SetWishlistId gets a reference to the given string and assigns it to the WishlistId field.
func (o *WishlistAreItemsInWishlistsRequest) SetWishlistId(v string) {
	o.WishlistId = &v
}

// GetCustomerGrn returns the CustomerGrn field value if set, zero value otherwise.
func (o *WishlistAreItemsInWishlistsRequest) GetCustomerGrn() string {
	if o == nil || IsNil(o.CustomerGrn) {
		var ret string
		return ret
	}
	return *o.CustomerGrn
}

// GetCustomerGrnOk returns a tuple with the CustomerGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistAreItemsInWishlistsRequest) GetCustomerGrnOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerGrn) {
		return nil, false
	}
	return o.CustomerGrn, true
}

// HasCustomerGrn returns a boolean if a field has been set.
func (o *WishlistAreItemsInWishlistsRequest) HasCustomerGrn() bool {
	if o != nil && !IsNil(o.CustomerGrn) {
		return true
	}

	return false
}

// SetCustomerGrn gets a reference to the given string and assigns it to the CustomerGrn field.
func (o *WishlistAreItemsInWishlistsRequest) SetCustomerGrn(v string) {
	o.CustomerGrn = &v
}

// GetItemGrns returns the ItemGrns field value
func (o *WishlistAreItemsInWishlistsRequest) GetItemGrns() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ItemGrns
}

// GetItemGrnsOk returns a tuple with the ItemGrns field value
// and a boolean to check if the value has been set.
func (o *WishlistAreItemsInWishlistsRequest) GetItemGrnsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ItemGrns, true
}

// SetItemGrns sets field value
func (o *WishlistAreItemsInWishlistsRequest) SetItemGrns(v []string) {
	o.ItemGrns = v
}

func (o WishlistAreItemsInWishlistsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WishlistAreItemsInWishlistsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	if !IsNil(o.WishlistId) {
		toSerialize["wishlistId"] = o.WishlistId
	}
	if !IsNil(o.CustomerGrn) {
		toSerialize["customerGrn"] = o.CustomerGrn
	}
	toSerialize["itemGrns"] = o.ItemGrns

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WishlistAreItemsInWishlistsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"itemGrns",
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

	varWishlistAreItemsInWishlistsRequest := _WishlistAreItemsInWishlistsRequest{}

	err = json.Unmarshal(data, &varWishlistAreItemsInWishlistsRequest)

	if err != nil {
		return err
	}

	*o = WishlistAreItemsInWishlistsRequest(varWishlistAreItemsInWishlistsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "wishlistId")
		delete(additionalProperties, "customerGrn")
		delete(additionalProperties, "itemGrns")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *WishlistAreItemsInWishlistsRequest) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *WishlistAreItemsInWishlistsRequest) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableWishlistAreItemsInWishlistsRequest struct {
	value *WishlistAreItemsInWishlistsRequest
	isSet bool
}

func (v NullableWishlistAreItemsInWishlistsRequest) Get() *WishlistAreItemsInWishlistsRequest {
	return v.value
}

func (v *NullableWishlistAreItemsInWishlistsRequest) Set(val *WishlistAreItemsInWishlistsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWishlistAreItemsInWishlistsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWishlistAreItemsInWishlistsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWishlistAreItemsInWishlistsRequest(val *WishlistAreItemsInWishlistsRequest) *NullableWishlistAreItemsInWishlistsRequest {
	return &NullableWishlistAreItemsInWishlistsRequest{value: val, isSet: true}
}

func (v NullableWishlistAreItemsInWishlistsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWishlistAreItemsInWishlistsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


