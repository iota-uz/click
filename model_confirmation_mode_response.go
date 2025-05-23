/*
CLICK-PASS-API

API for managing payments and fiscalization with CLICK Pass.

API version: 1.0.0
Contact: danil@iota.uz
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clickapi

import (
	"encoding/json"
)

// checks if the ConfirmationModeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfirmationModeResponse{}

// ConfirmationModeResponse struct for ConfirmationModeResponse
type ConfirmationModeResponse struct {
	// Error code
	ErrorCode *int32 `form:"error_code" json:"error_code,omitempty"` // Error code
	// Confirmation mode status
	ErrorNote *string `form:"error_note" json:"error_note,omitempty"` // Confirmation mode status
}

// NewConfirmationModeResponse instantiates a new ConfirmationModeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfirmationModeResponse() *ConfirmationModeResponse {
	this := ConfirmationModeResponse{}
	return &this
}

// NewConfirmationModeResponseWithDefaults instantiates a new ConfirmationModeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfirmationModeResponseWithDefaults() *ConfirmationModeResponse {
	this := ConfirmationModeResponse{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *ConfirmationModeResponse) GetErrorCode() int32 {
	if o == nil || IsNil(o.ErrorCode) {
		var ret int32
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfirmationModeResponse) GetErrorCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *ConfirmationModeResponse) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int32 and assigns it to the ErrorCode field.
func (o *ConfirmationModeResponse) SetErrorCode(v int32) {
	o.ErrorCode = &v
}

// GetErrorNote returns the ErrorNote field value if set, zero value otherwise.
func (o *ConfirmationModeResponse) GetErrorNote() string {
	if o == nil || IsNil(o.ErrorNote) {
		var ret string
		return ret
	}
	return *o.ErrorNote
}

// GetErrorNoteOk returns a tuple with the ErrorNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfirmationModeResponse) GetErrorNoteOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorNote) {
		return nil, false
	}
	return o.ErrorNote, true
}

// HasErrorNote returns a boolean if a field has been set.
func (o *ConfirmationModeResponse) HasErrorNote() bool {
	if o != nil && !IsNil(o.ErrorNote) {
		return true
	}

	return false
}

// SetErrorNote gets a reference to the given string and assigns it to the ErrorNote field.
func (o *ConfirmationModeResponse) SetErrorNote(v string) {
	o.ErrorNote = &v
}

func (o ConfirmationModeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfirmationModeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorCode) {
		toSerialize["error_code"] = o.ErrorCode
	}
	if !IsNil(o.ErrorNote) {
		toSerialize["error_note"] = o.ErrorNote
	}
	return toSerialize, nil
}

type NullableConfirmationModeResponse struct {
	value *ConfirmationModeResponse
	isSet bool
}

func (v NullableConfirmationModeResponse) Get() *ConfirmationModeResponse {
	return v.value
}

func (v *NullableConfirmationModeResponse) Set(val *ConfirmationModeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfirmationModeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfirmationModeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfirmationModeResponse(val *ConfirmationModeResponse) *NullableConfirmationModeResponse {
	return &NullableConfirmationModeResponse{value: val, isSet: true}
}

func (v NullableConfirmationModeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfirmationModeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
