# \TechniquesAPI

All URIs are relative to *https://rudder.example.local/rudder/api/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTechnique**](TechniquesAPI.md#CreateTechnique) | **Put** /techniques | Create technique
[**DeleteTechnique**](TechniquesAPI.md#DeleteTechnique) | **Delete** /techniques/{techniqueId}/{techniqueVersion} | Delete technique
[**GetTechniqueAllVersion**](TechniquesAPI.md#GetTechniqueAllVersion) | **Get** /techniques/{techniqueId} | Technique metadata by ID
[**GetTechniqueAllVersionId**](TechniquesAPI.md#GetTechniqueAllVersionId) | **Get** /techniques/{techniqueId}/{techniqueVersion} | Technique metadata by version and ID
[**GetTechniquesResources**](TechniquesAPI.md#GetTechniquesResources) | **Get** /techniques/{techniqueId}/{techniqueVersion}/resources | Technique&#39;s resources
[**ListTechniqueVersionDirectives**](TechniquesAPI.md#ListTechniqueVersionDirectives) | **Get** /techniques/{techniqueId}/{techniqueVersion}/directives | List all directives based on a version of a technique
[**ListTechniques**](TechniquesAPI.md#ListTechniques) | **Get** /techniques | List all techniques
[**ListTechniquesDirectives**](TechniquesAPI.md#ListTechniquesDirectives) | **Get** /techniques/{techniqueId}/directives | List all directives based on a technique
[**ListTechniquesVersions**](TechniquesAPI.md#ListTechniquesVersions) | **Get** /techniques/versions | List versions
[**Methods**](TechniquesAPI.md#Methods) | **Get** /methods | List methods
[**ReloadMethods**](TechniquesAPI.md#ReloadMethods) | **Post** /methods/reload | Reload methods
[**TechniqueCategories**](TechniquesAPI.md#TechniqueCategories) | **Get** /techniques/categories | List categories
[**TechniqueRevisions**](TechniquesAPI.md#TechniqueRevisions) | **Get** /techniques/{techniqueId}/{techniqueVersion}/revisions | Technique&#39;s revisions
[**Techniques**](TechniquesAPI.md#Techniques) | **Post** /techniques/reload | Reload techniques
[**UpdateTechnique**](TechniquesAPI.md#UpdateTechnique) | **Post** /techniques/{techniqueId}/{techniqueVersion} | Update technique



## CreateTechnique

> CreateTechnique200Response CreateTechnique(ctx).EditorTechnique(editorTechnique).Execute()

Create technique



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
	editorTechnique := []openapiclient.EditorTechnique{*openapiclient.NewEditorTechnique()} // []EditorTechnique | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TechniquesAPI.CreateTechnique(context.Background()).EditorTechnique(editorTechnique).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TechniquesAPI.CreateTechnique``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTechnique`: CreateTechnique200Response
	fmt.Fprintf(os.Stdout, "Response from `TechniquesAPI.CreateTechnique`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTechniqueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **editorTechnique** | [**[]EditorTechnique**](EditorTechnique.md) |  | 

### Return type

[**CreateTechnique200Response**](CreateTechnique200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTechnique

> DeleteTechnique200Response DeleteTechnique(ctx, techniqueId, techniqueVersion).Execute()

Delete technique



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
	techniqueId := "userManagement" // string | Technique ID
	techniqueVersion := "6.0" // string | Technique version

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TechniquesAPI.DeleteTechnique(context.Background(), techniqueId, techniqueVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TechniquesAPI.DeleteTechnique``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTechnique`: DeleteTechnique200Response
	fmt.Fprintf(os.Stdout, "Response from `TechniquesAPI.DeleteTechnique`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**techniqueId** | **string** | Technique ID | 
**techniqueVersion** | **string** | Technique version | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTechniqueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteTechnique200Response**](DeleteTechnique200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTechniqueAllVersion

> GetTechniqueAllVersion200Response GetTechniqueAllVersion(ctx, techniqueId).Execute()

Technique metadata by ID



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
	techniqueId := "userManagement" // string | Technique ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TechniquesAPI.GetTechniqueAllVersion(context.Background(), techniqueId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TechniquesAPI.GetTechniqueAllVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTechniqueAllVersion`: GetTechniqueAllVersion200Response
	fmt.Fprintf(os.Stdout, "Response from `TechniquesAPI.GetTechniqueAllVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**techniqueId** | **string** | Technique ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTechniqueAllVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetTechniqueAllVersion200Response**](GetTechniqueAllVersion200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTechniqueAllVersionId

> GetTechniqueAllVersion200Response GetTechniqueAllVersionId(ctx, techniqueId, techniqueVersion).Execute()

Technique metadata by version and ID



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
	techniqueId := "userManagement" // string | Technique ID
	techniqueVersion := "6.0" // string | Technique version

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TechniquesAPI.GetTechniqueAllVersionId(context.Background(), techniqueId, techniqueVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TechniquesAPI.GetTechniqueAllVersionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTechniqueAllVersionId`: GetTechniqueAllVersion200Response
	fmt.Fprintf(os.Stdout, "Response from `TechniquesAPI.GetTechniqueAllVersionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**techniqueId** | **string** | Technique ID | 
**techniqueVersion** | **string** | Technique version | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTechniqueAllVersionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetTechniqueAllVersion200Response**](GetTechniqueAllVersion200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTechniquesResources

> GetTechniquesResources200Response GetTechniquesResources(ctx, techniqueId, techniqueVersion).Execute()

Technique's resources



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
	techniqueId := "userManagement" // string | Technique ID
	techniqueVersion := "6.0" // string | Technique version

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TechniquesAPI.GetTechniquesResources(context.Background(), techniqueId, techniqueVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TechniquesAPI.GetTechniquesResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTechniquesResources`: GetTechniquesResources200Response
	fmt.Fprintf(os.Stdout, "Response from `TechniquesAPI.GetTechniquesResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**techniqueId** | **string** | Technique ID | 
**techniqueVersion** | **string** | Technique version | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTechniquesResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetTechniquesResources200Response**](GetTechniquesResources200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTechniqueVersionDirectives

> ListTechniqueVersionDirectives200Response ListTechniqueVersionDirectives(ctx, techniqueId, techniqueVersion).Execute()

List all directives based on a version of a technique



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
	techniqueId := "userManagement" // string | Technique ID
	techniqueVersion := "6.0" // string | Technique version

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TechniquesAPI.ListTechniqueVersionDirectives(context.Background(), techniqueId, techniqueVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TechniquesAPI.ListTechniqueVersionDirectives``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTechniqueVersionDirectives`: ListTechniqueVersionDirectives200Response
	fmt.Fprintf(os.Stdout, "Response from `TechniquesAPI.ListTechniqueVersionDirectives`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**techniqueId** | **string** | Technique ID | 
**techniqueVersion** | **string** | Technique version | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTechniqueVersionDirectivesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListTechniqueVersionDirectives200Response**](ListTechniqueVersionDirectives200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTechniques

> ListTechniques200Response ListTechniques(ctx).Execute()

List all techniques



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
	resp, r, err := apiClient.TechniquesAPI.ListTechniques(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TechniquesAPI.ListTechniques``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTechniques`: ListTechniques200Response
	fmt.Fprintf(os.Stdout, "Response from `TechniquesAPI.ListTechniques`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTechniquesRequest struct via the builder pattern


### Return type

[**ListTechniques200Response**](ListTechniques200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTechniquesDirectives

> ListTechniquesDirectives200Response ListTechniquesDirectives(ctx, techniqueId).Execute()

List all directives based on a technique



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
	techniqueId := "userManagement" // string | Technique ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TechniquesAPI.ListTechniquesDirectives(context.Background(), techniqueId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TechniquesAPI.ListTechniquesDirectives``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTechniquesDirectives`: ListTechniquesDirectives200Response
	fmt.Fprintf(os.Stdout, "Response from `TechniquesAPI.ListTechniquesDirectives`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**techniqueId** | **string** | Technique ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTechniquesDirectivesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListTechniquesDirectives200Response**](ListTechniquesDirectives200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTechniquesVersions

> ListTechniquesVersions200Response ListTechniquesVersions(ctx).Execute()

List versions



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
	resp, r, err := apiClient.TechniquesAPI.ListTechniquesVersions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TechniquesAPI.ListTechniquesVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTechniquesVersions`: ListTechniquesVersions200Response
	fmt.Fprintf(os.Stdout, "Response from `TechniquesAPI.ListTechniquesVersions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTechniquesVersionsRequest struct via the builder pattern


### Return type

[**ListTechniquesVersions200Response**](ListTechniquesVersions200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Methods

> Methods200Response Methods(ctx).Execute()

List methods



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
	resp, r, err := apiClient.TechniquesAPI.Methods(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TechniquesAPI.Methods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Methods`: Methods200Response
	fmt.Fprintf(os.Stdout, "Response from `TechniquesAPI.Methods`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMethodsRequest struct via the builder pattern


### Return type

[**Methods200Response**](Methods200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReloadMethods

> Methods200Response ReloadMethods(ctx).Execute()

Reload methods



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
	resp, r, err := apiClient.TechniquesAPI.ReloadMethods(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TechniquesAPI.ReloadMethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReloadMethods`: Methods200Response
	fmt.Fprintf(os.Stdout, "Response from `TechniquesAPI.ReloadMethods`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReloadMethodsRequest struct via the builder pattern


### Return type

[**Methods200Response**](Methods200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TechniqueCategories

> TechniqueCategories200Response TechniqueCategories(ctx).Execute()

List categories



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
	resp, r, err := apiClient.TechniquesAPI.TechniqueCategories(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TechniquesAPI.TechniqueCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TechniqueCategories`: TechniqueCategories200Response
	fmt.Fprintf(os.Stdout, "Response from `TechniquesAPI.TechniqueCategories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTechniqueCategoriesRequest struct via the builder pattern


### Return type

[**TechniqueCategories200Response**](TechniqueCategories200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TechniqueRevisions

> TechniqueRevisions200Response TechniqueRevisions(ctx, techniqueId, techniqueVersion).Execute()

Technique's revisions



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
	techniqueId := "userManagement" // string | Technique ID
	techniqueVersion := "6.0" // string | Technique version

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TechniquesAPI.TechniqueRevisions(context.Background(), techniqueId, techniqueVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TechniquesAPI.TechniqueRevisions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TechniqueRevisions`: TechniqueRevisions200Response
	fmt.Fprintf(os.Stdout, "Response from `TechniquesAPI.TechniqueRevisions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**techniqueId** | **string** | Technique ID | 
**techniqueVersion** | **string** | Technique version | 

### Other Parameters

Other parameters are passed through a pointer to a apiTechniqueRevisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TechniqueRevisions200Response**](TechniqueRevisions200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Techniques

> ListTechniques200Response Techniques(ctx).Execute()

Reload techniques



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
	resp, r, err := apiClient.TechniquesAPI.Techniques(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TechniquesAPI.Techniques``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Techniques`: ListTechniques200Response
	fmt.Fprintf(os.Stdout, "Response from `TechniquesAPI.Techniques`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTechniquesRequest struct via the builder pattern


### Return type

[**ListTechniques200Response**](ListTechniques200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTechnique

> CreateTechnique200Response UpdateTechnique(ctx, techniqueId, techniqueVersion).EditorTechnique(editorTechnique).Execute()

Update technique



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
	techniqueId := "userManagement" // string | Technique ID
	techniqueVersion := "6.0" // string | Technique version
	editorTechnique := []openapiclient.EditorTechnique{*openapiclient.NewEditorTechnique()} // []EditorTechnique | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TechniquesAPI.UpdateTechnique(context.Background(), techniqueId, techniqueVersion).EditorTechnique(editorTechnique).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TechniquesAPI.UpdateTechnique``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTechnique`: CreateTechnique200Response
	fmt.Fprintf(os.Stdout, "Response from `TechniquesAPI.UpdateTechnique`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**techniqueId** | **string** | Technique ID | 
**techniqueVersion** | **string** | Technique version | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTechniqueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **editorTechnique** | [**[]EditorTechnique**](EditorTechnique.md) |  | 

### Return type

[**CreateTechnique200Response**](CreateTechnique200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

