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

// AccountNumbersResponseBody struct for AccountNumbersResponseBody
type AccountNumbersResponseBody struct {
	AccountNumbers *[]AccountNumberResponse `json:"account_numbers,omitempty"`
	Pagination *PaginationResponse `json:"pagination,omitempty"`
}

// NewAccountNumbersResponseBody instantiates a new AccountNumbersResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountNumbersResponseBody() *AccountNumbersResponseBody {
	this := AccountNumbersResponseBody{}
	return &this
}

// NewAccountNumbersResponseBodyWithDefaults instantiates a new AccountNumbersResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountNumbersResponseBodyWithDefaults() *AccountNumbersResponseBody {
	this := AccountNumbersResponseBody{}
	return &this
}

// GetAccountNumbers returns the AccountNumbers field value if set, zero value otherwise.
func (o *AccountNumbersResponseBody) GetAccountNumbers() []AccountNumberResponse {
	if o == nil || o.AccountNumbers == nil {
		var ret []AccountNumberResponse
		return ret
	}
	return *o.AccountNumbers
}

// GetAccountNumbersOk returns a tuple with the AccountNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountNumbersResponseBody) GetAccountNumbersOk() (*[]AccountNumberResponse, bool) {
	if o == nil || o.AccountNumbers == nil {
		return nil, false
	}
	return o.AccountNumbers, true
}

// HasAccountNumbers returns a boolean if a field has been set.
func (o *AccountNumbersResponseBody) HasAccountNumbers() bool {
	if o != nil && o.AccountNumbers != nil {
		return true
	}

	return false
}

// SetAccountNumbers gets a reference to the given []AccountNumberResponse and assigns it to the AccountNumbers field.
func (o *AccountNumbersResponseBody) SetAccountNumbers(v []AccountNumberResponse) {
	o.AccountNumbers = &v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *AccountNumbersResponseBody) GetPagination() PaginationResponse {
	if o == nil || o.Pagination == nil {
		var ret PaginationResponse
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountNumbersResponseBody) GetPaginationOk() (*PaginationResponse, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *AccountNumbersResponseBody) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given PaginationResponse and assigns it to the Pagination field.
func (o *AccountNumbersResponseBody) SetPagination(v PaginationResponse) {
	o.Pagination = &v
}

func (o AccountNumbersResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountNumbers != nil {
		toSerialize["account_numbers"] = o.AccountNumbers
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableAccountNumbersResponseBody struct {
	value *AccountNumbersResponseBody
	isSet bool
}

func (v NullableAccountNumbersResponseBody) Get() *AccountNumbersResponseBody {
	return v.value
}

func (v *NullableAccountNumbersResponseBody) Set(val *AccountNumbersResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountNumbersResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountNumbersResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountNumbersResponseBody(val *AccountNumbersResponseBody) *NullableAccountNumbersResponseBody {
	return &NullableAccountNumbersResponseBody{value: val, isSet: true}
}

func (v NullableAccountNumbersResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountNumbersResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


