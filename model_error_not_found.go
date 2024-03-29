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

// ErrorNotFound struct for ErrorNotFound
type ErrorNotFound struct {
	Errors *[]ErrorNotFoundErrors `json:"errors,omitempty"`
}

// NewErrorNotFound instantiates a new ErrorNotFound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorNotFound() *ErrorNotFound {
	this := ErrorNotFound{}
	return &this
}

// NewErrorNotFoundWithDefaults instantiates a new ErrorNotFound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorNotFoundWithDefaults() *ErrorNotFound {
	this := ErrorNotFound{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ErrorNotFound) GetErrors() []ErrorNotFoundErrors {
	if o == nil || o.Errors == nil {
		var ret []ErrorNotFoundErrors
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorNotFound) GetErrorsOk() (*[]ErrorNotFoundErrors, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ErrorNotFound) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ErrorNotFoundErrors and assigns it to the Errors field.
func (o *ErrorNotFound) SetErrors(v []ErrorNotFoundErrors) {
	o.Errors = &v
}

func (o ErrorNotFound) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableErrorNotFound struct {
	value *ErrorNotFound
	isSet bool
}

func (v NullableErrorNotFound) Get() *ErrorNotFound {
	return v.value
}

func (v *NullableErrorNotFound) Set(val *ErrorNotFound) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorNotFound) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorNotFound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorNotFound(val *ErrorNotFound) *NullableErrorNotFound {
	return &NullableErrorNotFound{value: val, isSet: true}
}

func (v NullableErrorNotFound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorNotFound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


