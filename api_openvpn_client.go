/*
TrueNAS RESTful API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package truenas

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

// OpenvpnClientApiService OpenvpnClientApi service
type OpenvpnClientApiService service

type OpenvpnClientApiOpenvpnClientAuthenticationAlgorithmChoicesGetRequest struct {
	ctx        context.Context
	ApiService *OpenvpnClientApiService
}

func (r OpenvpnClientApiOpenvpnClientAuthenticationAlgorithmChoicesGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.OpenvpnClientAuthenticationAlgorithmChoicesGetExecute(r)
}

/*
OpenvpnClientAuthenticationAlgorithmChoicesGet Method for OpenvpnClientAuthenticationAlgorithmChoicesGet

Returns a dictionary of valid authentication algorithms which can be used with OpenVPN server.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return OpenvpnClientApiOpenvpnClientAuthenticationAlgorithmChoicesGetRequest
*/
func (a *OpenvpnClientApiService) OpenvpnClientAuthenticationAlgorithmChoicesGet(ctx context.Context) OpenvpnClientApiOpenvpnClientAuthenticationAlgorithmChoicesGetRequest {
	return OpenvpnClientApiOpenvpnClientAuthenticationAlgorithmChoicesGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *OpenvpnClientApiService) OpenvpnClientAuthenticationAlgorithmChoicesGetExecute(r OpenvpnClientApiOpenvpnClientAuthenticationAlgorithmChoicesGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OpenvpnClientApiService.OpenvpnClientAuthenticationAlgorithmChoicesGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/openvpn/client/authentication_algorithm_choices"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type OpenvpnClientApiOpenvpnClientCipherChoicesGetRequest struct {
	ctx        context.Context
	ApiService *OpenvpnClientApiService
}

func (r OpenvpnClientApiOpenvpnClientCipherChoicesGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.OpenvpnClientCipherChoicesGetExecute(r)
}

/*
OpenvpnClientCipherChoicesGet Method for OpenvpnClientCipherChoicesGet

Returns a dictionary of valid ciphers which can be used with OpenVPN server.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return OpenvpnClientApiOpenvpnClientCipherChoicesGetRequest
*/
func (a *OpenvpnClientApiService) OpenvpnClientCipherChoicesGet(ctx context.Context) OpenvpnClientApiOpenvpnClientCipherChoicesGetRequest {
	return OpenvpnClientApiOpenvpnClientCipherChoicesGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *OpenvpnClientApiService) OpenvpnClientCipherChoicesGetExecute(r OpenvpnClientApiOpenvpnClientCipherChoicesGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OpenvpnClientApiService.OpenvpnClientCipherChoicesGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/openvpn/client/cipher_choices"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type OpenvpnClientApiOpenvpnClientGetRequest struct {
	ctx        context.Context
	ApiService *OpenvpnClientApiService
}

func (r OpenvpnClientApiOpenvpnClientGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.OpenvpnClientGetExecute(r)
}

/*
OpenvpnClientGet Method for OpenvpnClientGet



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return OpenvpnClientApiOpenvpnClientGetRequest
*/
func (a *OpenvpnClientApiService) OpenvpnClientGet(ctx context.Context) OpenvpnClientApiOpenvpnClientGetRequest {
	return OpenvpnClientApiOpenvpnClientGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *OpenvpnClientApiService) OpenvpnClientGetExecute(r OpenvpnClientApiOpenvpnClientGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OpenvpnClientApiService.OpenvpnClientGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/openvpn/client"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type OpenvpnClientApiOpenvpnClientPutRequest struct {
	ctx                  context.Context
	ApiService           *OpenvpnClientApiService
	openvpnClientUpdate0 *OpenvpnClientUpdate0
}

func (r OpenvpnClientApiOpenvpnClientPutRequest) OpenvpnClientUpdate0(openvpnClientUpdate0 OpenvpnClientUpdate0) OpenvpnClientApiOpenvpnClientPutRequest {
	r.openvpnClientUpdate0 = &openvpnClientUpdate0
	return r
}

func (r OpenvpnClientApiOpenvpnClientPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.OpenvpnClientPutExecute(r)
}

/*
OpenvpnClientPut Method for OpenvpnClientPut

Update OpenVPN Client configuration.

`remote` can be a valid ip address / domain which openvpn will try to connect to.

`nobind` must be enabled if OpenVPN client / server are to run concurrently.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return OpenvpnClientApiOpenvpnClientPutRequest
*/
func (a *OpenvpnClientApiService) OpenvpnClientPut(ctx context.Context) OpenvpnClientApiOpenvpnClientPutRequest {
	return OpenvpnClientApiOpenvpnClientPutRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *OpenvpnClientApiService) OpenvpnClientPutExecute(r OpenvpnClientApiOpenvpnClientPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OpenvpnClientApiService.OpenvpnClientPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/openvpn/client"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.openvpnClientUpdate0
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
