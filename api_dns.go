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

// DnsApiService DnsApi service
type DnsApiService service

type DnsApiDnsQueryGetRequest struct {
	ctx        context.Context
	ApiService *DnsApiService
	limit      *int32
	offset     *int32
	count      *bool
	sort       *string
}

func (r DnsApiDnsQueryGetRequest) Limit(limit int32) DnsApiDnsQueryGetRequest {
	r.limit = &limit
	return r
}

func (r DnsApiDnsQueryGetRequest) Offset(offset int32) DnsApiDnsQueryGetRequest {
	r.offset = &offset
	return r
}

func (r DnsApiDnsQueryGetRequest) Count(count bool) DnsApiDnsQueryGetRequest {
	r.count = &count
	return r
}

func (r DnsApiDnsQueryGetRequest) Sort(sort string) DnsApiDnsQueryGetRequest {
	r.sort = &sort
	return r
}

func (r DnsApiDnsQueryGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.DnsQueryGetExecute(r)
}

/*
DnsQueryGet Method for DnsQueryGet

Query Name Servers with `query-filters` and `query-options`.

`query-options.extra` can be specified as query parameters with prefixing them with `extra.` prefix. For example, `extra.retrieve_properties=false` will pass `retrieve_properties` as an extra argument to pool/dataset endpoint.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return DnsApiDnsQueryGetRequest
*/
func (a *DnsApiService) DnsQueryGet(ctx context.Context) DnsApiDnsQueryGetRequest {
	return DnsApiDnsQueryGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *DnsApiService) DnsQueryGetExecute(r DnsApiDnsQueryGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DnsApiService.DnsQueryGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dns/query"

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
