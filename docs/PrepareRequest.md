# PrepareRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClickTransId** | **int64** | ID of transaction (iteration) in CLICK system, i.e. attempt to make a payment. | 
**ServiceId** | **int64** | ID of the service. | 
**ClickPaydocId** | **int64** | Payment ID in CLICK system. Displayed to the customer in SMS when paying. | 
**MerchantTransId** | **string** | Order ID / personal account / login in the supplier billing system. | 
**Amount** | **float64** | Payment amount (in soums). | 
**Action** | **int32** | Action to perform. 0 – for Prepare stage. | 
**Error** | **int32** | Status code for completion of payment. 0 – success. Otherwise, an error code. | 
**ErrorNote** | **string** | Description of the error code or result. | 
**SignTime** | **string** | Payment date in format: &#39;YYYY-MM-DD HH:mm:ss&#39;. | 
**SignString** | **string** | Verification string. MD5 hash of: click_trans_id + service_id + SECRET_KEY + merchant_trans_id + amount + action + sign_time. | 

## Methods

### NewPrepareRequest

`func NewPrepareRequest(clickTransId int64, serviceId int64, clickPaydocId int64, merchantTransId string, amount float64, action int32, error_ int32, errorNote string, signTime string, signString string, ) *PrepareRequest`

NewPrepareRequest instantiates a new PrepareRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrepareRequestWithDefaults

`func NewPrepareRequestWithDefaults() *PrepareRequest`

NewPrepareRequestWithDefaults instantiates a new PrepareRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClickTransId

`func (o *PrepareRequest) GetClickTransId() int64`

GetClickTransId returns the ClickTransId field if non-nil, zero value otherwise.

### GetClickTransIdOk

`func (o *PrepareRequest) GetClickTransIdOk() (*int64, bool)`

GetClickTransIdOk returns a tuple with the ClickTransId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickTransId

`func (o *PrepareRequest) SetClickTransId(v int64)`

SetClickTransId sets ClickTransId field to given value.


### GetServiceId

`func (o *PrepareRequest) GetServiceId() int64`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *PrepareRequest) GetServiceIdOk() (*int64, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *PrepareRequest) SetServiceId(v int64)`

SetServiceId sets ServiceId field to given value.


### GetClickPaydocId

`func (o *PrepareRequest) GetClickPaydocId() int64`

GetClickPaydocId returns the ClickPaydocId field if non-nil, zero value otherwise.

### GetClickPaydocIdOk

`func (o *PrepareRequest) GetClickPaydocIdOk() (*int64, bool)`

GetClickPaydocIdOk returns a tuple with the ClickPaydocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickPaydocId

`func (o *PrepareRequest) SetClickPaydocId(v int64)`

SetClickPaydocId sets ClickPaydocId field to given value.


### GetMerchantTransId

`func (o *PrepareRequest) GetMerchantTransId() string`

GetMerchantTransId returns the MerchantTransId field if non-nil, zero value otherwise.

### GetMerchantTransIdOk

`func (o *PrepareRequest) GetMerchantTransIdOk() (*string, bool)`

GetMerchantTransIdOk returns a tuple with the MerchantTransId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantTransId

`func (o *PrepareRequest) SetMerchantTransId(v string)`

SetMerchantTransId sets MerchantTransId field to given value.


### GetAmount

`func (o *PrepareRequest) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PrepareRequest) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PrepareRequest) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetAction

`func (o *PrepareRequest) GetAction() int32`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PrepareRequest) GetActionOk() (*int32, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PrepareRequest) SetAction(v int32)`

SetAction sets Action field to given value.


### GetError

`func (o *PrepareRequest) GetError() int32`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PrepareRequest) GetErrorOk() (*int32, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PrepareRequest) SetError(v int32)`

SetError sets Error field to given value.


### GetErrorNote

`func (o *PrepareRequest) GetErrorNote() string`

GetErrorNote returns the ErrorNote field if non-nil, zero value otherwise.

### GetErrorNoteOk

`func (o *PrepareRequest) GetErrorNoteOk() (*string, bool)`

GetErrorNoteOk returns a tuple with the ErrorNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorNote

`func (o *PrepareRequest) SetErrorNote(v string)`

SetErrorNote sets ErrorNote field to given value.


### GetSignTime

`func (o *PrepareRequest) GetSignTime() string`

GetSignTime returns the SignTime field if non-nil, zero value otherwise.

### GetSignTimeOk

`func (o *PrepareRequest) GetSignTimeOk() (*string, bool)`

GetSignTimeOk returns a tuple with the SignTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignTime

`func (o *PrepareRequest) SetSignTime(v string)`

SetSignTime sets SignTime field to given value.


### GetSignString

`func (o *PrepareRequest) GetSignString() string`

GetSignString returns the SignString field if non-nil, zero value otherwise.

### GetSignStringOk

`func (o *PrepareRequest) GetSignStringOk() (*string, bool)`

GetSignStringOk returns a tuple with the SignString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignString

`func (o *PrepareRequest) SetSignString(v string)`

SetSignString sets SignString field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


