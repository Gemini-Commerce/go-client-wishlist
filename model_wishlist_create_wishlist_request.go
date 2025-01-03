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

// checks if the WishlistCreateWishlistRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WishlistCreateWishlistRequest{}

// WishlistCreateWishlistRequest struct for WishlistCreateWishlistRequest
type WishlistCreateWishlistRequest struct {
	TenantId    string                 `json:"tenantId"`
	Privacy     WishlistPrivacy        `json:"privacy"`
	Label       *WishlistLocalizedText `json:"label,omitempty"`
	Description *WishlistLocalizedText `json:"description,omitempty"`
	// If the customer GRN is set on JWT, it will be used as default. Otherwise, it will be used the customer_grn field.
	CustomerGrn          *string `json:"customerGrn,omitempty"`
	IsDefault            *bool   `json:"isDefault,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WishlistCreateWishlistRequest WishlistCreateWishlistRequest

// NewWishlistCreateWishlistRequest instantiates a new WishlistCreateWishlistRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWishlistCreateWishlistRequest(tenantId string, privacy WishlistPrivacy) *WishlistCreateWishlistRequest {
	this := WishlistCreateWishlistRequest{}
	this.TenantId = tenantId
	this.Privacy = privacy
	return &this
}

// NewWishlistCreateWishlistRequestWithDefaults instantiates a new WishlistCreateWishlistRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWishlistCreateWishlistRequestWithDefaults() *WishlistCreateWishlistRequest {
	this := WishlistCreateWishlistRequest{}
	var privacy WishlistPrivacy = WISHLISTPRIVACY_UNKNOWN
	this.Privacy = privacy
	return &this
}

// GetTenantId returns the TenantId field value
func (o *WishlistCreateWishlistRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *WishlistCreateWishlistRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *WishlistCreateWishlistRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetPrivacy returns the Privacy field value
func (o *WishlistCreateWishlistRequest) GetPrivacy() WishlistPrivacy {
	if o == nil {
		var ret WishlistPrivacy
		return ret
	}

	return o.Privacy
}

// GetPrivacyOk returns a tuple with the Privacy field value
// and a boolean to check if the value has been set.
func (o *WishlistCreateWishlistRequest) GetPrivacyOk() (*WishlistPrivacy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Privacy, true
}

// SetPrivacy sets field value
func (o *WishlistCreateWishlistRequest) SetPrivacy(v WishlistPrivacy) {
	o.Privacy = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WishlistCreateWishlistRequest) GetLabel() WishlistLocalizedText {
	if o == nil || IsNil(o.Label) {
		var ret WishlistLocalizedText
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistCreateWishlistRequest) GetLabelOk() (*WishlistLocalizedText, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WishlistCreateWishlistRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given WishlistLocalizedText and assigns it to the Label field.
func (o *WishlistCreateWishlistRequest) SetLabel(v WishlistLocalizedText) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WishlistCreateWishlistRequest) GetDescription() WishlistLocalizedText {
	if o == nil || IsNil(o.Description) {
		var ret WishlistLocalizedText
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistCreateWishlistRequest) GetDescriptionOk() (*WishlistLocalizedText, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WishlistCreateWishlistRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given WishlistLocalizedText and assigns it to the Description field.
func (o *WishlistCreateWishlistRequest) SetDescription(v WishlistLocalizedText) {
	o.Description = &v
}

// GetCustomerGrn returns the CustomerGrn field value if set, zero value otherwise.
func (o *WishlistCreateWishlistRequest) GetCustomerGrn() string {
	if o == nil || IsNil(o.CustomerGrn) {
		var ret string
		return ret
	}
	return *o.CustomerGrn
}

// GetCustomerGrnOk returns a tuple with the CustomerGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistCreateWishlistRequest) GetCustomerGrnOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerGrn) {
		return nil, false
	}
	return o.CustomerGrn, true
}

// HasCustomerGrn returns a boolean if a field has been set.
func (o *WishlistCreateWishlistRequest) HasCustomerGrn() bool {
	if o != nil && !IsNil(o.CustomerGrn) {
		return true
	}

	return false
}

// SetCustomerGrn gets a reference to the given string and assigns it to the CustomerGrn field.
func (o *WishlistCreateWishlistRequest) SetCustomerGrn(v string) {
	o.CustomerGrn = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *WishlistCreateWishlistRequest) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistCreateWishlistRequest) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *WishlistCreateWishlistRequest) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *WishlistCreateWishlistRequest) SetIsDefault(v bool) {
	o.IsDefault = &v
}

func (o WishlistCreateWishlistRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WishlistCreateWishlistRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["privacy"] = o.Privacy
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.CustomerGrn) {
		toSerialize["customerGrn"] = o.CustomerGrn
	}
	if !IsNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WishlistCreateWishlistRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"privacy",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varWishlistCreateWishlistRequest := _WishlistCreateWishlistRequest{}

	err = json.Unmarshal(data, &varWishlistCreateWishlistRequest)

	if err != nil {
		return err
	}

	*o = WishlistCreateWishlistRequest(varWishlistCreateWishlistRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "privacy")
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		delete(additionalProperties, "customerGrn")
		delete(additionalProperties, "isDefault")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *WishlistCreateWishlistRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *WishlistCreateWishlistRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableWishlistCreateWishlistRequest struct {
	value *WishlistCreateWishlistRequest
	isSet bool
}

func (v NullableWishlistCreateWishlistRequest) Get() *WishlistCreateWishlistRequest {
	return v.value
}

func (v *NullableWishlistCreateWishlistRequest) Set(val *WishlistCreateWishlistRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWishlistCreateWishlistRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWishlistCreateWishlistRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWishlistCreateWishlistRequest(val *WishlistCreateWishlistRequest) *NullableWishlistCreateWishlistRequest {
	return &NullableWishlistCreateWishlistRequest{value: val, isSet: true}
}

func (v NullableWishlistCreateWishlistRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWishlistCreateWishlistRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
