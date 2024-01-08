# \InventoriesAPI

All URIs are relative to *https://rudder.example.local/rudder/api/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FileWatcherRestart**](InventoriesAPI.md#FileWatcherRestart) | **Post** /inventories/watcher/restart | Restart inventory watcher
[**FileWatcherStart**](InventoriesAPI.md#FileWatcherStart) | **Post** /inventories/watcher/start | Start inventory watcher
[**FileWatcherStop**](InventoriesAPI.md#FileWatcherStop) | **Post** /inventories/watcher/stop | Stop inventory watcher
[**QueueInformation**](InventoriesAPI.md#QueueInformation) | **Get** /inventories/info | Get information about inventory processing queue
[**UploadInventory**](InventoriesAPI.md#UploadInventory) | **Post** /inventories/upload | Upload an inventory for processing



## FileWatcherRestart

> FileWatcherRestart200Response FileWatcherRestart(ctx).Execute()

Restart inventory watcher



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
	resp, r, err := apiClient.InventoriesAPI.FileWatcherRestart(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoriesAPI.FileWatcherRestart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FileWatcherRestart`: FileWatcherRestart200Response
	fmt.Fprintf(os.Stdout, "Response from `InventoriesAPI.FileWatcherRestart`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFileWatcherRestartRequest struct via the builder pattern


### Return type

[**FileWatcherRestart200Response**](FileWatcherRestart200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FileWatcherStart

> FileWatcherStart200Response FileWatcherStart(ctx).Execute()

Start inventory watcher



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
	resp, r, err := apiClient.InventoriesAPI.FileWatcherStart(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoriesAPI.FileWatcherStart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FileWatcherStart`: FileWatcherStart200Response
	fmt.Fprintf(os.Stdout, "Response from `InventoriesAPI.FileWatcherStart`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFileWatcherStartRequest struct via the builder pattern


### Return type

[**FileWatcherStart200Response**](FileWatcherStart200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FileWatcherStop

> FileWatcherStop200Response FileWatcherStop(ctx).Execute()

Stop inventory watcher



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
	resp, r, err := apiClient.InventoriesAPI.FileWatcherStop(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoriesAPI.FileWatcherStop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FileWatcherStop`: FileWatcherStop200Response
	fmt.Fprintf(os.Stdout, "Response from `InventoriesAPI.FileWatcherStop`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFileWatcherStopRequest struct via the builder pattern


### Return type

[**FileWatcherStop200Response**](FileWatcherStop200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueueInformation

> QueueInformation200Response QueueInformation(ctx).Execute()

Get information about inventory processing queue



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
	resp, r, err := apiClient.InventoriesAPI.QueueInformation(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoriesAPI.QueueInformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueueInformation`: QueueInformation200Response
	fmt.Fprintf(os.Stdout, "Response from `InventoriesAPI.QueueInformation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiQueueInformationRequest struct via the builder pattern


### Return type

[**QueueInformation200Response**](QueueInformation200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadInventory

> UploadInventory200Response UploadInventory(ctx).File(file).Signature(signature).Execute()

Upload an inventory for processing



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
	file := os.NewFile(1234, "some_file") // *os.File | The inventory file. The original file name is used to check extension, that should be .xml[.gz] or .ocs[.gz] (optional)
	signature := os.NewFile(1234, "some_file") // *os.File | The signature file. The original file name is used to check extension, that should be ${originalInventoryFileName}.sign[.gz] (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoriesAPI.UploadInventory(context.Background()).File(file).Signature(signature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoriesAPI.UploadInventory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadInventory`: UploadInventory200Response
	fmt.Fprintf(os.Stdout, "Response from `InventoriesAPI.UploadInventory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | The inventory file. The original file name is used to check extension, that should be .xml[.gz] or .ocs[.gz] | 
 **signature** | ***os.File** | The signature file. The original file name is used to check extension, that should be ${originalInventoryFileName}.sign[.gz] | 

### Return type

[**UploadInventory200Response**](UploadInventory200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

