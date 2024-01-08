# \DirectivesAPI

All URIs are relative to *https://rudder.example.local/rudder/api/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckDirective**](DirectivesAPI.md#CheckDirective) | **Post** /directives/{directiveId}/check | Check that update on a directive is valid
[**CreateDirective**](DirectivesAPI.md#CreateDirective) | **Put** /directives | Create a directive
[**DeleteDirective**](DirectivesAPI.md#DeleteDirective) | **Delete** /directives/{directiveId} | Delete a directive
[**DirectiveDetails**](DirectivesAPI.md#DirectiveDetails) | **Get** /directives/{directiveId} | Get directive details
[**ListDirectives**](DirectivesAPI.md#ListDirectives) | **Get** /directives | List all directives
[**UpdateDirective**](DirectivesAPI.md#UpdateDirective) | **Post** /directives/{directiveId} | Update a directive details



## CheckDirective

> CheckDirective200Response CheckDirective(ctx, directiveId).Directive(directive).Execute()

Check that update on a directive is valid



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
	directiveId := "9a1773c9-0889-40b6-be89-f6504443ac1b" // string | Id of the directive
	directive := *openapiclient.NewDirective() // Directive | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DirectivesAPI.CheckDirective(context.Background(), directiveId).Directive(directive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectivesAPI.CheckDirective``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckDirective`: CheckDirective200Response
	fmt.Fprintf(os.Stdout, "Response from `DirectivesAPI.CheckDirective`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directiveId** | **string** | Id of the directive | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckDirectiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **directive** | [**Directive**](Directive.md) |  | 

### Return type

[**CheckDirective200Response**](CheckDirective200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDirective

> CreateDirective200Response CreateDirective(ctx).DirectiveNew(directiveNew).Execute()

Create a directive



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
	directiveNew := *openapiclient.NewDirectiveNew() // DirectiveNew |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DirectivesAPI.CreateDirective(context.Background()).DirectiveNew(directiveNew).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectivesAPI.CreateDirective``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDirective`: CreateDirective200Response
	fmt.Fprintf(os.Stdout, "Response from `DirectivesAPI.CreateDirective`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDirectiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **directiveNew** | [**DirectiveNew**](DirectiveNew.md) |  | 

### Return type

[**CreateDirective200Response**](CreateDirective200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDirective

> DeleteDirective200Response DeleteDirective(ctx, directiveId).Execute()

Delete a directive



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
	directiveId := "9a1773c9-0889-40b6-be89-f6504443ac1b" // string | Id of the directive

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DirectivesAPI.DeleteDirective(context.Background(), directiveId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectivesAPI.DeleteDirective``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDirective`: DeleteDirective200Response
	fmt.Fprintf(os.Stdout, "Response from `DirectivesAPI.DeleteDirective`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directiveId** | **string** | Id of the directive | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDirectiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteDirective200Response**](DeleteDirective200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectiveDetails

> DirectiveDetails200Response DirectiveDetails(ctx, directiveId).Execute()

Get directive details



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
	directiveId := "9a1773c9-0889-40b6-be89-f6504443ac1b" // string | Id of the directive

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DirectivesAPI.DirectiveDetails(context.Background(), directiveId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectivesAPI.DirectiveDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DirectiveDetails`: DirectiveDetails200Response
	fmt.Fprintf(os.Stdout, "Response from `DirectivesAPI.DirectiveDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directiveId** | **string** | Id of the directive | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectiveDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DirectiveDetails200Response**](DirectiveDetails200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDirectives

> ListDirectives200Response ListDirectives(ctx).Execute()

List all directives



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
	resp, r, err := apiClient.DirectivesAPI.ListDirectives(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectivesAPI.ListDirectives``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDirectives`: ListDirectives200Response
	fmt.Fprintf(os.Stdout, "Response from `DirectivesAPI.ListDirectives`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDirectivesRequest struct via the builder pattern


### Return type

[**ListDirectives200Response**](ListDirectives200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDirective

> UpdateDirective200Response UpdateDirective(ctx, directiveId).Directive(directive).Execute()

Update a directive details



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
	directiveId := "9a1773c9-0889-40b6-be89-f6504443ac1b" // string | Id of the directive
	directive := *openapiclient.NewDirective() // Directive | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DirectivesAPI.UpdateDirective(context.Background(), directiveId).Directive(directive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectivesAPI.UpdateDirective``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDirective`: UpdateDirective200Response
	fmt.Fprintf(os.Stdout, "Response from `DirectivesAPI.UpdateDirective`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directiveId** | **string** | Id of the directive | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDirectiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **directive** | [**Directive**](Directive.md) |  | 

### Return type

[**UpdateDirective200Response**](UpdateDirective200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

