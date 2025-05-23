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

// checks if the CardTokenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardTokenRequest{}

// CardTokenRequest struct for CardTokenRequest
type CardTokenRequest struct {
	// Service ID
	ServiceId int64 `form:"service_id" json:"service_id"` // Service ID
	// Card number
	CardNumber string `form:"card_number" json:"card_number"` // Card number
	// Card expire date
	ExpireDate string `form:"expire_date" json:"expire_date"` // Card expire date
	// Create temporary card
	Temporary int32 `form:"temporary" json:"temporary"` // Create temporary card
}

type _CardTokenRequest CardTokenRequest

// NewCardTokenRequest instantiates a new CardTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardTokenRequest(serviceId int64, cardNumber string, expireDate string, temporary int32) *CardTokenRequest {
	this := CardTokenRequest{}
	this.ServiceId = serviceId
	this.CardNumber = cardNumber
	this.ExpireDate = expireDate
	this.Temporary = temporary
	return &this
}

// NewCardTokenRequestWithDefaults instantiates a new CardTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardTokenRequestWithDefaults() *CardTokenRequest {
	this := CardTokenRequest{}
	return &this
}

// GetServiceId returns the ServiceId field value
func (o *CardTokenRequest) GetServiceId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *CardTokenRequest) GetServiceIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *CardTokenRequest) SetServiceId(v int64) {
	o.ServiceId = v
}

// GetCardNumber returns the CardNumber field value
func (o *CardTokenRequest) GetCardNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardNumber
}

// GetCardNumberOk returns a tuple with the CardNumber field value
// and a boolean to check if the value has been set.
func (o *CardTokenRequest) GetCardNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardNumber, true
}

// SetCardNumber sets field value
func (o *CardTokenRequest) SetCardNumber(v string) {
	o.CardNumber = v
}

// GetExpireDate returns the ExpireDate field value
func (o *CardTokenRequest) GetExpireDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpireDate
}

// GetExpireDateOk returns a tuple with the ExpireDate field value
// and a boolean to check if the value has been set.
func (o *CardTokenRequest) GetExpireDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpireDate, true
}

// SetExpireDate sets field value
func (o *CardTokenRequest) SetExpireDate(v string) {
	o.ExpireDate = v
}

// GetTemporary returns the Temporary field value
func (o *CardTokenRequest) GetTemporary() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Temporary
}

// GetTemporaryOk returns a tuple with the Temporary field value
// and a boolean to check if the value has been set.
func (o *CardTokenRequest) GetTemporaryOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Temporary, true
}

// SetTemporary sets field value
func (o *CardTokenRequest) SetTemporary(v int32) {
	o.Temporary = v
}

func (o CardTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardTokenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["service_id"] = o.ServiceId
	toSerialize["card_number"] = o.CardNumber
	toSerialize["expire_date"] = o.ExpireDate
	toSerialize["temporary"] = o.Temporary
	return toSerialize, nil
}

func (o *CardTokenRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"service_id",
		"card_number",
		"expire_date",
		"temporary",
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

	varCardTokenRequest := _CardTokenRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCardTokenRequest)

	if err != nil {
		return err
	}

	*o = CardTokenRequest(varCardTokenRequest)

	return err
}

type NullableCardTokenRequest struct {
	value *CardTokenRequest
	isSet bool
}

func (v NullableCardTokenRequest) Get() *CardTokenRequest {
	return v.value
}

func (v *NullableCardTokenRequest) Set(val *CardTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCardTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCardTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardTokenRequest(val *CardTokenRequest) *NullableCardTokenRequest {
	return &NullableCardTokenRequest{value: val, isSet: true}
}

func (v NullableCardTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
