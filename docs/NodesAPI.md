# \NodesAPI

All URIs are relative to *https://rudder.example.local/rudder/api/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyPolicy**](NodesAPI.md#ApplyPolicy) | **Post** /nodes/{nodeId}/applyPolicy | Trigger an agent run
[**ApplyPolicyAllNodes**](NodesAPI.md#ApplyPolicyAllNodes) | **Post** /nodes/applyPolicy | Trigger an agent run on all nodes
[**ChangePendingNodeStatus**](NodesAPI.md#ChangePendingNodeStatus) | **Post** /nodes/pending/{nodeId} | Update pending Node status
[**CreateNodes**](NodesAPI.md#CreateNodes) | **Put** /nodes | Create one or several new nodes
[**DeleteNode**](NodesAPI.md#DeleteNode) | **Delete** /nodes/{nodeId} | Delete a node
[**GetNodesStatus**](NodesAPI.md#GetNodesStatus) | **Get** /nodes/status | Get nodes acceptation status
[**ListAcceptedNodes**](NodesAPI.md#ListAcceptedNodes) | **Get** /nodes | List managed nodes
[**ListPendingNodes**](NodesAPI.md#ListPendingNodes) | **Get** /nodes/pending | List pending nodes
[**NodeDetails**](NodesAPI.md#NodeDetails) | **Get** /nodes/{nodeId} | Get information about a node
[**NodeInheritedProperties**](NodesAPI.md#NodeInheritedProperties) | **Get** /nodes/{nodeId}/inheritedProperties | Get inherited node properties for a node
[**UpdateNode**](NodesAPI.md#UpdateNode) | **Post** /nodes/{nodeId} | Update node settings and properties



## ApplyPolicy

> string ApplyPolicy(ctx, nodeId).Execute()

Trigger an agent run



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
	resp, r, err := apiClient.NodesAPI.ApplyPolicy(context.Background(), nodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.ApplyPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyPolicy`: string
	fmt.Fprintf(os.Stdout, "Response from `NodesAPI.ApplyPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | Id of the target node | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplyPolicyAllNodes

> ApplyPolicyAllNodes200Response ApplyPolicyAllNodes(ctx).Execute()

Trigger an agent run on all nodes



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
	resp, r, err := apiClient.NodesAPI.ApplyPolicyAllNodes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.ApplyPolicyAllNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyPolicyAllNodes`: ApplyPolicyAllNodes200Response
	fmt.Fprintf(os.Stdout, "Response from `NodesAPI.ApplyPolicyAllNodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApplyPolicyAllNodesRequest struct via the builder pattern


### Return type

[**ApplyPolicyAllNodes200Response**](ApplyPolicyAllNodes200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangePendingNodeStatus

> ChangePendingNodeStatus200Response ChangePendingNodeStatus(ctx, nodeId).ChangePendingNodeStatusRequest(changePendingNodeStatusRequest).Execute()

Update pending Node status



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
	changePendingNodeStatusRequest := *openapiclient.NewChangePendingNodeStatusRequest() // ChangePendingNodeStatusRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodesAPI.ChangePendingNodeStatus(context.Background(), nodeId).ChangePendingNodeStatusRequest(changePendingNodeStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.ChangePendingNodeStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangePendingNodeStatus`: ChangePendingNodeStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `NodesAPI.ChangePendingNodeStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | Id of the target node | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangePendingNodeStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **changePendingNodeStatusRequest** | [**ChangePendingNodeStatusRequest**](ChangePendingNodeStatusRequest.md) |  | 

### Return type

[**ChangePendingNodeStatus200Response**](ChangePendingNodeStatus200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNodes

> CreateNodes200Response CreateNodes(ctx).NodeAddInner(nodeAddInner).Execute()

Create one or several new nodes



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
	nodeAddInner := []openapiclient.NodeAddInner{*openapiclient.NewNodeAddInner("378740d3-c4a9-4474-8485-478e7e52db52", "my.node.hostname.local", "Status_example", *openapiclient.NewOs("Type_example", "Name_example", "9.5", "Debian GNU/Linux 9 (stretch)"), "MachineType_example", []openapiclient.NodeFullPropertiesInner{*openapiclient.NewNodeFullPropertiesInner("datacenter", interface{}(AMS2))}, []string{"192.168.180.90"})} // []NodeAddInner | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodesAPI.CreateNodes(context.Background()).NodeAddInner(nodeAddInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.CreateNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNodes`: CreateNodes200Response
	fmt.Fprintf(os.Stdout, "Response from `NodesAPI.CreateNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeAddInner** | [**[]NodeAddInner**](NodeAddInner.md) |  | 

### Return type

[**CreateNodes200Response**](CreateNodes200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNode

> DeleteNode200Response DeleteNode(ctx, nodeId).Mode(mode).Execute()

Delete a node



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
	mode := "move" // string | Deletion mode to use, either move to trash ('move', default) or erase ('erase') (optional) (default to "move")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodesAPI.DeleteNode(context.Background(), nodeId).Mode(mode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.DeleteNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNode`: DeleteNode200Response
	fmt.Fprintf(os.Stdout, "Response from `NodesAPI.DeleteNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | Id of the target node | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mode** | **string** | Deletion mode to use, either move to trash (&#39;move&#39;, default) or erase (&#39;erase&#39;) | [default to &quot;move&quot;]

### Return type

[**DeleteNode200Response**](DeleteNode200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodesStatus

> GetNodesStatus200Response GetNodesStatus(ctx).Ids(ids).Execute()

Get nodes acceptation status



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
	ids := "8403353b-42c4-46f5-a08d-bc77a1a0bad9,root" // string | Comma separated list of node Ids (default to "default")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodesAPI.GetNodesStatus(context.Background()).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.GetNodesStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNodesStatus`: GetNodesStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `NodesAPI.GetNodesStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNodesStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** | Comma separated list of node Ids | [default to &quot;default&quot;]

### Return type

[**GetNodesStatus200Response**](GetNodesStatus200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAcceptedNodes

> ListAcceptedNodes200Response ListAcceptedNodes(ctx).Include(include).Query(query).Where(where).Composition(composition).Select_(select_).Execute()

List managed nodes



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
	include := "minimal" // string | Level of information to include from the node inventory. Some base levels are defined (**minimal**, **default**, **full**). You can add fields you want to a base level by adding them to the list, possible values are keys from json answer. If you don't provide a base level, they will be added to `default` level, so if you only want os details, use `minimal,os` as the value for this parameter. * **minimal** includes: `id`, `hostname` and `status` * **default** includes **minimal** plus `architectureDescription`, `description`, `ipAddresses`, `lastRunDate`, `lastInventoryDate`, `machine`, `os`, `managementTechnology`, `policyServerId`, `properties` (be careful! Only node own properties, if you also need inherited properties, look at the dedicated `/nodes/{id}/inheritedProperties` endpoint), `policyMode `, `ram` and `timezone` * **full** includes: **default** plus `accounts`, `bios`, `controllers`, `environmentVariables`, `fileSystems`, `managementTechnologyDetails`, `memories`, `networkInterfaces`, `ports`, `processes`, `processors`, `slots`, `software`, `softwareUpdate`, `sound`, `storage`, `videos` and `virtualMachines` (optional) (default to "default")
	query := TODO // map[string]interface{} | The criterion you want to find for your nodes. Replaces the `where`, `composition` and `select` parameters in a single parameter. (optional)
	where := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | The criterion you want to find for your nodes (optional)
	composition := "and" // string | Boolean operator to use between each  `where` criteria. (optional) (default to "and")
	select_ := "select__example" // string | What kind of data we want to include. Here we can get policy servers/relay by setting `nodeAndPolicyServer`. Only used if `where` is defined. (optional) (default to "node")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodesAPI.ListAcceptedNodes(context.Background()).Include(include).Query(query).Where(where).Composition(composition).Select_(select_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.ListAcceptedNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAcceptedNodes`: ListAcceptedNodes200Response
	fmt.Fprintf(os.Stdout, "Response from `NodesAPI.ListAcceptedNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAcceptedNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **string** | Level of information to include from the node inventory. Some base levels are defined (**minimal**, **default**, **full**). You can add fields you want to a base level by adding them to the list, possible values are keys from json answer. If you don&#39;t provide a base level, they will be added to &#x60;default&#x60; level, so if you only want os details, use &#x60;minimal,os&#x60; as the value for this parameter. * **minimal** includes: &#x60;id&#x60;, &#x60;hostname&#x60; and &#x60;status&#x60; * **default** includes **minimal** plus &#x60;architectureDescription&#x60;, &#x60;description&#x60;, &#x60;ipAddresses&#x60;, &#x60;lastRunDate&#x60;, &#x60;lastInventoryDate&#x60;, &#x60;machine&#x60;, &#x60;os&#x60;, &#x60;managementTechnology&#x60;, &#x60;policyServerId&#x60;, &#x60;properties&#x60; (be careful! Only node own properties, if you also need inherited properties, look at the dedicated &#x60;/nodes/{id}/inheritedProperties&#x60; endpoint), &#x60;policyMode &#x60;, &#x60;ram&#x60; and &#x60;timezone&#x60; * **full** includes: **default** plus &#x60;accounts&#x60;, &#x60;bios&#x60;, &#x60;controllers&#x60;, &#x60;environmentVariables&#x60;, &#x60;fileSystems&#x60;, &#x60;managementTechnologyDetails&#x60;, &#x60;memories&#x60;, &#x60;networkInterfaces&#x60;, &#x60;ports&#x60;, &#x60;processes&#x60;, &#x60;processors&#x60;, &#x60;slots&#x60;, &#x60;software&#x60;, &#x60;softwareUpdate&#x60;, &#x60;sound&#x60;, &#x60;storage&#x60;, &#x60;videos&#x60; and &#x60;virtualMachines&#x60; | [default to &quot;default&quot;]
 **query** | [**map[string]interface{}**](map[string]interface{}.md) | The criterion you want to find for your nodes. Replaces the &#x60;where&#x60;, &#x60;composition&#x60; and &#x60;select&#x60; parameters in a single parameter. | 
 **where** | **[]map[string]interface{}** | The criterion you want to find for your nodes | 
 **composition** | **string** | Boolean operator to use between each  &#x60;where&#x60; criteria. | [default to &quot;and&quot;]
 **select_** | **string** | What kind of data we want to include. Here we can get policy servers/relay by setting &#x60;nodeAndPolicyServer&#x60;. Only used if &#x60;where&#x60; is defined. | [default to &quot;node&quot;]

### Return type

[**ListAcceptedNodes200Response**](ListAcceptedNodes200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPendingNodes

> ListPendingNodes200Response ListPendingNodes(ctx).Include(include).Query(query).Where(where).Composition(composition).Select_(select_).Execute()

List pending nodes



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
	include := "minimal" // string | Level of information to include from the node inventory. Some base levels are defined (**minimal**, **default**, **full**). You can add fields you want to a base level by adding them to the list, possible values are keys from json answer. If you don't provide a base level, they will be added to `default` level, so if you only want os details, use `minimal,os` as the value for this parameter. * **minimal** includes: `id`, `hostname` and `status` * **default** includes **minimal** plus `architectureDescription`, `description`, `ipAddresses`, `lastRunDate`, `lastInventoryDate`, `machine`, `os`, `managementTechnology`, `policyServerId`, `properties` (be careful! Only node own properties, if you also need inherited properties, look at the dedicated `/nodes/{id}/inheritedProperties` endpoint), `policyMode `, `ram` and `timezone` * **full** includes: **default** plus `accounts`, `bios`, `controllers`, `environmentVariables`, `fileSystems`, `managementTechnologyDetails`, `memories`, `networkInterfaces`, `ports`, `processes`, `processors`, `slots`, `software`, `softwareUpdate`, `sound`, `storage`, `videos` and `virtualMachines` (optional) (default to "default")
	query := TODO // map[string]interface{} | The criterion you want to find for your nodes. Replaces the `where`, `composition` and `select` parameters in a single parameter. (optional)
	where := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | The criterion you want to find for your nodes (optional)
	composition := "and" // string | Boolean operator to use between each  `where` criteria. (optional) (default to "and")
	select_ := "select__example" // string | What kind of data we want to include. Here we can get policy servers/relay by setting `nodeAndPolicyServer`. Only used if `where` is defined. (optional) (default to "node")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodesAPI.ListPendingNodes(context.Background()).Include(include).Query(query).Where(where).Composition(composition).Select_(select_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.ListPendingNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPendingNodes`: ListPendingNodes200Response
	fmt.Fprintf(os.Stdout, "Response from `NodesAPI.ListPendingNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPendingNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **string** | Level of information to include from the node inventory. Some base levels are defined (**minimal**, **default**, **full**). You can add fields you want to a base level by adding them to the list, possible values are keys from json answer. If you don&#39;t provide a base level, they will be added to &#x60;default&#x60; level, so if you only want os details, use &#x60;minimal,os&#x60; as the value for this parameter. * **minimal** includes: &#x60;id&#x60;, &#x60;hostname&#x60; and &#x60;status&#x60; * **default** includes **minimal** plus &#x60;architectureDescription&#x60;, &#x60;description&#x60;, &#x60;ipAddresses&#x60;, &#x60;lastRunDate&#x60;, &#x60;lastInventoryDate&#x60;, &#x60;machine&#x60;, &#x60;os&#x60;, &#x60;managementTechnology&#x60;, &#x60;policyServerId&#x60;, &#x60;properties&#x60; (be careful! Only node own properties, if you also need inherited properties, look at the dedicated &#x60;/nodes/{id}/inheritedProperties&#x60; endpoint), &#x60;policyMode &#x60;, &#x60;ram&#x60; and &#x60;timezone&#x60; * **full** includes: **default** plus &#x60;accounts&#x60;, &#x60;bios&#x60;, &#x60;controllers&#x60;, &#x60;environmentVariables&#x60;, &#x60;fileSystems&#x60;, &#x60;managementTechnologyDetails&#x60;, &#x60;memories&#x60;, &#x60;networkInterfaces&#x60;, &#x60;ports&#x60;, &#x60;processes&#x60;, &#x60;processors&#x60;, &#x60;slots&#x60;, &#x60;software&#x60;, &#x60;softwareUpdate&#x60;, &#x60;sound&#x60;, &#x60;storage&#x60;, &#x60;videos&#x60; and &#x60;virtualMachines&#x60; | [default to &quot;default&quot;]
 **query** | [**map[string]interface{}**](map[string]interface{}.md) | The criterion you want to find for your nodes. Replaces the &#x60;where&#x60;, &#x60;composition&#x60; and &#x60;select&#x60; parameters in a single parameter. | 
 **where** | **[]map[string]interface{}** | The criterion you want to find for your nodes | 
 **composition** | **string** | Boolean operator to use between each  &#x60;where&#x60; criteria. | [default to &quot;and&quot;]
 **select_** | **string** | What kind of data we want to include. Here we can get policy servers/relay by setting &#x60;nodeAndPolicyServer&#x60;. Only used if &#x60;where&#x60; is defined. | [default to &quot;node&quot;]

### Return type

[**ListPendingNodes200Response**](ListPendingNodes200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NodeDetails

> NodeDetails200Response NodeDetails(ctx, nodeId).Include(include).Execute()

Get information about a node



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
	include := "minimal" // string | Level of information to include from the node inventory. Some base levels are defined (**minimal**, **default**, **full**). You can add fields you want to a base level by adding them to the list, possible values are keys from json answer. If you don't provide a base level, they will be added to `default` level, so if you only want os details, use `minimal,os` as the value for this parameter. * **minimal** includes: `id`, `hostname` and `status` * **default** includes **minimal** plus `architectureDescription`, `description`, `ipAddresses`, `lastRunDate`, `lastInventoryDate`, `machine`, `os`, `managementTechnology`, `policyServerId`, `properties` (be careful! Only node own properties, if you also need inherited properties, look at the dedicated `/nodes/{id}/inheritedProperties` endpoint), `policyMode `, `ram` and `timezone` * **full** includes: **default** plus `accounts`, `bios`, `controllers`, `environmentVariables`, `fileSystems`, `managementTechnologyDetails`, `memories`, `networkInterfaces`, `ports`, `processes`, `processors`, `slots`, `software`, `softwareUpdate`, `sound`, `storage`, `videos` and `virtualMachines` (optional) (default to "default")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodesAPI.NodeDetails(context.Background(), nodeId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.NodeDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NodeDetails`: NodeDetails200Response
	fmt.Fprintf(os.Stdout, "Response from `NodesAPI.NodeDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | Id of the target node | 

### Other Parameters

Other parameters are passed through a pointer to a apiNodeDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **string** | Level of information to include from the node inventory. Some base levels are defined (**minimal**, **default**, **full**). You can add fields you want to a base level by adding them to the list, possible values are keys from json answer. If you don&#39;t provide a base level, they will be added to &#x60;default&#x60; level, so if you only want os details, use &#x60;minimal,os&#x60; as the value for this parameter. * **minimal** includes: &#x60;id&#x60;, &#x60;hostname&#x60; and &#x60;status&#x60; * **default** includes **minimal** plus &#x60;architectureDescription&#x60;, &#x60;description&#x60;, &#x60;ipAddresses&#x60;, &#x60;lastRunDate&#x60;, &#x60;lastInventoryDate&#x60;, &#x60;machine&#x60;, &#x60;os&#x60;, &#x60;managementTechnology&#x60;, &#x60;policyServerId&#x60;, &#x60;properties&#x60; (be careful! Only node own properties, if you also need inherited properties, look at the dedicated &#x60;/nodes/{id}/inheritedProperties&#x60; endpoint), &#x60;policyMode &#x60;, &#x60;ram&#x60; and &#x60;timezone&#x60; * **full** includes: **default** plus &#x60;accounts&#x60;, &#x60;bios&#x60;, &#x60;controllers&#x60;, &#x60;environmentVariables&#x60;, &#x60;fileSystems&#x60;, &#x60;managementTechnologyDetails&#x60;, &#x60;memories&#x60;, &#x60;networkInterfaces&#x60;, &#x60;ports&#x60;, &#x60;processes&#x60;, &#x60;processors&#x60;, &#x60;slots&#x60;, &#x60;software&#x60;, &#x60;softwareUpdate&#x60;, &#x60;sound&#x60;, &#x60;storage&#x60;, &#x60;videos&#x60; and &#x60;virtualMachines&#x60; | [default to &quot;default&quot;]

### Return type

[**NodeDetails200Response**](NodeDetails200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NodeInheritedProperties

> NodeInheritedProperties200Response NodeInheritedProperties(ctx, nodeId).Execute()

Get inherited node properties for a node



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
	resp, r, err := apiClient.NodesAPI.NodeInheritedProperties(context.Background(), nodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.NodeInheritedProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NodeInheritedProperties`: NodeInheritedProperties200Response
	fmt.Fprintf(os.Stdout, "Response from `NodesAPI.NodeInheritedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | Id of the target node | 

### Other Parameters

Other parameters are passed through a pointer to a apiNodeInheritedPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NodeInheritedProperties200Response**](NodeInheritedProperties200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNode

> UpdateNode200Response UpdateNode(ctx, nodeId).NodeSettings(nodeSettings).Execute()

Update node settings and properties



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
	nodeSettings := *openapiclient.NewNodeSettings() // NodeSettings |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodesAPI.UpdateNode(context.Background(), nodeId).NodeSettings(nodeSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.UpdateNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNode`: UpdateNode200Response
	fmt.Fprintf(os.Stdout, "Response from `NodesAPI.UpdateNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | Id of the target node | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodeSettings** | [**NodeSettings**](NodeSettings.md) |  | 

### Return type

[**UpdateNode200Response**](UpdateNode200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

