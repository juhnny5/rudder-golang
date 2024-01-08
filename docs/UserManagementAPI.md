# \UserManagementAPI

All URIs are relative to *https://rudder.example.local/rudder/api/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUser**](UserManagementAPI.md#AddUser) | **Post** /usermanagement | Add user
[**DeleteUser**](UserManagementAPI.md#DeleteUser) | **Delete** /usermanagement/{username} | Delete an user
[**GetRole**](UserManagementAPI.md#GetRole) | **Get** /usermanagement/roles | List all roles
[**GetUserInfo**](UserManagementAPI.md#GetUserInfo) | **Get** /usermanagement/users | List all users
[**ReloadUserConf**](UserManagementAPI.md#ReloadUserConf) | **Get** /usermanagement/users/reload | Reload user
[**UpdateUser**](UserManagementAPI.md#UpdateUser) | **Post** /usermanagement/update/{username} | Update user&#39;s infos



## AddUser

> AddUser200Response AddUser(ctx).Users(users).Execute()

Add user



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
	users := *openapiclient.NewUsers() // Users | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.AddUser(context.Background()).Users(users).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.AddUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddUser`: AddUser200Response
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.AddUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **users** | [**Users**](Users.md) |  | 

### Return type

[**AddUser200Response**](AddUser200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteUser200Response DeleteUser(ctx, username).Execute()

Delete an user



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
	resp, r, err := apiClient.UserManagementAPI.DeleteUser(context.Background(), username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUser`: DeleteUser200Response
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.DeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username of an user (unique) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteUser200Response**](DeleteUser200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRole

> GetRole200Response GetRole(ctx).Execute()

List all roles



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
	resp, r, err := apiClient.UserManagementAPI.GetRole(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.GetRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRole`: GetRole200Response
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.GetRole`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleRequest struct via the builder pattern


### Return type

[**GetRole200Response**](GetRole200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserInfo

> GetUserInfo200Response GetUserInfo(ctx).Execute()

List all users



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
	resp, r, err := apiClient.UserManagementAPI.GetUserInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.GetUserInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserInfo`: GetUserInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.GetUserInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserInfoRequest struct via the builder pattern


### Return type

[**GetUserInfo200Response**](GetUserInfo200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReloadUserConf

> ReloadUserConf200Response ReloadUserConf(ctx).Execute()

Reload user



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
	resp, r, err := apiClient.UserManagementAPI.ReloadUserConf(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.ReloadUserConf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReloadUserConf`: ReloadUserConf200Response
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.ReloadUserConf`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReloadUserConfRequest struct via the builder pattern


### Return type

[**ReloadUserConf200Response**](ReloadUserConf200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> UpdateUser200Response UpdateUser(ctx, username).Users(users).Execute()

Update user's infos



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
	users := *openapiclient.NewUsers() // Users | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.UpdateUser(context.Background(), username).Users(users).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.UpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUser`: UpdateUser200Response
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username of an user (unique) | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **users** | [**Users**](Users.md) |  | 

### Return type

[**UpdateUser200Response**](UpdateUser200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

