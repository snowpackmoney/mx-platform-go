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

// AccountOwnersResponseBody struct for AccountOwnersResponseBody
type AccountOwnersResponseBody struct {
	AccountOwners *[]AccountOwnerResponse `json:"account_owners,omitempty"`
	Pagination *PaginationResponse `json:"pagination,omitempty"`
}

// NewAccountOwnersResponseBody instantiates a new AccountOwnersResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountOwnersResponseBody() *AccountOwnersResponseBody {
	this := AccountOwnersResponseBody{}
	return &this
}

// NewAccountOwnersResponseBodyWithDefaults instantiates a new AccountOwnersResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountOwnersResponseBodyWithDefaults() *AccountOwnersResponseBody {
	this := AccountOwnersResponseBody{}
	return &this
}

// GetAccountOwners returns the AccountOwners field value if set, zero value otherwise.
func (o *AccountOwnersResponseBody) GetAccountOwners() []AccountOwnerResponse {
	if o == nil || o.AccountOwners == nil {
		var ret []AccountOwnerResponse
		return ret
	}
	return *o.AccountOwners
}

// GetAccountOwnersOk returns a tuple with the AccountOwners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountOwnersResponseBody) GetAccountOwnersOk() (*[]AccountOwnerResponse, bool) {
	if o == nil || o.AccountOwners == nil {
		return nil, false
	}
	return o.AccountOwners, true
}

// HasAccountOwners returns a boolean if a field has been set.
func (o *AccountOwnersResponseBody) HasAccountOwners() bool {
	if o != nil && o.AccountOwners != nil {
		return true
	}

	return false
}

// SetAccountOwners gets a reference to the given []AccountOwnerResponse and assigns it to the AccountOwners field.
func (o *AccountOwnersResponseBody) SetAccountOwners(v []AccountOwnerResponse) {
	o.AccountOwners = &v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *AccountOwnersResponseBody) GetPagination() PaginationResponse {
	if o == nil || o.Pagination == nil {
		var ret PaginationResponse
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountOwnersResponseBody) GetPaginationOk() (*PaginationResponse, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *AccountOwnersResponseBody) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given PaginationResponse and assigns it to the Pagination field.
func (o *AccountOwnersResponseBody) SetPagination(v PaginationResponse) {
	o.Pagination = &v
}

func (o AccountOwnersResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountOwners != nil {
		toSerialize["account_owners"] = o.AccountOwners
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableAccountOwnersResponseBody struct {
	value *AccountOwnersResponseBody
	isSet bool
}

func (v NullableAccountOwnersResponseBody) Get() *AccountOwnersResponseBody {
	return v.value
}

func (v *NullableAccountOwnersResponseBody) Set(val *AccountOwnersResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountOwnersResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountOwnersResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountOwnersResponseBody(val *AccountOwnersResponseBody) *NullableAccountOwnersResponseBody {
	return &NullableAccountOwnersResponseBody{value: val, isSet: true}
}

func (v NullableAccountOwnersResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountOwnersResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


