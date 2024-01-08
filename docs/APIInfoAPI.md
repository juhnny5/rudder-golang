# \APIInfoAPI

All URIs are relative to *https://rudder.example.local/rudder/api/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiGeneralInformations**](APIInfoAPI.md#ApiGeneralInformations) | **Get** /info | List all endpoints
[**ApiInformations**](APIInfoAPI.md#ApiInformations) | **Get** /info/details/{endpointName} | Get information about one API endpoint
[**ApiSubInformations**](APIInfoAPI.md#ApiSubInformations) | **Get** /info/{sectionId} | Get information on endpoint in a section



## ApiGeneralInformations

> ApiGeneralInformations200Response ApiGeneralInformations(ctx).Execute()

List all endpoints



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
	resp, r, err := apiClient.APIInfoAPI.ApiGeneralInformations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIInfoAPI.ApiGeneralInformations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiGeneralInformations`: ApiGeneralInformations200Response
	fmt.Fprintf(os.Stdout, "Response from `APIInfoAPI.ApiGeneralInformations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiGeneralInformationsRequest struct via the builder pattern


### Return type

[**ApiGeneralInformations200Response**](ApiGeneralInformations200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiInformations

> ApiInformations200Response ApiInformations(ctx, endpointName).Execute()

Get information about one API endpoint



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
	endpointName := "listAcceptedNodes" // string | Name of the endpoint for which one wants information

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIInfoAPI.ApiInformations(context.Background(), endpointName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIInfoAPI.ApiInformations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiInformations`: ApiInformations200Response
	fmt.Fprintf(os.Stdout, "Response from `APIInfoAPI.ApiInformations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointName** | **string** | Name of the endpoint for which one wants information | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiInformationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiInformations200Response**](ApiInformations200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiSubInformations

> ApiSubInformations200Response ApiSubInformations(ctx, sectionId).Execute()

Get information on endpoint in a section



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
	sectionId := "nodes" // string | Id of the API section

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIInfoAPI.ApiSubInformations(context.Background(), sectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIInfoAPI.ApiSubInformations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiSubInformations`: ApiSubInformations200Response
	fmt.Fprintf(os.Stdout, "Response from `APIInfoAPI.ApiSubInformations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sectionId** | **string** | Id of the API section | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSubInformationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiSubInformations200Response**](ApiSubInformations200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

