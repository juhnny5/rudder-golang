# \CVEAPI

All URIs are relative to *https://rudder.example.local/rudder/api/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckCVE**](CVEAPI.md#CheckCVE) | **Post** /cve/check | Trigger a CVE check
[**GetAllCve**](CVEAPI.md#GetAllCve) | **Get** /cve | Get all CVE details
[**GetCVECheckConfiguration**](CVEAPI.md#GetCVECheckConfiguration) | **Get** /cve/check/config | Get CVE check config
[**GetCVEList**](CVEAPI.md#GetCVEList) | **Post** /cve/list | Get a list of CVE details
[**GetCve**](CVEAPI.md#GetCve) | **Get** /cve/{cveId} | Get a CVE details
[**GetLastCVECheck**](CVEAPI.md#GetLastCVECheck) | **Get** /cve/check/last | Get last CVE check result
[**ReadCVEfromFS**](CVEAPI.md#ReadCVEfromFS) | **Post** /cve/update/fs | Update CVE database from file system
[**UpdateCVE**](CVEAPI.md#UpdateCVE) | **Post** /cve/update | Update CVE database from remote source
[**UpdateCVECheckConfiguration**](CVEAPI.md#UpdateCVECheckConfiguration) | **Post** /cve/check/config | Update cve check config



## CheckCVE

> CheckCVE200Response CheckCVE(ctx).Execute()

Trigger a CVE check



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
	resp, r, err := apiClient.CVEAPI.CheckCVE(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CVEAPI.CheckCVE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckCVE`: CheckCVE200Response
	fmt.Fprintf(os.Stdout, "Response from `CVEAPI.CheckCVE`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCheckCVERequest struct via the builder pattern


### Return type

[**CheckCVE200Response**](CheckCVE200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllCve

> GetAllCve200Response GetAllCve(ctx).Execute()

Get all CVE details



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
	resp, r, err := apiClient.CVEAPI.GetAllCve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CVEAPI.GetAllCve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllCve`: GetAllCve200Response
	fmt.Fprintf(os.Stdout, "Response from `CVEAPI.GetAllCve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllCveRequest struct via the builder pattern


### Return type

[**GetAllCve200Response**](GetAllCve200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCVECheckConfiguration

> GetCVECheckConfiguration200Response GetCVECheckConfiguration(ctx).Execute()

Get CVE check config



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
	resp, r, err := apiClient.CVEAPI.GetCVECheckConfiguration(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CVEAPI.GetCVECheckConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCVECheckConfiguration`: GetCVECheckConfiguration200Response
	fmt.Fprintf(os.Stdout, "Response from `CVEAPI.GetCVECheckConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCVECheckConfigurationRequest struct via the builder pattern


### Return type

[**GetCVECheckConfiguration200Response**](GetCVECheckConfiguration200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCVEList

> GetCVEList200Response GetCVEList(ctx).GetCVEListRequest(getCVEListRequest).Execute()

Get a list of CVE details



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
	getCVEListRequest := *openapiclient.NewGetCVEListRequest() // GetCVEListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CVEAPI.GetCVEList(context.Background()).GetCVEListRequest(getCVEListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CVEAPI.GetCVEList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCVEList`: GetCVEList200Response
	fmt.Fprintf(os.Stdout, "Response from `CVEAPI.GetCVEList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCVEListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getCVEListRequest** | [**GetCVEListRequest**](GetCVEListRequest.md) |  | 

### Return type

[**GetCVEList200Response**](GetCVEList200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCve

> GetCve200Response GetCve(ctx, cveId).Execute()

Get a CVE details



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
	cveId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Id of the CVE

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CVEAPI.GetCve(context.Background(), cveId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CVEAPI.GetCve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCve`: GetCve200Response
	fmt.Fprintf(os.Stdout, "Response from `CVEAPI.GetCve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cveId** | **string** | Id of the CVE | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCve200Response**](GetCve200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLastCVECheck

> GetLastCVECheck200Response GetLastCVECheck(ctx).GroupId(groupId).NodeId(nodeId).CveId(cveId).Package_(package_).Severity(severity).Execute()

Get last CVE check result



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
	groupId := "groupId_example" // string | Id of node groups you want to get from last CVE check (optional)
	nodeId := "nodeId_example" // string | Id of nodes you want to get from last CVE check (optional)
	cveId := "cveId_example" // string | Id of CVE you want to get from last CVE check (optional)
	package_ := "package__example" // string | Name of packages you want to get from last CVE check (optional)
	severity := "severity_example" // string | Severity of the CVE you want to get from last CVE check (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CVEAPI.GetLastCVECheck(context.Background()).GroupId(groupId).NodeId(nodeId).CveId(cveId).Package_(package_).Severity(severity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CVEAPI.GetLastCVECheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLastCVECheck`: GetLastCVECheck200Response
	fmt.Fprintf(os.Stdout, "Response from `CVEAPI.GetLastCVECheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLastCVECheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string** | Id of node groups you want to get from last CVE check | 
 **nodeId** | **string** | Id of nodes you want to get from last CVE check | 
 **cveId** | **string** | Id of CVE you want to get from last CVE check | 
 **package_** | **string** | Name of packages you want to get from last CVE check | 
 **severity** | **string** | Severity of the CVE you want to get from last CVE check | 

### Return type

[**GetLastCVECheck200Response**](GetLastCVECheck200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadCVEfromFS

> ReadCVEfromFS200Response ReadCVEfromFS(ctx).Execute()

Update CVE database from file system



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
	resp, r, err := apiClient.CVEAPI.ReadCVEfromFS(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CVEAPI.ReadCVEfromFS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadCVEfromFS`: ReadCVEfromFS200Response
	fmt.Fprintf(os.Stdout, "Response from `CVEAPI.ReadCVEfromFS`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReadCVEfromFSRequest struct via the builder pattern


### Return type

[**ReadCVEfromFS200Response**](ReadCVEfromFS200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCVE

> UpdateCVE200Response UpdateCVE(ctx).UpdateCVERequest(updateCVERequest).Execute()

Update CVE database from remote source



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
	updateCVERequest := *openapiclient.NewUpdateCVERequest() // UpdateCVERequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CVEAPI.UpdateCVE(context.Background()).UpdateCVERequest(updateCVERequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CVEAPI.UpdateCVE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCVE`: UpdateCVE200Response
	fmt.Fprintf(os.Stdout, "Response from `CVEAPI.UpdateCVE`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCVERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateCVERequest** | [**UpdateCVERequest**](UpdateCVERequest.md) |  | 

### Return type

[**UpdateCVE200Response**](UpdateCVE200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCVECheckConfiguration

> UpdateCVECheckConfiguration200Response UpdateCVECheckConfiguration(ctx).UpdateCVECheckConfigurationRequest(updateCVECheckConfigurationRequest).Execute()

Update cve check config



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
	updateCVECheckConfigurationRequest := *openapiclient.NewUpdateCVECheckConfigurationRequest() // UpdateCVECheckConfigurationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CVEAPI.UpdateCVECheckConfiguration(context.Background()).UpdateCVECheckConfigurationRequest(updateCVECheckConfigurationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CVEAPI.UpdateCVECheckConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCVECheckConfiguration`: UpdateCVECheckConfiguration200Response
	fmt.Fprintf(os.Stdout, "Response from `CVEAPI.UpdateCVECheckConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCVECheckConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateCVECheckConfigurationRequest** | [**UpdateCVECheckConfigurationRequest**](UpdateCVECheckConfigurationRequest.md) |  | 

### Return type

[**UpdateCVECheckConfiguration200Response**](UpdateCVECheckConfiguration200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

