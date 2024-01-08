# ApiInformations200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Documentation** | **string** | Link to Rudder API description | 
**EndpointName** | **string** | The endpoint name as key and the endpoint description as value | 
**Endpoints** | [**[]ApiInformations200ResponseDataEndpointsInner**](ApiInformations200ResponseDataEndpointsInner.md) |  | 

## Methods

### NewApiInformations200ResponseData

`func NewApiInformations200ResponseData(documentation string, endpointName string, endpoints []ApiInformations200ResponseDataEndpointsInner, ) *ApiInformations200ResponseData`

NewApiInformations200ResponseData instantiates a new ApiInformations200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiInformations200ResponseDataWithDefaults

`func NewApiInformations200ResponseDataWithDefaults() *ApiInformations200ResponseData`

NewApiInformations200ResponseDataWithDefaults instantiates a new ApiInformations200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentation

`func (o *ApiInformations200ResponseData) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *ApiInformations200ResponseData) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *ApiInformations200ResponseData) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.


### GetEndpointName

`func (o *ApiInformations200ResponseData) GetEndpointName() string`

GetEndpointName returns the EndpointName field if non-nil, zero value otherwise.

### GetEndpointNameOk

`func (o *ApiInformations200ResponseData) GetEndpointNameOk() (*string, bool)`

GetEndpointNameOk returns a tuple with the EndpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointName

`func (o *ApiInformations200ResponseData) SetEndpointName(v string)`

SetEndpointName sets EndpointName field to given value.


### GetEndpoints

`func (o *ApiInformations200ResponseData) GetEndpoints() []ApiInformations200ResponseDataEndpointsInner`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *ApiInformations200ResponseData) GetEndpointsOk() (*[]ApiInformations200ResponseDataEndpointsInner, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *ApiInformations200ResponseData) SetEndpoints(v []ApiInformations200ResponseDataEndpointsInner)`

SetEndpoints sets Endpoints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


