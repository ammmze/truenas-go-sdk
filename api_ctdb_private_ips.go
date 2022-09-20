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
	"strings"
)

// CtdbPrivateIpsApiService CtdbPrivateIpsApi service
type CtdbPrivateIpsApiService service

type CtdbPrivateIpsApiCtdbPrivateIpsGetRequest struct {
	ctx        context.Context
	ApiService *CtdbPrivateIpsApiService
	limit      *int32
	offset     *int32
	count      *bool
	sort       *string
}

func (r CtdbPrivateIpsApiCtdbPrivateIpsGetRequest) Limit(limit int32) CtdbPrivateIpsApiCtdbPrivateIpsGetRequest {
	r.limit = &limit
	return r
}

func (r CtdbPrivateIpsApiCtdbPrivateIpsGetRequest) Offset(offset int32) CtdbPrivateIpsApiCtdbPrivateIpsGetRequest {
	r.offset = &offset
	return r
}

func (r CtdbPrivateIpsApiCtdbPrivateIpsGetRequest) Count(count bool) CtdbPrivateIpsApiCtdbPrivateIpsGetRequest {
	r.count = &count
	return r
}

func (r CtdbPrivateIpsApiCtdbPrivateIpsGetRequest) Sort(sort string) CtdbPrivateIpsApiCtdbPrivateIpsGetRequest {
	r.sort = &sort
	return r
}

func (r CtdbPrivateIpsApiCtdbPrivateIpsGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.CtdbPrivateIpsGetExecute(r)
}

/*
CtdbPrivateIpsGet Method for CtdbPrivateIpsGet

`query-options.extra` can be specified as query parameters with prefixing them with `extra.` prefix. For example, `extra.retrieve_properties=false` will pass `retrieve_properties` as an extra argument to pool/dataset endpoint.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return CtdbPrivateIpsApiCtdbPrivateIpsGetRequest
*/
func (a *CtdbPrivateIpsApiService) CtdbPrivateIpsGet(ctx context.Context) CtdbPrivateIpsApiCtdbPrivateIpsGetRequest {
	return CtdbPrivateIpsApiCtdbPrivateIpsGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *CtdbPrivateIpsApiService) CtdbPrivateIpsGetExecute(r CtdbPrivateIpsApiCtdbPrivateIpsGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CtdbPrivateIpsApiService.CtdbPrivateIpsGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ctdb/private/ips"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
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

type CtdbPrivateIpsApiCtdbPrivateIpsIdIdGetRequest struct {
	ctx        context.Context
	ApiService *CtdbPrivateIpsApiService
	id         int32
}

func (r CtdbPrivateIpsApiCtdbPrivateIpsIdIdGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.CtdbPrivateIpsIdIdGetExecute(r)
}

/*
CtdbPrivateIpsIdIdGet Method for CtdbPrivateIpsIdIdGet

`query-options.extra` can be specified as query parameters with prefixing them with `extra.` prefix. For example, `extra.retrieve_properties=false` will pass `retrieve_properties` as an extra argument to pool/dataset endpoint.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return CtdbPrivateIpsApiCtdbPrivateIpsIdIdGetRequest
*/
func (a *CtdbPrivateIpsApiService) CtdbPrivateIpsIdIdGet(ctx context.Context, id int32) CtdbPrivateIpsApiCtdbPrivateIpsIdIdGetRequest {
	return CtdbPrivateIpsApiCtdbPrivateIpsIdIdGetRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *CtdbPrivateIpsApiService) CtdbPrivateIpsIdIdGetExecute(r CtdbPrivateIpsApiCtdbPrivateIpsIdIdGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CtdbPrivateIpsApiService.CtdbPrivateIpsIdIdGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ctdb/private/ips/id/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

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

type CtdbPrivateIpsApiCtdbPrivateIpsIdIdPutRequest struct {
	ctx        context.Context
	ApiService *CtdbPrivateIpsApiService
	id         int32
}

func (r CtdbPrivateIpsApiCtdbPrivateIpsIdIdPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.CtdbPrivateIpsIdIdPutExecute(r)
}

/*
CtdbPrivateIpsIdIdPut Method for CtdbPrivateIpsIdIdPut

Update Private IP address from the ctdb cluster with pnn value of `id`.

`id` integer representing the PNN value for the node.
`enable` boolean. When True, enable the node else disable the node.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return CtdbPrivateIpsApiCtdbPrivateIpsIdIdPutRequest
*/
func (a *CtdbPrivateIpsApiService) CtdbPrivateIpsIdIdPut(ctx context.Context, id int32) CtdbPrivateIpsApiCtdbPrivateIpsIdIdPutRequest {
	return CtdbPrivateIpsApiCtdbPrivateIpsIdIdPutRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *CtdbPrivateIpsApiService) CtdbPrivateIpsIdIdPutExecute(r CtdbPrivateIpsApiCtdbPrivateIpsIdIdPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CtdbPrivateIpsApiService.CtdbPrivateIpsIdIdPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ctdb/private/ips/id/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

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

type CtdbPrivateIpsApiCtdbPrivateIpsPostRequest struct {
	ctx                   context.Context
	ApiService            *CtdbPrivateIpsApiService
	ctdbPrivateIpsCreate0 *CtdbPrivateIpsCreate0
}

func (r CtdbPrivateIpsApiCtdbPrivateIpsPostRequest) CtdbPrivateIpsCreate0(ctdbPrivateIpsCreate0 CtdbPrivateIpsCreate0) CtdbPrivateIpsApiCtdbPrivateIpsPostRequest {
	r.ctdbPrivateIpsCreate0 = &ctdbPrivateIpsCreate0
	return r
}

func (r CtdbPrivateIpsApiCtdbPrivateIpsPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.CtdbPrivateIpsPostExecute(r)
}

/*
CtdbPrivateIpsPost Method for CtdbPrivateIpsPost

Add a ctdb private address to the cluster

`ip` string representing an IP v4/v6 address

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return CtdbPrivateIpsApiCtdbPrivateIpsPostRequest
*/
func (a *CtdbPrivateIpsApiService) CtdbPrivateIpsPost(ctx context.Context) CtdbPrivateIpsApiCtdbPrivateIpsPostRequest {
	return CtdbPrivateIpsApiCtdbPrivateIpsPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *CtdbPrivateIpsApiService) CtdbPrivateIpsPostExecute(r CtdbPrivateIpsApiCtdbPrivateIpsPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CtdbPrivateIpsApiService.CtdbPrivateIpsPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ctdb/private/ips"

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
	localVarPostBody = r.ctdbPrivateIpsCreate0
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
