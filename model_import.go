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

// Import Import describes a mapping from another account into this one
type Import struct {
	Account *string `json:"account,omitempty"`
	// Subject is a string that represents a NATS subject
	LocalSubject *string `json:"local_subject,omitempty"`
	Name *string `json:"name,omitempty"`
	Share *bool `json:"share,omitempty"`
	// Subject is a string that represents a NATS subject
	Subject *string `json:"subject,omitempty"`
	// Subject is a string that represents a NATS subject
	To *string `json:"to,omitempty"`
	Token *string `json:"token,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewImport instantiates a new Import object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImport() *Import {
	this := Import{}
	return &this
}

// NewImportWithDefaults instantiates a new Import object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportWithDefaults() *Import {
	this := Import{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *Import) GetAccount() string {
	if o == nil || o.Account == nil {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Import) GetAccountOk() (*string, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *Import) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *Import) SetAccount(v string) {
	o.Account = &v
}

// GetLocalSubject returns the LocalSubject field value if set, zero value otherwise.
func (o *Import) GetLocalSubject() string {
	if o == nil || o.LocalSubject == nil {
		var ret string
		return ret
	}
	return *o.LocalSubject
}

// GetLocalSubjectOk returns a tuple with the LocalSubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Import) GetLocalSubjectOk() (*string, bool) {
	if o == nil || o.LocalSubject == nil {
		return nil, false
	}
	return o.LocalSubject, true
}

// HasLocalSubject returns a boolean if a field has been set.
func (o *Import) HasLocalSubject() bool {
	if o != nil && o.LocalSubject != nil {
		return true
	}

	return false
}

// SetLocalSubject gets a reference to the given string and assigns it to the LocalSubject field.
func (o *Import) SetLocalSubject(v string) {
	o.LocalSubject = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Import) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Import) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Import) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Import) SetName(v string) {
	o.Name = &v
}

// GetShare returns the Share field value if set, zero value otherwise.
func (o *Import) GetShare() bool {
	if o == nil || o.Share == nil {
		var ret bool
		return ret
	}
	return *o.Share
}

// GetShareOk returns a tuple with the Share field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Import) GetShareOk() (*bool, bool) {
	if o == nil || o.Share == nil {
		return nil, false
	}
	return o.Share, true
}

// HasShare returns a boolean if a field has been set.
func (o *Import) HasShare() bool {
	if o != nil && o.Share != nil {
		return true
	}

	return false
}

// SetShare gets a reference to the given bool and assigns it to the Share field.
func (o *Import) SetShare(v bool) {
	o.Share = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *Import) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Import) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *Import) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *Import) SetSubject(v string) {
	o.Subject = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *Import) GetTo() string {
	if o == nil || o.To == nil {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Import) GetToOk() (*string, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *Import) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *Import) SetTo(v string) {
	o.To = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *Import) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Import) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *Import) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *Import) SetToken(v string) {
	o.Token = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Import) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Import) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Import) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Import) SetType(v string) {
	o.Type = &v
}

func (o Import) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	if o.LocalSubject != nil {
		toSerialize["local_subject"] = o.LocalSubject
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Share != nil {
		toSerialize["share"] = o.Share
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.To != nil {
		toSerialize["to"] = o.To
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableImport struct {
	value *Import
	isSet bool
}

func (v NullableImport) Get() *Import {
	return v.value
}

func (v *NullableImport) Set(val *Import) {
	v.value = val
	v.isSet = true
}

func (v NullableImport) IsSet() bool {
	return v.isSet
}

func (v *NullableImport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImport(val *Import) *NullableImport {
	return &NullableImport{value: val, isSet: true}
}

func (v NullableImport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

