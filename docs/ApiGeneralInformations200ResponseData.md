# ApiGeneralInformations200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Documentation** | **string** | Link to Rudder API description | 
**AvailableVersions** | [**[]ApiVersions**](ApiVersions.md) |  | 
**Endpoints** | [**[][]ApiEndpointsInner**]([]ApiEndpointsInner.md) |  | 

## Methods

### NewApiGeneralInformations200ResponseData

`func NewApiGeneralInformations200ResponseData(documentation string, availableVersions []ApiVersions, endpoints [][]ApiEndpointsInner, ) *ApiGeneralInformations200ResponseData`

NewApiGeneralInformations200ResponseData instantiates a new ApiGeneralInformations200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiGeneralInformations200ResponseDataWithDefaults

`func NewApiGeneralInformations200ResponseDataWithDefaults() *ApiGeneralInformations200ResponseData`

NewApiGeneralInformations200ResponseDataWithDefaults instantiates a new ApiGeneralInformations200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentation

`func (o *ApiGeneralInformations200ResponseData) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *ApiGeneralInformations200ResponseData) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *ApiGeneralInformations200ResponseData) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.


### GetAvailableVersions

`func (o *ApiGeneralInformations200ResponseData) GetAvailableVersions() []ApiVersions`

GetAvailableVersions returns the AvailableVersions field if non-nil, zero value otherwise.

### GetAvailableVersionsOk

`func (o *ApiGeneralInformations200ResponseData) GetAvailableVersionsOk() (*[]ApiVersions, bool)`

GetAvailableVersionsOk returns a tuple with the AvailableVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableVersions

`func (o *ApiGeneralInformations200ResponseData) SetAvailableVersions(v []ApiVersions)`

SetAvailableVersions sets AvailableVersions field to given value.


### GetEndpoints

`func (o *ApiGeneralInformations200ResponseData) GetEndpoints() [][]ApiEndpointsInner`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *ApiGeneralInformations200ResponseData) GetEndpointsOk() (*[][]ApiEndpointsInner, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *ApiGeneralInformations200ResponseData) SetEndpoints(v [][]ApiEndpointsInner)`

SetEndpoints sets Endpoints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


