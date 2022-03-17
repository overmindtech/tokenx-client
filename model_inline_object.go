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

// InlineObject struct for InlineObject
type InlineObject struct {
	// The Public NKey of the user that is requesting a token
	UserPubKey *string `json:"user_pub_key,omitempty"`
	// Friendly user name
	UserName *string `json:"user_name,omitempty"`
}

// NewInlineObject instantiates a new InlineObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject() *InlineObject {
	this := InlineObject{}
	return &this
}

// NewInlineObjectWithDefaults instantiates a new InlineObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObjectWithDefaults() *InlineObject {
	this := InlineObject{}
	return &this
}

// GetUserPubKey returns the UserPubKey field value if set, zero value otherwise.
func (o *InlineObject) GetUserPubKey() string {
	if o == nil || o.UserPubKey == nil {
		var ret string
		return ret
	}
	return *o.UserPubKey
}

// GetUserPubKeyOk returns a tuple with the UserPubKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject) GetUserPubKeyOk() (*string, bool) {
	if o == nil || o.UserPubKey == nil {
		return nil, false
	}
	return o.UserPubKey, true
}

// HasUserPubKey returns a boolean if a field has been set.
func (o *InlineObject) HasUserPubKey() bool {
	if o != nil && o.UserPubKey != nil {
		return true
	}

	return false
}

// SetUserPubKey gets a reference to the given string and assigns it to the UserPubKey field.
func (o *InlineObject) SetUserPubKey(v string) {
	o.UserPubKey = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *InlineObject) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *InlineObject) HasUserName() bool {
	if o != nil && o.UserName != nil {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *InlineObject) SetUserName(v string) {
	o.UserName = &v
}

func (o InlineObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserPubKey != nil {
		toSerialize["user_pub_key"] = o.UserPubKey
	}
	if o.UserName != nil {
		toSerialize["user_name"] = o.UserName
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject struct {
	value *InlineObject
	isSet bool
}

func (v NullableInlineObject) Get() *InlineObject {
	return v.value
}

func (v *NullableInlineObject) Set(val *InlineObject) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject(val *InlineObject) *NullableInlineObject {
	return &NullableInlineObject{value: val, isSet: true}
}

func (v NullableInlineObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

