# \OpenSCAPAPI

All URIs are relative to *https://rudder.example.local/rudder/api/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpenscapReport**](OpenSCAPAPI.md#OpenscapReport) | **Get** /openscap/report/{nodeId} | Get an OpenSCAP report



## OpenscapReport

> string OpenscapReport(ctx, nodeId).Execute()

Get an OpenSCAP report



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
	nodeId := "9a1773c9-0889-40b6-be89-f6504443ac1b" // string | Id of the target node

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenSCAPAPI.OpenscapReport(context.Background(), nodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenSCAPAPI.OpenscapReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenscapReport`: string
	fmt.Fprintf(os.Stdout, "Response from `OpenSCAPAPI.OpenscapReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | Id of the target node | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpenscapReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

