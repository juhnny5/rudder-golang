# \ArchivesAPI

All URIs are relative to *https://rudder.example.local/rudder/api/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CallImport**](ArchivesAPI.md#CallImport) | **Post** /archives/import | Import a ZIP archive of policies into Rudder
[**Export**](ArchivesAPI.md#Export) | **Get** /archives/export | Get a ZIP archive of the requested items and their dependencies



## CallImport

> Import200Response CallImport(ctx).Archive(archive).Merge(merge).Execute()

Import a ZIP archive of policies into Rudder



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
	archive := os.NewFile(1234, "some_file") // *os.File | The ZIP archive file containing policies in a conventional layout and serialization format (optional)
	merge := "merge_example" // string | Optional merge algo of the import. Default `override-all` means what is in the archive is the new reality. `keep-rule-groups` will keep existing target definition for existing rules (ignore archive value). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArchivesAPI.CallImport(context.Background()).Archive(archive).Merge(merge).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArchivesAPI.CallImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CallImport`: Import200Response
	fmt.Fprintf(os.Stdout, "Response from `ArchivesAPI.CallImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **archive** | ***os.File** | The ZIP archive file containing policies in a conventional layout and serialization format | 
 **merge** | **string** | Optional merge algo of the import. Default &#x60;override-all&#x60; means what is in the archive is the new reality. &#x60;keep-rule-groups&#x60; will keep existing target definition for existing rules (ignore archive value). | 

### Return type

[**Import200Response**](Import200Response.md)

### Authorization

[API-Tokens](../README.md#API-Tokens)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Export

> *os.File Export(ctx).Rules(rules).Directives(directives).Techniques(techniques).Groups(groups).Include(include).Execute()

Get a ZIP archive of the requested items and their dependencies



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
	rules := []string{"Inner_example"} // []string | IDs (optionally with revision, '+' need to be escaped as '%2B') of rules to include (optional)
	directives := []string{"Inner_example"} // []string | IDs (optionally with revision, '+' need to be escaped as '%2B') of directives to include (optional)
	techniques := []string{"Inner_example"} // []string | IDs, ie technique name/technique version (optionally with revision, '+' need to be escaped as '%2B') of techniques to include (optional)
	groups := []string{"Inner_example"} // []string | IDs (optionally with revision, '+' need to be escaped as '%2B') of groups to include (optional)
	include := []string{"Include_example"} // []string | Scope of dependencies to include in archive, where rule as directives and groups dependencies, directives have techniques dependencies, and techniques and groups don't have dependencies. 'none' means no dependencies will be include, 'all' means that the whole tree will,  'directives' and 'groups' means to include them specifically, 'techniques' means to include both directives and techniques. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArchivesAPI.Export(context.Background()).Rules(rules).Directives(directives).Techniques(techniques).Groups(groups).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArchivesAPI.Export``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Export`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ArchivesAPI.Export`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rules** | **[]string** | IDs (optionally with revision, &#39;+&#39; need to be escaped as &#39;%2B&#39;) of rules to include | 
 **directives** | **[]string** | IDs (optionally with revision, &#39;+&#39; need to be escaped as &#39;%2B&#39;) of directives to include | 
 **techniques** | **[]string** | IDs, ie technique name/technique version (optionally with revision, &#39;+&#39; need to be escaped as &#39;%2B&#39;) of techniques to include | 
 **groups** | **[]string** | IDs (optionally with revision, &#39;+&#39; need to be escaped as &#39;%2B&#39;) of groups to include | 
 **include** | **[]string** | Scope of dependencies to include in archive, where rule as directives and groups dependencies, directives have techniques dependencies, and techniques and groups don&#39;t have dependencies. &#39;none&#39; means no dependencies will be include, &#39;all&#39; means that the whole tree will,  &#39;directives&#39; and &#39;groups&#39; means to include them specifically, &#39;techniques&#39; means to include both directives and techniques. | 

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

