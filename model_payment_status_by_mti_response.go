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

// checks if the PaymentStatusByMTIResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentStatusByMTIResponse{}

// PaymentStatusByMTIResponse struct for PaymentStatusByMTIResponse
type PaymentStatusByMTIResponse struct {
	// Error code
	ErrorCode *int32 `form:"error_code" json:"error_code,omitempty"` // Error code
	// Error description
	ErrorNote *string `form:"error_note" json:"error_note,omitempty"` // Error description
	// Payment Identifier
	PaymentId *int64 `form:"payment_id" json:"payment_id,omitempty"` // Payment Identifier
	// Merchant transaction ID
	MerchantTransId *string `form:"merchant_trans_id" json:"merchant_trans_id,omitempty"` // Merchant transaction ID
}

// NewPaymentStatusByMTIResponse instantiates a new PaymentStatusByMTIResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentStatusByMTIResponse() *PaymentStatusByMTIResponse {
	this := PaymentStatusByMTIResponse{}
	return &this
}

// NewPaymentStatusByMTIResponseWithDefaults instantiates a new PaymentStatusByMTIResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentStatusByMTIResponseWithDefaults() *PaymentStatusByMTIResponse {
	this := PaymentStatusByMTIResponse{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *PaymentStatusByMTIResponse) GetErrorCode() int32 {
	if o == nil || IsNil(o.ErrorCode) {
		var ret int32
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentStatusByMTIResponse) GetErrorCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *PaymentStatusByMTIResponse) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int32 and assigns it to the ErrorCode field.
func (o *PaymentStatusByMTIResponse) SetErrorCode(v int32) {
	o.ErrorCode = &v
}

// GetErrorNote returns the ErrorNote field value if set, zero value otherwise.
func (o *PaymentStatusByMTIResponse) GetErrorNote() string {
	if o == nil || IsNil(o.ErrorNote) {
		var ret string
		return ret
	}
	return *o.ErrorNote
}

// GetErrorNoteOk returns a tuple with the ErrorNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentStatusByMTIResponse) GetErrorNoteOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorNote) {
		return nil, false
	}
	return o.ErrorNote, true
}

// HasErrorNote returns a boolean if a field has been set.
func (o *PaymentStatusByMTIResponse) HasErrorNote() bool {
	if o != nil && !IsNil(o.ErrorNote) {
		return true
	}

	return false
}

// SetErrorNote gets a reference to the given string and assigns it to the ErrorNote field.
func (o *PaymentStatusByMTIResponse) SetErrorNote(v string) {
	o.ErrorNote = &v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *PaymentStatusByMTIResponse) GetPaymentId() int64 {
	if o == nil || IsNil(o.PaymentId) {
		var ret int64
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentStatusByMTIResponse) GetPaymentIdOk() (*int64, bool) {
	if o == nil || IsNil(o.PaymentId) {
		return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *PaymentStatusByMTIResponse) HasPaymentId() bool {
	if o != nil && !IsNil(o.PaymentId) {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given int64 and assigns it to the PaymentId field.
func (o *PaymentStatusByMTIResponse) SetPaymentId(v int64) {
	o.PaymentId = &v
}

// GetMerchantTransId returns the MerchantTransId field value if set, zero value otherwise.
func (o *PaymentStatusByMTIResponse) GetMerchantTransId() string {
	if o == nil || IsNil(o.MerchantTransId) {
		var ret string
		return ret
	}
	return *o.MerchantTransId
}

// GetMerchantTransIdOk returns a tuple with the MerchantTransId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentStatusByMTIResponse) GetMerchantTransIdOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantTransId) {
		return nil, false
	}
	return o.MerchantTransId, true
}

// HasMerchantTransId returns a boolean if a field has been set.
func (o *PaymentStatusByMTIResponse) HasMerchantTransId() bool {
	if o != nil && !IsNil(o.MerchantTransId) {
		return true
	}

	return false
}

// SetMerchantTransId gets a reference to the given string and assigns it to the MerchantTransId field.
func (o *PaymentStatusByMTIResponse) SetMerchantTransId(v string) {
	o.MerchantTransId = &v
}

func (o PaymentStatusByMTIResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentStatusByMTIResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.MerchantTransId) {
		toSerialize["merchant_trans_id"] = o.MerchantTransId
	}
	return toSerialize, nil
}

type NullablePaymentStatusByMTIResponse struct {
	value *PaymentStatusByMTIResponse
	isSet bool
}

func (v NullablePaymentStatusByMTIResponse) Get() *PaymentStatusByMTIResponse {
	return v.value
}

func (v *NullablePaymentStatusByMTIResponse) Set(val *PaymentStatusByMTIResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentStatusByMTIResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentStatusByMTIResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentStatusByMTIResponse(val *PaymentStatusByMTIResponse) *NullablePaymentStatusByMTIResponse {
	return &NullablePaymentStatusByMTIResponse{value: val, isSet: true}
}

func (v NullablePaymentStatusByMTIResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentStatusByMTIResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
