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

// GenericFields struct for GenericFields
type GenericFields struct {
	// TagList is a unique array of lower case strings All tag list methods lower case the strings in the arguments
	Tags []string `json:"tags,omitempty"`
	// ClaimType is used to indicate the type of JWT being stored in a Claim
	Type *string `json:"type,omitempty"`
	Version *int64 `json:"version,omitempty"`
}

// NewGenericFields instantiates a new GenericFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericFields() *GenericFields {
	this := GenericFields{}
	return &this
}

// NewGenericFieldsWithDefaults instantiates a new GenericFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericFieldsWithDefaults() *GenericFields {
	this := GenericFields{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GenericFields) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericFields) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GenericFields) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GenericFields) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GenericFields) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericFields) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GenericFields) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GenericFields) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *GenericFields) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericFields) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *GenericFields) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *GenericFields) SetVersion(v int64) {
	o.Version = &v
}

func (o GenericFields) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableGenericFields struct {
	value *GenericFields
	isSet bool
}

func (v NullableGenericFields) Get() *GenericFields {
	return v.value
}

func (v *NullableGenericFields) Set(val *GenericFields) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericFields) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericFields(val *GenericFields) *NullableGenericFields {
	return &NullableGenericFields{value: val, isSet: true}
}

func (v NullableGenericFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


