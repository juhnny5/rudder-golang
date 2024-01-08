# \ScaleOutRelayAPI

All URIs are relative to *https://rudder.example.local/rudder/api/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DemoteToNode**](ScaleOutRelayAPI.md#DemoteToNode) | **Post** /scaleoutrelay/demote/{nodeId} | Demote a relay to simple node
[**PromoteToRelay**](ScaleOutRelayAPI.md#PromoteToRelay) | **Post** /scaleoutrelay/promote/{nodeId} | Promote a node to relay



## DemoteToNode

> DemoteToNode200Response DemoteToNode(ctx, nodeId).Execute()

Demote a relay to simple node



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
	resp, r, err := apiClient.ScaleOutRelayAPI.DemoteToNode(context.Background(), nodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScaleOutRelayAPI.DemoteToNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DemoteToNode`: DemoteToNode200Response
	fmt.Fprintf(os.Stdout, "Response from `ScaleOutRelayAPI.DemoteToNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | Id of the target node | 

### Other Parameters

Other parameters are passed through a pointer to a apiDemoteToNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DemoteToNode200Response**](DemoteToNode200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PromoteToRelay

> PromoteToRelay200Response PromoteToRelay(ctx, nodeId).Execute()

Promote a node to relay



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
	resp, r, err := apiClient.ScaleOutRelayAPI.PromoteToRelay(context.Background(), nodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScaleOutRelayAPI.PromoteToRelay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PromoteToRelay`: PromoteToRelay200Response
	fmt.Fprintf(os.Stdout, "Response from `ScaleOutRelayAPI.PromoteToRelay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | Id of the target node | 

### Other Parameters

Other parameters are passed through a pointer to a apiPromoteToRelayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PromoteToRelay200Response**](PromoteToRelay200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

