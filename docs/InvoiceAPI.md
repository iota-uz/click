# \InvoiceAPI

All URIs are relative to *https://api.click.uz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckInvoiceStatus**](InvoiceAPI.md#CheckInvoiceStatus) | **Get** /v2/merchant/invoice/status/{service_id}/{invoice_id} | Invoice status check
[**CreateInvoice**](InvoiceAPI.md#CreateInvoice) | **Post** /v2/merchant/invoice/create | Create invoice



## CheckInvoiceStatus

> InvoiceStatusResponse CheckInvoiceStatus(ctx, serviceId, invoiceId).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceAPI.CheckInvoiceStatus(context.Background(), serviceId, invoiceId).Execute()
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



### Return type

[**InvoiceStatusResponse**](InvoiceStatusResponse.md)

### Authorization

[AuthHeader](../README.md#AuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInvoice

> InvoiceResponse CreateInvoice(ctx).InvoiceRequest(invoiceRequest).Execute()

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
	invoiceRequest := *openapiclient.NewInvoiceRequest(int64(123), float64(123), "PhoneNumber_example", "MerchantTransId_example") // InvoiceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceAPI.CreateInvoice(context.Background()).InvoiceRequest(invoiceRequest).Execute()
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
 **invoiceRequest** | [**InvoiceRequest**](InvoiceRequest.md) |  | 

### Return type

[**InvoiceResponse**](InvoiceResponse.md)

### Authorization

[AuthHeader](../README.md#AuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

