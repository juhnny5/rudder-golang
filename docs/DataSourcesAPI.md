# \DataSourcesAPI

All URIs are relative to *https://rudder.example.local/rudder/api/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDataSource**](DataSourcesAPI.md#CreateDataSource) | **Put** /datasources | Create a data source
[**DeleteDataSource**](DataSourcesAPI.md#DeleteDataSource) | **Delete** /datasources/{datasourceId} | Delete a data source
[**GetAllDataSources**](DataSourcesAPI.md#GetAllDataSources) | **Get** /datasources | List all data sources
[**GetDataSource**](DataSourcesAPI.md#GetDataSource) | **Get** /datasources/{datasourceId} | Get data source configuration
[**ReloadAllDatasourcesAllNodes**](DataSourcesAPI.md#ReloadAllDatasourcesAllNodes) | **Post** /datasources/reload | Update properties from data sources
[**ReloadAllDatasourcesOneNode**](DataSourcesAPI.md#ReloadAllDatasourcesOneNode) | **Post** /nodes/{nodeId}/fetchData | Update properties for one node from all data sources
[**ReloadOneDatasourceAllNodes**](DataSourcesAPI.md#ReloadOneDatasourceAllNodes) | **Post** /datasources/reload/{datasourceId} | Update properties from data sources
[**ReloadOneDatasourceOneNode**](DataSourcesAPI.md#ReloadOneDatasourceOneNode) | **Post** /nodes/{nodeId}/fetchData/{datasourceId} | Update properties for one node from a data source
[**UpdateDataSource**](DataSourcesAPI.md#UpdateDataSource) | **Post** /datasources/{datasourceId} | Update a data source configuration



## CreateDataSource

> CreateDataSource200Response CreateDataSource(ctx).Datasource(datasource).Execute()

Create a data source



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
	datasource := *openapiclient.NewDatasource() // Datasource |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.CreateDataSource(context.Background()).Datasource(datasource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.CreateDataSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDataSource`: CreateDataSource200Response
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.CreateDataSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasource** | [**Datasource**](Datasource.md) |  | 

### Return type

[**CreateDataSource200Response**](CreateDataSource200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDataSource

> DeleteDataSource200Response DeleteDataSource(ctx, datasourceId).Execute()

Delete a data source



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
	datasourceId := "test-data-source" // string | Id of the data source

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.DeleteDataSource(context.Background(), datasourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.DeleteDataSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDataSource`: DeleteDataSource200Response
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.DeleteDataSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datasourceId** | **string** | Id of the data source | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDataSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteDataSource200Response**](DeleteDataSource200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllDataSources

> GetAllDataSources200Response GetAllDataSources(ctx).Execute()

List all data sources



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
	resp, r, err := apiClient.DataSourcesAPI.GetAllDataSources(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.GetAllDataSources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllDataSources`: GetAllDataSources200Response
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.GetAllDataSources`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllDataSourcesRequest struct via the builder pattern


### Return type

[**GetAllDataSources200Response**](GetAllDataSources200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataSource

> GetDataSource200Response GetDataSource(ctx, datasourceId).Execute()

Get data source configuration



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
	datasourceId := "test-data-source" // string | Id of the data source

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.GetDataSource(context.Background(), datasourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.GetDataSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataSource`: GetDataSource200Response
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.GetDataSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datasourceId** | **string** | Id of the data source | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDataSource200Response**](GetDataSource200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReloadAllDatasourcesAllNodes

> ReloadAllDatasourcesAllNodes200Response ReloadAllDatasourcesAllNodes(ctx).Execute()

Update properties from data sources



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
	resp, r, err := apiClient.DataSourcesAPI.ReloadAllDatasourcesAllNodes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.ReloadAllDatasourcesAllNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReloadAllDatasourcesAllNodes`: ReloadAllDatasourcesAllNodes200Response
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.ReloadAllDatasourcesAllNodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReloadAllDatasourcesAllNodesRequest struct via the builder pattern


### Return type

[**ReloadAllDatasourcesAllNodes200Response**](ReloadAllDatasourcesAllNodes200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReloadAllDatasourcesOneNode

> ReloadAllDatasourcesOneNode200Response ReloadAllDatasourcesOneNode(ctx, nodeId).Execute()

Update properties for one node from all data sources



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
	resp, r, err := apiClient.DataSourcesAPI.ReloadAllDatasourcesOneNode(context.Background(), nodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.ReloadAllDatasourcesOneNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReloadAllDatasourcesOneNode`: ReloadAllDatasourcesOneNode200Response
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.ReloadAllDatasourcesOneNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | Id of the target node | 

### Other Parameters

Other parameters are passed through a pointer to a apiReloadAllDatasourcesOneNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReloadAllDatasourcesOneNode200Response**](ReloadAllDatasourcesOneNode200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReloadOneDatasourceAllNodes

> ReloadOneDatasourceAllNodes200Response ReloadOneDatasourceAllNodes(ctx, datasourceId).Execute()

Update properties from data sources



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
	datasourceId := "test-data-source" // string | Id of the data source

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.ReloadOneDatasourceAllNodes(context.Background(), datasourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.ReloadOneDatasourceAllNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReloadOneDatasourceAllNodes`: ReloadOneDatasourceAllNodes200Response
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.ReloadOneDatasourceAllNodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datasourceId** | **string** | Id of the data source | 

### Other Parameters

Other parameters are passed through a pointer to a apiReloadOneDatasourceAllNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReloadOneDatasourceAllNodes200Response**](ReloadOneDatasourceAllNodes200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReloadOneDatasourceOneNode

> ReloadOneDatasourceOneNode200Response ReloadOneDatasourceOneNode(ctx, nodeId, datasourceId).Execute()

Update properties for one node from a data source



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
	datasourceId := "test-data-source" // string | Id of the data source

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.ReloadOneDatasourceOneNode(context.Background(), nodeId, datasourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.ReloadOneDatasourceOneNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReloadOneDatasourceOneNode`: ReloadOneDatasourceOneNode200Response
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.ReloadOneDatasourceOneNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | Id of the target node | 
**datasourceId** | **string** | Id of the data source | 

### Other Parameters

Other parameters are passed through a pointer to a apiReloadOneDatasourceOneNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ReloadOneDatasourceOneNode200Response**](ReloadOneDatasourceOneNode200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDataSource

> UpdateDataSource200Response UpdateDataSource(ctx, datasourceId).Datasource(datasource).Execute()

Update a data source configuration



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
	datasourceId := "test-data-source" // string | Id of the data source
	datasource := *openapiclient.NewDatasource() // Datasource |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataSourcesAPI.UpdateDataSource(context.Background(), datasourceId).Datasource(datasource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesAPI.UpdateDataSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDataSource`: UpdateDataSource200Response
	fmt.Fprintf(os.Stdout, "Response from `DataSourcesAPI.UpdateDataSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datasourceId** | **string** | Id of the data source | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDataSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | [**Datasource**](Datasource.md) |  | 

### Return type

[**UpdateDataSource200Response**](UpdateDataSource200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

