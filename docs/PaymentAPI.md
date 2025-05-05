# \PaymentAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckPaymentStatus**](PaymentAPI.md#CheckPaymentStatus) | **Get** /v2/merchant/payment/status/{service_id}/{payment_id} | Payment status check
[**CheckPaymentStatusByMTI**](PaymentAPI.md#CheckPaymentStatusByMTI) | **Get** /v2/merchant/payment/status_by_mti/{service_id}/{merchant_trans_id} | Payment status check by merchant_trans_id
[**ConfirmPayment**](PaymentAPI.md#ConfirmPayment) | **Post** /v2/merchant/click_pass/confirm | Payment confirmation
[**CreatePaymentWithClickPass**](PaymentAPI.md#CreatePaymentWithClickPass) | **Post** /v2/merchant/click_pass/payment | Payment with CLICK Pass
[**DisableConfirmationMode**](PaymentAPI.md#DisableConfirmationMode) | **Delete** /v2/merchant/click_pass/confirmation/{service_id} | Disable confirmation mode
[**EnableConfirmationMode**](PaymentAPI.md#EnableConfirmationMode) | **Put** /v2/merchant/click_pass/confirmation/{service_id} | Enable confirmation mode
[**PartialRefund**](PaymentAPI.md#PartialRefund) | **Delete** /v2/merchant/payment/partial_reversal/{service_id}/{payment_id}/{amount} | Partial refund
[**PaymentWithToken**](PaymentAPI.md#PaymentWithToken) | **Post** /v2/merchant/card_token/payment | Payment with token
[**ReversePayment**](PaymentAPI.md#ReversePayment) | **Delete** /v2/merchant/payment/reversal/{service_id}/{payment_id} | Payment reversal (cancel)



## CheckPaymentStatus

> PaymentStatusResponse CheckPaymentStatus(ctx, serviceId, paymentId).Auth(auth).Execute()

Payment status check

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
	auth := "auth_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentAPI.CheckPaymentStatus(context.Background(), serviceId, paymentId).Auth(auth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.CheckPaymentStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckPaymentStatus`: PaymentStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.CheckPaymentStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **int64** |  | 
**paymentId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckPaymentStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **auth** | **string** |  | 

### Return type

[**PaymentStatusResponse**](PaymentStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckPaymentStatusByMTI

> PaymentStatusByMTIResponse CheckPaymentStatusByMTI(ctx, serviceId, merchantTransId).Auth(auth).Execute()

Payment status check by merchant_trans_id

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
	merchantTransId := "merchantTransId_example" // string | 
	auth := "auth_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentAPI.CheckPaymentStatusByMTI(context.Background(), serviceId, merchantTransId).Auth(auth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.CheckPaymentStatusByMTI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckPaymentStatusByMTI`: PaymentStatusByMTIResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.CheckPaymentStatusByMTI`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **int64** |  | 
**merchantTransId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckPaymentStatusByMTIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **auth** | **string** |  | 

### Return type

[**PaymentStatusByMTIResponse**](PaymentStatusByMTIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfirmPayment

> PaymentConfirmationResponse ConfirmPayment(ctx).Auth(auth).PaymentConfirmationRequest(paymentConfirmationRequest).Execute()

Payment confirmation

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
	paymentConfirmationRequest := *openapiclient.NewPaymentConfirmationRequest(int64(123), int64(123)) // PaymentConfirmationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentAPI.ConfirmPayment(context.Background()).Auth(auth).PaymentConfirmationRequest(paymentConfirmationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.ConfirmPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfirmPayment`: PaymentConfirmationResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.ConfirmPayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfirmPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **auth** | **string** |  | 
 **paymentConfirmationRequest** | [**PaymentConfirmationRequest**](PaymentConfirmationRequest.md) |  | 

### Return type

[**PaymentConfirmationResponse**](PaymentConfirmationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePaymentWithClickPass

> ClickPassPaymentResponse CreatePaymentWithClickPass(ctx).Auth(auth).ClickPassPaymentRequest(clickPassPaymentRequest).Execute()

Payment with CLICK Pass

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
	clickPassPaymentRequest := *openapiclient.NewClickPassPaymentRequest(int64(123), "OtpData_example", float32(123)) // ClickPassPaymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentAPI.CreatePaymentWithClickPass(context.Background()).Auth(auth).ClickPassPaymentRequest(clickPassPaymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.CreatePaymentWithClickPass``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePaymentWithClickPass`: ClickPassPaymentResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.CreatePaymentWithClickPass`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentWithClickPassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **auth** | **string** |  | 
 **clickPassPaymentRequest** | [**ClickPassPaymentRequest**](ClickPassPaymentRequest.md) |  | 

### Return type

[**ClickPassPaymentResponse**](ClickPassPaymentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableConfirmationMode

> ConfirmationModeResponse DisableConfirmationMode(ctx, serviceId).Auth(auth).Execute()

Disable confirmation mode

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
	auth := "auth_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentAPI.DisableConfirmationMode(context.Background(), serviceId).Auth(auth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.DisableConfirmationMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisableConfirmationMode`: ConfirmationModeResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.DisableConfirmationMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableConfirmationModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **auth** | **string** |  | 

### Return type

[**ConfirmationModeResponse**](ConfirmationModeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableConfirmationMode

> ConfirmationModeResponse EnableConfirmationMode(ctx, serviceId).Auth(auth).Execute()

Enable confirmation mode

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
	auth := "auth_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentAPI.EnableConfirmationMode(context.Background(), serviceId).Auth(auth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.EnableConfirmationMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableConfirmationMode`: ConfirmationModeResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.EnableConfirmationMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableConfirmationModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **auth** | **string** |  | 

### Return type

[**ConfirmationModeResponse**](ConfirmationModeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialRefund

> PartialRefundResponse PartialRefund(ctx, serviceId, paymentId, amount).Auth(auth).Execute()

Partial refund

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
	amount := float32(3.4) // float32 | 
	auth := "auth_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentAPI.PartialRefund(context.Background(), serviceId, paymentId, amount).Auth(auth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.PartialRefund``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialRefund`: PartialRefundResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.PartialRefund`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **int64** |  | 
**paymentId** | **int64** |  | 
**amount** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialRefundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **auth** | **string** |  | 

### Return type

[**PartialRefundResponse**](PartialRefundResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentWithToken

> TokenPaymentResponse PaymentWithToken(ctx).Auth(auth).TokenPaymentRequest(tokenPaymentRequest).Execute()

Payment with token

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
	tokenPaymentRequest := *openapiclient.NewTokenPaymentRequest(int64(123), "CardToken_example", float32(123), "TransactionParameter_example") // TokenPaymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentAPI.PaymentWithToken(context.Background()).Auth(auth).TokenPaymentRequest(tokenPaymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.PaymentWithToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentWithToken`: TokenPaymentResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.PaymentWithToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentWithTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **auth** | **string** |  | 
 **tokenPaymentRequest** | [**TokenPaymentRequest**](TokenPaymentRequest.md) |  | 

### Return type

[**TokenPaymentResponse**](TokenPaymentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReversePayment

> PaymentReversalResponse ReversePayment(ctx, serviceId, paymentId).Auth(auth).Execute()

Payment reversal (cancel)

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
	auth := "auth_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentAPI.ReversePayment(context.Background(), serviceId, paymentId).Auth(auth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.ReversePayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReversePayment`: PaymentReversalResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.ReversePayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **int64** |  | 
**paymentId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReversePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **auth** | **string** |  | 

### Return type

[**PaymentReversalResponse**](PaymentReversalResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

