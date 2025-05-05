# FiscalizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **int64** | Service ID | 
**PaymentId** | **int64** | Payment Identifier | 
**Items** | [**[]Item**](Item.md) | Items or services list | 
**ReceivedEcash** | Pointer to **int32** | Amount paid by e-cash in tiyins | [optional] 
**ReceivedCash** | Pointer to **int32** | Amount paid by cash in tiyins | [optional] 
**ReceivedCard** | Pointer to **int32** | Amount paid by card in tiyins | [optional] 

## Methods

### NewFiscalizationRequest

`func NewFiscalizationRequest(serviceId int64, paymentId int64, items []Item, ) *FiscalizationRequest`

NewFiscalizationRequest instantiates a new FiscalizationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiscalizationRequestWithDefaults

`func NewFiscalizationRequestWithDefaults() *FiscalizationRequest`

NewFiscalizationRequestWithDefaults instantiates a new FiscalizationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *FiscalizationRequest) GetServiceId() int64`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FiscalizationRequest) GetServiceIdOk() (*int64, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FiscalizationRequest) SetServiceId(v int64)`

SetServiceId sets ServiceId field to given value.


### GetPaymentId

`func (o *FiscalizationRequest) GetPaymentId() int64`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *FiscalizationRequest) GetPaymentIdOk() (*int64, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *FiscalizationRequest) SetPaymentId(v int64)`

SetPaymentId sets PaymentId field to given value.


### GetItems

`func (o *FiscalizationRequest) GetItems() []Item`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *FiscalizationRequest) GetItemsOk() (*[]Item, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *FiscalizationRequest) SetItems(v []Item)`

SetItems sets Items field to given value.


### GetReceivedEcash

`func (o *FiscalizationRequest) GetReceivedEcash() int32`

GetReceivedEcash returns the ReceivedEcash field if non-nil, zero value otherwise.

### GetReceivedEcashOk

`func (o *FiscalizationRequest) GetReceivedEcashOk() (*int32, bool)`

GetReceivedEcashOk returns a tuple with the ReceivedEcash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedEcash

`func (o *FiscalizationRequest) SetReceivedEcash(v int32)`

SetReceivedEcash sets ReceivedEcash field to given value.

### HasReceivedEcash

`func (o *FiscalizationRequest) HasReceivedEcash() bool`

HasReceivedEcash returns a boolean if a field has been set.

### GetReceivedCash

`func (o *FiscalizationRequest) GetReceivedCash() int32`

GetReceivedCash returns the ReceivedCash field if non-nil, zero value otherwise.

### GetReceivedCashOk

`func (o *FiscalizationRequest) GetReceivedCashOk() (*int32, bool)`

GetReceivedCashOk returns a tuple with the ReceivedCash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedCash

`func (o *FiscalizationRequest) SetReceivedCash(v int32)`

SetReceivedCash sets ReceivedCash field to given value.

### HasReceivedCash

`func (o *FiscalizationRequest) HasReceivedCash() bool`

HasReceivedCash returns a boolean if a field has been set.

### GetReceivedCard

`func (o *FiscalizationRequest) GetReceivedCard() int32`

GetReceivedCard returns the ReceivedCard field if non-nil, zero value otherwise.

### GetReceivedCardOk

`func (o *FiscalizationRequest) GetReceivedCardOk() (*int32, bool)`

GetReceivedCardOk returns a tuple with the ReceivedCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedCard

`func (o *FiscalizationRequest) SetReceivedCard(v int32)`

SetReceivedCard sets ReceivedCard field to given value.

### HasReceivedCard

`func (o *FiscalizationRequest) HasReceivedCard() bool`

HasReceivedCard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


