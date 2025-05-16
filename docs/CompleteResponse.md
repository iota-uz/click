# CompleteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClickTransId** | **int64** | Payment ID in CLICK system. | 
**MerchantTransId** | **string** | Order ID / personal account / login in the supplier billing system. | 
**MerchantConfirmId** | **int64** | Transaction ID to complete the payment in the billing system. May be 0 if not applicable. | 
**Error** | **int32** | Status code for completion of payment. 0 – success. Otherwise, an error code. | 
**ErrorNote** | **string** | Description of the error code or result. | 

## Methods

### NewCompleteResponse

`func NewCompleteResponse(clickTransId int64, merchantTransId string, merchantConfirmId int64, error_ int32, errorNote string, ) *CompleteResponse`

NewCompleteResponse instantiates a new CompleteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompleteResponseWithDefaults

`func NewCompleteResponseWithDefaults() *CompleteResponse`

NewCompleteResponseWithDefaults instantiates a new CompleteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClickTransId

`func (o *CompleteResponse) GetClickTransId() int64`

GetClickTransId returns the ClickTransId field if non-nil, zero value otherwise.

### GetClickTransIdOk

`func (o *CompleteResponse) GetClickTransIdOk() (*int64, bool)`

GetClickTransIdOk returns a tuple with the ClickTransId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickTransId

`func (o *CompleteResponse) SetClickTransId(v int64)`

SetClickTransId sets ClickTransId field to given value.


### GetMerchantTransId

`func (o *CompleteResponse) GetMerchantTransId() string`

GetMerchantTransId returns the MerchantTransId field if non-nil, zero value otherwise.

### GetMerchantTransIdOk

`func (o *CompleteResponse) GetMerchantTransIdOk() (*string, bool)`

GetMerchantTransIdOk returns a tuple with the MerchantTransId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantTransId

`func (o *CompleteResponse) SetMerchantTransId(v string)`

SetMerchantTransId sets MerchantTransId field to given value.


### GetMerchantConfirmId

`func (o *CompleteResponse) GetMerchantConfirmId() int64`

GetMerchantConfirmId returns the MerchantConfirmId field if non-nil, zero value otherwise.

### GetMerchantConfirmIdOk

`func (o *CompleteResponse) GetMerchantConfirmIdOk() (*int64, bool)`

GetMerchantConfirmIdOk returns a tuple with the MerchantConfirmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantConfirmId

`func (o *CompleteResponse) SetMerchantConfirmId(v int64)`

SetMerchantConfirmId sets MerchantConfirmId field to given value.


### GetError

`func (o *CompleteResponse) GetError() int32`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CompleteResponse) GetErrorOk() (*int32, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CompleteResponse) SetError(v int32)`

SetError sets Error field to given value.


### GetErrorNote

`func (o *CompleteResponse) GetErrorNote() string`

GetErrorNote returns the ErrorNote field if non-nil, zero value otherwise.

### GetErrorNoteOk

`func (o *CompleteResponse) GetErrorNoteOk() (*string, bool)`

GetErrorNoteOk returns a tuple with the ErrorNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorNote

`func (o *CompleteResponse) SetErrorNote(v string)`

SetErrorNote sets ErrorNote field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


