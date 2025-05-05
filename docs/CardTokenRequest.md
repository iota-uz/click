# CardTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **int64** | Service ID | 
**CardNumber** | **string** | Card number | 
**ExpireDate** | **string** | Card expire date | 
**Temporary** | **int32** | Create temporary card | 

## Methods

### NewCardTokenRequest

`func NewCardTokenRequest(serviceId int64, cardNumber string, expireDate string, temporary int32, ) *CardTokenRequest`

NewCardTokenRequest instantiates a new CardTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardTokenRequestWithDefaults

`func NewCardTokenRequestWithDefaults() *CardTokenRequest`

NewCardTokenRequestWithDefaults instantiates a new CardTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *CardTokenRequest) GetServiceId() int64`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CardTokenRequest) GetServiceIdOk() (*int64, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CardTokenRequest) SetServiceId(v int64)`

SetServiceId sets ServiceId field to given value.


### GetCardNumber

`func (o *CardTokenRequest) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *CardTokenRequest) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *CardTokenRequest) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.


### GetExpireDate

`func (o *CardTokenRequest) GetExpireDate() string`

GetExpireDate returns the ExpireDate field if non-nil, zero value otherwise.

### GetExpireDateOk

`func (o *CardTokenRequest) GetExpireDateOk() (*string, bool)`

GetExpireDateOk returns a tuple with the ExpireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDate

`func (o *CardTokenRequest) SetExpireDate(v string)`

SetExpireDate sets ExpireDate field to given value.


### GetTemporary

`func (o *CardTokenRequest) GetTemporary() int32`

GetTemporary returns the Temporary field if non-nil, zero value otherwise.

### GetTemporaryOk

`func (o *CardTokenRequest) GetTemporaryOk() (*int32, bool)`

GetTemporaryOk returns a tuple with the Temporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporary

`func (o *CardTokenRequest) SetTemporary(v int32)`

SetTemporary sets Temporary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


