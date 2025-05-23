/*
CLICK-PASS-API

API for managing payments and fiscalization with CLICK Pass.

API version: 1.0.0
Contact: danil@iota.uz
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clickapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the InvoiceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceRequest{}

// InvoiceRequest struct for InvoiceRequest
type InvoiceRequest struct {
	// Service ID
	ServiceId int64 `form:"service_id" json:"service_id"` // Service ID
	// Requested amount
	Amount float64 `form:"amount" json:"amount"` // Requested amount
	// Invoice receiver
	PhoneNumber string `form:"phone_number" json:"phone_number"` // Invoice receiver
	// Order ID or personal account
	MerchantTransId string `form:"merchant_trans_id" json:"merchant_trans_id"` // Order ID or personal account
}

type _InvoiceRequest InvoiceRequest

// NewInvoiceRequest instantiates a new InvoiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceRequest(serviceId int64, amount float64, phoneNumber string, merchantTransId string) *InvoiceRequest {
	this := InvoiceRequest{}
	this.ServiceId = serviceId
	this.Amount = amount
	this.PhoneNumber = phoneNumber
	this.MerchantTransId = merchantTransId
	return &this
}

// NewInvoiceRequestWithDefaults instantiates a new InvoiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceRequestWithDefaults() *InvoiceRequest {
	this := InvoiceRequest{}
	return &this
}

// GetServiceId returns the ServiceId field value
func (o *InvoiceRequest) GetServiceId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetServiceIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *InvoiceRequest) SetServiceId(v int64) {
	o.ServiceId = v
}

// GetAmount returns the Amount field value
func (o *InvoiceRequest) GetAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *InvoiceRequest) SetAmount(v float64) {
	o.Amount = v
}

// GetPhoneNumber returns the PhoneNumber field value
func (o *InvoiceRequest) GetPhoneNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhoneNumber, true
}

// SetPhoneNumber sets field value
func (o *InvoiceRequest) SetPhoneNumber(v string) {
	o.PhoneNumber = v
}

// GetMerchantTransId returns the MerchantTransId field value
func (o *InvoiceRequest) GetMerchantTransId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantTransId
}

// GetMerchantTransIdOk returns a tuple with the MerchantTransId field value
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetMerchantTransIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantTransId, true
}

// SetMerchantTransId sets field value
func (o *InvoiceRequest) SetMerchantTransId(v string) {
	o.MerchantTransId = v
}

func (o InvoiceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["service_id"] = o.ServiceId
	toSerialize["amount"] = o.Amount
	toSerialize["phone_number"] = o.PhoneNumber
	toSerialize["merchant_trans_id"] = o.MerchantTransId
	return toSerialize, nil
}

func (o *InvoiceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"service_id",
		"amount",
		"phone_number",
		"merchant_trans_id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varInvoiceRequest := _InvoiceRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInvoiceRequest)

	if err != nil {
		return err
	}

	*o = InvoiceRequest(varInvoiceRequest)

	return err
}

type NullableInvoiceRequest struct {
	value *InvoiceRequest
	isSet bool
}

func (v NullableInvoiceRequest) Get() *InvoiceRequest {
	return v.value
}

func (v *NullableInvoiceRequest) Set(val *InvoiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceRequest(val *InvoiceRequest) *NullableInvoiceRequest {
	return &NullableInvoiceRequest{value: val, isSet: true}
}

func (v NullableInvoiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
