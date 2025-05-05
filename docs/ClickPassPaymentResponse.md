# ClickPassPaymentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** | Error code | [optional] 
**ErrorNote** | Pointer to **string** | Error description | [optional] 
**PaymentId** | Pointer to **int64** | Payment Identifier | [optional] 
**PaymentStatus** | Pointer to **int32** | Payment status code | [optional] 
**ConfirmMode** | Pointer to **int32** | Confirmation mode status | [optional] 
**CardType** | Pointer to **string** | Card type | [optional] 
**ProcessingType** | Pointer to **string** | Card processing | [optional] 
**CardNumber** | Pointer to **string** | Masked card number | [optional] 
**PhoneNumber** | Pointer to **string** | Phone number | [optional] 

## Methods

### NewClickPassPaymentResponse

`func NewClickPassPaymentResponse() *ClickPassPaymentResponse`

NewClickPassPaymentResponse instantiates a new ClickPassPaymentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClickPassPaymentResponseWithDefaults

`func NewClickPassPaymentResponseWithDefaults() *ClickPassPaymentResponse`

NewClickPassPaymentResponseWithDefaults instantiates a new ClickPassPaymentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *ClickPassPaymentResponse) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ClickPassPaymentResponse) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ClickPassPaymentResponse) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ClickPassPaymentResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorNote

`func (o *ClickPassPaymentResponse) GetErrorNote() string`

GetErrorNote returns the ErrorNote field if non-nil, zero value otherwise.

### GetErrorNoteOk

`func (o *ClickPassPaymentResponse) GetErrorNoteOk() (*string, bool)`

GetErrorNoteOk returns a tuple with the ErrorNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorNote

`func (o *ClickPassPaymentResponse) SetErrorNote(v string)`

SetErrorNote sets ErrorNote field to given value.

### HasErrorNote

`func (o *ClickPassPaymentResponse) HasErrorNote() bool`

HasErrorNote returns a boolean if a field has been set.

### GetPaymentId

`func (o *ClickPassPaymentResponse) GetPaymentId() int64`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *ClickPassPaymentResponse) GetPaymentIdOk() (*int64, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *ClickPassPaymentResponse) SetPaymentId(v int64)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *ClickPassPaymentResponse) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *ClickPassPaymentResponse) GetPaymentStatus() int32`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *ClickPassPaymentResponse) GetPaymentStatusOk() (*int32, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *ClickPassPaymentResponse) SetPaymentStatus(v int32)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *ClickPassPaymentResponse) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetConfirmMode

`func (o *ClickPassPaymentResponse) GetConfirmMode() int32`

GetConfirmMode returns the ConfirmMode field if non-nil, zero value otherwise.

### GetConfirmModeOk

`func (o *ClickPassPaymentResponse) GetConfirmModeOk() (*int32, bool)`

GetConfirmModeOk returns a tuple with the ConfirmMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmMode

`func (o *ClickPassPaymentResponse) SetConfirmMode(v int32)`

SetConfirmMode sets ConfirmMode field to given value.

### HasConfirmMode

`func (o *ClickPassPaymentResponse) HasConfirmMode() bool`

HasConfirmMode returns a boolean if a field has been set.

### GetCardType

`func (o *ClickPassPaymentResponse) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *ClickPassPaymentResponse) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *ClickPassPaymentResponse) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *ClickPassPaymentResponse) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### GetProcessingType

`func (o *ClickPassPaymentResponse) GetProcessingType() string`

GetProcessingType returns the ProcessingType field if non-nil, zero value otherwise.

### GetProcessingTypeOk

`func (o *ClickPassPaymentResponse) GetProcessingTypeOk() (*string, bool)`

GetProcessingTypeOk returns a tuple with the ProcessingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingType

`func (o *ClickPassPaymentResponse) SetProcessingType(v string)`

SetProcessingType sets ProcessingType field to given value.

### HasProcessingType

`func (o *ClickPassPaymentResponse) HasProcessingType() bool`

HasProcessingType returns a boolean if a field has been set.

### GetCardNumber

`func (o *ClickPassPaymentResponse) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *ClickPassPaymentResponse) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *ClickPassPaymentResponse) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *ClickPassPaymentResponse) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *ClickPassPaymentResponse) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ClickPassPaymentResponse) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ClickPassPaymentResponse) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ClickPassPaymentResponse) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


