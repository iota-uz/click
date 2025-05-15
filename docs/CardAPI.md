# \CardAPI

All URIs are relative to *https://api.click.uz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCardToken**](CardAPI.md#CreateCardToken) | **Post** /v2/merchant/card_token/request | Create card token
[**DeleteCardToken**](CardAPI.md#DeleteCardToken) | **Delete** /v2/merchant/card_token/{service_id}/{card_token} | Delete card token
[**VerifyCardToken**](CardAPI.md#VerifyCardToken) | **Post** /v2/merchant/card_token/verify | Verify card token



## CreateCardToken

> CardTokenResponse CreateCardToken(ctx).CardTokenRequest(cardTokenRequest).Execute()

Create card token

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
	cardTokenRequest := *openapiclient.NewCardTokenRequest(int64(123), "CardNumber_example", "ExpireDate_example", int32(123)) // CardTokenRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CardAPI.CreateCardToken(context.Background()).CardTokenRequest(cardTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardAPI.CreateCardToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCardToken`: CardTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `CardAPI.CreateCardToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCardTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cardTokenRequest** | [**CardTokenRequest**](CardTokenRequest.md) |  | 

### Return type

[**CardTokenResponse**](CardTokenResponse.md)

### Authorization

[AuthHeader](../README.md#AuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCardToken

> DeleteCardTokenResponse DeleteCardToken(ctx, serviceId, cardToken).Execute()

Delete card token

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
	cardToken := "cardToken_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CardAPI.DeleteCardToken(context.Background(), serviceId, cardToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardAPI.DeleteCardToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCardToken`: DeleteCardTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `CardAPI.DeleteCardToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **int64** |  | 
**cardToken** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCardTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteCardTokenResponse**](DeleteCardTokenResponse.md)

### Authorization

[AuthHeader](../README.md#AuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyCardToken

> CardTokenVerifyResponse VerifyCardToken(ctx).CardTokenVerifyRequest(cardTokenVerifyRequest).Execute()

Verify card token

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
	cardTokenVerifyRequest := *openapiclient.NewCardTokenVerifyRequest(int64(123), "CardToken_example", int32(123)) // CardTokenVerifyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CardAPI.VerifyCardToken(context.Background()).CardTokenVerifyRequest(cardTokenVerifyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardAPI.VerifyCardToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyCardToken`: CardTokenVerifyResponse
	fmt.Fprintf(os.Stdout, "Response from `CardAPI.VerifyCardToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyCardTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cardTokenVerifyRequest** | [**CardTokenVerifyRequest**](CardTokenVerifyRequest.md) |  | 

### Return type

[**CardTokenVerifyResponse**](CardTokenVerifyResponse.md)

### Authorization

[AuthHeader](../README.md#AuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

