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

// CategoryCreateRequestBody struct for CategoryCreateRequestBody
type CategoryCreateRequestBody struct {
	Category *CategoryCreateRequest `json:"category,omitempty"`
}

// NewCategoryCreateRequestBody instantiates a new CategoryCreateRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryCreateRequestBody() *CategoryCreateRequestBody {
	this := CategoryCreateRequestBody{}
	return &this
}

// NewCategoryCreateRequestBodyWithDefaults instantiates a new CategoryCreateRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryCreateRequestBodyWithDefaults() *CategoryCreateRequestBody {
	this := CategoryCreateRequestBody{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *CategoryCreateRequestBody) GetCategory() CategoryCreateRequest {
	if o == nil || o.Category == nil {
		var ret CategoryCreateRequest
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryCreateRequestBody) GetCategoryOk() (*CategoryCreateRequest, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *CategoryCreateRequestBody) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given CategoryCreateRequest and assigns it to the Category field.
func (o *CategoryCreateRequestBody) SetCategory(v CategoryCreateRequest) {
	o.Category = &v
}

func (o CategoryCreateRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	return json.Marshal(toSerialize)
}

type NullableCategoryCreateRequestBody struct {
	value *CategoryCreateRequestBody
	isSet bool
}

func (v NullableCategoryCreateRequestBody) Get() *CategoryCreateRequestBody {
	return v.value
}

func (v *NullableCategoryCreateRequestBody) Set(val *CategoryCreateRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryCreateRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryCreateRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryCreateRequestBody(val *CategoryCreateRequestBody) *NullableCategoryCreateRequestBody {
	return &NullableCategoryCreateRequestBody{value: val, isSet: true}
}

func (v NullableCategoryCreateRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryCreateRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


