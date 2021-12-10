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

// AccountNumberResponse struct for AccountNumberResponse
type AccountNumberResponse struct {
	AccountGuid *string `json:"account_guid,omitempty"`
	AccountNumber NullableString `json:"account_number,omitempty"`
	Guid *string `json:"guid,omitempty"`
	InstitutionNumber NullableString `json:"institution_number,omitempty"`
	MemberGuid *string `json:"member_guid,omitempty"`
	RoutingNumber NullableString `json:"routing_number,omitempty"`
	TransitNumber NullableString `json:"transit_number,omitempty"`
	UserGuid *string `json:"user_guid,omitempty"`
}

// NewAccountNumberResponse instantiates a new AccountNumberResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountNumberResponse() *AccountNumberResponse {
	this := AccountNumberResponse{}
	return &this
}

// NewAccountNumberResponseWithDefaults instantiates a new AccountNumberResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountNumberResponseWithDefaults() *AccountNumberResponse {
	this := AccountNumberResponse{}
	return &this
}

// GetAccountGuid returns the AccountGuid field value if set, zero value otherwise.
func (o *AccountNumberResponse) GetAccountGuid() string {
	if o == nil || o.AccountGuid == nil {
		var ret string
		return ret
	}
	return *o.AccountGuid
}

// GetAccountGuidOk returns a tuple with the AccountGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountNumberResponse) GetAccountGuidOk() (*string, bool) {
	if o == nil || o.AccountGuid == nil {
		return nil, false
	}
	return o.AccountGuid, true
}

// HasAccountGuid returns a boolean if a field has been set.
func (o *AccountNumberResponse) HasAccountGuid() bool {
	if o != nil && o.AccountGuid != nil {
		return true
	}

	return false
}

// SetAccountGuid gets a reference to the given string and assigns it to the AccountGuid field.
func (o *AccountNumberResponse) SetAccountGuid(v string) {
	o.AccountGuid = &v
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountNumberResponse) GetAccountNumber() string {
	if o == nil || o.AccountNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccountNumber.Get()
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountNumberResponse) GetAccountNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountNumber.Get(), o.AccountNumber.IsSet()
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *AccountNumberResponse) HasAccountNumber() bool {
	if o != nil && o.AccountNumber.IsSet() {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given NullableString and assigns it to the AccountNumber field.
func (o *AccountNumberResponse) SetAccountNumber(v string) {
	o.AccountNumber.Set(&v)
}
// SetAccountNumberNil sets the value for AccountNumber to be an explicit nil
func (o *AccountNumberResponse) SetAccountNumberNil() {
	o.AccountNumber.Set(nil)
}

// UnsetAccountNumber ensures that no value is present for AccountNumber, not even an explicit nil
func (o *AccountNumberResponse) UnsetAccountNumber() {
	o.AccountNumber.Unset()
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *AccountNumberResponse) GetGuid() string {
	if o == nil || o.Guid == nil {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountNumberResponse) GetGuidOk() (*string, bool) {
	if o == nil || o.Guid == nil {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *AccountNumberResponse) HasGuid() bool {
	if o != nil && o.Guid != nil {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *AccountNumberResponse) SetGuid(v string) {
	o.Guid = &v
}

// GetInstitutionNumber returns the InstitutionNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountNumberResponse) GetInstitutionNumber() string {
	if o == nil || o.InstitutionNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.InstitutionNumber.Get()
}

// GetInstitutionNumberOk returns a tuple with the InstitutionNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountNumberResponse) GetInstitutionNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InstitutionNumber.Get(), o.InstitutionNumber.IsSet()
}

// HasInstitutionNumber returns a boolean if a field has been set.
func (o *AccountNumberResponse) HasInstitutionNumber() bool {
	if o != nil && o.InstitutionNumber.IsSet() {
		return true
	}

	return false
}

// SetInstitutionNumber gets a reference to the given NullableString and assigns it to the InstitutionNumber field.
func (o *AccountNumberResponse) SetInstitutionNumber(v string) {
	o.InstitutionNumber.Set(&v)
}
// SetInstitutionNumberNil sets the value for InstitutionNumber to be an explicit nil
func (o *AccountNumberResponse) SetInstitutionNumberNil() {
	o.InstitutionNumber.Set(nil)
}

// UnsetInstitutionNumber ensures that no value is present for InstitutionNumber, not even an explicit nil
func (o *AccountNumberResponse) UnsetInstitutionNumber() {
	o.InstitutionNumber.Unset()
}

// GetMemberGuid returns the MemberGuid field value if set, zero value otherwise.
func (o *AccountNumberResponse) GetMemberGuid() string {
	if o == nil || o.MemberGuid == nil {
		var ret string
		return ret
	}
	return *o.MemberGuid
}

// GetMemberGuidOk returns a tuple with the MemberGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountNumberResponse) GetMemberGuidOk() (*string, bool) {
	if o == nil || o.MemberGuid == nil {
		return nil, false
	}
	return o.MemberGuid, true
}

// HasMemberGuid returns a boolean if a field has been set.
func (o *AccountNumberResponse) HasMemberGuid() bool {
	if o != nil && o.MemberGuid != nil {
		return true
	}

	return false
}

// SetMemberGuid gets a reference to the given string and assigns it to the MemberGuid field.
func (o *AccountNumberResponse) SetMemberGuid(v string) {
	o.MemberGuid = &v
}

// GetRoutingNumber returns the RoutingNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountNumberResponse) GetRoutingNumber() string {
	if o == nil || o.RoutingNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.RoutingNumber.Get()
}

// GetRoutingNumberOk returns a tuple with the RoutingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountNumberResponse) GetRoutingNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RoutingNumber.Get(), o.RoutingNumber.IsSet()
}

