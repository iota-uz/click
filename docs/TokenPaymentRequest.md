# TokenPaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **int64** | Service ID | 
**CardToken** | **string** | Card token | 
**Amount** | **float64** | Payment amount | 
**TransactionParameter** | **string** | Merchant transaction ID | 

## Methods

### NewTokenPaymentRequest

`func NewTokenPaymentRequest(serviceId int64, cardToken string, amount float64, transactionParameter string, ) *TokenPaymentRequest`

NewTokenPaymentRequest instantiates a new TokenPaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenPaymentRequestWithDefaults

`func NewTokenPaymentRequestWithDefaults() *TokenPaymentRequest`

NewTokenPaymentRequestWithDefaults instantiates a new TokenPaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *TokenPaymentRequest) GetServiceId() int64`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *TokenPaymentRequest) GetServiceIdOk() (*int64, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *TokenPaymentRequest) SetServiceId(v int64)`

SetServiceId sets ServiceId field to given value.


### GetCardToken

`func (o *TokenPaymentRequest) GetCardToken() string`

GetCardToken returns the CardToken field if non-nil, zero value otherwise.

### GetCardTokenOk

`func (o *TokenPaymentRequest) GetCardTokenOk() (*string, bool)`

GetCardTokenOk returns a tuple with the CardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardToken

`func (o *TokenPaymentRequest) SetCardToken(v string)`

SetCardToken sets CardToken field to given value.


### GetAmount

`func (o *TokenPaymentRequest) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TokenPaymentRequest) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TokenPaymentRequest) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetTransactionParameter

`func (o *TokenPaymentRequest) GetTransactionParameter() string`

GetTransactionParameter returns the TransactionParameter field if non-nil, zero value otherwise.

### GetTransactionParameterOk

`func (o *TokenPaymentRequest) GetTransactionParameterOk() (*string, bool)`

GetTransactionParameterOk returns a tuple with the TransactionParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionParameter

`func (o *TokenPaymentRequest) SetTransactionParameter(v string)`

SetTransactionParameter sets TransactionParameter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


