# PrepareResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClickTransId** | **int64** | Payment ID in CLICK system. | 
**MerchantTransId** | **string** | Order ID / personal account / login in the supplier billing system. | 
**MerchantPrepareId** | **int64** | Payment ID in the supplier&#39;s billing system. | 
**Error** | **int32** | Status code for completion of payment. 0 – success. Otherwise, an error code. | 
**ErrorNote** | **string** | Description of the error code or result. | 

## Methods

### NewPrepareResponse

`func NewPrepareResponse(clickTransId int64, merchantTransId string, merchantPrepareId int64, error_ int32, errorNote string, ) *PrepareResponse`

NewPrepareResponse instantiates a new PrepareResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrepareResponseWithDefaults

`func NewPrepareResponseWithDefaults() *PrepareResponse`

NewPrepareResponseWithDefaults instantiates a new PrepareResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClickTransId

`func (o *PrepareResponse) GetClickTransId() int64`

GetClickTransId returns the ClickTransId field if non-nil, zero value otherwise.

### GetClickTransIdOk

`func (o *PrepareResponse) GetClickTransIdOk() (*int64, bool)`

GetClickTransIdOk returns a tuple with the ClickTransId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickTransId

`func (o *PrepareResponse) SetClickTransId(v int64)`

SetClickTransId sets ClickTransId field to given value.


### GetMerchantTransId

`func (o *PrepareResponse) GetMerchantTransId() string`

GetMerchantTransId returns the MerchantTransId field if non-nil, zero value otherwise.

### GetMerchantTransIdOk

`func (o *PrepareResponse) GetMerchantTransIdOk() (*string, bool)`

GetMerchantTransIdOk returns a tuple with the MerchantTransId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantTransId

`func (o *PrepareResponse) SetMerchantTransId(v string)`

SetMerchantTransId sets MerchantTransId field to given value.


### GetMerchantPrepareId

`func (o *PrepareResponse) GetMerchantPrepareId() int64`

GetMerchantPrepareId returns the MerchantPrepareId field if non-nil, zero value otherwise.

### GetMerchantPrepareIdOk

`func (o *PrepareResponse) GetMerchantPrepareIdOk() (*int64, bool)`

GetMerchantPrepareIdOk returns a tuple with the MerchantPrepareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantPrepareId

`func (o *PrepareResponse) SetMerchantPrepareId(v int64)`

SetMerchantPrepareId sets MerchantPrepareId field to given value.


### GetError

`func (o *PrepareResponse) GetError() int32`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PrepareResponse) GetErrorOk() (*int32, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PrepareResponse) SetError(v int32)`

SetError sets Error field to given value.


### GetErrorNote

`func (o *PrepareResponse) GetErrorNote() string`

GetErrorNote returns the ErrorNote field if non-nil, zero value otherwise.

### GetErrorNoteOk

`func (o *PrepareResponse) GetErrorNoteOk() (*string, bool)`

GetErrorNoteOk returns a tuple with the ErrorNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorNote

`func (o *PrepareResponse) SetErrorNote(v string)`

SetErrorNote sets ErrorNote field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


