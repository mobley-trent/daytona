/*
Daytona Server API

Daytona Server API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"encoding/json"
)

// checks if the CreateWorkspaceRequestProjectSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateWorkspaceRequestProjectSource{}

// CreateWorkspaceRequestProjectSource struct for CreateWorkspaceRequestProjectSource
type CreateWorkspaceRequestProjectSource struct {
	Repository *GitRepository `json:"repository,omitempty"`
}

// NewCreateWorkspaceRequestProjectSource instantiates a new CreateWorkspaceRequestProjectSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWorkspaceRequestProjectSource() *CreateWorkspaceRequestProjectSource {
	this := CreateWorkspaceRequestProjectSource{}
	return &this
}

// NewCreateWorkspaceRequestProjectSourceWithDefaults instantiates a new CreateWorkspaceRequestProjectSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWorkspaceRequestProjectSourceWithDefaults() *CreateWorkspaceRequestProjectSource {
	this := CreateWorkspaceRequestProjectSource{}
	return &this
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *CreateWorkspaceRequestProjectSource) GetRepository() GitRepository {
	if o == nil || IsNil(o.Repository) {
		var ret GitRepository
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWorkspaceRequestProjectSource) GetRepositoryOk() (*GitRepository, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *CreateWorkspaceRequestProjectSource) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given GitRepository and assigns it to the Repository field.
func (o *CreateWorkspaceRequestProjectSource) SetRepository(v GitRepository) {
	o.Repository = &v
}

func (o CreateWorkspaceRequestProjectSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWorkspaceRequestProjectSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	return toSerialize, nil
}

type NullableCreateWorkspaceRequestProjectSource struct {
	value *CreateWorkspaceRequestProjectSource
	isSet bool
}

func (v NullableCreateWorkspaceRequestProjectSource) Get() *CreateWorkspaceRequestProjectSource {
	return v.value
}

func (v *NullableCreateWorkspaceRequestProjectSource) Set(val *CreateWorkspaceRequestProjectSource) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWorkspaceRequestProjectSource) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWorkspaceRequestProjectSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWorkspaceRequestProjectSource(val *CreateWorkspaceRequestProjectSource) *NullableCreateWorkspaceRequestProjectSource {
	return &NullableCreateWorkspaceRequestProjectSource{value: val, isSet: true}
}

func (v NullableCreateWorkspaceRequestProjectSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWorkspaceRequestProjectSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


