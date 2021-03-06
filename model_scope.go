/*
NATS Token Exchange

Exchanges OAuth tokens for NATS tokens

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tokenx

import (
	"encoding/json"
)

// Scope struct for Scope
type Scope struct {
	SigningKey *string `json:"SigningKey,omitempty"`
}

// NewScope instantiates a new Scope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScope() *Scope {
	this := Scope{}
	return &this
}

// NewScopeWithDefaults instantiates a new Scope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScopeWithDefaults() *Scope {
	this := Scope{}
	return &this
}

// GetSigningKey returns the SigningKey field value if set, zero value otherwise.
func (o *Scope) GetSigningKey() string {
	if o == nil || o.SigningKey == nil {
		var ret string
		return ret
	}
	return *o.SigningKey
}

// GetSigningKeyOk returns a tuple with the SigningKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scope) GetSigningKeyOk() (*string, bool) {
	if o == nil || o.SigningKey == nil {
		return nil, false
	}
	return o.SigningKey, true
}

// HasSigningKey returns a boolean if a field has been set.
func (o *Scope) HasSigningKey() bool {
	if o != nil && o.SigningKey != nil {
		return true
	}

	return false
}

// SetSigningKey gets a reference to the given string and assigns it to the SigningKey field.
func (o *Scope) SetSigningKey(v string) {
	o.SigningKey = &v
}

func (o Scope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SigningKey != nil {
		toSerialize["SigningKey"] = o.SigningKey
	}
	return json.Marshal(toSerialize)
}

type NullableScope struct {
	value *Scope
	isSet bool
}

func (v NullableScope) Get() *Scope {
	return v.value
}

func (v *NullableScope) Set(val *Scope) {
	v.value = val
	v.isSet = true
}

func (v NullableScope) IsSet() bool {
	return v.isSet
}

func (v *NullableScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScope(val *Scope) *NullableScope {
	return &NullableScope{value: val, isSet: true}
}

func (v NullableScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


