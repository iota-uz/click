# PaymentConfirmationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** | Error code | [optional] 
**ErrorNote** | Pointer to **string** | Payment confirmed | [optional] 

## Methods

### NewPaymentConfirmationResponse

`func NewPaymentConfirmationResponse() *PaymentConfirmationResponse`

NewPaymentConfirmationResponse instantiates a new PaymentConfirmationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentConfirmationResponseWithDefaults

`func NewPaymentConfirmationResponseWithDefaults() *PaymentConfirmationResponse`

NewPaymentConfirmationResponseWithDefaults instantiates a new PaymentConfirmationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *PaymentConfirmationResponse) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *PaymentConfirmationResponse) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *PaymentConfirmationResponse) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *PaymentConfirmationResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorNote

`func (o *PaymentConfirmationResponse) GetErrorNote() string`

GetErrorNote returns the ErrorNote field if non-nil, zero value otherwise.

### GetErrorNoteOk

`func (o *PaymentConfirmationResponse) GetErrorNoteOk() (*string, bool)`

GetErrorNoteOk returns a tuple with the ErrorNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorNote

`func (o *PaymentConfirmationResponse) SetErrorNote(v string)`

SetErrorNote sets ErrorNote field to given value.

### HasErrorNote

`func (o *PaymentConfirmationResponse) HasErrorNote() bool`

HasErrorNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


