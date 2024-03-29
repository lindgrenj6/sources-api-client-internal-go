/*
 * Sources Internal
 *
 * Sources Internal
 *
 * API version: 2.0.0
 * Contact: support@redhat.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sourcesapiinternal

import (
	"encoding/json"
)

// AuthenticationExtraAzure struct for AuthenticationExtraAzure
type AuthenticationExtraAzure struct {
	TenantId *string `json:"tenant_id,omitempty"`
}

// NewAuthenticationExtraAzure instantiates a new AuthenticationExtraAzure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationExtraAzure() *AuthenticationExtraAzure {
	this := AuthenticationExtraAzure{}
	return &this
}

// NewAuthenticationExtraAzureWithDefaults instantiates a new AuthenticationExtraAzure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationExtraAzureWithDefaults() *AuthenticationExtraAzure {
	this := AuthenticationExtraAzure{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *AuthenticationExtraAzure) GetTenantId() string {
	if o == nil || o.TenantId == nil {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationExtraAzure) GetTenantIdOk() (*string, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *AuthenticationExtraAzure) HasTenantId() bool {
	if o != nil && o.TenantId != nil {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *AuthenticationExtraAzure) SetTenantId(v string) {
	o.TenantId = &v
}

func (o AuthenticationExtraAzure) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TenantId != nil {
		toSerialize["tenant_id"] = o.TenantId
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticationExtraAzure struct {
	value *AuthenticationExtraAzure
	isSet bool
}

func (v NullableAuthenticationExtraAzure) Get() *AuthenticationExtraAzure {
	return v.value
}

func (v *NullableAuthenticationExtraAzure) Set(val *AuthenticationExtraAzure) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationExtraAzure) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationExtraAzure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationExtraAzure(val *AuthenticationExtraAzure) *NullableAuthenticationExtraAzure {
	return &NullableAuthenticationExtraAzure{value: val, isSet: true}
}

func (v NullableAuthenticationExtraAzure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationExtraAzure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


