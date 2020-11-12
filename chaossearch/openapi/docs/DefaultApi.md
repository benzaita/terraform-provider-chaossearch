# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BucketCreateObjectGroupPost**](DefaultApi.md#BucketCreateObjectGroupPost) | **Post** /Bucket/createObjectGroup | Create a new Object Group
[**BucketCreateViewPost**](DefaultApi.md#BucketCreateViewPost) | **Post** /Bucket/createView | Create a new View
[**BucketMetadataPost**](DefaultApi.md#BucketMetadataPost) | **Post** /Bucket/metadata | Fetch bucket metadata that includes information regarding index statistics
[**BucketModelPost**](DefaultApi.md#BucketModelPost) | **Post** /Bucket/model | Send a control signal to the system to operate indexing
[**BucketPartitionKeysPost**](DefaultApi.md#BucketPartitionKeysPost) | **Post** /Bucket/partitionKeys | Fetch the partition keys for a View, Object Group, or index
[**BucketUpdateObjectGroupPost**](DefaultApi.md#BucketUpdateObjectGroupPost) | **Post** /Bucket/updateObjectGroup | Update an existing Object Group
[**BucketUpdateViewPost**](DefaultApi.md#BucketUpdateViewPost) | **Post** /Bucket/updateView | Update an existing View
[**UserCreateSubAccountPost**](DefaultApi.md#UserCreateSubAccountPost) | **Post** /user/createSubAccount | Create or Edit a sub-account under the given user context
[**UserDeleteSubAccountPost**](DefaultApi.md#UserDeleteSubAccountPost) | **Post** /user/deleteSubAccount | Delete a sub-account under the given user context
[**UserGroupIdDelete**](DefaultApi.md#UserGroupIdDelete) | **Delete** /user/group/{id} | Delete the group data for the group ID
[**UserGroupIdGet**](DefaultApi.md#UserGroupIdGet) | **Get** /user/group/{id} | Fetch the group data for the group ID
[**UserGroupsDelete**](DefaultApi.md#UserGroupsDelete) | **Delete** /user/groups | Delete groups of this user context
[**UserGroupsGet**](DefaultApi.md#UserGroupsGet) | **Get** /user/groups | Fetch all groups of this user context
[**UserGroupsPost**](DefaultApi.md#UserGroupsPost) | **Post** /user/groups | Create new groups for this user context
[**UserGroupsPut**](DefaultApi.md#UserGroupsPut) | **Put** /user/groups | Update existing groups of this user context
[**UserManifestPost**](DefaultApi.md#UserManifestPost) | **Post** /user/manifest | Provide a manifest of all user related resources
[**V1Get**](DefaultApi.md#V1Get) | **Get** /V1/ | List existing Object Groups
[**V1KeyDelete**](DefaultApi.md#V1KeyDelete) | **Delete** /V1/{key} | Delete Views/Object Groups



## BucketCreateObjectGroupPost

> BucketCreateObjectGroupPost(ctx).InlineObject(inlineObject).Execute()

Create a new Object Group



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
    inlineObject := *openapiclient.NewInlineObject() // InlineObject |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.BucketCreateObjectGroupPost(context.Background()).InlineObject(inlineObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.BucketCreateObjectGroupPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBucketCreateObjectGroupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject** | [**InlineObject**](InlineObject.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BucketCreateViewPost

> BucketCreateViewPost(ctx).InlineObject2(inlineObject2).Execute()

Create a new View



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
    inlineObject2 := *openapiclient.Newinline_object_2("Bucket_example", []string{"Sources_example"), "IndexPattern_example") // InlineObject2 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.BucketCreateViewPost(context.Background()).InlineObject2(inlineObject2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.BucketCreateViewPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBucketCreateViewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject2** | [**InlineObject2**](InlineObject2.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BucketMetadataPost

> BucketMetadataPost(ctx).InlineObject6(inlineObject6).Execute()

Fetch bucket metadata that includes information regarding index statistics



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
    inlineObject6 := *openapiclient.Newinline_object_6() // InlineObject6 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.BucketMetadataPost(context.Background()).InlineObject6(inlineObject6).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.BucketMetadataPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBucketMetadataPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject6** | [**InlineObject6**](InlineObject6.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BucketModelPost

> BucketModelPost(ctx).InlineObject5(inlineObject5).Execute()

Send a control signal to the system to operate indexing



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
    inlineObject5 := *openapiclient.Newinline_object_5() // InlineObject5 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.BucketModelPost(context.Background()).InlineObject5(inlineObject5).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.BucketModelPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBucketModelPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject5** | [**InlineObject5**](InlineObject5.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BucketPartitionKeysPost

> InlineResponse200 BucketPartitionKeysPost(ctx).InlineObject4(inlineObject4).Execute()

Fetch the partition keys for a View, Object Group, or index



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
    inlineObject4 := *openapiclient.Newinline_object_4() // InlineObject4 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.BucketPartitionKeysPost(context.Background()).InlineObject4(inlineObject4).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.BucketPartitionKeysPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BucketPartitionKeysPost`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.BucketPartitionKeysPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBucketPartitionKeysPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject4** | [**InlineObject4**](InlineObject4.md) |  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BucketUpdateObjectGroupPost

> BucketUpdateObjectGroupPost(ctx).InlineObject1(inlineObject1).Execute()

Update an existing Object Group



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
    inlineObject1 := *openapiclient.Newinline_object_1("Bucket_example") // InlineObject1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.BucketUpdateObjectGroupPost(context.Background()).InlineObject1(inlineObject1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.BucketUpdateObjectGroupPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBucketUpdateObjectGroupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject1** | [**InlineObject1**](InlineObject1.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BucketUpdateViewPost

> BucketUpdateViewPost(ctx).InlineObject3(inlineObject3).Execute()

Update an existing View



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
    inlineObject3 := *openapiclient.Newinline_object_3() // InlineObject3 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.BucketUpdateViewPost(context.Background()).InlineObject3(inlineObject3).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.BucketUpdateViewPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBucketUpdateViewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject3** | [**InlineObject3**](InlineObject3.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserCreateSubAccountPost

> InlineResponse201 UserCreateSubAccountPost(ctx).InlineObject7(inlineObject7).Execute()

Create or Edit a sub-account under the given user context



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
    inlineObject7 := *openapiclient.Newinline_object_7() // InlineObject7 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UserCreateSubAccountPost(context.Background()).InlineObject7(inlineObject7).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UserCreateSubAccountPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserCreateSubAccountPost`: InlineResponse201
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UserCreateSubAccountPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserCreateSubAccountPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject7** | [**InlineObject7**](InlineObject7.md) |  | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserDeleteSubAccountPost

> InlineResponse2002 UserDeleteSubAccountPost(ctx).InlineObject8(inlineObject8).Execute()

Delete a sub-account under the given user context



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
    inlineObject8 := *openapiclient.Newinline_object_8() // InlineObject8 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UserDeleteSubAccountPost(context.Background()).InlineObject8(inlineObject8).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UserDeleteSubAccountPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserDeleteSubAccountPost`: InlineResponse2002
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UserDeleteSubAccountPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserDeleteSubAccountPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject8** | [**InlineObject8**](InlineObject8.md) |  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGroupIdDelete

> []string UserGroupIdDelete(ctx, id).Execute()

Delete the group data for the group ID



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
    id := "id_example" // string | {id} is the ID of a Group 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UserGroupIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UserGroupIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGroupIdDelete`: []string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UserGroupIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | {id} is the ID of a Group  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserGroupIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGroupIdGet

> InlineResponse2011 UserGroupIdGet(ctx, id).Execute()

Fetch the group data for the group ID



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
    id := "id_example" // string | {id} is the ID of a Group 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UserGroupIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UserGroupIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGroupIdGet`: InlineResponse2011
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UserGroupIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | {id} is the ID of a Group  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserGroupIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2011**](inline_response_201_1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGroupsDelete

> UserGroupsDelete(ctx).Execute()

Delete groups of this user context



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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UserGroupsDelete(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UserGroupsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserGroupsDeleteRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGroupsGet

> []InlineResponse2003 UserGroupsGet(ctx).Execute()

Fetch all groups of this user context



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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UserGroupsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UserGroupsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGroupsGet`: []InlineResponse2003
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UserGroupsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserGroupsGetRequest struct via the builder pattern


### Return type

[**[]InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGroupsPost

> InlineResponse2011 UserGroupsPost(ctx).InlineObject(inlineObject).Execute()

Create new groups for this user context



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
    inlineObject := []InlineObject{*openapiclient.NewInlineObject()} // []InlineObject |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UserGroupsPost(context.Background()).InlineObject(inlineObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UserGroupsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGroupsPost`: InlineResponse2011
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UserGroupsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserGroupsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject** | [**[]InlineObject**](InlineObject.md) |  | 

### Return type

[**InlineResponse2011**](inline_response_201_1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGroupsPut

> InlineResponse2011 UserGroupsPut(ctx).InlineObject(inlineObject).Execute()

Update existing groups of this user context



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
    inlineObject := []InlineObject{} // []InlineObject |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UserGroupsPut(context.Background()).InlineObject(inlineObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UserGroupsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGroupsPut`: InlineResponse2011
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UserGroupsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserGroupsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject** | [**[]InlineObject**](InlineObject.md) |  | 

### Return type

[**InlineResponse2011**](inline_response_201_1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserManifestPost

> InlineResponse2001 UserManifestPost(ctx).Execute()

Provide a manifest of all user related resources



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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UserManifestPost(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UserManifestPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserManifestPost`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UserManifestPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserManifestPostRequest struct via the builder pattern


### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1Get

> XmlResponse V1Get(ctx).Execute()

List existing Object Groups



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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.V1Get(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1Get`: XmlResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.V1Get`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetRequest struct via the builder pattern


### Return type

[**XmlResponse**](XmlResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1KeyDelete

> V1KeyDelete(ctx, key).Execute()

Delete Views/Object Groups



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
    key := "key_example" // string | {key} can be a View Name, or an Object Group Name 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.V1KeyDelete(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1KeyDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | {key} can be a View Name, or an Object Group Name  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1KeyDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

