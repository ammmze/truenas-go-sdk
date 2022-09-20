# \GlusterFuseApi

All URIs are relative to *https://truenas.example.com/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GlusterFuseIsMountedPost**](GlusterFuseApi.md#GlusterFuseIsMountedPost) | **Post** /gluster/fuse/is_mounted | 
[**GlusterFuseMountPost**](GlusterFuseApi.md#GlusterFuseMountPost) | **Post** /gluster/fuse/mount | 
[**GlusterFuseUmountPost**](GlusterFuseApi.md#GlusterFuseUmountPost) | **Post** /gluster/fuse/umount | 



## GlusterFuseIsMountedPost

> GlusterFuseIsMountedPost(ctx).GlusterFuseIsMounted0(glusterFuseIsMounted0).Execute()





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
    glusterFuseIsMounted0 := *openapiclient.NewGlusterFuseIsMounted0() // GlusterFuseIsMounted0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlusterFuseApi.GlusterFuseIsMountedPost(context.Background()).GlusterFuseIsMounted0(glusterFuseIsMounted0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlusterFuseApi.GlusterFuseIsMountedPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlusterFuseIsMountedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **glusterFuseIsMounted0** | [**GlusterFuseIsMounted0**](GlusterFuseIsMounted0.md) |  | 

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


## GlusterFuseMountPost

> GlusterFuseMountPost(ctx).GlusterFuseMount0(glusterFuseMount0).Execute()





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
    glusterFuseMount0 := *openapiclient.NewGlusterFuseMount0() // GlusterFuseMount0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlusterFuseApi.GlusterFuseMountPost(context.Background()).GlusterFuseMount0(glusterFuseMount0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlusterFuseApi.GlusterFuseMountPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlusterFuseMountPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **glusterFuseMount0** | [**GlusterFuseMount0**](GlusterFuseMount0.md) |  | 

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


## GlusterFuseUmountPost

> GlusterFuseUmountPost(ctx).GlusterFuseUmount0(glusterFuseUmount0).Execute()





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
    glusterFuseUmount0 := *openapiclient.NewGlusterFuseUmount0() // GlusterFuseUmount0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlusterFuseApi.GlusterFuseUmountPost(context.Background()).GlusterFuseUmount0(glusterFuseUmount0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlusterFuseApi.GlusterFuseUmountPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlusterFuseUmountPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **glusterFuseUmount0** | [**GlusterFuseUmount0**](GlusterFuseUmount0.md) |  | 

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