// HasRoutingNumber returns a boolean if a field has been set.
func (o *AccountNumberResponse) HasRoutingNumber() bool {
	if o != nil && o.RoutingNumber.IsSet() {
		return true
	}

	return false
}

// SetRoutingNumber gets a reference to the given NullableString and assigns it to the RoutingNumber field.
func (o *AccountNumberResponse) SetRoutingNumber(v string) {
	o.RoutingNumber.Set(&v)
}
// SetRoutingNumberNil sets the value for RoutingNumber to be an explicit nil
func (o *AccountNumberResponse) SetRoutingNumberNil() {
	o.RoutingNumber.Set(nil)
}

// UnsetRoutingNumber ensures that no value is present for RoutingNumber, not even an explicit nil
func (o *AccountNumberResponse) UnsetRoutingNumber() {
	o.RoutingNumber.Unset()
}

// GetTransitNumber returns the TransitNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountNumberResponse) GetTransitNumber() string {
	if o == nil || o.TransitNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.TransitNumber.Get()
}

// GetTransitNumberOk returns a tuple with the TransitNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountNumberResponse) GetTransitNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TransitNumber.Get(), o.TransitNumber.IsSet()
}

// HasTransitNumber returns a boolean if a field has been set.
func (o *AccountNumberResponse) HasTransitNumber() bool {
	if o != nil && o.TransitNumber.IsSet() {
		return true
	}

	return false
}

// SetTransitNumber gets a reference to the given NullableString and assigns it to the TransitNumber field.
func (o *AccountNumberResponse) SetTransitNumber(v string) {
	o.TransitNumber.Set(&v)
}
// SetTransitNumberNil sets the value for TransitNumber to be an explicit nil
func (o *AccountNumberResponse) SetTransitNumberNil() {
	o.TransitNumber.Set(nil)
}

// UnsetTransitNumber ensures that no value is present for TransitNumber, not even an explicit nil
func (o *AccountNumberResponse) UnsetTransitNumber() {
	o.TransitNumber.Unset()
}

// GetUserGuid returns the UserGuid field value if set, zero value otherwise.
func (o *AccountNumberResponse) GetUserGuid() string {
	if o == nil || o.UserGuid == nil {
		var ret string
		return ret
	}
	return *o.UserGuid
}

// GetUserGuidOk returns a tuple with the UserGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountNumberResponse) GetUserGuidOk() (*string, bool) {
	if o == nil || o.UserGuid == nil {
		return nil, false
	}
	return o.UserGuid, true
}

// HasUserGuid returns a boolean if a field has been set.
func (o *AccountNumberResponse) HasUserGuid() bool {
	if o != nil && o.UserGuid != nil {
		return true
	}

	return false
}

// SetUserGuid gets a reference to the given string and assigns it to the UserGuid field.
func (o *AccountNumberResponse) SetUserGuid(v string) {
	o.UserGuid = &v
}

func (o AccountNumberResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountGuid != nil {
		toSerialize["account_guid"] = o.AccountGuid
	}
	if o.AccountNumber.IsSet() {
		toSerialize["account_number"] = o.AccountNumber.Get()
	}
	if o.Guid != nil {
		toSerialize["guid"] = o.Guid
	}
	if o.InstitutionNumber.IsSet() {
		toSerialize["institution_number"] = o.InstitutionNumber.Get()
	}
	if o.MemberGuid != nil {
		toSerialize["member_guid"] = o.MemberGuid
	}
	if o.RoutingNumber.IsSet() {
		toSerialize["routing_number"] = o.RoutingNumber.Get()
	}
	if o.TransitNumber.IsSet() {
		toSerialize["transit_number"] = o.TransitNumber.Get()
	}
	if o.UserGuid != nil {
		toSerialize["user_guid"] = o.UserGuid
	}
	return json.Marshal(toSerialize)
}

type NullableAccountNumberResponse struct {
	value *AccountNumberResponse
	isSet bool
}

func (v NullableAccountNumberResponse) Get() *AccountNumberResponse {
	return v.value
}

func (v *NullableAccountNumberResponse) Set(val *AccountNumberResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountNumberResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountNumberResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountNumberResponse(val *AccountNumberResponse) *NullableAccountNumberResponse {
	return &NullableAccountNumberResponse{value: val, isSet: true}
}

func (v NullableAccountNumberResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountNumberResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


