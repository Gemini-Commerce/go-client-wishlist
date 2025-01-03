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
	"time"
)

// checks if the WishlistWishlistResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WishlistWishlistResponse{}

// WishlistWishlistResponse struct for WishlistWishlistResponse
type WishlistWishlistResponse struct {
	Id                   *string                `json:"id,omitempty"`
	Grn                  *string                `json:"grn,omitempty"`
	SharedCode           *string                `json:"sharedCode,omitempty"`
	Privacy              *WishlistPrivacy       `json:"privacy,omitempty"`
	Label                *WishlistLocalizedText `json:"label,omitempty"`
	Description          *WishlistLocalizedText `json:"description,omitempty"`
	CustomerGrn          *string                `json:"customerGrn,omitempty"`
	IsDefault            *bool                  `json:"isDefault,omitempty"`
	ItemsCount           *string                `json:"itemsCount,omitempty"`
	CreatedAt            *time.Time             `json:"createdAt,omitempty"`
	UpdatedAt            *time.Time             `json:"updatedAt,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WishlistWishlistResponse WishlistWishlistResponse

// NewWishlistWishlistResponse instantiates a new WishlistWishlistResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWishlistWishlistResponse() *WishlistWishlistResponse {
	this := WishlistWishlistResponse{}
	var privacy WishlistPrivacy = WISHLISTPRIVACY_UNKNOWN
	this.Privacy = &privacy
	return &this
}

// NewWishlistWishlistResponseWithDefaults instantiates a new WishlistWishlistResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWishlistWishlistResponseWithDefaults() *WishlistWishlistResponse {
	this := WishlistWishlistResponse{}
	var privacy WishlistPrivacy = WISHLISTPRIVACY_UNKNOWN
	this.Privacy = &privacy
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WishlistWishlistResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistWishlistResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WishlistWishlistResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WishlistWishlistResponse) SetId(v string) {
	o.Id = &v
}

// GetGrn returns the Grn field value if set, zero value otherwise.
func (o *WishlistWishlistResponse) GetGrn() string {
	if o == nil || IsNil(o.Grn) {
		var ret string
		return ret
	}
	return *o.Grn
}

// GetGrnOk returns a tuple with the Grn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistWishlistResponse) GetGrnOk() (*string, bool) {
	if o == nil || IsNil(o.Grn) {
		return nil, false
	}
	return o.Grn, true
}

// HasGrn returns a boolean if a field has been set.
func (o *WishlistWishlistResponse) HasGrn() bool {
	if o != nil && !IsNil(o.Grn) {
		return true
	}

	return false
}

// SetGrn gets a reference to the given string and assigns it to the Grn field.
func (o *WishlistWishlistResponse) SetGrn(v string) {
	o.Grn = &v
}

// GetSharedCode returns the SharedCode field value if set, zero value otherwise.
func (o *WishlistWishlistResponse) GetSharedCode() string {
	if o == nil || IsNil(o.SharedCode) {
		var ret string
		return ret
	}
	return *o.SharedCode
}

// GetSharedCodeOk returns a tuple with the SharedCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistWishlistResponse) GetSharedCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SharedCode) {
		return nil, false
	}
	return o.SharedCode, true
}

// HasSharedCode returns a boolean if a field has been set.
func (o *WishlistWishlistResponse) HasSharedCode() bool {
	if o != nil && !IsNil(o.SharedCode) {
		return true
	}

	return false
}

// SetSharedCode gets a reference to the given string and assigns it to the SharedCode field.
func (o *WishlistWishlistResponse) SetSharedCode(v string) {
	o.SharedCode = &v
}

// GetPrivacy returns the Privacy field value if set, zero value otherwise.
func (o *WishlistWishlistResponse) GetPrivacy() WishlistPrivacy {
	if o == nil || IsNil(o.Privacy) {
		var ret WishlistPrivacy
		return ret
	}
	return *o.Privacy
}

// GetPrivacyOk returns a tuple with the Privacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistWishlistResponse) GetPrivacyOk() (*WishlistPrivacy, bool) {
	if o == nil || IsNil(o.Privacy) {
		return nil, false
	}
	return o.Privacy, true
}

// HasPrivacy returns a boolean if a field has been set.
func (o *WishlistWishlistResponse) HasPrivacy() bool {
	if o != nil && !IsNil(o.Privacy) {
		return true
	}

	return false
}

// SetPrivacy gets a reference to the given WishlistPrivacy and assigns it to the Privacy field.
func (o *WishlistWishlistResponse) SetPrivacy(v WishlistPrivacy) {
	o.Privacy = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WishlistWishlistResponse) GetLabel() WishlistLocalizedText {
	if o == nil || IsNil(o.Label) {
		var ret WishlistLocalizedText
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistWishlistResponse) GetLabelOk() (*WishlistLocalizedText, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WishlistWishlistResponse) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given WishlistLocalizedText and assigns it to the Label field.
func (o *WishlistWishlistResponse) SetLabel(v WishlistLocalizedText) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WishlistWishlistResponse) GetDescription() WishlistLocalizedText {
	if o == nil || IsNil(o.Description) {
		var ret WishlistLocalizedText
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistWishlistResponse) GetDescriptionOk() (*WishlistLocalizedText, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WishlistWishlistResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given WishlistLocalizedText and assigns it to the Description field.
func (o *WishlistWishlistResponse) SetDescription(v WishlistLocalizedText) {
	o.Description = &v
}

// GetCustomerGrn returns the CustomerGrn field value if set, zero value otherwise.
func (o *WishlistWishlistResponse) GetCustomerGrn() string {
	if o == nil || IsNil(o.CustomerGrn) {
		var ret string
		return ret
	}
	return *o.CustomerGrn
}

// GetCustomerGrnOk returns a tuple with the CustomerGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistWishlistResponse) GetCustomerGrnOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerGrn) {
		return nil, false
	}
	return o.CustomerGrn, true
}

// HasCustomerGrn returns a boolean if a field has been set.
func (o *WishlistWishlistResponse) HasCustomerGrn() bool {
	if o != nil && !IsNil(o.CustomerGrn) {
		return true
	}

	return false
}

// SetCustomerGrn gets a reference to the given string and assigns it to the CustomerGrn field.
func (o *WishlistWishlistResponse) SetCustomerGrn(v string) {
	o.CustomerGrn = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *WishlistWishlistResponse) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistWishlistResponse) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *WishlistWishlistResponse) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *WishlistWishlistResponse) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetItemsCount returns the ItemsCount field value if set, zero value otherwise.
func (o *WishlistWishlistResponse) GetItemsCount() string {
	if o == nil || IsNil(o.ItemsCount) {
		var ret string
		return ret
	}
	return *o.ItemsCount
}

// GetItemsCountOk returns a tuple with the ItemsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistWishlistResponse) GetItemsCountOk() (*string, bool) {
	if o == nil || IsNil(o.ItemsCount) {
		return nil, false
	}
	return o.ItemsCount, true
}

// HasItemsCount returns a boolean if a field has been set.
func (o *WishlistWishlistResponse) HasItemsCount() bool {
	if o != nil && !IsNil(o.ItemsCount) {
		return true
	}

	return false
}

// SetItemsCount gets a reference to the given string and assigns it to the ItemsCount field.
func (o *WishlistWishlistResponse) SetItemsCount(v string) {
	o.ItemsCount = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *WishlistWishlistResponse) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistWishlistResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *WishlistWishlistResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *WishlistWishlistResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *WishlistWishlistResponse) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WishlistWishlistResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *WishlistWishlistResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *WishlistWishlistResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o WishlistWishlistResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WishlistWishlistResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Grn) {
		toSerialize["grn"] = o.Grn
	}
	if !IsNil(o.SharedCode) {
		toSerialize["sharedCode"] = o.SharedCode
	}
	if !IsNil(o.Privacy) {
		toSerialize["privacy"] = o.Privacy
	}
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
	if !IsNil(o.ItemsCount) {
		toSerialize["itemsCount"] = o.ItemsCount
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WishlistWishlistResponse) UnmarshalJSON(data []byte) (err error) {
	varWishlistWishlistResponse := _WishlistWishlistResponse{}

	err = json.Unmarshal(data, &varWishlistWishlistResponse)

	if err != nil {
		return err
	}

	*o = WishlistWishlistResponse(varWishlistWishlistResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "grn")
		delete(additionalProperties, "sharedCode")
		delete(additionalProperties, "privacy")
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		delete(additionalProperties, "customerGrn")
		delete(additionalProperties, "isDefault")
		delete(additionalProperties, "itemsCount")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *WishlistWishlistResponse) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *WishlistWishlistResponse) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableWishlistWishlistResponse struct {
	value *WishlistWishlistResponse
	isSet bool
}

func (v NullableWishlistWishlistResponse) Get() *WishlistWishlistResponse {
	return v.value
}

func (v *NullableWishlistWishlistResponse) Set(val *WishlistWishlistResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWishlistWishlistResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWishlistWishlistResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWishlistWishlistResponse(val *WishlistWishlistResponse) *NullableWishlistWishlistResponse {
	return &NullableWishlistWishlistResponse{value: val, isSet: true}
}

func (v NullableWishlistWishlistResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWishlistWishlistResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
