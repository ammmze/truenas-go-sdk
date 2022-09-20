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

// SmbApiService SmbApi service
type SmbApiService service

type SmbApiSmbBindipChoicesGetRequest struct {
	ctx        context.Context
	ApiService *SmbApiService
}

func (r SmbApiSmbBindipChoicesGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.SmbBindipChoicesGetExecute(r)
}

/*
SmbBindipChoicesGet Method for SmbBindipChoicesGet

List of valid choices for IP addresses to which to bind the SMB service.
Addresses assigned by DHCP are excluded from the results.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SmbApiSmbBindipChoicesGetRequest
*/
func (a *SmbApiService) SmbBindipChoicesGet(ctx context.Context) SmbApiSmbBindipChoicesGetRequest {
	return SmbApiSmbBindipChoicesGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *SmbApiService) SmbBindipChoicesGetExecute(r SmbApiSmbBindipChoicesGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SmbApiService.SmbBindipChoicesGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/smb/bindip_choices"

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

type SmbApiSmbClientCountGetRequest struct {
	ctx        context.Context
	ApiService *SmbApiService
}

func (r SmbApiSmbClientCountGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.SmbClientCountGetExecute(r)
}

/*
SmbClientCountGet Method for SmbClientCountGet

Return currently connected clients count.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SmbApiSmbClientCountGetRequest
*/
func (a *SmbApiService) SmbClientCountGet(ctx context.Context) SmbApiSmbClientCountGetRequest {
	return SmbApiSmbClientCountGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *SmbApiService) SmbClientCountGetExecute(r SmbApiSmbClientCountGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SmbApiService.SmbClientCountGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/smb/client_count"

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

type SmbApiSmbDomainChoicesGetRequest struct {
	ctx        context.Context
	ApiService *SmbApiService
}

func (r SmbApiSmbDomainChoicesGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.SmbDomainChoicesGetExecute(r)
}

/*
SmbDomainChoicesGet Method for SmbDomainChoicesGet

List of domains visible to winbindd. Returns empty list if winbindd is
stopped.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SmbApiSmbDomainChoicesGetRequest
*/
func (a *SmbApiService) SmbDomainChoicesGet(ctx context.Context) SmbApiSmbDomainChoicesGetRequest {
	return SmbApiSmbDomainChoicesGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *SmbApiService) SmbDomainChoicesGetExecute(r SmbApiSmbDomainChoicesGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SmbApiService.SmbDomainChoicesGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/smb/domain_choices"

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

type SmbApiSmbGetRequest struct {
	ctx        context.Context
	ApiService *SmbApiService
}

func (r SmbApiSmbGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.SmbGetExecute(r)
}

/*
SmbGet Method for SmbGet



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SmbApiSmbGetRequest
*/
func (a *SmbApiService) SmbGet(ctx context.Context) SmbApiSmbGetRequest {
	return SmbApiSmbGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *SmbApiService) SmbGetExecute(r SmbApiSmbGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SmbApiService.SmbGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/smb"

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

type SmbApiSmbGetRemoteAclPostRequest struct {
	ctx              context.Context
	ApiService       *SmbApiService
	smbGetRemoteAcl0 *SmbGetRemoteAcl0
}

func (r SmbApiSmbGetRemoteAclPostRequest) SmbGetRemoteAcl0(smbGetRemoteAcl0 SmbGetRemoteAcl0) SmbApiSmbGetRemoteAclPostRequest {
	r.smbGetRemoteAcl0 = &smbGetRemoteAcl0
	return r
}

func (r SmbApiSmbGetRemoteAclPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.SmbGetRemoteAclPostExecute(r)
}

/*
SmbGetRemoteAclPost Method for SmbGetRemoteAclPost

Retrieves an ACL from a remote SMB server.

`server` IP Address or hostname of the remote server

`share` Share name

`path` path on the remote SMB server. Use "" to separate path components

`username` username to use for authentication

`password` password to use for authentication

`use_kerberos` use credentials to get a kerberos ticket for authentication.
AD only.

`output_format` format for resulting ACL data. Choices are either 'SMB',
which will present the information as a Windows SD or 'LOCAL', which formats
the ACL information according local filesystem of the TrueNAS server.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SmbApiSmbGetRemoteAclPostRequest
*/
func (a *SmbApiService) SmbGetRemoteAclPost(ctx context.Context) SmbApiSmbGetRemoteAclPostRequest {
	return SmbApiSmbGetRemoteAclPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *SmbApiService) SmbGetRemoteAclPostExecute(r SmbApiSmbGetRemoteAclPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SmbApiService.SmbGetRemoteAclPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/smb/get_remote_acl"

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
	localVarPostBody = r.smbGetRemoteAcl0
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

type SmbApiSmbPutRequest struct {
	ctx        context.Context
	ApiService *SmbApiService
	smbUpdate0 *SmbUpdate0
}

func (r SmbApiSmbPutRequest) SmbUpdate0(smbUpdate0 SmbUpdate0) SmbApiSmbPutRequest {
	r.smbUpdate0 = &smbUpdate0
	return r
}

func (r SmbApiSmbPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.SmbPutExecute(r)
}

/*
SmbPut Method for SmbPut

Update SMB Service Configuration.

`netbiosname` defaults to the original hostname of the system.

`netbiosalias` a list of netbios aliases. If Server is joined to an AD domain, additional Kerberos
Service Principal Names will be generated for these aliases.

`workgroup` specifies the NetBIOS workgroup to which the TrueNAS server belongs. This will be automatically
set to the correct value during the process of joining an AD domain. `workgroup` and `netbiosname` should have different values.

`enable_smb1` allows legacy SMB clients to connect to the server when enabled.

`aapl_extensions` enables support for SMB2 protocol extensions for MacOS clients. This is not a requirement for MacOS support,
but is currently a requirement for time machine support.

`localmaster` when set, determines if the system participates in a browser election.

`guest` attribute is specified to select the account to be used for guest access. It defaults to "nobody".

The group specified as the SMB `admin_group` will be automatically added as a foreign group member of S-1-5-32-544 (builtindmins).
This will afford the group all privileges granted to a local admin. Any SMB group may be selected (including AD groups).

`ntlmv1_auth` enables a legacy and insecure authentication method, which may be required for legacy or poorly-implemented
SMB clients.

`smb_options` smb.conf parameters that are not covered by the above supported configuration options may be added as
an smb_option. Not all options are tested or supported, and behavior of smb_options may change between releases. Stability of
smb.conf options is not guaranteed.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SmbApiSmbPutRequest
*/
func (a *SmbApiService) SmbPut(ctx context.Context) SmbApiSmbPutRequest {
	return SmbApiSmbPutRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *SmbApiService) SmbPutExecute(r SmbApiSmbPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SmbApiService.SmbPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/smb"

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
	localVarPostBody = r.smbUpdate0
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

type SmbApiSmbStatusPostRequest struct {
	ctx        context.Context
	ApiService *SmbApiService
	smbStatus  *SmbStatus
}

func (r SmbApiSmbStatusPostRequest) SmbStatus(smbStatus SmbStatus) SmbApiSmbStatusPostRequest {
	r.smbStatus = &smbStatus
	return r
}

func (r SmbApiSmbStatusPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.SmbStatusPostExecute(r)
}

/*
SmbStatusPost Method for SmbStatusPost

Returns SMB server status (sessions, open files, locks, notifications).

`info_level` type of information requests. Defaults to ALL.

`status_options` additional options to filter query results. Supported
values are as follows: `verbose` gives more verbose status output
`fast` causes smbstatus to not check if the status data is valid by
checking if the processes that the status data refer to all still
exist. This speeds up execution on busy systems and clusters but
might display stale data of processes that died without cleaning up
properly. `restrict_user` specifies the limits results to the specified
user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SmbApiSmbStatusPostRequest
*/
func (a *SmbApiService) SmbStatusPost(ctx context.Context) SmbApiSmbStatusPostRequest {
	return SmbApiSmbStatusPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *SmbApiService) SmbStatusPostExecute(r SmbApiSmbStatusPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SmbApiService.SmbStatusPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/smb/status"

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
	localVarPostBody = r.smbStatus
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

type SmbApiSmbUnixcharsetChoicesGetRequest struct {
	ctx        context.Context
	ApiService *SmbApiService
}

func (r SmbApiSmbUnixcharsetChoicesGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.SmbUnixcharsetChoicesGetExecute(r)
}

/*
SmbUnixcharsetChoicesGet Method for SmbUnixcharsetChoicesGet



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SmbApiSmbUnixcharsetChoicesGetRequest
*/
func (a *SmbApiService) SmbUnixcharsetChoicesGet(ctx context.Context) SmbApiSmbUnixcharsetChoicesGetRequest {
	return SmbApiSmbUnixcharsetChoicesGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *SmbApiService) SmbUnixcharsetChoicesGetExecute(r SmbApiSmbUnixcharsetChoicesGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SmbApiService.SmbUnixcharsetChoicesGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/smb/unixcharset_choices"

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
