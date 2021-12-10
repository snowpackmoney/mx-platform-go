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

// WidgetRequestBody struct for WidgetRequestBody
type WidgetRequestBody struct {
	WidgetUrl *WidgetRequest `json:"widget_url,omitempty"`
}

// NewWidgetRequestBody instantiates a new WidgetRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWidgetRequestBody() *WidgetRequestBody {
	this := WidgetRequestBody{}
	return &this
}

// NewWidgetRequestBodyWithDefaults instantiates a new WidgetRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWidgetRequestBodyWithDefaults() *WidgetRequestBody {
	this := WidgetRequestBody{}
	return &this
}

// GetWidgetUrl returns the WidgetUrl field value if set, zero value otherwise.
func (o *WidgetRequestBody) GetWidgetUrl() WidgetRequest {
	if o == nil || o.WidgetUrl == nil {
		var ret WidgetRequest
		return ret
	}
	return *o.WidgetUrl
}

// GetWidgetUrlOk returns a tuple with the WidgetUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetRequestBody) GetWidgetUrlOk() (*WidgetRequest, bool) {
	if o == nil || o.WidgetUrl == nil {
		return nil, false
	}
	return o.WidgetUrl, true
}

// HasWidgetUrl returns a boolean if a field has been set.
func (o *WidgetRequestBody) HasWidgetUrl() bool {
	if o != nil && o.WidgetUrl != nil {
		return true
	}

	return false
}

// SetWidgetUrl gets a reference to the given WidgetRequest and assigns it to the WidgetUrl field.
func (o *WidgetRequestBody) SetWidgetUrl(v WidgetRequest) {
	o.WidgetUrl = &v
}

func (o WidgetRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.WidgetUrl != nil {
		toSerialize["widget_url"] = o.WidgetUrl
	}
	return json.Marshal(toSerialize)
}

type NullableWidgetRequestBody struct {
	value *WidgetRequestBody
	isSet bool
}

func (v NullableWidgetRequestBody) Get() *WidgetRequestBody {
	return v.value
}

func (v *NullableWidgetRequestBody) Set(val *WidgetRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetRequestBody(val *WidgetRequestBody) *NullableWidgetRequestBody {
	return &NullableWidgetRequestBody{value: val, isSet: true}
}

func (v NullableWidgetRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


