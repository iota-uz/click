# CardTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** | Error code | [optional] 
**ErrorNote** | Pointer to **string** | Error description | [optional] 
**CardToken** | Pointer to **string** | Card token | [optional] 
**PhoneNumber** | Pointer to **string** | User phone number | [optional] 
**Temporary** | Pointer to **int32** | Type of created token | [optional] 

## Methods

### NewCardTokenResponse

`func NewCardTokenResponse() *CardTokenResponse`

NewCardTokenResponse instantiates a new CardTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardTokenResponseWithDefaults

`func NewCardTokenResponseWithDefaults() *CardTokenResponse`

NewCardTokenResponseWithDefaults instantiates a new CardTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *CardTokenResponse) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *CardTokenResponse) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *CardTokenResponse) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *CardTokenResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorNote

`func (o *CardTokenResponse) GetErrorNote() string`

GetErrorNote returns the ErrorNote field if non-nil, zero value otherwise.

### GetErrorNoteOk

`func (o *CardTokenResponse) GetErrorNoteOk() (*string, bool)`

GetErrorNoteOk returns a tuple with the ErrorNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorNote

`func (o *CardTokenResponse) SetErrorNote(v string)`

SetErrorNote sets ErrorNote field to given value.

### HasErrorNote

`func (o *CardTokenResponse) HasErrorNote() bool`

HasErrorNote returns a boolean if a field has been set.

### GetCardToken

`func (o *CardTokenResponse) GetCardToken() string`

GetCardToken returns the CardToken field if non-nil, zero value otherwise.

### GetCardTokenOk

`func (o *CardTokenResponse) GetCardTokenOk() (*string, bool)`

GetCardTokenOk returns a tuple with the CardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardToken

`func (o *CardTokenResponse) SetCardToken(v string)`

SetCardToken sets CardToken field to given value.

### HasCardToken

`func (o *CardTokenResponse) HasCardToken() bool`

HasCardToken returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *CardTokenResponse) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *CardTokenResponse) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *CardTokenResponse) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *CardTokenResponse) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetTemporary

`func (o *CardTokenResponse) GetTemporary() int32`

GetTemporary returns the Temporary field if non-nil, zero value otherwise.

### GetTemporaryOk

`func (o *CardTokenResponse) GetTemporaryOk() (*int32, bool)`

GetTemporaryOk returns a tuple with the Temporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporary

`func (o *CardTokenResponse) SetTemporary(v int32)`

SetTemporary sets Temporary field to given value.

### HasTemporary

`func (o *CardTokenResponse) HasTemporary() bool`

HasTemporary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


