# \InterfaceCapabilitiesApi

All URIs are relative to *https://truenas.example.com/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InterfaceCapabilitiesGetPost**](InterfaceCapabilitiesApi.md#InterfaceCapabilitiesGetPost) | **Post** /interface/capabilities/get | 
[**InterfaceCapabilitiesSetPost**](InterfaceCapabilitiesApi.md#InterfaceCapabilitiesSetPost) | **Post** /interface/capabilities/set | 



## InterfaceCapabilitiesGetPost

> InterfaceCapabilitiesGetPost(ctx).Body(body).Execute()





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
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InterfaceCapabilitiesApi.InterfaceCapabilitiesGetPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterfaceCapabilitiesApi.InterfaceCapabilitiesGetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInterfaceCapabilitiesGetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

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


## InterfaceCapabilitiesSetPost

> InterfaceCapabilitiesSetPost(ctx).InterfaceCapabilitiesSet0(interfaceCapabilitiesSet0).Execute()





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
    interfaceCapabilitiesSet0 := *openapiclient.NewInterfaceCapabilitiesSet0() // InterfaceCapabilitiesSet0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InterfaceCapabilitiesApi.InterfaceCapabilitiesSetPost(context.Background()).InterfaceCapabilitiesSet0(interfaceCapabilitiesSet0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterfaceCapabilitiesApi.InterfaceCapabilitiesSetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInterfaceCapabilitiesSetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **interfaceCapabilitiesSet0** | [**InterfaceCapabilitiesSet0**](InterfaceCapabilitiesSet0.md) |  | 

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

