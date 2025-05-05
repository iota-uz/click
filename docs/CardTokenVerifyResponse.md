# CardTokenVerifyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** | Error code | [optional] 
**ErrorNote** | Pointer to **string** | Error description | [optional] 
**CardNumber** | Pointer to **string** | Card number | [optional] 

## Methods

### NewCardTokenVerifyResponse

`func NewCardTokenVerifyResponse() *CardTokenVerifyResponse`

NewCardTokenVerifyResponse instantiates a new CardTokenVerifyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardTokenVerifyResponseWithDefaults

`func NewCardTokenVerifyResponseWithDefaults() *CardTokenVerifyResponse`

NewCardTokenVerifyResponseWithDefaults instantiates a new CardTokenVerifyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *CardTokenVerifyResponse) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *CardTokenVerifyResponse) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *CardTokenVerifyResponse) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *CardTokenVerifyResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorNote

`func (o *CardTokenVerifyResponse) GetErrorNote() string`

GetErrorNote returns the ErrorNote field if non-nil, zero value otherwise.

### GetErrorNoteOk

`func (o *CardTokenVerifyResponse) GetErrorNoteOk() (*string, bool)`

GetErrorNoteOk returns a tuple with the ErrorNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorNote

`func (o *CardTokenVerifyResponse) SetErrorNote(v string)`

SetErrorNote sets ErrorNote field to given value.

### HasErrorNote

`func (o *CardTokenVerifyResponse) HasErrorNote() bool`

HasErrorNote returns a boolean if a field has been set.

### GetCardNumber

`func (o *CardTokenVerifyResponse) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *CardTokenVerifyResponse) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *CardTokenVerifyResponse) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *CardTokenVerifyResponse) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


