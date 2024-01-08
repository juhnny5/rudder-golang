# \SecretManagementAPI

All URIs are relative to *https://rudder.example.local/rudder/api/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSecret**](SecretManagementAPI.md#AddSecret) | **Put** /secret | Create a secret
[**DeleteSecret**](SecretManagementAPI.md#DeleteSecret) | **Delete** /secret/{name} | Delete a secret
[**GetAllSecrets**](SecretManagementAPI.md#GetAllSecrets) | **Get** /secret | List all secrets
[**GetSecret**](SecretManagementAPI.md#GetSecret) | **Get** /secret/{name} | Get one secret
[**UpdateSecret**](SecretManagementAPI.md#UpdateSecret) | **Post** /secret | Update a secret



## AddSecret

> AddSecret200Response AddSecret(ctx).Secrets(secrets).Execute()

Create a secret



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
	secrets := *openapiclient.NewSecrets() // Secrets | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretManagementAPI.AddSecret(context.Background()).Secrets(secrets).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretManagementAPI.AddSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddSecret`: AddSecret200Response
	fmt.Fprintf(os.Stdout, "Response from `SecretManagementAPI.AddSecret`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **secrets** | [**Secrets**](Secrets.md) |  | 

### Return type

[**AddSecret200Response**](AddSecret200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSecret

> DeleteSecret200Response DeleteSecret(ctx, name).Execute()

Delete a secret



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
	name := "name_example" // string | Unique name of the secret

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretManagementAPI.DeleteSecret(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretManagementAPI.DeleteSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSecret`: DeleteSecret200Response
	fmt.Fprintf(os.Stdout, "Response from `SecretManagementAPI.DeleteSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Unique name of the secret | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteSecret200Response**](DeleteSecret200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllSecrets

> GetAllSecrets200Response GetAllSecrets(ctx).Execute()

List all secrets



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
	resp, r, err := apiClient.SecretManagementAPI.GetAllSecrets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretManagementAPI.GetAllSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllSecrets`: GetAllSecrets200Response
	fmt.Fprintf(os.Stdout, "Response from `SecretManagementAPI.GetAllSecrets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSecretsRequest struct via the builder pattern


### Return type

[**GetAllSecrets200Response**](GetAllSecrets200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecret

> GetSecret200Response GetSecret(ctx, name).Execute()

Get one secret



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
	name := "name_example" // string | Unique name of the secret

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretManagementAPI.GetSecret(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretManagementAPI.GetSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecret`: GetSecret200Response
	fmt.Fprintf(os.Stdout, "Response from `SecretManagementAPI.GetSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Unique name of the secret | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSecret200Response**](GetSecret200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSecret

> UpdateSecret200Response UpdateSecret(ctx).Secrets(secrets).Execute()

Update a secret



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
	secrets := *openapiclient.NewSecrets() // Secrets | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretManagementAPI.UpdateSecret(context.Background()).Secrets(secrets).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretManagementAPI.UpdateSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSecret`: UpdateSecret200Response
	fmt.Fprintf(os.Stdout, "Response from `SecretManagementAPI.UpdateSecret`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **secrets** | [**Secrets**](Secrets.md) |  | 

### Return type

[**UpdateSecret200Response**](UpdateSecret200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

