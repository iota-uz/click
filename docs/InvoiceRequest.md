# InvoiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **int64** | Service ID | 
**Amount** | **float64** | Requested amount | 
**PhoneNumber** | **string** | Invoice receiver | 
**MerchantTransId** | **string** | Order ID or personal account | 

## Methods

### NewInvoiceRequest

`func NewInvoiceRequest(serviceId int64, amount float64, phoneNumber string, merchantTransId string, ) *InvoiceRequest`

NewInvoiceRequest instantiates a new InvoiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceRequestWithDefaults

`func NewInvoiceRequestWithDefaults() *InvoiceRequest`

NewInvoiceRequestWithDefaults instantiates a new InvoiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *InvoiceRequest) GetServiceId() int64`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *InvoiceRequest) GetServiceIdOk() (*int64, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *InvoiceRequest) SetServiceId(v int64)`

SetServiceId sets ServiceId field to given value.


### GetAmount

`func (o *InvoiceRequest) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InvoiceRequest) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InvoiceRequest) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetPhoneNumber

`func (o *InvoiceRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *InvoiceRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *InvoiceRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetMerchantTransId

`func (o *InvoiceRequest) GetMerchantTransId() string`

GetMerchantTransId returns the MerchantTransId field if non-nil, zero value otherwise.

### GetMerchantTransIdOk

`func (o *InvoiceRequest) GetMerchantTransIdOk() (*string, bool)`

GetMerchantTransIdOk returns a tuple with the MerchantTransId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantTransId

`func (o *InvoiceRequest) SetMerchantTransId(v string)`

SetMerchantTransId sets MerchantTransId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


