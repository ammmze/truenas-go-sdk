# \GlusterRebalanceApi

All URIs are relative to *https://truenas.example.com/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GlusterRebalanceFixLayoutPost**](GlusterRebalanceApi.md#GlusterRebalanceFixLayoutPost) | **Post** /gluster/rebalance/fix_layout | 
[**GlusterRebalanceStartPost**](GlusterRebalanceApi.md#GlusterRebalanceStartPost) | **Post** /gluster/rebalance/start | 
[**GlusterRebalanceStatusPost**](GlusterRebalanceApi.md#GlusterRebalanceStatusPost) | **Post** /gluster/rebalance/status | 
[**GlusterRebalanceStopPost**](GlusterRebalanceApi.md#GlusterRebalanceStopPost) | **Post** /gluster/rebalance/stop | 



## GlusterRebalanceFixLayoutPost

> GlusterRebalanceFixLayoutPost(ctx).GlusterRebalanceFixLayout0(glusterRebalanceFixLayout0).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    glusterRebalanceFixLayout0 := *openapiclient.NewGlusterRebalanceFixLayout0() // GlusterRebalanceFixLayout0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlusterRebalanceApi.GlusterRebalanceFixLayoutPost(context.Background()).GlusterRebalanceFixLayout0(glusterRebalanceFixLayout0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlusterRebalanceApi.GlusterRebalanceFixLayoutPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlusterRebalanceFixLayoutPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **glusterRebalanceFixLayout0** | [**GlusterRebalanceFixLayout0**](GlusterRebalanceFixLayout0.md) |  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GlusterRebalanceStartPost

> GlusterRebalanceStartPost(ctx).GlusterRebalanceStart0(glusterRebalanceStart0).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    glusterRebalanceStart0 := *openapiclient.NewGlusterRebalanceStart0() // GlusterRebalanceStart0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlusterRebalanceApi.GlusterRebalanceStartPost(context.Background()).GlusterRebalanceStart0(glusterRebalanceStart0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlusterRebalanceApi.GlusterRebalanceStartPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlusterRebalanceStartPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **glusterRebalanceStart0** | [**GlusterRebalanceStart0**](GlusterRebalanceStart0.md) |  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GlusterRebalanceStatusPost

> GlusterRebalanceStatusPost(ctx).GlusterRebalanceStatus0(glusterRebalanceStatus0).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    glusterRebalanceStatus0 := *openapiclient.NewGlusterRebalanceStatus0() // GlusterRebalanceStatus0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlusterRebalanceApi.GlusterRebalanceStatusPost(context.Background()).GlusterRebalanceStatus0(glusterRebalanceStatus0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlusterRebalanceApi.GlusterRebalanceStatusPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlusterRebalanceStatusPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **glusterRebalanceStatus0** | [**GlusterRebalanceStatus0**](GlusterRebalanceStatus0.md) |  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GlusterRebalanceStopPost

> GlusterRebalanceStopPost(ctx).GlusterRebalanceStop0(glusterRebalanceStop0).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    glusterRebalanceStop0 := *openapiclient.NewGlusterRebalanceStop0() // GlusterRebalanceStop0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlusterRebalanceApi.GlusterRebalanceStopPost(context.Background()).GlusterRebalanceStop0(glusterRebalanceStop0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlusterRebalanceApi.GlusterRebalanceStopPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlusterRebalanceStopPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **glusterRebalanceStop0** | [**GlusterRebalanceStop0**](GlusterRebalanceStop0.md) |  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

