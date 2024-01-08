# \StatusAPI

All URIs are relative to *https://rudder.example.local/rudder/api/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**None**](StatusAPI.md#None) | **Get** /status | Check if Rudder is alive



## None

> string None(ctx).Execute()

Check if Rudder is alive



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
	resp, r, err := apiClient.StatusAPI.None(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusAPI.None``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `None`: string
	fmt.Fprintf(os.Stdout, "Response from `StatusAPI.None`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNoneRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

