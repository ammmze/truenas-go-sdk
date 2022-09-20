# \CtdbGeneralApi

All URIs are relative to *https://truenas.example.com/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CtdbGeneralHealthyGet**](CtdbGeneralApi.md#CtdbGeneralHealthyGet) | **Get** /ctdb/general/healthy | 
[**CtdbGeneralIpsPost**](CtdbGeneralApi.md#CtdbGeneralIpsPost) | **Post** /ctdb/general/ips | 
[**CtdbGeneralListnodesGet**](CtdbGeneralApi.md#CtdbGeneralListnodesGet) | **Get** /ctdb/general/listnodes | 
[**CtdbGeneralPnnGet**](CtdbGeneralApi.md#CtdbGeneralPnnGet) | **Get** /ctdb/general/pnn | 
[**CtdbGeneralStatusPost**](CtdbGeneralApi.md#CtdbGeneralStatusPost) | **Post** /ctdb/general/status | 



## CtdbGeneralHealthyGet

> CtdbGeneralHealthyGet(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CtdbGeneralApi.CtdbGeneralHealthyGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CtdbGeneralApi.CtdbGeneralHealthyGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCtdbGeneralHealthyGetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CtdbGeneralIpsPost

> CtdbGeneralIpsPost(ctx).CtdbGeneralIps0(ctdbGeneralIps0).Execute()





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
    ctdbGeneralIps0 := *openapiclient.NewCtdbGeneralIps0() // CtdbGeneralIps0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CtdbGeneralApi.CtdbGeneralIpsPost(context.Background()).CtdbGeneralIps0(ctdbGeneralIps0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CtdbGeneralApi.CtdbGeneralIpsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCtdbGeneralIpsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctdbGeneralIps0** | [**CtdbGeneralIps0**](CtdbGeneralIps0.md) |  | 

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


## CtdbGeneralListnodesGet

> CtdbGeneralListnodesGet(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CtdbGeneralApi.CtdbGeneralListnodesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CtdbGeneralApi.CtdbGeneralListnodesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCtdbGeneralListnodesGetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CtdbGeneralPnnGet

> CtdbGeneralPnnGet(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CtdbGeneralApi.CtdbGeneralPnnGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CtdbGeneralApi.CtdbGeneralPnnGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCtdbGeneralPnnGetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CtdbGeneralStatusPost

> CtdbGeneralStatusPost(ctx).CtdbGeneralStatus0(ctdbGeneralStatus0).Execute()





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
    ctdbGeneralStatus0 := *openapiclient.NewCtdbGeneralStatus0() // CtdbGeneralStatus0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CtdbGeneralApi.CtdbGeneralStatusPost(context.Background()).CtdbGeneralStatus0(ctdbGeneralStatus0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CtdbGeneralApi.CtdbGeneralStatusPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCtdbGeneralStatusPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctdbGeneralStatus0** | [**CtdbGeneralStatus0**](CtdbGeneralStatus0.md) |  | 

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

