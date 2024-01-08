# \ChangeRequestsAPI

All URIs are relative to *https://rudder.example.local/rudder/api/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptChangeRequest**](ChangeRequestsAPI.md#AcceptChangeRequest) | **Post** /changeRequests/{changeRequestId}/accept | Accept a request details
[**ChangeRequestDetails**](ChangeRequestsAPI.md#ChangeRequestDetails) | **Get** /changeRequests/{changeRequestId} | Get a change request details
[**DeclineChangeRequest**](ChangeRequestsAPI.md#DeclineChangeRequest) | **Delete** /changeRequests/{changeRequestId} | Decline a request details
[**ListChangeRequests**](ChangeRequestsAPI.md#ListChangeRequests) | **Get** /api/changeRequests | List all change requests
[**ListUsers**](ChangeRequestsAPI.md#ListUsers) | **Get** /users | List user
[**RemoveValidatedUser**](ChangeRequestsAPI.md#RemoveValidatedUser) | **Delete** /validatedUsers/{username} | Remove an user from validated user list
[**SaveWorkflowUser**](ChangeRequestsAPI.md#SaveWorkflowUser) | **Post** /validatedUsers | Update validated user list
[**UpdateChangeRequest**](ChangeRequestsAPI.md#UpdateChangeRequest) | **Post** /changeRequests/{changeRequestId} | Update a request details



## AcceptChangeRequest

> AcceptChangeRequest200Response AcceptChangeRequest(ctx, changeRequestId).AcceptChangeRequestRequest(acceptChangeRequestRequest).Execute()

Accept a request details



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
	changeRequestId := int32(37) // int32 | 
	acceptChangeRequestRequest := *openapiclient.NewAcceptChangeRequestRequest() // AcceptChangeRequestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChangeRequestsAPI.AcceptChangeRequest(context.Background(), changeRequestId).AcceptChangeRequestRequest(acceptChangeRequestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeRequestsAPI.AcceptChangeRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcceptChangeRequest`: AcceptChangeRequest200Response
	fmt.Fprintf(os.Stdout, "Response from `ChangeRequestsAPI.AcceptChangeRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**changeRequestId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcceptChangeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptChangeRequestRequest** | [**AcceptChangeRequestRequest**](AcceptChangeRequestRequest.md) |  | 

### Return type

[**AcceptChangeRequest200Response**](AcceptChangeRequest200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangeRequestDetails

> ChangeRequestDetails200Response ChangeRequestDetails(ctx, changeRequestId).Execute()

Get a change request details



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
	changeRequestId := int32(37) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChangeRequestsAPI.ChangeRequestDetails(context.Background(), changeRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeRequestsAPI.ChangeRequestDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeRequestDetails`: ChangeRequestDetails200Response
	fmt.Fprintf(os.Stdout, "Response from `ChangeRequestsAPI.ChangeRequestDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**changeRequestId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeRequestDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChangeRequestDetails200Response**](ChangeRequestDetails200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeclineChangeRequest

> DeclineChangeRequest200Response DeclineChangeRequest(ctx, changeRequestId).Execute()

Decline a request details



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
	changeRequestId := int32(37) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChangeRequestsAPI.DeclineChangeRequest(context.Background(), changeRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeRequestsAPI.DeclineChangeRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeclineChangeRequest`: DeclineChangeRequest200Response
	fmt.Fprintf(os.Stdout, "Response from `ChangeRequestsAPI.DeclineChangeRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**changeRequestId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeclineChangeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeclineChangeRequest200Response**](DeclineChangeRequest200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListChangeRequests

> ListChangeRequests200Response ListChangeRequests(ctx).Execute()

List all change requests



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
	resp, r, err := apiClient.ChangeRequestsAPI.ListChangeRequests(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeRequestsAPI.ListChangeRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListChangeRequests`: ListChangeRequests200Response
	fmt.Fprintf(os.Stdout, "Response from `ChangeRequestsAPI.ListChangeRequests`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListChangeRequestsRequest struct via the builder pattern


### Return type

[**ListChangeRequests200Response**](ListChangeRequests200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsers

> ListUsers200Response ListUsers(ctx).Execute()

List user



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
	resp, r, err := apiClient.ChangeRequestsAPI.ListUsers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeRequestsAPI.ListUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUsers`: ListUsers200Response
	fmt.Fprintf(os.Stdout, "Response from `ChangeRequestsAPI.ListUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


### Return type

[**ListUsers200Response**](ListUsers200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveValidatedUser

> RemoveValidatedUser200Response RemoveValidatedUser(ctx, username).Execute()

Remove an user from validated user list



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
	username := "JaneDoe" // string | Username of an user (unique)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChangeRequestsAPI.RemoveValidatedUser(context.Background(), username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeRequestsAPI.RemoveValidatedUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveValidatedUser`: RemoveValidatedUser200Response
	fmt.Fprintf(os.Stdout, "Response from `ChangeRequestsAPI.RemoveValidatedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username of an user (unique) | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveValidatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RemoveValidatedUser200Response**](RemoveValidatedUser200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveWorkflowUser

> SaveWorkflowUser200Response SaveWorkflowUser(ctx).SaveWorkflowUserRequest(saveWorkflowUserRequest).Execute()

Update validated user list



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
	saveWorkflowUserRequest := *openapiclient.NewSaveWorkflowUserRequest([]string{"John Do"}) // SaveWorkflowUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChangeRequestsAPI.SaveWorkflowUser(context.Background()).SaveWorkflowUserRequest(saveWorkflowUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeRequestsAPI.SaveWorkflowUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveWorkflowUser`: SaveWorkflowUser200Response
	fmt.Fprintf(os.Stdout, "Response from `ChangeRequestsAPI.SaveWorkflowUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveWorkflowUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **saveWorkflowUserRequest** | [**SaveWorkflowUserRequest**](SaveWorkflowUserRequest.md) |  | 

### Return type

[**SaveWorkflowUser200Response**](SaveWorkflowUser200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChangeRequest

> UpdateChangeRequest200Response UpdateChangeRequest(ctx, changeRequestId).UpdateChangeRequestRequest(updateChangeRequestRequest).Execute()

Update a request details



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
	changeRequestId := int32(37) // int32 | 
	updateChangeRequestRequest := *openapiclient.NewUpdateChangeRequestRequest() // UpdateChangeRequestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChangeRequestsAPI.UpdateChangeRequest(context.Background(), changeRequestId).UpdateChangeRequestRequest(updateChangeRequestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeRequestsAPI.UpdateChangeRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateChangeRequest`: UpdateChangeRequest200Response
	fmt.Fprintf(os.Stdout, "Response from `ChangeRequestsAPI.UpdateChangeRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**changeRequestId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateChangeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateChangeRequestRequest** | [**UpdateChangeRequestRequest**](UpdateChangeRequestRequest.md) |  | 

### Return type

[**UpdateChangeRequest200Response**](UpdateChangeRequest200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

