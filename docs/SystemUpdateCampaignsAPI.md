# \SystemUpdateCampaignsAPI

All URIs are relative to *https://rudder.example.local/rudder/api/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCampaignEventResult**](SystemUpdateCampaignsAPI.md#GetCampaignEventResult) | **Get** /systemUpdate/events/{id}/result | Get a campaign event result
[**GetCampaignResults**](SystemUpdateCampaignsAPI.md#GetCampaignResults) | **Get** /systemUpdate/campaign/{id}/history | Get a campaign result history
[**GetSystemUpdateResultForNode**](SystemUpdateCampaignsAPI.md#GetSystemUpdateResultForNode) | **Get** /systemUpdate/events/{id}/result/{nodeId} | Get detailed campaign event result for a Node



## GetCampaignEventResult

> GetCampaignEventResult200Response GetCampaignEventResult(ctx, id).Execute()

Get a campaign event result



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
	id := "0076a379-f32d-4732-9e91-33ab219d8fde" // string | Id of the campaign event

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemUpdateCampaignsAPI.GetCampaignEventResult(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemUpdateCampaignsAPI.GetCampaignEventResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignEventResult`: GetCampaignEventResult200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemUpdateCampaignsAPI.GetCampaignEventResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the campaign event | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignEventResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCampaignEventResult200Response**](GetCampaignEventResult200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignResults

> GetCampaignResults200Response GetCampaignResults(ctx, id).Limit(limit).Offset(offset).Before(before).After(after).Order(order).Asc(asc).Execute()

Get a campaign result history



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/juhnny5/rudder-golang"
)

func main() {
	id := "0076a379-f32d-4732-9e91-33ab219d8fde" // string | Id of the campaign
	limit := int32(56) // int32 | Max number of elements in response (optional)
	offset := int32(56) // int32 | Offset of data in response (skip X elements) (optional)
	before := time.Now() // string |  (optional)
	after := time.Now() // string |  (optional)
	order := "order_example" // string |  (optional)
	asc := "asc_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemUpdateCampaignsAPI.GetCampaignResults(context.Background(), id).Limit(limit).Offset(offset).Before(before).After(after).Order(order).Asc(asc).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemUpdateCampaignsAPI.GetCampaignResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignResults`: GetCampaignResults200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemUpdateCampaignsAPI.GetCampaignResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the campaign | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Max number of elements in response | 
 **offset** | **int32** | Offset of data in response (skip X elements) | 
 **before** | **string** |  | 
 **after** | **string** |  | 
 **order** | **string** |  | 
 **asc** | **string** |  | 

### Return type

[**GetCampaignResults200Response**](GetCampaignResults200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemUpdateResultForNode

> GetSystemUpdateResultForNode200Response GetSystemUpdateResultForNode(ctx, id, nodeId).Execute()

Get detailed campaign event result for a Node



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
	id := "0076a379-f32d-4732-9e91-33ab219d8fde" // string | Id of the campaign event
	nodeId := "9a1773c9-0889-40b6-be89-f6504443ac1b" // string | Id of the target node

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemUpdateCampaignsAPI.GetSystemUpdateResultForNode(context.Background(), id, nodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemUpdateCampaignsAPI.GetSystemUpdateResultForNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemUpdateResultForNode`: GetSystemUpdateResultForNode200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemUpdateCampaignsAPI.GetSystemUpdateResultForNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the campaign event | 
**nodeId** | **string** | Id of the target node | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemUpdateResultForNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetSystemUpdateResultForNode200Response**](GetSystemUpdateResultForNode200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

