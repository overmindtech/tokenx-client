/*
NATS Token Exchange

Exchanges OAuth tokens for NATS tokens

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// TokenRequestData struct for TokenRequestData
type TokenRequestData struct {
	// The Public NKey of the user that is requesting a token
	UserPubKey *string `json:"user_pub_key,omitempty"`
	// Friendly user name
	UserName *string `json:"user_name,omitempty"`
}

// NewTokenRequestData instantiates a new TokenRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenRequestData() *TokenRequestData {
	this := TokenRequestData{}
	return &this
}

// NewTokenRequestDataWithDefaults instantiates a new TokenRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenRequestDataWithDefaults() *TokenRequestData {
	this := TokenRequestData{}
	return &this
}

// GetUserPubKey returns the UserPubKey field value if set, zero value otherwise.
func (o *TokenRequestData) GetUserPubKey() string {
	if o == nil || o.UserPubKey == nil {
		var ret string
		return ret
	}
	return *o.UserPubKey
}

// GetUserPubKeyOk returns a tuple with the UserPubKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRequestData) GetUserPubKeyOk() (*string, bool) {
	if o == nil || o.UserPubKey == nil {
		return nil, false
	}
	return o.UserPubKey, true
}

// HasUserPubKey returns a boolean if a field has been set.
func (o *TokenRequestData) HasUserPubKey() bool {
	if o != nil && o.UserPubKey != nil {
		return true
	}

	return false
}

// SetUserPubKey gets a reference to the given string and assigns it to the UserPubKey field.
func (o *TokenRequestData) SetUserPubKey(v string) {
	o.UserPubKey = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *TokenRequestData) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRequestData) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *TokenRequestData) HasUserName() bool {
	if o != nil && o.UserName != nil {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *TokenRequestData) SetUserName(v string) {
	o.UserName = &v
}

func (o TokenRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserPubKey != nil {
		toSerialize["user_pub_key"] = o.UserPubKey
	}
	if o.UserName != nil {
		toSerialize["user_name"] = o.UserName
	}
	return json.Marshal(toSerialize)
}

type NullableTokenRequestData struct {
	value *TokenRequestData
	isSet bool
}

func (v NullableTokenRequestData) Get() *TokenRequestData {
	return v.value
}

func (v *NullableTokenRequestData) Set(val *TokenRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenRequestData(val *TokenRequestData) *NullableTokenRequestData {
	return &NullableTokenRequestData{value: val, isSet: true}
}

func (v NullableTokenRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

