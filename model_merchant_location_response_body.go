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

// MerchantLocationResponseBody struct for MerchantLocationResponseBody
type MerchantLocationResponseBody struct {
	MerchantLocation *MerchantLocationResponse `json:"merchant_location,omitempty"`
}

// NewMerchantLocationResponseBody instantiates a new MerchantLocationResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantLocationResponseBody() *MerchantLocationResponseBody {
	this := MerchantLocationResponseBody{}
	return &this
}

// NewMerchantLocationResponseBodyWithDefaults instantiates a new MerchantLocationResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantLocationResponseBodyWithDefaults() *MerchantLocationResponseBody {
	this := MerchantLocationResponseBody{}
	return &this
}

// GetMerchantLocation returns the MerchantLocation field value if set, zero value otherwise.
func (o *MerchantLocationResponseBody) GetMerchantLocation() MerchantLocationResponse {
	if o == nil || o.MerchantLocation == nil {
		var ret MerchantLocationResponse
		return ret
	}
	return *o.MerchantLocation
}

// GetMerchantLocationOk returns a tuple with the MerchantLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantLocationResponseBody) GetMerchantLocationOk() (*MerchantLocationResponse, bool) {
	if o == nil || o.MerchantLocation == nil {
		return nil, false
	}
	return o.MerchantLocation, true
}

// HasMerchantLocation returns a boolean if a field has been set.
func (o *MerchantLocationResponseBody) HasMerchantLocation() bool {
	if o != nil && o.MerchantLocation != nil {
		return true
	}

	return false
}

// SetMerchantLocation gets a reference to the given MerchantLocationResponse and assigns it to the MerchantLocation field.
func (o *MerchantLocationResponseBody) SetMerchantLocation(v MerchantLocationResponse) {
	o.MerchantLocation = &v
}

func (o MerchantLocationResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MerchantLocation != nil {
		toSerialize["merchant_location"] = o.MerchantLocation
	}
	return json.Marshal(toSerialize)
}

type NullableMerchantLocationResponseBody struct {
	value *MerchantLocationResponseBody
	isSet bool
}

func (v NullableMerchantLocationResponseBody) Get() *MerchantLocationResponseBody {
	return v.value
}

func (v *NullableMerchantLocationResponseBody) Set(val *MerchantLocationResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantLocationResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantLocationResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantLocationResponseBody(val *MerchantLocationResponseBody) *NullableMerchantLocationResponseBody {
	return &NullableMerchantLocationResponseBody{value: val, isSet: true}
}

func (v NullableMerchantLocationResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantLocationResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


