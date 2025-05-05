# \InvoiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckInvoiceStatus**](InvoiceAPI.md#CheckInvoiceStatus) | **Get** /v2/merchant/invoice/status/{service_id}/{invoice_id} | Invoice status check
[**CreateInvoice**](InvoiceAPI.md#CreateInvoice) | **Post** /v2/merchant/invoice/create | Create invoice



## CheckInvoiceStatus

> InvoiceStatusResponse CheckInvoiceStatus(ctx, serviceId, invoiceId).Auth(auth).Execute()

Invoice status check

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/iota-uz/click/clickapi"
)

func main() {
	serviceId := int64(789) // int64 | 
	invoiceId := int64(789) // int64 | 
	auth := "auth_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceAPI.CheckInvoiceStatus(context.Background(), serviceId, invoiceId).Auth(auth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceAPI.CheckInvoiceStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckInvoiceStatus`: InvoiceStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoiceAPI.CheckInvoiceStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **int64** |  | 
**invoiceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckInvoiceStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **auth** | **string** |  | 

### Return type

[**InvoiceStatusResponse**](InvoiceStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInvoice

> InvoiceResponse CreateInvoice(ctx).Auth(auth).InvoiceRequest(invoiceRequest).Execute()

Create invoice

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/iota-uz/click/clickapi"
)

func main() {
	auth := "auth_example" // string | 
	invoiceRequest := *openapiclient.NewInvoiceRequest(int64(123), float32(123), "PhoneNumber_example", "MerchantTransId_example") // InvoiceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceAPI.CreateInvoice(context.Background()).Auth(auth).InvoiceRequest(invoiceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceAPI.CreateInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInvoice`: InvoiceResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoiceAPI.CreateInvoice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **auth** | **string** |  | 
 **invoiceRequest** | [**InvoiceRequest**](InvoiceRequest.md) |  | 

### Return type

[**InvoiceResponse**](InvoiceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

