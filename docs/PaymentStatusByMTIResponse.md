# PaymentStatusByMTIResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** | Error code | [optional] 
**ErrorNote** | Pointer to **string** | Error description | [optional] 
**PaymentId** | Pointer to **int64** | Payment Identifier | [optional] 
**MerchantTransId** | Pointer to **string** | Merchant transaction ID | [optional] 

## Methods

### NewPaymentStatusByMTIResponse

`func NewPaymentStatusByMTIResponse() *PaymentStatusByMTIResponse`

NewPaymentStatusByMTIResponse instantiates a new PaymentStatusByMTIResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentStatusByMTIResponseWithDefaults

`func NewPaymentStatusByMTIResponseWithDefaults() *PaymentStatusByMTIResponse`

NewPaymentStatusByMTIResponseWithDefaults instantiates a new PaymentStatusByMTIResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *PaymentStatusByMTIResponse) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *PaymentStatusByMTIResponse) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *PaymentStatusByMTIResponse) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *PaymentStatusByMTIResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorNote

`func (o *PaymentStatusByMTIResponse) GetErrorNote() string`

GetErrorNote returns the ErrorNote field if non-nil, zero value otherwise.

### GetErrorNoteOk

`func (o *PaymentStatusByMTIResponse) GetErrorNoteOk() (*string, bool)`

GetErrorNoteOk returns a tuple with the ErrorNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorNote

`func (o *PaymentStatusByMTIResponse) SetErrorNote(v string)`

SetErrorNote sets ErrorNote field to given value.

### HasErrorNote

`func (o *PaymentStatusByMTIResponse) HasErrorNote() bool`

HasErrorNote returns a boolean if a field has been set.

### GetPaymentId

`func (o *PaymentStatusByMTIResponse) GetPaymentId() int64`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentStatusByMTIResponse) GetPaymentIdOk() (*int64, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentStatusByMTIResponse) SetPaymentId(v int64)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *PaymentStatusByMTIResponse) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetMerchantTransId

`func (o *PaymentStatusByMTIResponse) GetMerchantTransId() string`

GetMerchantTransId returns the MerchantTransId field if non-nil, zero value otherwise.

### GetMerchantTransIdOk

`func (o *PaymentStatusByMTIResponse) GetMerchantTransIdOk() (*string, bool)`

GetMerchantTransIdOk returns a tuple with the MerchantTransId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantTransId

`func (o *PaymentStatusByMTIResponse) SetMerchantTransId(v string)`

SetMerchantTransId sets MerchantTransId field to given value.

### HasMerchantTransId

`func (o *PaymentStatusByMTIResponse) HasMerchantTransId() bool`

HasMerchantTransId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


