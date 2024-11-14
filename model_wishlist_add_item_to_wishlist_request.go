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

// checks if the WishlistAddItemToWishlistRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WishlistAddItemToWishlistRequest{}

// WishlistAddItemToWishlistRequest struct for WishlistAddItemToWishlistRequest
type WishlistAddItemToWishlistRequest struct {
	TenantId string `json:"tenantId"`
	WishlistId string `json:"wishlistId"`
	ItemGrn string `json:"itemGrn"`
	PreferredQuantity *string `json:"preferredQuantity,omitempty"`
	Description *WishlistLocalizedText `json:"description,omitempty"`
	CustomerGrn *string `json:"customerGrn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WishlistAddItemToWishlistRequest WishlistAddItemToWishlistRequest

// NewWishlistAddItemToWishlistRequest instantiates a new WishlistAddItemToWishlistRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWishlistAddItemToWishlistRequest(tenantId string, wishlistId string, itemGrn string) *WishlistAddItemToWishlistRequest {
	this := WishlistAddItemToWishlistRequest{}
	this.TenantId = tenantId
	this.WishlistId = wishlistId
	this.ItemGrn = itemGrn
	return &this
}

// NewWishlistAddItemToWishlistRequestWithDefaults instantiates a new WishlistAddItemToWishlistRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWishlistAddItemToWishlistRequestWithDefaults() *WishlistAddItemToWishlistRequest {
	this := WishlistAddItemToWishlistRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *WishlistAddItemToWishlistRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *WishlistAddItemToWishlistRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *WishlistAddItemToWishlistRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetWishlistId returns the WishlistId field value
func (o *WishlistAddItemToWishlistRequest) GetWishlistId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WishlistId
}

// GetWishlistIdOk returns a tuple with the WishlistId field value
// and a boolean to check if the value has been set.
func (o *WishlistAddItemToWishlistRequest) GetWishlistIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WishlistId, true
}

// SetWishlistId sets field value
func (o *WishlistAddItemToWishlistRequest) SetWishlistId(v string) {
	o.WishlistId = v
}

// GetItemGrn returns the ItemGrn field value
func (o *WishlistAddItemToWishlistRequest) GetItemGrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemGrn
}

// GetItemGrnOk returns a tuple with the ItemGrn field value
// and a boolean to check if the value has been set.
func (o *WishlistAddItemToWishlistRequest) GetItemGrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemGrn, true
}

// SetItemGrn sets field value
func (o *WishlistAddItemToWishlistRequest) SetItemGrn(v string) {
	o.ItemGrn = v
}

// GetPreferredQuantity returns the PreferredQuantity field value if set, zero value otherwise.
func (o *WishlistAddItemToWishlistRequest) GetPreferredQuantity() string {
	if o == nil || IsNil(o.PreferredQuantity) {
		var ret string
		return ret
	}
	return *o.PreferredQuantity
}

// GetPreferredQuantityOk returns a tuple with the PreferredQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistAddItemToWishlistRequest) GetPreferredQuantityOk() (*string, bool) {
	if o == nil || IsNil(o.PreferredQuantity) {
		return nil, false
	}
	return o.PreferredQuantity, true
}

// &#39;Has&#39;PreferredQuantity returns a boolean if a field has been set.
func (o *WishlistAddItemToWishlistRequest) &#39;Has&#39;PreferredQuantity() bool {
	if o != nil && !IsNil(o.PreferredQuantity) {
		return true
	}

	return false
}

// SetPreferredQuantity gets a reference to the given string and assigns it to the PreferredQuantity field.
func (o *WishlistAddItemToWishlistRequest) SetPreferredQuantity(v string) {
	o.PreferredQuantity = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WishlistAddItemToWishlistRequest) GetDescription() WishlistLocalizedText {
	if o == nil || IsNil(o.Description) {
		var ret WishlistLocalizedText
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistAddItemToWishlistRequest) GetDescriptionOk() (*WishlistLocalizedText, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// &#39;Has&#39;Description returns a boolean if a field has been set.
func (o *WishlistAddItemToWishlistRequest) &#39;Has&#39;Description() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given WishlistLocalizedText and assigns it to the Description field.
func (o *WishlistAddItemToWishlistRequest) SetDescription(v WishlistLocalizedText) {
	o.Description = &v
}

// GetCustomerGrn returns the CustomerGrn field value if set, zero value otherwise.
func (o *WishlistAddItemToWishlistRequest) GetCustomerGrn() string {
	if o == nil || IsNil(o.CustomerGrn) {
		var ret string
		return ret
	}
	return *o.CustomerGrn
}

// GetCustomerGrnOk returns a tuple with the CustomerGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistAddItemToWishlistRequest) GetCustomerGrnOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerGrn) {
		return nil, false
	}
	return o.CustomerGrn, true
}

// &#39;Has&#39;CustomerGrn returns a boolean if a field has been set.
func (o *WishlistAddItemToWishlistRequest) &#39;Has&#39;CustomerGrn() bool {
	if o != nil && !IsNil(o.CustomerGrn) {
		return true
	}

	return false
}

// SetCustomerGrn gets a reference to the given string and assigns it to the CustomerGrn field.
func (o *WishlistAddItemToWishlistRequest) SetCustomerGrn(v string) {
	o.CustomerGrn = &v
}

func (o WishlistAddItemToWishlistRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WishlistAddItemToWishlistRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["wishlistId"] = o.WishlistId
	toSerialize["itemGrn"] = o.ItemGrn
	if !IsNil(o.PreferredQuantity) {
		toSerialize["preferredQuantity"] = o.PreferredQuantity
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.CustomerGrn) {
		toSerialize["customerGrn"] = o.CustomerGrn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WishlistAddItemToWishlistRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"wishlistId",
		"itemGrn",
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

	varWishlistAddItemToWishlistRequest := _WishlistAddItemToWishlistRequest{}

	err = json.Unmarshal(data, &varWishlistAddItemToWishlistRequest)

	if err != nil {
		return err
	}

	*o = WishlistAddItemToWishlistRequest(varWishlistAddItemToWishlistRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "wishlistId")
		delete(additionalProperties, "itemGrn")
		delete(additionalProperties, "preferredQuantity")
		delete(additionalProperties, "description")
		delete(additionalProperties, "customerGrn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *WishlistAddItemToWishlistRequest) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *WishlistAddItemToWishlistRequest) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableWishlistAddItemToWishlistRequest struct {
	value *WishlistAddItemToWishlistRequest
	isSet bool
}

func (v NullableWishlistAddItemToWishlistRequest) Get() *WishlistAddItemToWishlistRequest {
	return v.value
}

func (v *NullableWishlistAddItemToWishlistRequest) Set(val *WishlistAddItemToWishlistRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWishlistAddItemToWishlistRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWishlistAddItemToWishlistRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWishlistAddItemToWishlistRequest(val *WishlistAddItemToWishlistRequest) *NullableWishlistAddItemToWishlistRequest {
	return &NullableWishlistAddItemToWishlistRequest{value: val, isSet: true}
}

func (v NullableWishlistAddItemToWishlistRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWishlistAddItemToWishlistRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


