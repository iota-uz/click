# CardTokenVerifyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **int64** | Service ID | 
**CardToken** | **string** | Card token | 
**SmsCode** | **int32** | Received SMS code | 

## Methods

### NewCardTokenVerifyRequest

`func NewCardTokenVerifyRequest(serviceId int64, cardToken string, smsCode int32, ) *CardTokenVerifyRequest`

NewCardTokenVerifyRequest instantiates a new CardTokenVerifyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardTokenVerifyRequestWithDefaults

`func NewCardTokenVerifyRequestWithDefaults() *CardTokenVerifyRequest`

NewCardTokenVerifyRequestWithDefaults instantiates a new CardTokenVerifyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *CardTokenVerifyRequest) GetServiceId() int64`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CardTokenVerifyRequest) GetServiceIdOk() (*int64, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CardTokenVerifyRequest) SetServiceId(v int64)`

SetServiceId sets ServiceId field to given value.


### GetCardToken

`func (o *CardTokenVerifyRequest) GetCardToken() string`

GetCardToken returns the CardToken field if non-nil, zero value otherwise.

### GetCardTokenOk

`func (o *CardTokenVerifyRequest) GetCardTokenOk() (*string, bool)`

GetCardTokenOk returns a tuple with the CardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardToken

`func (o *CardTokenVerifyRequest) SetCardToken(v string)`

SetCardToken sets CardToken field to given value.


### GetSmsCode

`func (o *CardTokenVerifyRequest) GetSmsCode() int32`

GetSmsCode returns the SmsCode field if non-nil, zero value otherwise.

### GetSmsCodeOk

`func (o *CardTokenVerifyRequest) GetSmsCodeOk() (*int32, bool)`

GetSmsCodeOk returns a tuple with the SmsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsCode

`func (o *CardTokenVerifyRequest) SetSmsCode(v int32)`

SetSmsCode sets SmsCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


