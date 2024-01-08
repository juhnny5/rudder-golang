# \ComplianceAPI

All URIs are relative to *https://rudder.example.local/rudder/api/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDirectiveComplianceId**](ComplianceAPI.md#GetDirectiveComplianceId) | **Get** /compliance/directives/{directiveId} | Compliance details by directive
[**GetDirectivesCompliance**](ComplianceAPI.md#GetDirectivesCompliance) | **Get** /compliance/directives | Compliance details for all directives
[**GetGlobalCompliance**](ComplianceAPI.md#GetGlobalCompliance) | **Get** /compliance | Global compliance
[**GetNodeCompliance**](ComplianceAPI.md#GetNodeCompliance) | **Get** /compliance/nodes/{nodeId} | Compliance details by node
[**GetNodesCompliance**](ComplianceAPI.md#GetNodesCompliance) | **Get** /compliance/nodes | Compliance details for all nodes
[**GetRuleCompliance**](ComplianceAPI.md#GetRuleCompliance) | **Get** /compliance/rules/{ruleId} | Compliance details by rule
[**GetRulesCompliance**](ComplianceAPI.md#GetRulesCompliance) | **Get** /compliance/rules | Compliance details for all rules



## GetDirectiveComplianceId

> GetDirectiveComplianceId200Response GetDirectiveComplianceId(ctx, directiveId).Format(format).Execute()

Compliance details by directive



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
	format := "["csv"]" // string | format of export (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.GetDirectiveComplianceId(context.Background(), directiveId).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.GetDirectiveComplianceId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDirectiveComplianceId`: GetDirectiveComplianceId200Response
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.GetDirectiveComplianceId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directiveId** | **string** | Id of the directive | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDirectiveComplianceIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **string** | format of export | 

### Return type

[**GetDirectiveComplianceId200Response**](GetDirectiveComplianceId200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDirectivesCompliance

> GetDirectivesCompliance200Response GetDirectivesCompliance(ctx).Execute()

Compliance details for all directives



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
	resp, r, err := apiClient.ComplianceAPI.GetDirectivesCompliance(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.GetDirectivesCompliance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDirectivesCompliance`: GetDirectivesCompliance200Response
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.GetDirectivesCompliance`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDirectivesComplianceRequest struct via the builder pattern


### Return type

[**GetDirectivesCompliance200Response**](GetDirectivesCompliance200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalCompliance

> GetGlobalCompliance200Response GetGlobalCompliance(ctx).Precision(precision).Execute()

Global compliance



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
	precision := int32(0) // int32 | Number of digits after comma in compliance percent figures (optional) (default to 2)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.GetGlobalCompliance(context.Background()).Precision(precision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.GetGlobalCompliance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGlobalCompliance`: GetGlobalCompliance200Response
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.GetGlobalCompliance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalComplianceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **precision** | **int32** | Number of digits after comma in compliance percent figures | [default to 2]

### Return type

[**GetGlobalCompliance200Response**](GetGlobalCompliance200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeCompliance

> GetNodeCompliance200Response GetNodeCompliance(ctx, nodeId).Level(level).Precision(precision).Execute()

Compliance details by node



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
	level := int32(4) // int32 | Number of depth level of compliance objects to display (1:rules, 2:directives, 3:components, 4:nodes, 5:values, 6:reports) (optional) (default to 10)
	precision := int32(0) // int32 | Number of digits after comma in compliance percent figures (optional) (default to 2)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.GetNodeCompliance(context.Background(), nodeId).Level(level).Precision(precision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.GetNodeCompliance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNodeCompliance`: GetNodeCompliance200Response
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.GetNodeCompliance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | Id of the target node | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeComplianceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **level** | **int32** | Number of depth level of compliance objects to display (1:rules, 2:directives, 3:components, 4:nodes, 5:values, 6:reports) | [default to 10]
 **precision** | **int32** | Number of digits after comma in compliance percent figures | [default to 2]

### Return type

[**GetNodeCompliance200Response**](GetNodeCompliance200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodesCompliance

> GetNodesCompliance200Response GetNodesCompliance(ctx).Level(level).Precision(precision).Execute()

Compliance details for all nodes



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
	level := int32(4) // int32 | Number of depth level of compliance objects to display (1:rules, 2:directives, 3:components, 4:nodes, 5:values, 6:reports) (optional) (default to 10)
	precision := int32(0) // int32 | Number of digits after comma in compliance percent figures (optional) (default to 2)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.GetNodesCompliance(context.Background()).Level(level).Precision(precision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.GetNodesCompliance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNodesCompliance`: GetNodesCompliance200Response
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.GetNodesCompliance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNodesComplianceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **level** | **int32** | Number of depth level of compliance objects to display (1:rules, 2:directives, 3:components, 4:nodes, 5:values, 6:reports) | [default to 10]
 **precision** | **int32** | Number of digits after comma in compliance percent figures | [default to 2]

### Return type

[**GetNodesCompliance200Response**](GetNodesCompliance200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleCompliance

> GetRuleCompliance200Response GetRuleCompliance(ctx, ruleId).Level(level).Precision(precision).Execute()

Compliance details by rule



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
	ruleId := "9a1773c9-0889-40b6-be89-f6504443ac1b" // string | Id of the target rule
	level := int32(4) // int32 | Number of depth level of compliance objects to display (1:rules, 2:directives, 3:components, 4:nodes, 5:values, 6:reports) (optional) (default to 10)
	precision := int32(0) // int32 | Number of digits after comma in compliance percent figures (optional) (default to 2)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.GetRuleCompliance(context.Background(), ruleId).Level(level).Precision(precision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.GetRuleCompliance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRuleCompliance`: GetRuleCompliance200Response
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.GetRuleCompliance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** | Id of the target rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleComplianceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **level** | **int32** | Number of depth level of compliance objects to display (1:rules, 2:directives, 3:components, 4:nodes, 5:values, 6:reports) | [default to 10]
 **precision** | **int32** | Number of digits after comma in compliance percent figures | [default to 2]

### Return type

[**GetRuleCompliance200Response**](GetRuleCompliance200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRulesCompliance

> GetRulesCompliance200Response GetRulesCompliance(ctx).Level(level).Precision(precision).Execute()

Compliance details for all rules



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
	level := int32(4) // int32 | Number of depth level of compliance objects to display (1:rules, 2:directives, 3:components, 4:nodes, 5:values, 6:reports) (optional) (default to 10)
	precision := int32(0) // int32 | Number of digits after comma in compliance percent figures (optional) (default to 2)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.GetRulesCompliance(context.Background()).Level(level).Precision(precision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.GetRulesCompliance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRulesCompliance`: GetRulesCompliance200Response
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.GetRulesCompliance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRulesComplianceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **level** | **int32** | Number of depth level of compliance objects to display (1:rules, 2:directives, 3:components, 4:nodes, 5:values, 6:reports) | [default to 10]
 **precision** | **int32** | Number of digits after comma in compliance percent figures | [default to 2]

### Return type

[**GetRulesCompliance200Response**](GetRulesCompliance200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

