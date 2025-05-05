# TokenPaymentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** | Error code | [optional] 
**ErrorNote** | Pointer to **string** | Error description | [optional] 
**PaymentId** | Pointer to **int64** | Payment Identifier | [optional] 
**PaymentStatus** | Pointer to **int32** | Payment status code | [optional] 

## Methods

### NewTokenPaymentResponse

`func NewTokenPaymentResponse() *TokenPaymentResponse`

NewTokenPaymentResponse instantiates a new TokenPaymentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenPaymentResponseWithDefaults

`func NewTokenPaymentResponseWithDefaults() *TokenPaymentResponse`

NewTokenPaymentResponseWithDefaults instantiates a new TokenPaymentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *TokenPaymentResponse) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *TokenPaymentResponse) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *TokenPaymentResponse) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *TokenPaymentResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorNote

`func (o *TokenPaymentResponse) GetErrorNote() string`

GetErrorNote returns the ErrorNote field if non-nil, zero value otherwise.

### GetErrorNoteOk

`func (o *TokenPaymentResponse) GetErrorNoteOk() (*string, bool)`

GetErrorNoteOk returns a tuple with the ErrorNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorNote

`func (o *TokenPaymentResponse) SetErrorNote(v string)`

SetErrorNote sets ErrorNote field to given value.

### HasErrorNote

`func (o *TokenPaymentResponse) HasErrorNote() bool`

HasErrorNote returns a boolean if a field has been set.

### GetPaymentId

`func (o *TokenPaymentResponse) GetPaymentId() int64`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *TokenPaymentResponse) GetPaymentIdOk() (*int64, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *TokenPaymentResponse) SetPaymentId(v int64)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *TokenPaymentResponse) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *TokenPaymentResponse) GetPaymentStatus() int32`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *TokenPaymentResponse) GetPaymentStatusOk() (*int32, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *TokenPaymentResponse) SetPaymentStatus(v int32)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *TokenPaymentResponse) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


