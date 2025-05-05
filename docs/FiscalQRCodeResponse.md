# FiscalQRCodeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** | Error code | [optional] 
**ErrorNote** | Pointer to **string** | Error description | [optional] 

## Methods

### NewFiscalQRCodeResponse

`func NewFiscalQRCodeResponse() *FiscalQRCodeResponse`

NewFiscalQRCodeResponse instantiates a new FiscalQRCodeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiscalQRCodeResponseWithDefaults

`func NewFiscalQRCodeResponseWithDefaults() *FiscalQRCodeResponse`

NewFiscalQRCodeResponseWithDefaults instantiates a new FiscalQRCodeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *FiscalQRCodeResponse) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *FiscalQRCodeResponse) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *FiscalQRCodeResponse) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *FiscalQRCodeResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorNote

`func (o *FiscalQRCodeResponse) GetErrorNote() string`

GetErrorNote returns the ErrorNote field if non-nil, zero value otherwise.

### GetErrorNoteOk

`func (o *FiscalQRCodeResponse) GetErrorNoteOk() (*string, bool)`

GetErrorNoteOk returns a tuple with the ErrorNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorNote

`func (o *FiscalQRCodeResponse) SetErrorNote(v string)`

SetErrorNote sets ErrorNote field to given value.

### HasErrorNote

`func (o *FiscalQRCodeResponse) HasErrorNote() bool`

HasErrorNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


