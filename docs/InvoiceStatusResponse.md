# InvoiceStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** | Error code | [optional] 
**ErrorNote** | Pointer to **string** | Error description | [optional] 
**InvoiceStatus** | Pointer to **int64** | Invoice status code | [optional] 
**InvoiceStatusNote** | Pointer to **string** | Status description | [optional] 

## Methods

### NewInvoiceStatusResponse

`func NewInvoiceStatusResponse() *InvoiceStatusResponse`

NewInvoiceStatusResponse instantiates a new InvoiceStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceStatusResponseWithDefaults

`func NewInvoiceStatusResponseWithDefaults() *InvoiceStatusResponse`

NewInvoiceStatusResponseWithDefaults instantiates a new InvoiceStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *InvoiceStatusResponse) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *InvoiceStatusResponse) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *InvoiceStatusResponse) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *InvoiceStatusResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorNote

`func (o *InvoiceStatusResponse) GetErrorNote() string`

GetErrorNote returns the ErrorNote field if non-nil, zero value otherwise.

### GetErrorNoteOk

`func (o *InvoiceStatusResponse) GetErrorNoteOk() (*string, bool)`

GetErrorNoteOk returns a tuple with the ErrorNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorNote

`func (o *InvoiceStatusResponse) SetErrorNote(v string)`

SetErrorNote sets ErrorNote field to given value.

### HasErrorNote

`func (o *InvoiceStatusResponse) HasErrorNote() bool`

HasErrorNote returns a boolean if a field has been set.

### GetInvoiceStatus

`func (o *InvoiceStatusResponse) GetInvoiceStatus() int64`

GetInvoiceStatus returns the InvoiceStatus field if non-nil, zero value otherwise.

### GetInvoiceStatusOk

`func (o *InvoiceStatusResponse) GetInvoiceStatusOk() (*int64, bool)`

GetInvoiceStatusOk returns a tuple with the InvoiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceStatus

`func (o *InvoiceStatusResponse) SetInvoiceStatus(v int64)`

SetInvoiceStatus sets InvoiceStatus field to given value.

### HasInvoiceStatus

`func (o *InvoiceStatusResponse) HasInvoiceStatus() bool`

HasInvoiceStatus returns a boolean if a field has been set.

### GetInvoiceStatusNote

`func (o *InvoiceStatusResponse) GetInvoiceStatusNote() string`

GetInvoiceStatusNote returns the InvoiceStatusNote field if non-nil, zero value otherwise.

### GetInvoiceStatusNoteOk

`func (o *InvoiceStatusResponse) GetInvoiceStatusNoteOk() (*string, bool)`

GetInvoiceStatusNoteOk returns a tuple with the InvoiceStatusNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceStatusNote

`func (o *InvoiceStatusResponse) SetInvoiceStatusNote(v string)`

SetInvoiceStatusNote sets InvoiceStatusNote field to given value.

### HasInvoiceStatusNote

`func (o *InvoiceStatusResponse) HasInvoiceStatusNote() bool`

HasInvoiceStatusNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


