# \FilesystemApi

All URIs are relative to *https://truenas.example.com/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesystemAclIsTrivialPost**](FilesystemApi.md#FilesystemAclIsTrivialPost) | **Post** /filesystem/acl_is_trivial | 
[**FilesystemCanAccessAsUserPost**](FilesystemApi.md#FilesystemCanAccessAsUserPost) | **Post** /filesystem/can_access_as_user | 
[**FilesystemChownPost**](FilesystemApi.md#FilesystemChownPost) | **Post** /filesystem/chown | 
[**FilesystemDefaultAclChoicesPost**](FilesystemApi.md#FilesystemDefaultAclChoicesPost) | **Post** /filesystem/default_acl_choices | 
[**FilesystemGetDefaultAclPost**](FilesystemApi.md#FilesystemGetDefaultAclPost) | **Post** /filesystem/get_default_acl | 
[**FilesystemGetDosmodePost**](FilesystemApi.md#FilesystemGetDosmodePost) | **Post** /filesystem/get_dosmode | 
[**FilesystemGetPost**](FilesystemApi.md#FilesystemGetPost) | **Post** /filesystem/get | 
[**FilesystemGetaclPost**](FilesystemApi.md#FilesystemGetaclPost) | **Post** /filesystem/getacl | 
[**FilesystemIsImmutablePost**](FilesystemApi.md#FilesystemIsImmutablePost) | **Post** /filesystem/is_immutable | 
[**FilesystemListdirPost**](FilesystemApi.md#FilesystemListdirPost) | **Post** /filesystem/listdir | 
[**FilesystemMkdirPost**](FilesystemApi.md#FilesystemMkdirPost) | **Post** /filesystem/mkdir | 
[**FilesystemPutPost**](FilesystemApi.md#FilesystemPutPost) | **Post** /filesystem/put | 
[**FilesystemSetDosmodePost**](FilesystemApi.md#FilesystemSetDosmodePost) | **Post** /filesystem/set_dosmode | 
[**FilesystemSetImmutablePost**](FilesystemApi.md#FilesystemSetImmutablePost) | **Post** /filesystem/set_immutable | 
[**FilesystemSetaclPost**](FilesystemApi.md#FilesystemSetaclPost) | **Post** /filesystem/setacl | 
[**FilesystemSetpermPost**](FilesystemApi.md#FilesystemSetpermPost) | **Post** /filesystem/setperm | 
[**FilesystemStatPost**](FilesystemApi.md#FilesystemStatPost) | **Post** /filesystem/stat | 
[**FilesystemStatfsPost**](FilesystemApi.md#FilesystemStatfsPost) | **Post** /filesystem/statfs | 



## FilesystemAclIsTrivialPost

> FilesystemAclIsTrivialPost(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.FilesystemApi.FilesystemAclIsTrivialPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesystemApi.FilesystemAclIsTrivialPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesystemAclIsTrivialPostRequest struct via the builder pattern


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


## FilesystemCanAccessAsUserPost

> FilesystemCanAccessAsUserPost(ctx).FilesystemCanAccessAsUser(filesystemCanAccessAsUser).Execute()





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
    filesystemCanAccessAsUser := *openapiclient.NewFilesystemCanAccessAsUser() // FilesystemCanAccessAsUser |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesystemApi.FilesystemCanAccessAsUserPost(context.Background()).FilesystemCanAccessAsUser(filesystemCanAccessAsUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesystemApi.FilesystemCanAccessAsUserPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesystemCanAccessAsUserPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filesystemCanAccessAsUser** | [**FilesystemCanAccessAsUser**](FilesystemCanAccessAsUser.md) |  | 

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


## FilesystemChownPost

> FilesystemChownPost(ctx).FilesystemChown0(filesystemChown0).Execute()





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
    filesystemChown0 := *openapiclient.NewFilesystemChown0() // FilesystemChown0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesystemApi.FilesystemChownPost(context.Background()).FilesystemChown0(filesystemChown0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesystemApi.FilesystemChownPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesystemChownPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filesystemChown0** | [**FilesystemChown0**](FilesystemChown0.md) |  | 

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


## FilesystemDefaultAclChoicesPost

> FilesystemDefaultAclChoicesPost(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.FilesystemApi.FilesystemDefaultAclChoicesPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesystemApi.FilesystemDefaultAclChoicesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesystemDefaultAclChoicesPostRequest struct via the builder pattern


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


## FilesystemGetDefaultAclPost

> FilesystemGetDefaultAclPost(ctx).FilesystemGetDefaultAcl(filesystemGetDefaultAcl).Execute()





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
    filesystemGetDefaultAcl := *openapiclient.NewFilesystemGetDefaultAcl() // FilesystemGetDefaultAcl |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesystemApi.FilesystemGetDefaultAclPost(context.Background()).FilesystemGetDefaultAcl(filesystemGetDefaultAcl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesystemApi.FilesystemGetDefaultAclPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesystemGetDefaultAclPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filesystemGetDefaultAcl** | [**FilesystemGetDefaultAcl**](FilesystemGetDefaultAcl.md) |  | 

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


## FilesystemGetDosmodePost

> FilesystemGetDosmodePost(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.FilesystemApi.FilesystemGetDosmodePost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesystemApi.FilesystemGetDosmodePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesystemGetDosmodePostRequest struct via the builder pattern


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


## FilesystemGetPost

> FilesystemGetPost(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.FilesystemApi.FilesystemGetPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesystemApi.FilesystemGetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesystemGetPostRequest struct via the builder pattern


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


## FilesystemGetaclPost

> FilesystemGetaclPost(ctx).FilesystemGetacl(filesystemGetacl).Execute()





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
    filesystemGetacl := *openapiclient.NewFilesystemGetacl() // FilesystemGetacl |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesystemApi.FilesystemGetaclPost(context.Background()).FilesystemGetacl(filesystemGetacl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesystemApi.FilesystemGetaclPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesystemGetaclPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filesystemGetacl** | [**FilesystemGetacl**](FilesystemGetacl.md) |  | 

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


## FilesystemIsImmutablePost

> FilesystemIsImmutablePost(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.FilesystemApi.FilesystemIsImmutablePost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesystemApi.FilesystemIsImmutablePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesystemIsImmutablePostRequest struct via the builder pattern


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


## FilesystemListdirPost

> FilesystemListdirPost(ctx).FilesystemListdir(filesystemListdir).Execute()





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
    filesystemListdir := *openapiclient.NewFilesystemListdir() // FilesystemListdir |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesystemApi.FilesystemListdirPost(context.Background()).FilesystemListdir(filesystemListdir).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesystemApi.FilesystemListdirPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesystemListdirPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filesystemListdir** | [**FilesystemListdir**](FilesystemListdir.md) |  | 

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


## FilesystemMkdirPost

> FilesystemMkdirPost(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.FilesystemApi.FilesystemMkdirPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesystemApi.FilesystemMkdirPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesystemMkdirPostRequest struct via the builder pattern


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


## FilesystemPutPost

> FilesystemPutPost(ctx).FilesystemPut(filesystemPut).Execute()





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
    filesystemPut := *openapiclient.NewFilesystemPut() // FilesystemPut |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesystemApi.FilesystemPutPost(context.Background()).FilesystemPut(filesystemPut).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesystemApi.FilesystemPutPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesystemPutPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filesystemPut** | [**FilesystemPut**](FilesystemPut.md) |  | 

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


## FilesystemSetDosmodePost

> FilesystemSetDosmodePost(ctx).FilesystemSetDosmode0(filesystemSetDosmode0).Execute()





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
    filesystemSetDosmode0 := *openapiclient.NewFilesystemSetDosmode0() // FilesystemSetDosmode0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesystemApi.FilesystemSetDosmodePost(context.Background()).FilesystemSetDosmode0(filesystemSetDosmode0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesystemApi.FilesystemSetDosmodePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesystemSetDosmodePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filesystemSetDosmode0** | [**FilesystemSetDosmode0**](FilesystemSetDosmode0.md) |  | 

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


## FilesystemSetImmutablePost

> FilesystemSetImmutablePost(ctx).FilesystemSetImmutable(filesystemSetImmutable).Execute()





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
    filesystemSetImmutable := *openapiclient.NewFilesystemSetImmutable() // FilesystemSetImmutable |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesystemApi.FilesystemSetImmutablePost(context.Background()).FilesystemSetImmutable(filesystemSetImmutable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesystemApi.FilesystemSetImmutablePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesystemSetImmutablePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filesystemSetImmutable** | [**FilesystemSetImmutable**](FilesystemSetImmutable.md) |  | 

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


## FilesystemSetaclPost

> FilesystemSetaclPost(ctx).FilesystemSetacl0(filesystemSetacl0).Execute()





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
    filesystemSetacl0 := *openapiclient.NewFilesystemSetacl0() // FilesystemSetacl0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesystemApi.FilesystemSetaclPost(context.Background()).FilesystemSetacl0(filesystemSetacl0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesystemApi.FilesystemSetaclPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesystemSetaclPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filesystemSetacl0** | [**FilesystemSetacl0**](FilesystemSetacl0.md) |  | 

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


## FilesystemSetpermPost

> FilesystemSetpermPost(ctx).FilesystemSetperm0(filesystemSetperm0).Execute()





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
    filesystemSetperm0 := *openapiclient.NewFilesystemSetperm0() // FilesystemSetperm0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesystemApi.FilesystemSetpermPost(context.Background()).FilesystemSetperm0(filesystemSetperm0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesystemApi.FilesystemSetpermPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesystemSetpermPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filesystemSetperm0** | [**FilesystemSetperm0**](FilesystemSetperm0.md) |  | 

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


## FilesystemStatPost

> FilesystemStatPost(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.FilesystemApi.FilesystemStatPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesystemApi.FilesystemStatPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesystemStatPostRequest struct via the builder pattern


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


## FilesystemStatfsPost

> FilesystemStatfsPost(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.FilesystemApi.FilesystemStatfsPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesystemApi.FilesystemStatfsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesystemStatfsPostRequest struct via the builder pattern


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

