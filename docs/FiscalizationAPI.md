# \FiscalizationAPI

All URIs are relative to *https://api.click.uz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFiscalData**](FiscalizationAPI.md#GetFiscalData) | **Get** /v2/merchant/payment/ofd_data/{service_id}/{payment_id} | Retrieving fiscal data (URL)
[**SubmitFiscalItems**](FiscalizationAPI.md#SubmitFiscalItems) | **Post** /v2/merchant/payment/ofd_data/submit_items | Fiscalization of goods and services
[**SubmitFiscalQRCode**](FiscalizationAPI.md#SubmitFiscalQRCode) | **Post** /v2/merchant/payment/ofd_data/submit_qrcode | Registering already fiscalized check



## GetFiscalData

> FiscalDataResponse GetFiscalData(ctx, serviceId, paymentId).Execute()

Retrieving fiscal data (URL)

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
	paymentId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalizationAPI.GetFiscalData(context.Background(), serviceId, paymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalizationAPI.GetFiscalData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFiscalData`: FiscalDataResponse
	fmt.Fprintf(os.Stdout, "Response from `FiscalizationAPI.GetFiscalData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **int64** |  | 
**paymentId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFiscalDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FiscalDataResponse**](FiscalDataResponse.md)

### Authorization

[AuthHeader](../README.md#AuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitFiscalItems

> FiscalizationResponse SubmitFiscalItems(ctx).FiscalizationRequest(fiscalizationRequest).Execute()

Fiscalization of goods and services

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
	fiscalizationRequest := *openapiclient.NewFiscalizationRequest(int64(123), int64(123), []openapiclient.Item{*openapiclient.NewItem("Name_example", "SPIC_example", int64(123), int64(123))}) // FiscalizationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalizationAPI.SubmitFiscalItems(context.Background()).FiscalizationRequest(fiscalizationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalizationAPI.SubmitFiscalItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitFiscalItems`: FiscalizationResponse
	fmt.Fprintf(os.Stdout, "Response from `FiscalizationAPI.SubmitFiscalItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitFiscalItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fiscalizationRequest** | [**FiscalizationRequest**](FiscalizationRequest.md) |  | 

### Return type

[**FiscalizationResponse**](FiscalizationResponse.md)

### Authorization

[AuthHeader](../README.md#AuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitFiscalQRCode

> FiscalQRCodeResponse SubmitFiscalQRCode(ctx).FiscalQRCodeRequest(fiscalQRCodeRequest).Execute()

Registering already fiscalized check

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
	fiscalQRCodeRequest := *openapiclient.NewFiscalQRCodeRequest(int64(123), int64(123), "Qrcode_example") // FiscalQRCodeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalizationAPI.SubmitFiscalQRCode(context.Background()).FiscalQRCodeRequest(fiscalQRCodeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalizationAPI.SubmitFiscalQRCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitFiscalQRCode`: FiscalQRCodeResponse
	fmt.Fprintf(os.Stdout, "Response from `FiscalizationAPI.SubmitFiscalQRCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitFiscalQRCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fiscalQRCodeRequest** | [**FiscalQRCodeRequest**](FiscalQRCodeRequest.md) |  | 

### Return type

[**FiscalQRCodeResponse**](FiscalQRCodeResponse.md)

### Authorization

[AuthHeader](../README.md#AuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

