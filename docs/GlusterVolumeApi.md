# \GlusterVolumeApi

All URIs are relative to *https://truenas.example.com/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GlusterVolumeAddbrickPost**](GlusterVolumeApi.md#GlusterVolumeAddbrickPost) | **Post** /gluster/volume/addbrick | 
[**GlusterVolumeGet**](GlusterVolumeApi.md#GlusterVolumeGet) | **Get** /gluster/volume | 
[**GlusterVolumeIdIdDelete**](GlusterVolumeApi.md#GlusterVolumeIdIdDelete) | **Delete** /gluster/volume/id/{id} | 
[**GlusterVolumeIdIdGet**](GlusterVolumeApi.md#GlusterVolumeIdIdGet) | **Get** /gluster/volume/id/{id} | 
[**GlusterVolumeInfoPost**](GlusterVolumeApi.md#GlusterVolumeInfoPost) | **Post** /gluster/volume/info | 
[**GlusterVolumeListGet**](GlusterVolumeApi.md#GlusterVolumeListGet) | **Get** /gluster/volume/list | 
[**GlusterVolumeOptresetPost**](GlusterVolumeApi.md#GlusterVolumeOptresetPost) | **Post** /gluster/volume/optreset | 
[**GlusterVolumeOptsetPost**](GlusterVolumeApi.md#GlusterVolumeOptsetPost) | **Post** /gluster/volume/optset | 
[**GlusterVolumePost**](GlusterVolumeApi.md#GlusterVolumePost) | **Post** /gluster/volume | 
[**GlusterVolumeQuotaPost**](GlusterVolumeApi.md#GlusterVolumeQuotaPost) | **Post** /gluster/volume/quota | 
[**GlusterVolumeRemovebrickPost**](GlusterVolumeApi.md#GlusterVolumeRemovebrickPost) | **Post** /gluster/volume/removebrick | 
[**GlusterVolumeReplacebrickPost**](GlusterVolumeApi.md#GlusterVolumeReplacebrickPost) | **Post** /gluster/volume/replacebrick | 
[**GlusterVolumeRestartPost**](GlusterVolumeApi.md#GlusterVolumeRestartPost) | **Post** /gluster/volume/restart | 
[**GlusterVolumeStartPost**](GlusterVolumeApi.md#GlusterVolumeStartPost) | **Post** /gluster/volume/start | 
[**GlusterVolumeStatusPost**](GlusterVolumeApi.md#GlusterVolumeStatusPost) | **Post** /gluster/volume/status | 
[**GlusterVolumeStopPost**](GlusterVolumeApi.md#GlusterVolumeStopPost) | **Post** /gluster/volume/stop | 



## GlusterVolumeAddbrickPost

> GlusterVolumeAddbrickPost(ctx).GlusterVolumeAddbrick0(glusterVolumeAddbrick0).Execute()





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
    glusterVolumeAddbrick0 := *openapiclient.NewGlusterVolumeAddbrick0() // GlusterVolumeAddbrick0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlusterVolumeApi.GlusterVolumeAddbrickPost(context.Background()).GlusterVolumeAddbrick0(glusterVolumeAddbrick0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlusterVolumeApi.GlusterVolumeAddbrickPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlusterVolumeAddbrickPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **glusterVolumeAddbrick0** | [**GlusterVolumeAddbrick0**](GlusterVolumeAddbrick0.md) |  | 

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


## GlusterVolumeGet

> GlusterVolumeGet(ctx).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()





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
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    count := true // bool |  (optional)
    sort := "sort_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlusterVolumeApi.GlusterVolumeGet(context.Background()).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlusterVolumeApi.GlusterVolumeGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlusterVolumeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **count** | **bool** |  | 
 **sort** | **string** |  | 

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


## GlusterVolumeIdIdDelete

> GlusterVolumeIdIdDelete(ctx, id).Execute()





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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlusterVolumeApi.GlusterVolumeIdIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlusterVolumeApi.GlusterVolumeIdIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGlusterVolumeIdIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GlusterVolumeIdIdGet

> GlusterVolumeIdIdGet(ctx, id).Execute()





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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlusterVolumeApi.GlusterVolumeIdIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlusterVolumeApi.GlusterVolumeIdIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGlusterVolumeIdIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GlusterVolumeInfoPost

> GlusterVolumeInfoPost(ctx).GlusterVolumeInfo0(glusterVolumeInfo0).Execute()





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
    glusterVolumeInfo0 := *openapiclient.NewGlusterVolumeInfo0() // GlusterVolumeInfo0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlusterVolumeApi.GlusterVolumeInfoPost(context.Background()).GlusterVolumeInfo0(glusterVolumeInfo0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlusterVolumeApi.GlusterVolumeInfoPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlusterVolumeInfoPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **glusterVolumeInfo0** | [**GlusterVolumeInfo0**](GlusterVolumeInfo0.md) |  | 

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


## GlusterVolumeListGet

> GlusterVolumeListGet(ctx).Execute()





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
    resp, r, err := apiClient.GlusterVolumeApi.GlusterVolumeListGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlusterVolumeApi.GlusterVolumeListGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGlusterVolumeListGetRequest struct via the builder pattern


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


## GlusterVolumeOptresetPost

> GlusterVolumeOptresetPost(ctx).GlusterVolumeOptreset0(glusterVolumeOptreset0).Execute()





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
    glusterVolumeOptreset0 := *openapiclient.NewGlusterVolumeOptreset0() // GlusterVolumeOptreset0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlusterVolumeApi.GlusterVolumeOptresetPost(context.Background()).GlusterVolumeOptreset0(glusterVolumeOptreset0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlusterVolumeApi.GlusterVolumeOptresetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlusterVolumeOptresetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **glusterVolumeOptreset0** | [**GlusterVolumeOptreset0**](GlusterVolumeOptreset0.md) |  | 

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


## GlusterVolumeOptsetPost

> GlusterVolumeOptsetPost(ctx).GlusterVolumeOptset0(glusterVolumeOptset0).Execute()





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
    glusterVolumeOptset0 := *openapiclient.NewGlusterVolumeOptset0() // GlusterVolumeOptset0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlusterVolumeApi.GlusterVolumeOptsetPost(context.Background()).GlusterVolumeOptset0(glusterVolumeOptset0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlusterVolumeApi.GlusterVolumeOptsetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlusterVolumeOptsetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **glusterVolumeOptset0** | [**GlusterVolumeOptset0**](GlusterVolumeOptset0.md) |  | 

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


## GlusterVolumePost

> GlusterVolumePost(ctx).GlusterVolumeCreate0(glusterVolumeCreate0).Execute()





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
    glusterVolumeCreate0 := *openapiclient.NewGlusterVolumeCreate0() // GlusterVolumeCreate0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlusterVolumeApi.GlusterVolumePost(context.Background()).GlusterVolumeCreate0(glusterVolumeCreate0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlusterVolumeApi.GlusterVolumePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlusterVolumePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **glusterVolumeCreate0** | [**GlusterVolumeCreate0**](GlusterVolumeCreate0.md) |  | 

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


## GlusterVolumeQuotaPost

> GlusterVolumeQuotaPost(ctx).GlusterVolumeQuota0(glusterVolumeQuota0).Execute()





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
    glusterVolumeQuota0 := *openapiclient.NewGlusterVolumeQuota0() // GlusterVolumeQuota0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlusterVolumeApi.GlusterVolumeQuotaPost(context.Background()).GlusterVolumeQuota0(glusterVolumeQuota0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlusterVolumeApi.GlusterVolumeQuotaPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlusterVolumeQuotaPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **glusterVolumeQuota0** | [**GlusterVolumeQuota0**](GlusterVolumeQuota0.md) |  | 

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


## GlusterVolumeRemovebrickPost

> GlusterVolumeRemovebrickPost(ctx).GlusterVolumeRemovebrick0(glusterVolumeRemovebrick0).Execute()





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
    glusterVolumeRemovebrick0 := *openapiclient.NewGlusterVolumeRemovebrick0() // GlusterVolumeRemovebrick0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlusterVolumeApi.GlusterVolumeRemovebrickPost(context.Background()).GlusterVolumeRemovebrick0(glusterVolumeRemovebrick0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlusterVolumeApi.GlusterVolumeRemovebrickPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlusterVolumeRemovebrickPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **glusterVolumeRemovebrick0** | [**GlusterVolumeRemovebrick0**](GlusterVolumeRemovebrick0.md) |  | 

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


## GlusterVolumeReplacebrickPost

> GlusterVolumeReplacebrickPost(ctx).GlusterVolumeReplacebrick0(glusterVolumeReplacebrick0).Execute()





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
    glusterVolumeReplacebrick0 := *openapiclient.NewGlusterVolumeReplacebrick0() // GlusterVolumeReplacebrick0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlusterVolumeApi.GlusterVolumeReplacebrickPost(context.Background()).GlusterVolumeReplacebrick0(glusterVolumeReplacebrick0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlusterVolumeApi.GlusterVolumeReplacebrickPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlusterVolumeReplacebrickPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **glusterVolumeReplacebrick0** | [**GlusterVolumeReplacebrick0**](GlusterVolumeReplacebrick0.md) |  | 

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


## GlusterVolumeRestartPost

> GlusterVolumeRestartPost(ctx).GlusterVolumeRestart0(glusterVolumeRestart0).Execute()





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
    glusterVolumeRestart0 := *openapiclient.NewGlusterVolumeRestart0() // GlusterVolumeRestart0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlusterVolumeApi.GlusterVolumeRestartPost(context.Background()).GlusterVolumeRestart0(glusterVolumeRestart0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlusterVolumeApi.GlusterVolumeRestartPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlusterVolumeRestartPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **glusterVolumeRestart0** | [**GlusterVolumeRestart0**](GlusterVolumeRestart0.md) |  | 

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


## GlusterVolumeStartPost

> GlusterVolumeStartPost(ctx).GlusterVolumeStart0(glusterVolumeStart0).Execute()





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
    glusterVolumeStart0 := *openapiclient.NewGlusterVolumeStart0() // GlusterVolumeStart0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlusterVolumeApi.GlusterVolumeStartPost(context.Background()).GlusterVolumeStart0(glusterVolumeStart0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlusterVolumeApi.GlusterVolumeStartPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlusterVolumeStartPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **glusterVolumeStart0** | [**GlusterVolumeStart0**](GlusterVolumeStart0.md) |  | 

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


## GlusterVolumeStatusPost

> GlusterVolumeStatusPost(ctx).GlusterVolumeStatus0(glusterVolumeStatus0).Execute()





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
    glusterVolumeStatus0 := *openapiclient.NewGlusterVolumeStatus0() // GlusterVolumeStatus0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlusterVolumeApi.GlusterVolumeStatusPost(context.Background()).GlusterVolumeStatus0(glusterVolumeStatus0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlusterVolumeApi.GlusterVolumeStatusPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlusterVolumeStatusPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **glusterVolumeStatus0** | [**GlusterVolumeStatus0**](GlusterVolumeStatus0.md) |  | 

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


## GlusterVolumeStopPost

> GlusterVolumeStopPost(ctx).GlusterVolumeStop0(glusterVolumeStop0).Execute()





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
    glusterVolumeStop0 := *openapiclient.NewGlusterVolumeStop0() // GlusterVolumeStop0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlusterVolumeApi.GlusterVolumeStopPost(context.Background()).GlusterVolumeStop0(glusterVolumeStop0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlusterVolumeApi.GlusterVolumeStopPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGlusterVolumeStopPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **glusterVolumeStop0** | [**GlusterVolumeStop0**](GlusterVolumeStop0.md) |  | 

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

