# \SettingsAPI

All URIs are relative to *https://rudder.example.local/rudder/api/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllSettings**](SettingsAPI.md#GetAllSettings) | **Get** /settings | List all settings
[**GetAllowedNetworks**](SettingsAPI.md#GetAllowedNetworks) | **Get** /settings/allowed_networks/{nodeId} | Get allowed networks for a policy server
[**GetSetting**](SettingsAPI.md#GetSetting) | **Get** /settings/{settingId} | Get the value of a setting
[**ModifyAllowedNetworks**](SettingsAPI.md#ModifyAllowedNetworks) | **Post** /settings/allowed_networks/{nodeId}/diff | Modify allowed networks for a policy server
[**ModifySetting**](SettingsAPI.md#ModifySetting) | **Post** /settings/{settingId} | Set the value of a setting
[**SetAllowedNetworks**](SettingsAPI.md#SetAllowedNetworks) | **Post** /settings/allowed_networks/{nodeId} | Set allowed networks for a policy server



## GetAllSettings

> GetAllSettings200Response GetAllSettings(ctx).Execute()

List all settings



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
	resp, r, err := apiClient.SettingsAPI.GetAllSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetAllSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllSettings`: GetAllSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetAllSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSettingsRequest struct via the builder pattern


### Return type

[**GetAllSettings200Response**](GetAllSettings200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllowedNetworks

> GetAllowedNetworks200Response GetAllowedNetworks(ctx, nodeId).Execute()

Get allowed networks for a policy server



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
	nodeId := "9a1773c9-0889-40b6-be89-f6504443ac1b" // string | Policy server ID for which you want to manage allowed networks.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.GetAllowedNetworks(context.Background(), nodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetAllowedNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllowedNetworks`: GetAllowedNetworks200Response
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetAllowedNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | Policy server ID for which you want to manage allowed networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllowedNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAllowedNetworks200Response**](GetAllowedNetworks200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSetting

> GetSetting200Response GetSetting(ctx, settingId).Execute()

Get the value of a setting



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
	settingId := "global_policy_mode" // string | Id of the setting to set

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.GetSetting(context.Background(), settingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSetting`: GetSetting200Response
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settingId** | **string** | Id of the setting to set | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSetting200Response**](GetSetting200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyAllowedNetworks

> ModifyAllowedNetworks200Response ModifyAllowedNetworks(ctx, nodeId).ModifyAllowedNetworksRequest(modifyAllowedNetworksRequest).Execute()

Modify allowed networks for a policy server



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
	nodeId := "9a1773c9-0889-40b6-be89-f6504443ac1b" // string | Policy server ID for which you want to manage allowed networks.
	modifyAllowedNetworksRequest := *openapiclient.NewModifyAllowedNetworksRequest() // ModifyAllowedNetworksRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.ModifyAllowedNetworks(context.Background(), nodeId).ModifyAllowedNetworksRequest(modifyAllowedNetworksRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.ModifyAllowedNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyAllowedNetworks`: ModifyAllowedNetworks200Response
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.ModifyAllowedNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | Policy server ID for which you want to manage allowed networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyAllowedNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modifyAllowedNetworksRequest** | [**ModifyAllowedNetworksRequest**](ModifyAllowedNetworksRequest.md) |  | 

### Return type

[**ModifyAllowedNetworks200Response**](ModifyAllowedNetworks200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifySetting

> ModifySetting200Response ModifySetting(ctx, settingId).ModifySettingRequest(modifySettingRequest).Execute()

Set the value of a setting



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
	settingId := "global_policy_mode" // string | Id of the setting to set
	modifySettingRequest := *openapiclient.NewModifySettingRequest() // ModifySettingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.ModifySetting(context.Background(), settingId).ModifySettingRequest(modifySettingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.ModifySetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySetting`: ModifySetting200Response
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.ModifySetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settingId** | **string** | Id of the setting to set | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modifySettingRequest** | [**ModifySettingRequest**](ModifySettingRequest.md) |  | 

### Return type

[**ModifySetting200Response**](ModifySetting200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAllowedNetworks

> SetAllowedNetworks200Response SetAllowedNetworks(ctx, nodeId).SetAllowedNetworksRequest(setAllowedNetworksRequest).Execute()

Set allowed networks for a policy server



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
	nodeId := "9a1773c9-0889-40b6-be89-f6504443ac1b" // string | Policy server ID for which you want to manage allowed networks.
	setAllowedNetworksRequest := *openapiclient.NewSetAllowedNetworksRequest() // SetAllowedNetworksRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SetAllowedNetworks(context.Background(), nodeId).SetAllowedNetworksRequest(setAllowedNetworksRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SetAllowedNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetAllowedNetworks`: SetAllowedNetworks200Response
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SetAllowedNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | Policy server ID for which you want to manage allowed networks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetAllowedNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setAllowedNetworksRequest** | [**SetAllowedNetworksRequest**](SetAllowedNetworksRequest.md) |  | 

### Return type

[**SetAllowedNetworks200Response**](SetAllowedNetworks200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

