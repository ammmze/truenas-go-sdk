# \GlusterEventsdApi

All URIs are relative to *https://truenas.example.com/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GlusterEventsdCreatePost**](GlusterEventsdApi.md#GlusterEventsdCreatePost) | **Post** /gluster/eventsd/create | 
[**GlusterEventsdDeletePost**](GlusterEventsdApi.md#GlusterEventsdDeletePost) | **Post** /gluster/eventsd/delete | 
[**GlusterEventsdSyncGet**](GlusterEventsdApi.md#GlusterEventsdSyncGet) | **Get** /gluster/eventsd/sync | 
[**GlusterEventsdWebhooksGet**](GlusterEventsdApi.md#GlusterEventsdWebhooksGet) | **Get** /gluster/eventsd/webhooks | 



## GlusterEventsdCreatePost

> GlusterEventsdCreatePost(ctx).GlusterEventsdCreate0(glusterEventsdCreate0).Execute()





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
    glusterEventsdCreate0 := *openapiclient.NewGlusterEventsdCreate0() // GlusterEventsdCreate0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlusterEventsdApi.GlusterEventsdCreatePost(context.Background()).GlusterEventsdCreate0(glusterEventsdCreate0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlusterEventsdApi.GlusterEventsdCreatePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlusterEventsdCreatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **glusterEventsdCreate0** | [**GlusterEventsdCreate0**](GlusterEventsdCreate0.md) |  | 

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


## GlusterEventsdDeletePost

> GlusterEventsdDeletePost(ctx).GlusterEventsdDelete0(glusterEventsdDelete0).Execute()





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
    glusterEventsdDelete0 := *openapiclient.NewGlusterEventsdDelete0() // GlusterEventsdDelete0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlusterEventsdApi.GlusterEventsdDeletePost(context.Background()).GlusterEventsdDelete0(glusterEventsdDelete0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlusterEventsdApi.GlusterEventsdDeletePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlusterEventsdDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **glusterEventsdDelete0** | [**GlusterEventsdDelete0**](GlusterEventsdDelete0.md) |  | 

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


## GlusterEventsdSyncGet

> GlusterEventsdSyncGet(ctx).Execute()





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
    resp, r, err := apiClient.GlusterEventsdApi.GlusterEventsdSyncGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlusterEventsdApi.GlusterEventsdSyncGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGlusterEventsdSyncGetRequest struct via the builder pattern


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


## GlusterEventsdWebhooksGet

> GlusterEventsdWebhooksGet(ctx).Execute()





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
    resp, r, err := apiClient.GlusterEventsdApi.GlusterEventsdWebhooksGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlusterEventsdApi.GlusterEventsdWebhooksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGlusterEventsdWebhooksGetRequest struct via the builder pattern


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

