# \SystemAPI

All URIs are relative to *https://rudder.example.local/rudder/api/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateArchive**](SystemAPI.md#CreateArchive) | **Post** /system/archives/{archiveKind} | Create an archive
[**GetHealthcheckResult**](SystemAPI.md#GetHealthcheckResult) | **Get** /system/healthcheck | Get healthcheck
[**GetStatus**](SystemAPI.md#GetStatus) | **Get** /system/status | Get server status
[**GetSystemInfo**](SystemAPI.md#GetSystemInfo) | **Get** /system/info | Get server information
[**GetZipArchive**](SystemAPI.md#GetZipArchive) | **Get** /system/archives/{archiveKind}/zip/{commitId} | Get an archive as a ZIP
[**ListArchives**](SystemAPI.md#ListArchives) | **Get** /system/archives/{archiveKind} | List archives
[**PurgeSoftware**](SystemAPI.md#PurgeSoftware) | **Post** /system/maintenance/purgeSoftware | Trigger batch for cleaning unreferenced software
[**RegeneratePolicies**](SystemAPI.md#RegeneratePolicies) | **Post** /system/regenerate/policies | Trigger a new policy generation
[**ReloadAll**](SystemAPI.md#ReloadAll) | **Post** /system/reload | Reload both techniques and dynamic groups
[**ReloadGroups**](SystemAPI.md#ReloadGroups) | **Post** /system/reload/groups | Reload dynamic groups
[**ReloadTechniques**](SystemAPI.md#ReloadTechniques) | **Post** /system/reload/techniques | Reload techniques
[**RestoreArchive**](SystemAPI.md#RestoreArchive) | **Post** /system/archives/{archiveKind}/restore/{archiveRestoreKind} | Restore an archive
[**UpdatePolicies**](SystemAPI.md#UpdatePolicies) | **Post** /system/update/policies | Trigger update of policies



## CreateArchive

> CreateArchive200Response CreateArchive(ctx, archiveKind).Execute()

Create an archive



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
	archiveKind := "full" // string | Type of archive to make

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.CreateArchive(context.Background(), archiveKind).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.CreateArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateArchive`: CreateArchive200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.CreateArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**archiveKind** | **string** | Type of archive to make | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateArchive200Response**](CreateArchive200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHealthcheckResult

> GetHealthcheckResult200Response GetHealthcheckResult(ctx).Execute()

Get healthcheck



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
	resp, r, err := apiClient.SystemAPI.GetHealthcheckResult(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.GetHealthcheckResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHealthcheckResult`: GetHealthcheckResult200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.GetHealthcheckResult`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHealthcheckResultRequest struct via the builder pattern


### Return type

[**GetHealthcheckResult200Response**](GetHealthcheckResult200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatus

> GetStatus200Response GetStatus(ctx).Execute()

Get server status



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
	resp, r, err := apiClient.SystemAPI.GetStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.GetStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatus`: GetStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.GetStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusRequest struct via the builder pattern


### Return type

[**GetStatus200Response**](GetStatus200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemInfo

> GetSystemInfo200Response GetSystemInfo(ctx).Execute()

Get server information



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
	resp, r, err := apiClient.SystemAPI.GetSystemInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.GetSystemInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemInfo`: GetSystemInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.GetSystemInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemInfoRequest struct via the builder pattern


### Return type

[**GetSystemInfo200Response**](GetSystemInfo200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetZipArchive

> *os.File GetZipArchive(ctx, archiveKind, commitId).Execute()

Get an archive as a ZIP



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
	archiveKind := "full" // string | Type of archive to make
	commitId := "commitId_example" // string | commit ID of the archive to get as a ZIP file

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.GetZipArchive(context.Background(), archiveKind, commitId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.GetZipArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetZipArchive`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.GetZipArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**archiveKind** | **string** | Type of archive to make | 
**commitId** | **string** | commit ID of the archive to get as a ZIP file | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZipArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListArchives

> ListArchives200Response ListArchives(ctx, archiveKind).Execute()

List archives



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
	archiveKind := "full" // string | Type of archive to make

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.ListArchives(context.Background(), archiveKind).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.ListArchives``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListArchives`: ListArchives200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.ListArchives`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**archiveKind** | **string** | Type of archive to make | 

### Other Parameters

Other parameters are passed through a pointer to a apiListArchivesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListArchives200Response**](ListArchives200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PurgeSoftware

> PurgeSoftware200Response PurgeSoftware(ctx).Execute()

Trigger batch for cleaning unreferenced software



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
	resp, r, err := apiClient.SystemAPI.PurgeSoftware(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.PurgeSoftware``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PurgeSoftware`: PurgeSoftware200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.PurgeSoftware`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPurgeSoftwareRequest struct via the builder pattern


### Return type

[**PurgeSoftware200Response**](PurgeSoftware200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegeneratePolicies

> RegeneratePolicies200Response RegeneratePolicies(ctx).Execute()

Trigger a new policy generation



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
	resp, r, err := apiClient.SystemAPI.RegeneratePolicies(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.RegeneratePolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegeneratePolicies`: RegeneratePolicies200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.RegeneratePolicies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRegeneratePoliciesRequest struct via the builder pattern


### Return type

[**RegeneratePolicies200Response**](RegeneratePolicies200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReloadAll

> ReloadAll200Response ReloadAll(ctx).Execute()

Reload both techniques and dynamic groups



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
	resp, r, err := apiClient.SystemAPI.ReloadAll(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.ReloadAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReloadAll`: ReloadAll200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.ReloadAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReloadAllRequest struct via the builder pattern


### Return type

[**ReloadAll200Response**](ReloadAll200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReloadGroups

> ReloadGroups200Response ReloadGroups(ctx).Execute()

Reload dynamic groups



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
	resp, r, err := apiClient.SystemAPI.ReloadGroups(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.ReloadGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReloadGroups`: ReloadGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.ReloadGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReloadGroupsRequest struct via the builder pattern


### Return type

[**ReloadGroups200Response**](ReloadGroups200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReloadTechniques

> ReloadTechniques200Response ReloadTechniques(ctx).Execute()

Reload techniques



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
	resp, r, err := apiClient.SystemAPI.ReloadTechniques(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.ReloadTechniques``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReloadTechniques`: ReloadTechniques200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.ReloadTechniques`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReloadTechniquesRequest struct via the builder pattern


### Return type

[**ReloadTechniques200Response**](ReloadTechniques200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreArchive

> RestoreArchive200Response RestoreArchive(ctx, archiveKind, archiveRestoreKind).Execute()

Restore an archive



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
	archiveKind := "full" // string | Type of archive to make
	archiveRestoreKind := "latestCommit" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.RestoreArchive(context.Background(), archiveKind, archiveRestoreKind).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.RestoreArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreArchive`: RestoreArchive200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.RestoreArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**archiveKind** | **string** | Type of archive to make | 
**archiveRestoreKind** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RestoreArchive200Response**](RestoreArchive200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePolicies

> UpdatePolicies200Response UpdatePolicies(ctx).Execute()

Trigger update of policies



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
	resp, r, err := apiClient.SystemAPI.UpdatePolicies(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.UpdatePolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePolicies`: UpdatePolicies200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.UpdatePolicies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePoliciesRequest struct via the builder pattern


### Return type

[**UpdatePolicies200Response**](UpdatePolicies200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

