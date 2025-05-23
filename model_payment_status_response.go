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

// checks if the PaymentStatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentStatusResponse{}

// PaymentStatusResponse struct for PaymentStatusResponse
type PaymentStatusResponse struct {
	// Error code
	ErrorCode *int32 `form:"error_code" json:"error_code,omitempty"` // Error code
	// Error description
	ErrorNote *string `form:"error_note" json:"error_note,omitempty"` // Error description
	// Payment Identifier
	PaymentId *string `form:"payment_id" json:"payment_id,omitempty"` // Payment Identifier
	// Payment status code
	PaymentStatus *int32 `form:"payment_status" json:"payment_status,omitempty"` // Payment status code
}

// NewPaymentStatusResponse instantiates a new PaymentStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentStatusResponse() *PaymentStatusResponse {
	this := PaymentStatusResponse{}
	return &this
}

// NewPaymentStatusResponseWithDefaults instantiates a new PaymentStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentStatusResponseWithDefaults() *PaymentStatusResponse {
	this := PaymentStatusResponse{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *PaymentStatusResponse) GetErrorCode() int32 {
	if o == nil || IsNil(o.ErrorCode) {
		var ret int32
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentStatusResponse) GetErrorCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *PaymentStatusResponse) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int32 and assigns it to the ErrorCode field.
func (o *PaymentStatusResponse) SetErrorCode(v int32) {
	o.ErrorCode = &v
}

// GetErrorNote returns the ErrorNote field value if set, zero value otherwise.
func (o *PaymentStatusResponse) GetErrorNote() string {
	if o == nil || IsNil(o.ErrorNote) {
		var ret string
		return ret
	}
	return *o.ErrorNote
}

// GetErrorNoteOk returns a tuple with the ErrorNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentStatusResponse) GetErrorNoteOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorNote) {
		return nil, false
	}
	return o.ErrorNote, true
}

// HasErrorNote returns a boolean if a field has been set.
func (o *PaymentStatusResponse) HasErrorNote() bool {
	if o != nil && !IsNil(o.ErrorNote) {
		return true
	}

	return false
}

// SetErrorNote gets a reference to the given string and assigns it to the ErrorNote field.
func (o *PaymentStatusResponse) SetErrorNote(v string) {
	o.ErrorNote = &v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *PaymentStatusResponse) GetPaymentId() string {
	if o == nil || IsNil(o.PaymentId) {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentStatusResponse) GetPaymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentId) {
		return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *PaymentStatusResponse) HasPaymentId() bool {
	if o != nil && !IsNil(o.PaymentId) {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *PaymentStatusResponse) SetPaymentId(v string) {
	o.PaymentId = &v
}

// GetPaymentStatus returns the PaymentStatus field value if set, zero value otherwise.
func (o *PaymentStatusResponse) GetPaymentStatus() int32 {
	if o == nil || IsNil(o.PaymentStatus) {
		var ret int32
		return ret
	}
	return *o.PaymentStatus
}

// GetPaymentStatusOk returns a tuple with the PaymentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentStatusResponse) GetPaymentStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.PaymentStatus) {
		return nil, false
	}
	return o.PaymentStatus, true
}

// HasPaymentStatus returns a boolean if a field has been set.
func (o *PaymentStatusResponse) HasPaymentStatus() bool {
	if o != nil && !IsNil(o.PaymentStatus) {
		return true
	}

	return false
}

// SetPaymentStatus gets a reference to the given int32 and assigns it to the PaymentStatus field.
func (o *PaymentStatusResponse) SetPaymentStatus(v int32) {
	o.PaymentStatus = &v
}

func (o PaymentStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorCode) {
		toSerialize["error_code"] = o.ErrorCode
	}
	if !IsNil(o.ErrorNote) {
		toSerialize["error_note"] = o.ErrorNote
	}
	if !IsNil(o.PaymentId) {
		toSerialize["payment_id"] = o.PaymentId
	}
	if !IsNil(o.PaymentStatus) {
		toSerialize["payment_status"] = o.PaymentStatus
	}
	return toSerialize, nil
}

type NullablePaymentStatusResponse struct {
	value *PaymentStatusResponse
	isSet bool
}

func (v NullablePaymentStatusResponse) Get() *PaymentStatusResponse {
	return v.value
}

func (v *NullablePaymentStatusResponse) Set(val *PaymentStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentStatusResponse(val *PaymentStatusResponse) *NullablePaymentStatusResponse {
	return &NullablePaymentStatusResponse{value: val, isSet: true}
}

func (v NullablePaymentStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
