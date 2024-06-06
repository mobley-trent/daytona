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

// checks if the FRPSConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FRPSConfig{}

// FRPSConfig struct for FRPSConfig
type FRPSConfig struct {
	Domain   *string `json:"domain,omitempty"`
	Port     *int32  `json:"port,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
}

// NewFRPSConfig instantiates a new FRPSConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFRPSConfig() *FRPSConfig {
	this := FRPSConfig{}
	return &this
}

// NewFRPSConfigWithDefaults instantiates a new FRPSConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFRPSConfigWithDefaults() *FRPSConfig {
	this := FRPSConfig{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *FRPSConfig) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FRPSConfig) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *FRPSConfig) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *FRPSConfig) SetDomain(v string) {
	o.Domain = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *FRPSConfig) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FRPSConfig) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *FRPSConfig) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *FRPSConfig) SetPort(v int32) {
	o.Port = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *FRPSConfig) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FRPSConfig) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *FRPSConfig) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *FRPSConfig) SetProtocol(v string) {
	o.Protocol = &v
}

func (o FRPSConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FRPSConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	return toSerialize, nil
}

type NullableFRPSConfig struct {
	value *FRPSConfig
	isSet bool
}

func (v NullableFRPSConfig) Get() *FRPSConfig {
	return v.value
}

func (v *NullableFRPSConfig) Set(val *FRPSConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableFRPSConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableFRPSConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFRPSConfig(val *FRPSConfig) *NullableFRPSConfig {
	return &NullableFRPSConfig{value: val, isSet: true}
}

func (v NullableFRPSConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFRPSConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
