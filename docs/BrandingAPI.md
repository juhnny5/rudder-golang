# \BrandingAPI

All URIs are relative to *https://rudder.example.local/rudder/api/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBrandingConf**](BrandingAPI.md#GetBrandingConf) | **Get** /branding | Get branding configuration
[**ReloadBrandingConf**](BrandingAPI.md#ReloadBrandingConf) | **Post** /branding/reload | Reload branding file
[**UpdateBRandingConf**](BrandingAPI.md#UpdateBRandingConf) | **Post** /branding | Update web interface customization



## GetBrandingConf

> GetBrandingConf200Response GetBrandingConf(ctx).Execute()

Get branding configuration



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
	resp, r, err := apiClient.BrandingAPI.GetBrandingConf(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BrandingAPI.GetBrandingConf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBrandingConf`: GetBrandingConf200Response
	fmt.Fprintf(os.Stdout, "Response from `BrandingAPI.GetBrandingConf`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandingConfRequest struct via the builder pattern


### Return type

[**GetBrandingConf200Response**](GetBrandingConf200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReloadBrandingConf

> GetBrandingConf200Response ReloadBrandingConf(ctx).Execute()

Reload branding file



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
	resp, r, err := apiClient.BrandingAPI.ReloadBrandingConf(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BrandingAPI.ReloadBrandingConf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReloadBrandingConf`: GetBrandingConf200Response
	fmt.Fprintf(os.Stdout, "Response from `BrandingAPI.ReloadBrandingConf`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReloadBrandingConfRequest struct via the builder pattern


### Return type

[**GetBrandingConf200Response**](GetBrandingConf200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBRandingConf

> UpdateBRandingConf200Response UpdateBRandingConf(ctx).BrandingConf(brandingConf).Execute()

Update web interface customization



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
	brandingConf := *openapiclient.NewBrandingConf(false, false, "Production", *openapiclient.NewColor(float32(0.2), float32(0.235), float32(0.01), float32(0.5)), *openapiclient.NewColor(float32(0.2), float32(0.235), float32(0.01), float32(0.5)), *openapiclient.NewLogo(false), *openapiclient.NewLogo(false), false, false, "Welcome, please sign in:") // BrandingConf | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BrandingAPI.UpdateBRandingConf(context.Background()).BrandingConf(brandingConf).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BrandingAPI.UpdateBRandingConf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBRandingConf`: UpdateBRandingConf200Response
	fmt.Fprintf(os.Stdout, "Response from `BrandingAPI.UpdateBRandingConf`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBRandingConfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **brandingConf** | [**BrandingConf**](BrandingConf.md) |  | 

### Return type

[**UpdateBRandingConf200Response**](UpdateBRandingConf200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

