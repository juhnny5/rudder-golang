# \RulesAPI

All URIs are relative to *https://rudder.example.local/rudder/api/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRule**](RulesAPI.md#CreateRule) | **Put** /rules | Create a rule
[**CreateRuleCategory**](RulesAPI.md#CreateRuleCategory) | **Put** /rules/categories | Create a rule category
[**DeleteRule**](RulesAPI.md#DeleteRule) | **Delete** /rules/{ruleId} | Delete a rule
[**DeleteRuleCategory**](RulesAPI.md#DeleteRuleCategory) | **Delete** /rules/categories/{ruleCategoryId} | Delete group category
[**GetRuleCategoryDetails**](RulesAPI.md#GetRuleCategoryDetails) | **Get** /rules/categories/{ruleCategoryId} | Get rule category details
[**GetRuleTree**](RulesAPI.md#GetRuleTree) | **Get** /rules/tree | Get rules tree
[**ListRules**](RulesAPI.md#ListRules) | **Get** /rules | List all rules
[**RuleDetails**](RulesAPI.md#RuleDetails) | **Get** /rules/{ruleId} | Get a rule details
[**UpdateRule**](RulesAPI.md#UpdateRule) | **Post** /rules/{ruleId} | Update a rule details
[**UpdateRuleCategory**](RulesAPI.md#UpdateRuleCategory) | **Post** /rules/categories/{ruleCategoryId} | Update rule category details



## CreateRule

> CreateRule200Response CreateRule(ctx).RuleNew(ruleNew).Execute()

Create a rule



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
	ruleNew := *openapiclient.NewRuleNew() // RuleNew |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.CreateRule(context.Background()).RuleNew(ruleNew).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.CreateRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRule`: CreateRule200Response
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.CreateRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ruleNew** | [**RuleNew**](RuleNew.md) |  | 

### Return type

[**CreateRule200Response**](CreateRule200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRuleCategory

> CreateRuleCategory200Response CreateRuleCategory(ctx).RuleCategory(ruleCategory).Execute()

Create a rule category



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
	ruleCategory := *openapiclient.NewRuleCategory("b9f6d98a-28bc-4d80-90f7-d2f14269e215", "Security policies") // RuleCategory | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.CreateRuleCategory(context.Background()).RuleCategory(ruleCategory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.CreateRuleCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRuleCategory`: CreateRuleCategory200Response
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.CreateRuleCategory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRuleCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ruleCategory** | [**RuleCategory**](RuleCategory.md) |  | 

### Return type

[**CreateRuleCategory200Response**](CreateRuleCategory200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRule

> DeleteRule200Response DeleteRule(ctx, ruleId).Execute()

Delete a rule



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.DeleteRule(context.Background(), ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.DeleteRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRule`: DeleteRule200Response
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.DeleteRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** | Id of the target rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteRule200Response**](DeleteRule200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRuleCategory

> DeleteRuleCategory200Response DeleteRuleCategory(ctx, ruleCategoryId).Execute()

Delete group category



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
	ruleCategoryId := "e0a311fa-f7b2-4f9e-89a9-db517b9c6b90" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.DeleteRuleCategory(context.Background(), ruleCategoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.DeleteRuleCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRuleCategory`: DeleteRuleCategory200Response
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.DeleteRuleCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleCategoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRuleCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteRuleCategory200Response**](DeleteRuleCategory200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleCategoryDetails

> GetRuleCategoryDetails200Response GetRuleCategoryDetails(ctx, ruleCategoryId).Execute()

Get rule category details



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
	ruleCategoryId := "e0a311fa-f7b2-4f9e-89a9-db517b9c6b90" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.GetRuleCategoryDetails(context.Background(), ruleCategoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.GetRuleCategoryDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRuleCategoryDetails`: GetRuleCategoryDetails200Response
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.GetRuleCategoryDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleCategoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleCategoryDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRuleCategoryDetails200Response**](GetRuleCategoryDetails200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleTree

> GetRuleTree200Response GetRuleTree(ctx).Execute()

Get rules tree



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
	resp, r, err := apiClient.RulesAPI.GetRuleTree(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.GetRuleTree``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRuleTree`: GetRuleTree200Response
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.GetRuleTree`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleTreeRequest struct via the builder pattern


### Return type

[**GetRuleTree200Response**](GetRuleTree200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRules

> ListRules200Response ListRules(ctx).Execute()

List all rules



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
	resp, r, err := apiClient.RulesAPI.ListRules(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.ListRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRules`: ListRules200Response
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.ListRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRulesRequest struct via the builder pattern


### Return type

[**ListRules200Response**](ListRules200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RuleDetails

> RuleDetails200Response RuleDetails(ctx, ruleId).Execute()

Get a rule details



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RuleDetails(context.Background(), ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RuleDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RuleDetails`: RuleDetails200Response
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RuleDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** | Id of the target rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiRuleDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RuleDetails200Response**](RuleDetails200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRule

> UpdateRule200Response UpdateRule(ctx, ruleId).RuleWithCategory(ruleWithCategory).Execute()

Update a rule details



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
	ruleWithCategory := *openapiclient.NewRuleWithCategory() // RuleWithCategory | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.UpdateRule(context.Background(), ruleId).RuleWithCategory(ruleWithCategory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.UpdateRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRule`: UpdateRule200Response
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.UpdateRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** | Id of the target rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ruleWithCategory** | [**RuleWithCategory**](RuleWithCategory.md) |  | 

### Return type

[**UpdateRule200Response**](UpdateRule200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRuleCategory

> UpdateRuleCategory200Response UpdateRuleCategory(ctx, ruleCategoryId).RuleCategoryUpdate(ruleCategoryUpdate).Execute()

Update rule category details



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
	ruleCategoryId := "e0a311fa-f7b2-4f9e-89a9-db517b9c6b90" // string | 
	ruleCategoryUpdate := *openapiclient.NewRuleCategoryUpdate("b9f6d98a-28bc-4d80-90f7-d2f14269e215", "Security policies") // RuleCategoryUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.UpdateRuleCategory(context.Background(), ruleCategoryId).RuleCategoryUpdate(ruleCategoryUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.UpdateRuleCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRuleCategory`: UpdateRuleCategory200Response
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.UpdateRuleCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleCategoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRuleCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ruleCategoryUpdate** | [**RuleCategoryUpdate**](RuleCategoryUpdate.md) |  | 

### Return type

[**UpdateRuleCategory200Response**](UpdateRuleCategory200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

