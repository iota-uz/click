# PaymentConfirmationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **int64** | Service ID | 
**PaymentId** | **int64** | Payment Identifier | 

## Methods

### NewPaymentConfirmationRequest

`func NewPaymentConfirmationRequest(serviceId int64, paymentId int64, ) *PaymentConfirmationRequest`

NewPaymentConfirmationRequest instantiates a new PaymentConfirmationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentConfirmationRequestWithDefaults

`func NewPaymentConfirmationRequestWithDefaults() *PaymentConfirmationRequest`

NewPaymentConfirmationRequestWithDefaults instantiates a new PaymentConfirmationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *PaymentConfirmationRequest) GetServiceId() int64`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *PaymentConfirmationRequest) GetServiceIdOk() (*int64, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *PaymentConfirmationRequest) SetServiceId(v int64)`

SetServiceId sets ServiceId field to given value.


### GetPaymentId

`func (o *PaymentConfirmationRequest) GetPaymentId() int64`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentConfirmationRequest) GetPaymentIdOk() (*int64, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentConfirmationRequest) SetPaymentId(v int64)`

SetPaymentId sets PaymentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


