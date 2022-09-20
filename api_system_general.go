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

// SystemGeneralApiService SystemGeneralApi service
type SystemGeneralApiService service

type SystemGeneralApiSystemGeneralCountryChoicesGetRequest struct {
	ctx        context.Context
	ApiService *SystemGeneralApiService
}

func (r SystemGeneralApiSystemGeneralCountryChoicesGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.SystemGeneralCountryChoicesGetExecute(r)
}

/*
SystemGeneralCountryChoicesGet Method for SystemGeneralCountryChoicesGet

Returns country choices.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemGeneralApiSystemGeneralCountryChoicesGetRequest
*/
func (a *SystemGeneralApiService) SystemGeneralCountryChoicesGet(ctx context.Context) SystemGeneralApiSystemGeneralCountryChoicesGetRequest {
	return SystemGeneralApiSystemGeneralCountryChoicesGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *SystemGeneralApiService) SystemGeneralCountryChoicesGetExecute(r SystemGeneralApiSystemGeneralCountryChoicesGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemGeneralApiService.SystemGeneralCountryChoicesGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/general/country_choices"

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

type SystemGeneralApiSystemGeneralGetRequest struct {
	ctx        context.Context
	ApiService *SystemGeneralApiService
}

func (r SystemGeneralApiSystemGeneralGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.SystemGeneralGetExecute(r)
}

/*
SystemGeneralGet Method for SystemGeneralGet



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemGeneralApiSystemGeneralGetRequest
*/
func (a *SystemGeneralApiService) SystemGeneralGet(ctx context.Context) SystemGeneralApiSystemGeneralGetRequest {
	return SystemGeneralApiSystemGeneralGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *SystemGeneralApiService) SystemGeneralGetExecute(r SystemGeneralApiSystemGeneralGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemGeneralApiService.SystemGeneralGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/general"

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

type SystemGeneralApiSystemGeneralKbdmapChoicesGetRequest struct {
	ctx        context.Context
	ApiService *SystemGeneralApiService
}

func (r SystemGeneralApiSystemGeneralKbdmapChoicesGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.SystemGeneralKbdmapChoicesGetExecute(r)
}

/*
SystemGeneralKbdmapChoicesGet Method for SystemGeneralKbdmapChoicesGet

Returns kbdmap choices.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemGeneralApiSystemGeneralKbdmapChoicesGetRequest
*/
func (a *SystemGeneralApiService) SystemGeneralKbdmapChoicesGet(ctx context.Context) SystemGeneralApiSystemGeneralKbdmapChoicesGetRequest {
	return SystemGeneralApiSystemGeneralKbdmapChoicesGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *SystemGeneralApiService) SystemGeneralKbdmapChoicesGetExecute(r SystemGeneralApiSystemGeneralKbdmapChoicesGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemGeneralApiService.SystemGeneralKbdmapChoicesGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/general/kbdmap_choices"

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

type SystemGeneralApiSystemGeneralLanguageChoicesGetRequest struct {
	ctx        context.Context
	ApiService *SystemGeneralApiService
}

func (r SystemGeneralApiSystemGeneralLanguageChoicesGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.SystemGeneralLanguageChoicesGetExecute(r)
}

/*
SystemGeneralLanguageChoicesGet Method for SystemGeneralLanguageChoicesGet

Returns language choices.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemGeneralApiSystemGeneralLanguageChoicesGetRequest
*/
func (a *SystemGeneralApiService) SystemGeneralLanguageChoicesGet(ctx context.Context) SystemGeneralApiSystemGeneralLanguageChoicesGetRequest {
	return SystemGeneralApiSystemGeneralLanguageChoicesGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *SystemGeneralApiService) SystemGeneralLanguageChoicesGetExecute(r SystemGeneralApiSystemGeneralLanguageChoicesGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemGeneralApiService.SystemGeneralLanguageChoicesGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/general/language_choices"

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

type SystemGeneralApiSystemGeneralLocalUrlGetRequest struct {
	ctx        context.Context
	ApiService *SystemGeneralApiService
}

func (r SystemGeneralApiSystemGeneralLocalUrlGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.SystemGeneralLocalUrlGetExecute(r)
}

/*
SystemGeneralLocalUrlGet Method for SystemGeneralLocalUrlGet

Returns configured local url in the format of protocol://host:port

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemGeneralApiSystemGeneralLocalUrlGetRequest
*/
func (a *SystemGeneralApiService) SystemGeneralLocalUrlGet(ctx context.Context) SystemGeneralApiSystemGeneralLocalUrlGetRequest {
	return SystemGeneralApiSystemGeneralLocalUrlGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *SystemGeneralApiService) SystemGeneralLocalUrlGetExecute(r SystemGeneralApiSystemGeneralLocalUrlGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemGeneralApiService.SystemGeneralLocalUrlGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/general/local_url"

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

type SystemGeneralApiSystemGeneralPutRequest struct {
	ctx                  context.Context
	ApiService           *SystemGeneralApiService
	systemGeneralUpdate0 *SystemGeneralUpdate0
}

func (r SystemGeneralApiSystemGeneralPutRequest) SystemGeneralUpdate0(systemGeneralUpdate0 SystemGeneralUpdate0) SystemGeneralApiSystemGeneralPutRequest {
	r.systemGeneralUpdate0 = &systemGeneralUpdate0
	return r
}

func (r SystemGeneralApiSystemGeneralPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.SystemGeneralPutExecute(r)
}

/*
SystemGeneralPut Method for SystemGeneralPut

Update System General Service Configuration.

`ui_certificate` is used to enable HTTPS access to the system. If `ui_certificate` is not configured on boot,
it is automatically created by the system.

`ui_httpsredirect` when set, makes sure that all HTTP requests are converted to HTTPS requests to better
enhance security.

`ui_address` and `ui_v6address` are a list of valid ipv4/ipv6 addresses respectively which the system will
listen on.

`syslogserver` and `sysloglevel` are deprecated fields as of 11.3
and will be permanently moved to system.advanced.update for 12.0

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemGeneralApiSystemGeneralPutRequest
*/
func (a *SystemGeneralApiService) SystemGeneralPut(ctx context.Context) SystemGeneralApiSystemGeneralPutRequest {
	return SystemGeneralApiSystemGeneralPutRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *SystemGeneralApiService) SystemGeneralPutExecute(r SystemGeneralApiSystemGeneralPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemGeneralApiService.SystemGeneralPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/general"

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
	localVarPostBody = r.systemGeneralUpdate0
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

type SystemGeneralApiSystemGeneralTimezoneChoicesGetRequest struct {
	ctx        context.Context
	ApiService *SystemGeneralApiService
}

func (r SystemGeneralApiSystemGeneralTimezoneChoicesGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.SystemGeneralTimezoneChoicesGetExecute(r)
}

/*
SystemGeneralTimezoneChoicesGet Method for SystemGeneralTimezoneChoicesGet

Returns time zone choices.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemGeneralApiSystemGeneralTimezoneChoicesGetRequest
*/
func (a *SystemGeneralApiService) SystemGeneralTimezoneChoicesGet(ctx context.Context) SystemGeneralApiSystemGeneralTimezoneChoicesGetRequest {
	return SystemGeneralApiSystemGeneralTimezoneChoicesGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *SystemGeneralApiService) SystemGeneralTimezoneChoicesGetExecute(r SystemGeneralApiSystemGeneralTimezoneChoicesGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemGeneralApiService.SystemGeneralTimezoneChoicesGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/general/timezone_choices"

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

type SystemGeneralApiSystemGeneralUiAddressChoicesGetRequest struct {
	ctx        context.Context
	ApiService *SystemGeneralApiService
}

func (r SystemGeneralApiSystemGeneralUiAddressChoicesGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.SystemGeneralUiAddressChoicesGetExecute(r)
}

/*
SystemGeneralUiAddressChoicesGet Method for SystemGeneralUiAddressChoicesGet

Returns UI ipv4 address choices.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemGeneralApiSystemGeneralUiAddressChoicesGetRequest
*/
func (a *SystemGeneralApiService) SystemGeneralUiAddressChoicesGet(ctx context.Context) SystemGeneralApiSystemGeneralUiAddressChoicesGetRequest {
	return SystemGeneralApiSystemGeneralUiAddressChoicesGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *SystemGeneralApiService) SystemGeneralUiAddressChoicesGetExecute(r SystemGeneralApiSystemGeneralUiAddressChoicesGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemGeneralApiService.SystemGeneralUiAddressChoicesGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/general/ui_address_choices"

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

type SystemGeneralApiSystemGeneralUiCertificateChoicesGetRequest struct {
	ctx        context.Context
	ApiService *SystemGeneralApiService
}

func (r SystemGeneralApiSystemGeneralUiCertificateChoicesGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.SystemGeneralUiCertificateChoicesGetExecute(r)
}

/*
SystemGeneralUiCertificateChoicesGet Method for SystemGeneralUiCertificateChoicesGet

Return choices of certificates which can be used for `ui_certificate`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemGeneralApiSystemGeneralUiCertificateChoicesGetRequest
*/
func (a *SystemGeneralApiService) SystemGeneralUiCertificateChoicesGet(ctx context.Context) SystemGeneralApiSystemGeneralUiCertificateChoicesGetRequest {
	return SystemGeneralApiSystemGeneralUiCertificateChoicesGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *SystemGeneralApiService) SystemGeneralUiCertificateChoicesGetExecute(r SystemGeneralApiSystemGeneralUiCertificateChoicesGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemGeneralApiService.SystemGeneralUiCertificateChoicesGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/general/ui_certificate_choices"

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

type SystemGeneralApiSystemGeneralUiHttpsprotocolsChoicesGetRequest struct {
	ctx        context.Context
	ApiService *SystemGeneralApiService
}

func (r SystemGeneralApiSystemGeneralUiHttpsprotocolsChoicesGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.SystemGeneralUiHttpsprotocolsChoicesGetExecute(r)
}

/*
SystemGeneralUiHttpsprotocolsChoicesGet Method for SystemGeneralUiHttpsprotocolsChoicesGet

Returns available HTTPS protocols.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemGeneralApiSystemGeneralUiHttpsprotocolsChoicesGetRequest
*/
func (a *SystemGeneralApiService) SystemGeneralUiHttpsprotocolsChoicesGet(ctx context.Context) SystemGeneralApiSystemGeneralUiHttpsprotocolsChoicesGetRequest {
	return SystemGeneralApiSystemGeneralUiHttpsprotocolsChoicesGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *SystemGeneralApiService) SystemGeneralUiHttpsprotocolsChoicesGetExecute(r SystemGeneralApiSystemGeneralUiHttpsprotocolsChoicesGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemGeneralApiService.SystemGeneralUiHttpsprotocolsChoicesGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/general/ui_httpsprotocols_choices"

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

type SystemGeneralApiSystemGeneralUiRestartGetRequest struct {
	ctx        context.Context
	ApiService *SystemGeneralApiService
	body       *int32
}

func (r SystemGeneralApiSystemGeneralUiRestartGetRequest) Body(body int32) SystemGeneralApiSystemGeneralUiRestartGetRequest {
	r.body = &body
	return r
}

func (r SystemGeneralApiSystemGeneralUiRestartGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.SystemGeneralUiRestartGetExecute(r)
}

/*
SystemGeneralUiRestartGet Method for SystemGeneralUiRestartGet

Restart HTTP server to use latest UI settings.

HTTP server will be restarted after `delay` seconds.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemGeneralApiSystemGeneralUiRestartGetRequest
*/
func (a *SystemGeneralApiService) SystemGeneralUiRestartGet(ctx context.Context) SystemGeneralApiSystemGeneralUiRestartGetRequest {
	return SystemGeneralApiSystemGeneralUiRestartGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *SystemGeneralApiService) SystemGeneralUiRestartGetExecute(r SystemGeneralApiSystemGeneralUiRestartGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemGeneralApiService.SystemGeneralUiRestartGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/general/ui_restart"

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
	localVarPostBody = r.body
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

type SystemGeneralApiSystemGeneralUiRestartPostRequest struct {
	ctx        context.Context
	ApiService *SystemGeneralApiService
	body       *int32
}

func (r SystemGeneralApiSystemGeneralUiRestartPostRequest) Body(body int32) SystemGeneralApiSystemGeneralUiRestartPostRequest {
	r.body = &body
	return r
}

func (r SystemGeneralApiSystemGeneralUiRestartPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.SystemGeneralUiRestartPostExecute(r)
}

/*
SystemGeneralUiRestartPost Method for SystemGeneralUiRestartPost

Restart HTTP server to use latest UI settings.

HTTP server will be restarted after `delay` seconds.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemGeneralApiSystemGeneralUiRestartPostRequest
*/
func (a *SystemGeneralApiService) SystemGeneralUiRestartPost(ctx context.Context) SystemGeneralApiSystemGeneralUiRestartPostRequest {
	return SystemGeneralApiSystemGeneralUiRestartPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *SystemGeneralApiService) SystemGeneralUiRestartPostExecute(r SystemGeneralApiSystemGeneralUiRestartPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemGeneralApiService.SystemGeneralUiRestartPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/general/ui_restart"

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
	localVarPostBody = r.body
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

type SystemGeneralApiSystemGeneralUiV6addressChoicesGetRequest struct {
	ctx        context.Context
	ApiService *SystemGeneralApiService
}

func (r SystemGeneralApiSystemGeneralUiV6addressChoicesGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.SystemGeneralUiV6addressChoicesGetExecute(r)
}

/*
SystemGeneralUiV6addressChoicesGet Method for SystemGeneralUiV6addressChoicesGet

Returns UI ipv6 address choices.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemGeneralApiSystemGeneralUiV6addressChoicesGetRequest
*/
func (a *SystemGeneralApiService) SystemGeneralUiV6addressChoicesGet(ctx context.Context) SystemGeneralApiSystemGeneralUiV6addressChoicesGetRequest {
	return SystemGeneralApiSystemGeneralUiV6addressChoicesGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *SystemGeneralApiService) SystemGeneralUiV6addressChoicesGetExecute(r SystemGeneralApiSystemGeneralUiV6addressChoicesGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemGeneralApiService.SystemGeneralUiV6addressChoicesGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/general/ui_v6address_choices"

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
