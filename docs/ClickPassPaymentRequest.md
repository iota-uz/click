# ClickPassPaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **int64** | Service ID | 
**OtpData** | **string** | QR code contents | 
**Amount** | **float32** | Payment amount | 
**CashboxCode** | Pointer to **string** | Cashbox identifier | [optional] 
**TransactionId** | Pointer to **string** | Merchant transaction ID | [optional] 

## Methods

### NewClickPassPaymentRequest

`func NewClickPassPaymentRequest(serviceId int64, otpData string, amount float32, ) *ClickPassPaymentRequest`

NewClickPassPaymentRequest instantiates a new ClickPassPaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClickPassPaymentRequestWithDefaults

`func NewClickPassPaymentRequestWithDefaults() *ClickPassPaymentRequest`

NewClickPassPaymentRequestWithDefaults instantiates a new ClickPassPaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *ClickPassPaymentRequest) GetServiceId() int64`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ClickPassPaymentRequest) GetServiceIdOk() (*int64, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ClickPassPaymentRequest) SetServiceId(v int64)`

SetServiceId sets ServiceId field to given value.


### GetOtpData

`func (o *ClickPassPaymentRequest) GetOtpData() string`

GetOtpData returns the OtpData field if non-nil, zero value otherwise.

### GetOtpDataOk

`func (o *ClickPassPaymentRequest) GetOtpDataOk() (*string, bool)`

GetOtpDataOk returns a tuple with the OtpData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpData

`func (o *ClickPassPaymentRequest) SetOtpData(v string)`

SetOtpData sets OtpData field to given value.


### GetAmount

`func (o *ClickPassPaymentRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ClickPassPaymentRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ClickPassPaymentRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCashboxCode

`func (o *ClickPassPaymentRequest) GetCashboxCode() string`

GetCashboxCode returns the CashboxCode field if non-nil, zero value otherwise.

### GetCashboxCodeOk

`func (o *ClickPassPaymentRequest) GetCashboxCodeOk() (*string, bool)`

GetCashboxCodeOk returns a tuple with the CashboxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashboxCode

`func (o *ClickPassPaymentRequest) SetCashboxCode(v string)`

SetCashboxCode sets CashboxCode field to given value.

### HasCashboxCode

`func (o *ClickPassPaymentRequest) HasCashboxCode() bool`

HasCashboxCode returns a boolean if a field has been set.

### GetTransactionId

`func (o *ClickPassPaymentRequest) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ClickPassPaymentRequest) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ClickPassPaymentRequest) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *ClickPassPaymentRequest) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


