# \ParametersAPI

All URIs are relative to *https://rudder.example.local/rudder/api/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateParameter**](ParametersAPI.md#CreateParameter) | **Put** /parameters | Create a new property
[**DeleteParameter**](ParametersAPI.md#DeleteParameter) | **Delete** /parameters/{parameterId} | Delete a global parameter
[**ListParameters**](ParametersAPI.md#ListParameters) | **Get** /parameters | List all global properties
[**ParameterDetails**](ParametersAPI.md#ParameterDetails) | **Get** /parameters/{parameterId} | Get the value of a global property
[**UpdateParameter**](ParametersAPI.md#UpdateParameter) | **Post** /parameters/{parameterId} | Update a global property&#39;s value



## CreateParameter

> CreateParameter200Response CreateParameter(ctx).Parameter(parameter).Execute()

Create a new property



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
	parameter := *openapiclient.NewParameter("rudder_file_edit_footer") // Parameter | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParametersAPI.CreateParameter(context.Background()).Parameter(parameter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParametersAPI.CreateParameter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateParameter`: CreateParameter200Response
	fmt.Fprintf(os.Stdout, "Response from `ParametersAPI.CreateParameter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateParameterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parameter** | [**Parameter**](Parameter.md) |  | 

### Return type

[**CreateParameter200Response**](CreateParameter200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteParameter

> DeleteParameter200Response DeleteParameter(ctx, parameterId).Execute()

Delete a global parameter



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
	parameterId := "rudder_file_edit_header" // string | Id of the parameter to manage

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParametersAPI.DeleteParameter(context.Background(), parameterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParametersAPI.DeleteParameter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteParameter`: DeleteParameter200Response
	fmt.Fprintf(os.Stdout, "Response from `ParametersAPI.DeleteParameter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parameterId** | **string** | Id of the parameter to manage | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteParameterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteParameter200Response**](DeleteParameter200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListParameters

> ListParameters200Response ListParameters(ctx).Execute()

List all global properties



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
	resp, r, err := apiClient.ParametersAPI.ListParameters(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParametersAPI.ListParameters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListParameters`: ListParameters200Response
	fmt.Fprintf(os.Stdout, "Response from `ParametersAPI.ListParameters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListParametersRequest struct via the builder pattern


### Return type

[**ListParameters200Response**](ListParameters200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ParameterDetails

> ParameterDetails200Response ParameterDetails(ctx, parameterId).Execute()

Get the value of a global property



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
	parameterId := "rudder_file_edit_header" // string | Id of the parameter to manage

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParametersAPI.ParameterDetails(context.Background(), parameterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParametersAPI.ParameterDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParameterDetails`: ParameterDetails200Response
	fmt.Fprintf(os.Stdout, "Response from `ParametersAPI.ParameterDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parameterId** | **string** | Id of the parameter to manage | 

### Other Parameters

Other parameters are passed through a pointer to a apiParameterDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ParameterDetails200Response**](ParameterDetails200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateParameter

> UpdateParameter200Response UpdateParameter(ctx, parameterId).Execute()

Update a global property's value



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
	parameterId := "rudder_file_edit_header" // string | Id of the parameter to manage

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParametersAPI.UpdateParameter(context.Background(), parameterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParametersAPI.UpdateParameter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateParameter`: UpdateParameter200Response
	fmt.Fprintf(os.Stdout, "Response from `ParametersAPI.UpdateParameter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parameterId** | **string** | Id of the parameter to manage | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateParameterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UpdateParameter200Response**](UpdateParameter200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

