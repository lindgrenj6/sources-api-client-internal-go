# \DefaultApi

All URIs are relative to *http://localhost:3000//internal/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSources**](DefaultApi.md#ListSources) | **Get** /sources | List Sources
[**ListTenants**](DefaultApi.md#ListTenants) | **Get** /tenants | List Tenants
[**ShowAuthentication**](DefaultApi.md#ShowAuthentication) | **Get** /authentications/{id} | Show an existing Authentication
[**ShowTenant**](DefaultApi.md#ShowTenant) | **Get** /tenants/{id} | Show an existing Tenant



## ListSources

> SourcesCollection ListSources(ctx).Limit(limit).Offset(offset).Filter(filter).SortBy(sortBy).Execute()

List Sources



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
    limit := int32(56) // int32 | The numbers of items to return per page. (optional) (default to 100)
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set. (optional) (default to 0)
    filter := TODO // map[string]interface{} | Filter for querying collections. (optional)
    sortBy := TODO // OneOfstringarray | The list of attribute and order to sort the result set by. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSources(context.Background()).Limit(limit).Offset(offset).Filter(filter).SortBy(sortBy).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSources`: SourcesCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The numbers of items to return per page. | [default to 100]
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | [default to 0]
 **filter** | [**map[string]interface{}**](map[string]interface{}.md) | Filter for querying collections. | 
 **sortBy** | [**OneOfstringarray**](OneOfstringarray.md) | The list of attribute and order to sort the result set by. | 

### Return type

[**SourcesCollection**](SourcesCollection.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTenants

> TenantsCollection ListTenants(ctx).Limit(limit).Offset(offset).Filter(filter).SortBy(sortBy).Execute()

List Tenants



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
    limit := int32(56) // int32 | The numbers of items to return per page. (optional) (default to 100)
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set. (optional) (default to 0)
    filter := TODO // map[string]interface{} | Filter for querying collections. (optional)
    sortBy := TODO // OneOfstringarray | The list of attribute and order to sort the result set by. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListTenants(context.Background()).Limit(limit).Offset(offset).Filter(filter).SortBy(sortBy).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListTenants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTenants`: TenantsCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListTenants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTenantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The numbers of items to return per page. | [default to 100]
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | [default to 0]
 **filter** | [**map[string]interface{}**](map[string]interface{}.md) | Filter for querying collections. | 
 **sortBy** | [**OneOfstringarray**](OneOfstringarray.md) | The list of attribute and order to sort the result set by. | 

### Return type

[**TenantsCollection**](TenantsCollection.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowAuthentication

> Authentication ShowAuthentication(ctx, id).Execute()

Show an existing Authentication



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
    id := "id_example" // string | ID of the resource

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ShowAuthentication(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ShowAuthentication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowAuthentication`: Authentication
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ShowAuthentication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowAuthenticationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Authentication**](Authentication.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowTenant

> Tenant ShowTenant(ctx, id).Execute()

Show an existing Tenant



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
    id := "id_example" // string | ID of the resource

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ShowTenant(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ShowTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowTenant`: Tenant
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ShowTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Tenant**](Tenant.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

