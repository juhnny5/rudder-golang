# \CampaignsAPI

All URIs are relative to *https://rudder.example.local/rudder/api/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllCampaigns**](CampaignsAPI.md#AllCampaigns) | **Get** /campaigns | Get all campaigns details
[**DeleteCampaign**](CampaignsAPI.md#DeleteCampaign) | **Delete** /campaigns/{id} | Delete a campaign
[**DeleteCampaignEvent**](CampaignsAPI.md#DeleteCampaignEvent) | **Delete** /campaigns/events/{id} | Delete a campaign event details
[**GetAllCampaignEvents**](CampaignsAPI.md#GetAllCampaignEvents) | **Get** /campaigns/events | Get all campaign events
[**GetCampaign**](CampaignsAPI.md#GetCampaign) | **Get** /campaigns/{id} | Get a campaign details
[**GetCampaignEvent**](CampaignsAPI.md#GetCampaignEvent) | **Get** /campaigns/events/{id} | Get a campaign event details
[**GetEventsCampaign**](CampaignsAPI.md#GetEventsCampaign) | **Get** /campaigns/{id}/events | Get campaign events for a campaign
[**SaveCampaign**](CampaignsAPI.md#SaveCampaign) | **Post** /campaigns | Save a campaign
[**SaveCampaignEvent**](CampaignsAPI.md#SaveCampaignEvent) | **Post** /campaigns/events/{id} | Update an existing event
[**ScheduleCampaign**](CampaignsAPI.md#ScheduleCampaign) | **Post** /campaigns/{id}/schedule | Schedule a campaign event for a campaign



## AllCampaigns

> AllCampaigns200Response AllCampaigns(ctx).CampaignType(campaignType).Status(status).Execute()

Get all campaigns details



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
	campaignType := "system-update" // string | Type of the campaigns we want (optional)
	status := "enabled" // string | Status of the campaigns we want (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.AllCampaigns(context.Background()).CampaignType(campaignType).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.AllCampaigns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AllCampaigns`: AllCampaigns200Response
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.AllCampaigns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAllCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignType** | **string** | Type of the campaigns we want | 
 **status** | **string** | Status of the campaigns we want | 

### Return type

[**AllCampaigns200Response**](AllCampaigns200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCampaign

> DeleteCampaign200Response DeleteCampaign(ctx, id).Execute()

Delete a campaign



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
	id := "0076a379-f32d-4732-9e91-33ab219d8fde" // string | Id of the campaign

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.DeleteCampaign(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.DeleteCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCampaign`: DeleteCampaign200Response
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.DeleteCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the campaign | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteCampaign200Response**](DeleteCampaign200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCampaignEvent

> DeleteCampaignEvent200Response DeleteCampaignEvent(ctx, id).Execute()

Delete a campaign event details



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
	resp, r, err := apiClient.CampaignsAPI.DeleteCampaignEvent(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.DeleteCampaignEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCampaignEvent`: DeleteCampaignEvent200Response
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.DeleteCampaignEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the campaign event | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCampaignEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteCampaignEvent200Response**](DeleteCampaignEvent200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllCampaignEvents

> GetAllCampaignEvents200Response GetAllCampaignEvents(ctx).CampaignType(campaignType).State(state).CampaignId(campaignId).Limit(limit).Offset(offset).Before(before).After(after).Order(order).Asc(asc).Execute()

Get all campaign events



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
	campaignType := "system-update" // string | Type of the campaigns we want (optional)
	state := "enabled" // string | Status of the campaign events we want (optional)
	campaignId := "system-update" // string | id of the campaigns we want (optional)
	limit := int32(56) // int32 | Max number of elements in response (optional)
	offset := int32(56) // int32 | Offset of data in response (skip X elements) (optional)
	before := time.Now() // string |  (optional)
	after := time.Now() // string |  (optional)
	order := "order_example" // string |  (optional)
	asc := "asc_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.GetAllCampaignEvents(context.Background()).CampaignType(campaignType).State(state).CampaignId(campaignId).Limit(limit).Offset(offset).Before(before).After(after).Order(order).Asc(asc).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.GetAllCampaignEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllCampaignEvents`: GetAllCampaignEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.GetAllCampaignEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllCampaignEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignType** | **string** | Type of the campaigns we want | 
 **state** | **string** | Status of the campaign events we want | 
 **campaignId** | **string** | id of the campaigns we want | 
 **limit** | **int32** | Max number of elements in response | 
 **offset** | **int32** | Offset of data in response (skip X elements) | 
 **before** | **string** |  | 
 **after** | **string** |  | 
 **order** | **string** |  | 
 **asc** | **string** |  | 

### Return type

[**GetAllCampaignEvents200Response**](GetAllCampaignEvents200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaign

> GetCampaign200Response GetCampaign(ctx, id).Execute()

Get a campaign details



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
	id := "0076a379-f32d-4732-9e91-33ab219d8fde" // string | Id of the campaign

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.GetCampaign(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.GetCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaign`: GetCampaign200Response
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.GetCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the campaign | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCampaign200Response**](GetCampaign200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignEvent

> GetAllCampaignEvents200Response GetCampaignEvent(ctx, id).Execute()

Get a campaign event details



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
	resp, r, err := apiClient.CampaignsAPI.GetCampaignEvent(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.GetCampaignEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignEvent`: GetAllCampaignEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.GetCampaignEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the campaign event | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAllCampaignEvents200Response**](GetAllCampaignEvents200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventsCampaign

> GetEventsCampaign200Response GetEventsCampaign(ctx, id).CampaignType(campaignType).State(state).Limit(limit).Offset(offset).Before(before).After(after).Order(order).Asc(asc).Execute()

Get campaign events for a campaign



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
	campaignType := "system-update" // string | Type of the campaigns we want (optional)
	state := "enabled" // string | Status of the campaign events we want (optional)
	limit := int32(56) // int32 | Max number of elements in response (optional)
	offset := int32(56) // int32 | Offset of data in response (skip X elements) (optional)
	before := time.Now() // string |  (optional)
	after := time.Now() // string |  (optional)
	order := "order_example" // string |  (optional)
	asc := "asc_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.GetEventsCampaign(context.Background(), id).CampaignType(campaignType).State(state).Limit(limit).Offset(offset).Before(before).After(after).Order(order).Asc(asc).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.GetEventsCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventsCampaign`: GetEventsCampaign200Response
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.GetEventsCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the campaign | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **campaignType** | **string** | Type of the campaigns we want | 
 **state** | **string** | Status of the campaign events we want | 
 **limit** | **int32** | Max number of elements in response | 
 **offset** | **int32** | Offset of data in response (skip X elements) | 
 **before** | **string** |  | 
 **after** | **string** |  | 
 **order** | **string** |  | 
 **asc** | **string** |  | 

### Return type

[**GetEventsCampaign200Response**](GetEventsCampaign200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveCampaign

> SaveCampaign200Response SaveCampaign(ctx).CampaignDetails(campaignDetails).Execute()

Save a campaign



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
	campaignDetails := *openapiclient.NewCampaignDetails() // CampaignDetails | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.SaveCampaign(context.Background()).CampaignDetails(campaignDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.SaveCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveCampaign`: SaveCampaign200Response
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.SaveCampaign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignDetails** | [**CampaignDetails**](CampaignDetails.md) |  | 

### Return type

[**SaveCampaign200Response**](SaveCampaign200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveCampaignEvent

> SaveCampaignEvent200Response SaveCampaignEvent(ctx, id).Execute()

Update an existing event



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
	resp, r, err := apiClient.CampaignsAPI.SaveCampaignEvent(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.SaveCampaignEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveCampaignEvent`: SaveCampaignEvent200Response
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.SaveCampaignEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the campaign event | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveCampaignEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SaveCampaignEvent200Response**](SaveCampaignEvent200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduleCampaign

> ScheduleCampaign200Response ScheduleCampaign(ctx, id).Execute()

Schedule a campaign event for a campaign



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
	id := "0076a379-f32d-4732-9e91-33ab219d8fde" // string | Id of the campaign

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.ScheduleCampaign(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.ScheduleCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScheduleCampaign`: ScheduleCampaign200Response
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.ScheduleCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the campaign | 

### Other Parameters

Other parameters are passed through a pointer to a apiScheduleCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScheduleCampaign200Response**](ScheduleCampaign200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

