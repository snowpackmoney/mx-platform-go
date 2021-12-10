/*
MX Platform API

The MX Platform API is a powerful, fully-featured API designed to make aggregating and enhancing financial data easy and reliable. It can seamlessly connect your app or website to tens of thousands of financial institutions.

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mxplatformgo

import (
	"encoding/json"
)

// ManagedAccountUpdateRequestBody struct for ManagedAccountUpdateRequestBody
type ManagedAccountUpdateRequestBody struct {
	Account *ManagedAccountUpdateRequest `json:"account,omitempty"`
}

// NewManagedAccountUpdateRequestBody instantiates a new ManagedAccountUpdateRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedAccountUpdateRequestBody() *ManagedAccountUpdateRequestBody {
	this := ManagedAccountUpdateRequestBody{}
	return &this
}

// NewManagedAccountUpdateRequestBodyWithDefaults instantiates a new ManagedAccountUpdateRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedAccountUpdateRequestBodyWithDefaults() *ManagedAccountUpdateRequestBody {
	this := ManagedAccountUpdateRequestBody{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ManagedAccountUpdateRequestBody) GetAccount() ManagedAccountUpdateRequest {
	if o == nil || o.Account == nil {
		var ret ManagedAccountUpdateRequest
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedAccountUpdateRequestBody) GetAccountOk() (*ManagedAccountUpdateRequest, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ManagedAccountUpdateRequestBody) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given ManagedAccountUpdateRequest and assigns it to the Account field.
func (o *ManagedAccountUpdateRequestBody) SetAccount(v ManagedAccountUpdateRequest) {
	o.Account = &v
}

func (o ManagedAccountUpdateRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	return json.Marshal(toSerialize)
}

type NullableManagedAccountUpdateRequestBody struct {
	value *ManagedAccountUpdateRequestBody
	isSet bool
}

func (v NullableManagedAccountUpdateRequestBody) Get() *ManagedAccountUpdateRequestBody {
	return v.value
}

func (v *NullableManagedAccountUpdateRequestBody) Set(val *ManagedAccountUpdateRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedAccountUpdateRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedAccountUpdateRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedAccountUpdateRequestBody(val *ManagedAccountUpdateRequestBody) *NullableManagedAccountUpdateRequestBody {
	return &NullableManagedAccountUpdateRequestBody{value: val, isSet: true}
}

func (v NullableManagedAccountUpdateRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedAccountUpdateRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


