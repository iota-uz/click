# \CardAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCardToken**](CardAPI.md#CreateCardToken) | **Post** /v2/merchant/card_token/request | Create card token
[**DeleteCardToken**](CardAPI.md#DeleteCardToken) | **Delete** /v2/merchant/card_token/{service_id}/{card_token} | Delete card token
[**VerifyCardToken**](CardAPI.md#VerifyCardToken) | **Post** /v2/merchant/card_token/verify | Verify card token



## CreateCardToken

> CardTokenResponse CreateCardToken(ctx).Auth(auth).CardTokenRequest(cardTokenRequest).Execute()

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
	auth := "auth_example" // string | 
	cardTokenRequest := *openapiclient.NewCardTokenRequest(int64(123), "CardNumber_example", "ExpireDate_example", int32(123)) // CardTokenRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CardAPI.CreateCardToken(context.Background()).Auth(auth).CardTokenRequest(cardTokenRequest).Execute()
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
 **auth** | **string** |  | 
 **cardTokenRequest** | [**CardTokenRequest**](CardTokenRequest.md) |  | 

### Return type

[**CardTokenResponse**](CardTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCardToken

> DeleteCardTokenResponse DeleteCardToken(ctx, serviceId, cardToken).Auth(auth).Execute()

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
	auth := "auth_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CardAPI.DeleteCardToken(context.Background(), serviceId, cardToken).Auth(auth).Execute()
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


 **auth** | **string** |  | 

### Return type

[**DeleteCardTokenResponse**](DeleteCardTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyCardToken

> CardTokenVerifyResponse VerifyCardToken(ctx).Auth(auth).CardTokenVerifyRequest(cardTokenVerifyRequest).Execute()

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
	auth := "auth_example" // string | 
	cardTokenVerifyRequest := *openapiclient.NewCardTokenVerifyRequest(int64(123), "CardToken_example", int32(123)) // CardTokenVerifyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CardAPI.VerifyCardToken(context.Background()).Auth(auth).CardTokenVerifyRequest(cardTokenVerifyRequest).Execute()
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
 **auth** | **string** |  | 
 **cardTokenVerifyRequest** | [**CardTokenVerifyRequest**](CardTokenVerifyRequest.md) |  | 

### Return type

[**CardTokenVerifyResponse**](CardTokenVerifyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

