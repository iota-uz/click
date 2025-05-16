# CompleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClickTransId** | **int64** | Payment ID in CLICK system. | 
**ServiceId** | **int32** | ID of the service. | 
**ClickPaydocId** | **int64** | Payment number in CLICK system. Shown in SMS to customer. | 
**MerchantTransId** | **string** | Order ID / personal account / login in the supplier billing system. | 
**MerchantPrepareId** | **int32** | Payment ID received during Prepare stage for confirmation. | 
**Amount** | **float32** | Payment amount (in soums). | 
**Action** | **int32** | Action to perform. 1 – for Complete stage. | 
**Error** | **int32** | Status code for completion of payment. 0 – success. Otherwise, an error code. | 
**ErrorNote** | **string** | Description of the error code or result. | 
**SignTime** | **string** | Payment date in format: &#39;YYYY-MM-DD HH:mm:ss&#39;. | 
**SignString** | **string** | Verification string. MD5 hash of: click_trans_id + service_id + SECRET_KEY + merchant_trans_id + merchant_prepare_id + amount + action + sign_time. | 

## Methods

### NewCompleteRequest

`func NewCompleteRequest(clickTransId int64, serviceId int32, clickPaydocId int64, merchantTransId string, merchantPrepareId int32, amount float32, action int32, error_ int32, errorNote string, signTime string, signString string, ) *CompleteRequest`

NewCompleteRequest instantiates a new CompleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompleteRequestWithDefaults

`func NewCompleteRequestWithDefaults() *CompleteRequest`

NewCompleteRequestWithDefaults instantiates a new CompleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClickTransId

`func (o *CompleteRequest) GetClickTransId() int64`

GetClickTransId returns the ClickTransId field if non-nil, zero value otherwise.

### GetClickTransIdOk

`func (o *CompleteRequest) GetClickTransIdOk() (*int64, bool)`

GetClickTransIdOk returns a tuple with the ClickTransId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickTransId

`func (o *CompleteRequest) SetClickTransId(v int64)`

SetClickTransId sets ClickTransId field to given value.


### GetServiceId

`func (o *CompleteRequest) GetServiceId() int32`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CompleteRequest) GetServiceIdOk() (*int32, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CompleteRequest) SetServiceId(v int32)`

SetServiceId sets ServiceId field to given value.


### GetClickPaydocId

`func (o *CompleteRequest) GetClickPaydocId() int64`

GetClickPaydocId returns the ClickPaydocId field if non-nil, zero value otherwise.

### GetClickPaydocIdOk

`func (o *CompleteRequest) GetClickPaydocIdOk() (*int64, bool)`

GetClickPaydocIdOk returns a tuple with the ClickPaydocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickPaydocId

`func (o *CompleteRequest) SetClickPaydocId(v int64)`

SetClickPaydocId sets ClickPaydocId field to given value.


### GetMerchantTransId

`func (o *CompleteRequest) GetMerchantTransId() string`

GetMerchantTransId returns the MerchantTransId field if non-nil, zero value otherwise.

### GetMerchantTransIdOk

`func (o *CompleteRequest) GetMerchantTransIdOk() (*string, bool)`

GetMerchantTransIdOk returns a tuple with the MerchantTransId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantTransId

`func (o *CompleteRequest) SetMerchantTransId(v string)`

SetMerchantTransId sets MerchantTransId field to given value.


### GetMerchantPrepareId

`func (o *CompleteRequest) GetMerchantPrepareId() int32`

GetMerchantPrepareId returns the MerchantPrepareId field if non-nil, zero value otherwise.

### GetMerchantPrepareIdOk

`func (o *CompleteRequest) GetMerchantPrepareIdOk() (*int32, bool)`

GetMerchantPrepareIdOk returns a tuple with the MerchantPrepareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantPrepareId

`func (o *CompleteRequest) SetMerchantPrepareId(v int32)`

SetMerchantPrepareId sets MerchantPrepareId field to given value.


### GetAmount

`func (o *CompleteRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CompleteRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CompleteRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetAction

`func (o *CompleteRequest) GetAction() int32`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CompleteRequest) GetActionOk() (*int32, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CompleteRequest) SetAction(v int32)`

SetAction sets Action field to given value.


### GetError

`func (o *CompleteRequest) GetError() int32`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CompleteRequest) GetErrorOk() (*int32, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CompleteRequest) SetError(v int32)`

SetError sets Error field to given value.


### GetErrorNote

`func (o *CompleteRequest) GetErrorNote() string`

GetErrorNote returns the ErrorNote field if non-nil, zero value otherwise.

### GetErrorNoteOk

`func (o *CompleteRequest) GetErrorNoteOk() (*string, bool)`

GetErrorNoteOk returns a tuple with the ErrorNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorNote

`func (o *CompleteRequest) SetErrorNote(v string)`

SetErrorNote sets ErrorNote field to given value.


### GetSignTime

`func (o *CompleteRequest) GetSignTime() string`

GetSignTime returns the SignTime field if non-nil, zero value otherwise.

### GetSignTimeOk

`func (o *CompleteRequest) GetSignTimeOk() (*string, bool)`

GetSignTimeOk returns a tuple with the SignTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignTime

`func (o *CompleteRequest) SetSignTime(v string)`

SetSignTime sets SignTime field to given value.


### GetSignString

`func (o *CompleteRequest) GetSignString() string`

GetSignString returns the SignString field if non-nil, zero value otherwise.

### GetSignStringOk

`func (o *CompleteRequest) GetSignStringOk() (*string, bool)`

GetSignStringOk returns a tuple with the SignString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignString

`func (o *CompleteRequest) SetSignString(v string)`

SetSignString sets SignString field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


