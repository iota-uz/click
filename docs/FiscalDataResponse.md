# FiscalDataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentId** | Pointer to **int64** | Payment Id | [optional] 
**QrCodeURL** | Pointer to **string** | URL | [optional] 

## Methods

### NewFiscalDataResponse

`func NewFiscalDataResponse() *FiscalDataResponse`

NewFiscalDataResponse instantiates a new FiscalDataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiscalDataResponseWithDefaults

`func NewFiscalDataResponseWithDefaults() *FiscalDataResponse`

NewFiscalDataResponseWithDefaults instantiates a new FiscalDataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentId

`func (o *FiscalDataResponse) GetPaymentId() int64`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *FiscalDataResponse) GetPaymentIdOk() (*int64, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *FiscalDataResponse) SetPaymentId(v int64)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *FiscalDataResponse) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetQrCodeURL

`func (o *FiscalDataResponse) GetQrCodeURL() string`

GetQrCodeURL returns the QrCodeURL field if non-nil, zero value otherwise.

### GetQrCodeURLOk

`func (o *FiscalDataResponse) GetQrCodeURLOk() (*string, bool)`

GetQrCodeURLOk returns a tuple with the QrCodeURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCodeURL

`func (o *FiscalDataResponse) SetQrCodeURL(v string)`

SetQrCodeURL sets QrCodeURL field to given value.

### HasQrCodeURL

`func (o *FiscalDataResponse) HasQrCodeURL() bool`

HasQrCodeURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


