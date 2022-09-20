# \GlusterLocaleventsApi

All URIs are relative to *https://truenas.example.com/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GlusterLocaleventsAddJwtSecretPost**](GlusterLocaleventsApi.md#GlusterLocaleventsAddJwtSecretPost) | **Post** /gluster/localevents/add_jwt_secret | 
[**GlusterLocaleventsGetSetJwtSecretGet**](GlusterLocaleventsApi.md#GlusterLocaleventsGetSetJwtSecretGet) | **Get** /gluster/localevents/get_set_jwt_secret | 



## GlusterLocaleventsAddJwtSecretPost

> GlusterLocaleventsAddJwtSecretPost(ctx).GlusterLocaleventsAddJwtSecret0(glusterLocaleventsAddJwtSecret0).Execute()





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
    glusterLocaleventsAddJwtSecret0 := *openapiclient.NewGlusterLocaleventsAddJwtSecret0() // GlusterLocaleventsAddJwtSecret0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlusterLocaleventsApi.GlusterLocaleventsAddJwtSecretPost(context.Background()).GlusterLocaleventsAddJwtSecret0(glusterLocaleventsAddJwtSecret0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlusterLocaleventsApi.GlusterLocaleventsAddJwtSecretPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlusterLocaleventsAddJwtSecretPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **glusterLocaleventsAddJwtSecret0** | [**GlusterLocaleventsAddJwtSecret0**](GlusterLocaleventsAddJwtSecret0.md) |  | 

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


## GlusterLocaleventsGetSetJwtSecretGet

> GlusterLocaleventsGetSetJwtSecretGet(ctx).Execute()





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
    resp, r, err := apiClient.GlusterLocaleventsApi.GlusterLocaleventsGetSetJwtSecretGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlusterLocaleventsApi.GlusterLocaleventsGetSetJwtSecretGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGlusterLocaleventsGetSetJwtSecretGetRequest struct via the builder pattern


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

