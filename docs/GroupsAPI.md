# \GroupsAPI

All URIs are relative to *https://rudder.example.local/rudder/api/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGroup**](GroupsAPI.md#CreateGroup) | **Put** /groups | Create a group
[**CreateGroupCategory**](GroupsAPI.md#CreateGroupCategory) | **Put** /groups/categories | Create a group category
[**DeleteGroup**](GroupsAPI.md#DeleteGroup) | **Delete** /groups/{groupId} | Delete a group
[**DeleteGroupCategory**](GroupsAPI.md#DeleteGroupCategory) | **Delete** /groups/categories/{groupCategoryId} | Delete group category
[**GetGroupCategoryDetails**](GroupsAPI.md#GetGroupCategoryDetails) | **Get** /groups/categories/{groupCategoryId} | Get group category details
[**GetGroupTree**](GroupsAPI.md#GetGroupTree) | **Get** /groups/tree | Get groups tree
[**GroupDetails**](GroupsAPI.md#GroupDetails) | **Get** /groups/{groupId} | Get group details
[**ListGroups**](GroupsAPI.md#ListGroups) | **Get** /groups | List all groups
[**ReloadGroup**](GroupsAPI.md#ReloadGroup) | **Post** /groups/{groupId}/reload | Reload a group
[**UpdateGroup**](GroupsAPI.md#UpdateGroup) | **Post** /groups/{groupId} | Update group details
[**UpdateGroupCategory**](GroupsAPI.md#UpdateGroupCategory) | **Post** /groups/categories/{groupCategoryId} | Update group category details



## CreateGroup

> CreateGroup200Response CreateGroup(ctx).GroupNew(groupNew).Execute()

Create a group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/juhnny5/rudder-golang"
)

func main() {
	groupNew := *openapiclient.NewGroupNew("e17ecf6a-a9f2-44de-a97c-116d24d30ff4", "Ubuntu 18.04 nodes") // GroupNew |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.CreateGroup(context.Background()).GroupNew(groupNew).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.CreateGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGroup`: CreateGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.CreateGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupNew** | [**GroupNew**](GroupNew.md) |  | 

### Return type

[**CreateGroup200Response**](CreateGroup200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroupCategory

> CreateGroupCategory200Response CreateGroupCategory(ctx).GroupCategory(groupCategory).Execute()

Create a group category



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/juhnny5/rudder-golang"
)

func main() {
	groupCategory := *openapiclient.NewGroupCategory("b9f6d98a-28bc-4d80-90f7-d2f14269e215", "Hardware groups") // GroupCategory | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.CreateGroupCategory(context.Background()).GroupCategory(groupCategory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.CreateGroupCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGroupCategory`: CreateGroupCategory200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.CreateGroupCategory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupCategory** | [**GroupCategory**](GroupCategory.md) |  | 

### Return type

[**CreateGroupCategory200Response**](CreateGroupCategory200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroup

> DeleteGroup200Response DeleteGroup(ctx, groupId).Execute()

Delete a group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/juhnny5/rudder-golang"
)

func main() {
	groupId := "9a1773c9-0889-40b6-be89-f6504443ac1b" // string | Id of the group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.DeleteGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.DeleteGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGroup`: DeleteGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.DeleteGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Id of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteGroup200Response**](DeleteGroup200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroupCategory

> DeleteGroupCategory200Response DeleteGroupCategory(ctx, groupCategoryId).Execute()

Delete group category



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/juhnny5/rudder-golang"
)

func main() {
	groupCategoryId := "e0a311fa-f7b2-4f9e-89a9-db517b9c6b90" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.DeleteGroupCategory(context.Background(), groupCategoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.DeleteGroupCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGroupCategory`: DeleteGroupCategory200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.DeleteGroupCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupCategoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteGroupCategory200Response**](DeleteGroupCategory200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupCategoryDetails

> GetGroupCategoryDetails200Response GetGroupCategoryDetails(ctx, groupCategoryId).Execute()

Get group category details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/juhnny5/rudder-golang"
)

func main() {
	groupCategoryId := "e0a311fa-f7b2-4f9e-89a9-db517b9c6b90" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupCategoryDetails(context.Background(), groupCategoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupCategoryDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupCategoryDetails`: GetGroupCategoryDetails200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupCategoryDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupCategoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupCategoryDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetGroupCategoryDetails200Response**](GetGroupCategoryDetails200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupTree

> GetGroupTree200Response GetGroupTree(ctx).Execute()

Get groups tree



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/juhnny5/rudder-golang"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupTree(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupTree``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupTree`: GetGroupTree200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupTree`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupTreeRequest struct via the builder pattern


### Return type

[**GetGroupTree200Response**](GetGroupTree200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupDetails

> GroupDetails200Response GroupDetails(ctx, groupId).Execute()

Get group details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/juhnny5/rudder-golang"
)

func main() {
	groupId := "9a1773c9-0889-40b6-be89-f6504443ac1b" // string | Id of the group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GroupDetails(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GroupDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupDetails`: GroupDetails200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GroupDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Id of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupDetails200Response**](GroupDetails200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroups

> ListGroups200Response ListGroups(ctx).Execute()

List all groups



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/juhnny5/rudder-golang"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.ListGroups(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.ListGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGroups`: ListGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.ListGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupsRequest struct via the builder pattern


### Return type

[**ListGroups200Response**](ListGroups200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReloadGroup

> ReloadGroup200Response ReloadGroup(ctx, groupId).Execute()

Reload a group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/juhnny5/rudder-golang"
)

func main() {
	groupId := "9a1773c9-0889-40b6-be89-f6504443ac1b" // string | Id of the group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.ReloadGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.ReloadGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReloadGroup`: ReloadGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.ReloadGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Id of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiReloadGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReloadGroup200Response**](ReloadGroup200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroup

> UpdateGroup200Response UpdateGroup(ctx, groupId).GroupUpdate(groupUpdate).Execute()

Update group details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/juhnny5/rudder-golang"
)

func main() {
	groupId := "9a1773c9-0889-40b6-be89-f6504443ac1b" // string | Id of the group
	groupUpdate := *openapiclient.NewGroupUpdate() // GroupUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.UpdateGroup(context.Background(), groupId).GroupUpdate(groupUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.UpdateGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGroup`: UpdateGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.UpdateGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Id of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupUpdate** | [**GroupUpdate**](GroupUpdate.md) |  | 

### Return type

[**UpdateGroup200Response**](UpdateGroup200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupCategory

> UpdateGroupCategory200Response UpdateGroupCategory(ctx, groupCategoryId).GroupCategoryUpdate(groupCategoryUpdate).Execute()

Update group category details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/juhnny5/rudder-golang"
)

func main() {
	groupCategoryId := "e0a311fa-f7b2-4f9e-89a9-db517b9c6b90" // string | 
	groupCategoryUpdate := *openapiclient.NewGroupCategoryUpdate("b9f6d98a-28bc-4d80-90f7-d2f14269e215", "Hardware groups") // GroupCategoryUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.UpdateGroupCategory(context.Background(), groupCategoryId).GroupCategoryUpdate(groupCategoryUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.UpdateGroupCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGroupCategory`: UpdateGroupCategory200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.UpdateGroupCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupCategoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupCategoryUpdate** | [**GroupCategoryUpdate**](GroupCategoryUpdate.md) |  | 

### Return type

[**UpdateGroupCategory200Response**](UpdateGroupCategory200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

