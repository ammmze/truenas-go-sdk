# \ChartReleaseApi

All URIs are relative to *https://truenas.example.com/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChartReleaseCertificateAuthorityChoicesGet**](ChartReleaseApi.md#ChartReleaseCertificateAuthorityChoicesGet) | **Get** /chart/release/certificate_authority_choices | 
[**ChartReleaseCertificateChoicesGet**](ChartReleaseApi.md#ChartReleaseCertificateChoicesGet) | **Get** /chart/release/certificate_choices | 
[**ChartReleaseEventsPost**](ChartReleaseApi.md#ChartReleaseEventsPost) | **Post** /chart/release/events | 
[**ChartReleaseGet**](ChartReleaseApi.md#ChartReleaseGet) | **Get** /chart/release | 
[**ChartReleaseIdIdDelete**](ChartReleaseApi.md#ChartReleaseIdIdDelete) | **Delete** /chart/release/id/{id} | 
[**ChartReleaseIdIdGet**](ChartReleaseApi.md#ChartReleaseIdIdGet) | **Get** /chart/release/id/{id} | 
[**ChartReleaseIdIdPut**](ChartReleaseApi.md#ChartReleaseIdIdPut) | **Put** /chart/release/id/{id} | 
[**ChartReleaseNicChoicesGet**](ChartReleaseApi.md#ChartReleaseNicChoicesGet) | **Get** /chart/release/nic_choices | 
[**ChartReleasePodConsoleChoicesPost**](ChartReleaseApi.md#ChartReleasePodConsoleChoicesPost) | **Post** /chart/release/pod_console_choices | 
[**ChartReleasePodLogsChoicesPost**](ChartReleaseApi.md#ChartReleasePodLogsChoicesPost) | **Post** /chart/release/pod_logs_choices | 
[**ChartReleasePodLogsPost**](ChartReleaseApi.md#ChartReleasePodLogsPost) | **Post** /chart/release/pod_logs | 
[**ChartReleasePodStatusPost**](ChartReleaseApi.md#ChartReleasePodStatusPost) | **Post** /chart/release/pod_status | 
[**ChartReleasePost**](ChartReleaseApi.md#ChartReleasePost) | **Post** /chart/release | 
[**ChartReleasePullContainerImagesPost**](ChartReleaseApi.md#ChartReleasePullContainerImagesPost) | **Post** /chart/release/pull_container_images | 
[**ChartReleaseRedeployPost**](ChartReleaseApi.md#ChartReleaseRedeployPost) | **Post** /chart/release/redeploy | 
[**ChartReleaseRemoveIxVolumePost**](ChartReleaseApi.md#ChartReleaseRemoveIxVolumePost) | **Post** /chart/release/remove_ix_volume | 
[**ChartReleaseRollbackPost**](ChartReleaseApi.md#ChartReleaseRollbackPost) | **Post** /chart/release/rollback | 
[**ChartReleaseScalePost**](ChartReleaseApi.md#ChartReleaseScalePost) | **Post** /chart/release/scale | 
[**ChartReleaseScaleWorkloadsPost**](ChartReleaseApi.md#ChartReleaseScaleWorkloadsPost) | **Post** /chart/release/scale_workloads | 
[**ChartReleaseScaleableResourcesGet**](ChartReleaseApi.md#ChartReleaseScaleableResourcesGet) | **Get** /chart/release/scaleable_resources | 
[**ChartReleaseUpgradePost**](ChartReleaseApi.md#ChartReleaseUpgradePost) | **Post** /chart/release/upgrade | 
[**ChartReleaseUpgradeSummaryPost**](ChartReleaseApi.md#ChartReleaseUpgradeSummaryPost) | **Post** /chart/release/upgrade_summary | 
[**ChartReleaseUsedPortsGet**](ChartReleaseApi.md#ChartReleaseUsedPortsGet) | **Get** /chart/release/used_ports | 



## ChartReleaseCertificateAuthorityChoicesGet

> ChartReleaseCertificateAuthorityChoicesGet(ctx).Execute()





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
    resp, r, err := apiClient.ChartReleaseApi.ChartReleaseCertificateAuthorityChoicesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartReleaseApi.ChartReleaseCertificateAuthorityChoicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChartReleaseCertificateAuthorityChoicesGetRequest struct via the builder pattern


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


## ChartReleaseCertificateChoicesGet

> ChartReleaseCertificateChoicesGet(ctx).Execute()





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
    resp, r, err := apiClient.ChartReleaseApi.ChartReleaseCertificateChoicesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartReleaseApi.ChartReleaseCertificateChoicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChartReleaseCertificateChoicesGetRequest struct via the builder pattern


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


## ChartReleaseEventsPost

> ChartReleaseEventsPost(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.ChartReleaseApi.ChartReleaseEventsPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartReleaseApi.ChartReleaseEventsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChartReleaseEventsPostRequest struct via the builder pattern


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


## ChartReleaseGet

> ChartReleaseGet(ctx).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()





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
    resp, r, err := apiClient.ChartReleaseApi.ChartReleaseGet(context.Background()).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartReleaseApi.ChartReleaseGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChartReleaseGetRequest struct via the builder pattern


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


## ChartReleaseIdIdDelete

> ChartReleaseIdIdDelete(ctx, id).Execute()





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
    resp, r, err := apiClient.ChartReleaseApi.ChartReleaseIdIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartReleaseApi.ChartReleaseIdIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiChartReleaseIdIdDeleteRequest struct via the builder pattern


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


## ChartReleaseIdIdGet

> ChartReleaseIdIdGet(ctx, id).Execute()





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
    resp, r, err := apiClient.ChartReleaseApi.ChartReleaseIdIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartReleaseApi.ChartReleaseIdIdGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiChartReleaseIdIdGetRequest struct via the builder pattern


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


## ChartReleaseIdIdPut

> ChartReleaseIdIdPut(ctx, id).Execute()





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
    resp, r, err := apiClient.ChartReleaseApi.ChartReleaseIdIdPut(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartReleaseApi.ChartReleaseIdIdPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiChartReleaseIdIdPutRequest struct via the builder pattern


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


## ChartReleaseNicChoicesGet

> ChartReleaseNicChoicesGet(ctx).Execute()





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
    resp, r, err := apiClient.ChartReleaseApi.ChartReleaseNicChoicesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartReleaseApi.ChartReleaseNicChoicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChartReleaseNicChoicesGetRequest struct via the builder pattern


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


## ChartReleasePodConsoleChoicesPost

> ChartReleasePodConsoleChoicesPost(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.ChartReleaseApi.ChartReleasePodConsoleChoicesPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartReleaseApi.ChartReleasePodConsoleChoicesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChartReleasePodConsoleChoicesPostRequest struct via the builder pattern


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


## ChartReleasePodLogsChoicesPost

> ChartReleasePodLogsChoicesPost(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.ChartReleaseApi.ChartReleasePodLogsChoicesPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartReleaseApi.ChartReleasePodLogsChoicesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChartReleasePodLogsChoicesPostRequest struct via the builder pattern


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


## ChartReleasePodLogsPost

> ChartReleasePodLogsPost(ctx).ChartReleasePodLogs(chartReleasePodLogs).Execute()





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
    chartReleasePodLogs := *openapiclient.NewChartReleasePodLogs() // ChartReleasePodLogs |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartReleaseApi.ChartReleasePodLogsPost(context.Background()).ChartReleasePodLogs(chartReleasePodLogs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartReleaseApi.ChartReleasePodLogsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChartReleasePodLogsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chartReleasePodLogs** | [**ChartReleasePodLogs**](ChartReleasePodLogs.md) |  | 

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


## ChartReleasePodStatusPost

> ChartReleasePodStatusPost(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.ChartReleaseApi.ChartReleasePodStatusPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartReleaseApi.ChartReleasePodStatusPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChartReleasePodStatusPostRequest struct via the builder pattern


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


## ChartReleasePost

> ChartReleasePost(ctx).ChartReleaseCreate0(chartReleaseCreate0).Execute()





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
    chartReleaseCreate0 := *openapiclient.NewChartReleaseCreate0() // ChartReleaseCreate0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartReleaseApi.ChartReleasePost(context.Background()).ChartReleaseCreate0(chartReleaseCreate0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartReleaseApi.ChartReleasePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChartReleasePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chartReleaseCreate0** | [**ChartReleaseCreate0**](ChartReleaseCreate0.md) |  | 

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


## ChartReleasePullContainerImagesPost

> ChartReleasePullContainerImagesPost(ctx).ChartReleasePullContainerImages(chartReleasePullContainerImages).Execute()





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
    chartReleasePullContainerImages := *openapiclient.NewChartReleasePullContainerImages() // ChartReleasePullContainerImages |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartReleaseApi.ChartReleasePullContainerImagesPost(context.Background()).ChartReleasePullContainerImages(chartReleasePullContainerImages).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartReleaseApi.ChartReleasePullContainerImagesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChartReleasePullContainerImagesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chartReleasePullContainerImages** | [**ChartReleasePullContainerImages**](ChartReleasePullContainerImages.md) |  | 

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


## ChartReleaseRedeployPost

> ChartReleaseRedeployPost(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.ChartReleaseApi.ChartReleaseRedeployPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartReleaseApi.ChartReleaseRedeployPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChartReleaseRedeployPostRequest struct via the builder pattern


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


## ChartReleaseRemoveIxVolumePost

> ChartReleaseRemoveIxVolumePost(ctx).ChartReleaseRemoveIxVolume(chartReleaseRemoveIxVolume).Execute()





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
    chartReleaseRemoveIxVolume := *openapiclient.NewChartReleaseRemoveIxVolume() // ChartReleaseRemoveIxVolume |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartReleaseApi.ChartReleaseRemoveIxVolumePost(context.Background()).ChartReleaseRemoveIxVolume(chartReleaseRemoveIxVolume).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartReleaseApi.ChartReleaseRemoveIxVolumePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChartReleaseRemoveIxVolumePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chartReleaseRemoveIxVolume** | [**ChartReleaseRemoveIxVolume**](ChartReleaseRemoveIxVolume.md) |  | 

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


## ChartReleaseRollbackPost

> ChartReleaseRollbackPost(ctx).ChartReleaseRollback(chartReleaseRollback).Execute()





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
    chartReleaseRollback := *openapiclient.NewChartReleaseRollback() // ChartReleaseRollback |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartReleaseApi.ChartReleaseRollbackPost(context.Background()).ChartReleaseRollback(chartReleaseRollback).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartReleaseApi.ChartReleaseRollbackPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChartReleaseRollbackPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chartReleaseRollback** | [**ChartReleaseRollback**](ChartReleaseRollback.md) |  | 

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


## ChartReleaseScalePost

> ChartReleaseScalePost(ctx).ChartReleaseScale(chartReleaseScale).Execute()





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
    chartReleaseScale := *openapiclient.NewChartReleaseScale() // ChartReleaseScale |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartReleaseApi.ChartReleaseScalePost(context.Background()).ChartReleaseScale(chartReleaseScale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartReleaseApi.ChartReleaseScalePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChartReleaseScalePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chartReleaseScale** | [**ChartReleaseScale**](ChartReleaseScale.md) |  | 

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


## ChartReleaseScaleWorkloadsPost

> ChartReleaseScaleWorkloadsPost(ctx).ChartReleaseScaleWorkloads(chartReleaseScaleWorkloads).Execute()





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
    chartReleaseScaleWorkloads := *openapiclient.NewChartReleaseScaleWorkloads() // ChartReleaseScaleWorkloads |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartReleaseApi.ChartReleaseScaleWorkloadsPost(context.Background()).ChartReleaseScaleWorkloads(chartReleaseScaleWorkloads).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartReleaseApi.ChartReleaseScaleWorkloadsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChartReleaseScaleWorkloadsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chartReleaseScaleWorkloads** | [**ChartReleaseScaleWorkloads**](ChartReleaseScaleWorkloads.md) |  | 

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


## ChartReleaseScaleableResourcesGet

> ChartReleaseScaleableResourcesGet(ctx).Execute()





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
    resp, r, err := apiClient.ChartReleaseApi.ChartReleaseScaleableResourcesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartReleaseApi.ChartReleaseScaleableResourcesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChartReleaseScaleableResourcesGetRequest struct via the builder pattern


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


## ChartReleaseUpgradePost

> ChartReleaseUpgradePost(ctx).ChartReleaseUpgrade(chartReleaseUpgrade).Execute()





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
    chartReleaseUpgrade := *openapiclient.NewChartReleaseUpgrade() // ChartReleaseUpgrade |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartReleaseApi.ChartReleaseUpgradePost(context.Background()).ChartReleaseUpgrade(chartReleaseUpgrade).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartReleaseApi.ChartReleaseUpgradePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChartReleaseUpgradePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chartReleaseUpgrade** | [**ChartReleaseUpgrade**](ChartReleaseUpgrade.md) |  | 

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


## ChartReleaseUpgradeSummaryPost

> ChartReleaseUpgradeSummaryPost(ctx).ChartReleaseUpgradeSummary(chartReleaseUpgradeSummary).Execute()





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
    chartReleaseUpgradeSummary := *openapiclient.NewChartReleaseUpgradeSummary() // ChartReleaseUpgradeSummary |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChartReleaseApi.ChartReleaseUpgradeSummaryPost(context.Background()).ChartReleaseUpgradeSummary(chartReleaseUpgradeSummary).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartReleaseApi.ChartReleaseUpgradeSummaryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChartReleaseUpgradeSummaryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chartReleaseUpgradeSummary** | [**ChartReleaseUpgradeSummary**](ChartReleaseUpgradeSummary.md) |  | 

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


## ChartReleaseUsedPortsGet

> ChartReleaseUsedPortsGet(ctx).Execute()





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
    resp, r, err := apiClient.ChartReleaseApi.ChartReleaseUsedPortsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartReleaseApi.ChartReleaseUsedPortsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChartReleaseUsedPortsGetRequest struct via the builder pattern


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

