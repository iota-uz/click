/*
CLICK-PASS-API

API for managing payments and fiscalization with CLICK Pass.

API version: 1.0.0
Contact: danil@iota.uz
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clickapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type InvoiceAPI interface {

	/*
	CheckInvoiceStatus Invoice status check

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId
	@param invoiceId
	@return InvoiceAPICheckInvoiceStatusRequest
	*/
	CheckInvoiceStatus(ctx context.Context, serviceId int64, invoiceId int64) InvoiceAPICheckInvoiceStatusRequest

	// CheckInvoiceStatusExecute executes the request
	//  @return InvoiceStatusResponse
	CheckInvoiceStatusExecute(r InvoiceAPICheckInvoiceStatusRequest) (*InvoiceStatusResponse, *http.Response, error)

	/*
	CreateInvoice Create invoice

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return InvoiceAPICreateInvoiceRequest
	*/
	CreateInvoice(ctx context.Context) InvoiceAPICreateInvoiceRequest

	// CreateInvoiceExecute executes the request
	//  @return InvoiceResponse
	CreateInvoiceExecute(r InvoiceAPICreateInvoiceRequest) (*InvoiceResponse, *http.Response, error)
}

// InvoiceAPIService InvoiceAPI service
type InvoiceAPIService service

type InvoiceAPICheckInvoiceStatusRequest struct {
	ctx context.Context
	ApiService InvoiceAPI
	serviceId int64
	invoiceId int64
}

func (r InvoiceAPICheckInvoiceStatusRequest) Execute() (*InvoiceStatusResponse, *http.Response, error) {
	return r.ApiService.CheckInvoiceStatusExecute(r)
}

/*
CheckInvoiceStatus Invoice status check

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId
 @param invoiceId
 @return InvoiceAPICheckInvoiceStatusRequest
*/
func (a *InvoiceAPIService) CheckInvoiceStatus(ctx context.Context, serviceId int64, invoiceId int64) InvoiceAPICheckInvoiceStatusRequest {
	return InvoiceAPICheckInvoiceStatusRequest{
		ApiService: a,
		ctx: ctx,
		serviceId: serviceId,
		invoiceId: invoiceId,
	}
}

// Execute executes the request
//  @return InvoiceStatusResponse
func (a *InvoiceAPIService) CheckInvoiceStatusExecute(r InvoiceAPICheckInvoiceStatusRequest) (*InvoiceStatusResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InvoiceStatusResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoiceAPIService.CheckInvoiceStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/merchant/invoice/status/{service_id}/{invoice_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"service_id"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"invoice_id"+"}", url.PathEscape(parameterValueToString(r.invoiceId, "invoiceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["AuthHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Auth"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type InvoiceAPICreateInvoiceRequest struct {
	ctx context.Context
	ApiService InvoiceAPI
	invoiceRequest *InvoiceRequest
}

func (r InvoiceAPICreateInvoiceRequest) InvoiceRequest(invoiceRequest InvoiceRequest) InvoiceAPICreateInvoiceRequest {
	r.invoiceRequest = &invoiceRequest
	return r
}

func (r InvoiceAPICreateInvoiceRequest) Execute() (*InvoiceResponse, *http.Response, error) {
	return r.ApiService.CreateInvoiceExecute(r)
}

/*
CreateInvoice Create invoice

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return InvoiceAPICreateInvoiceRequest
*/
func (a *InvoiceAPIService) CreateInvoice(ctx context.Context) InvoiceAPICreateInvoiceRequest {
	return InvoiceAPICreateInvoiceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return InvoiceResponse
func (a *InvoiceAPIService) CreateInvoiceExecute(r InvoiceAPICreateInvoiceRequest) (*InvoiceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InvoiceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoiceAPIService.CreateInvoice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/merchant/invoice/create"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.invoiceRequest == nil {
		return localVarReturnValue, nil, reportError("invoiceRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.invoiceRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["AuthHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Auth"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
