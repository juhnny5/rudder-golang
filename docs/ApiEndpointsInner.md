# ApiEndpointsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointName** | Pointer to **string** | The endpoint name for key and its description for value | [optional] 
**HttpVerb** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewApiEndpointsInner

`func NewApiEndpointsInner() *ApiEndpointsInner`

NewApiEndpointsInner instantiates a new ApiEndpointsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiEndpointsInnerWithDefaults

`func NewApiEndpointsInnerWithDefaults() *ApiEndpointsInner`

NewApiEndpointsInnerWithDefaults instantiates a new ApiEndpointsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointName

`func (o *ApiEndpointsInner) GetEndpointName() string`

GetEndpointName returns the EndpointName field if non-nil, zero value otherwise.

### GetEndpointNameOk

`func (o *ApiEndpointsInner) GetEndpointNameOk() (*string, bool)`

GetEndpointNameOk returns a tuple with the EndpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointName

`func (o *ApiEndpointsInner) SetEndpointName(v string)`

SetEndpointName sets EndpointName field to given value.

### HasEndpointName

`func (o *ApiEndpointsInner) HasEndpointName() bool`

HasEndpointName returns a boolean if a field has been set.

### GetHttpVerb

`func (o *ApiEndpointsInner) GetHttpVerb() interface{}`

GetHttpVerb returns the HttpVerb field if non-nil, zero value otherwise.

### GetHttpVerbOk

`func (o *ApiEndpointsInner) GetHttpVerbOk() (*interface{}, bool)`

GetHttpVerbOk returns a tuple with the HttpVerb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpVerb

`func (o *ApiEndpointsInner) SetHttpVerb(v interface{})`

SetHttpVerb sets HttpVerb field to given value.

### HasHttpVerb

`func (o *ApiEndpointsInner) HasHttpVerb() bool`

HasHttpVerb returns a boolean if a field has been set.

### SetHttpVerbNil

`func (o *ApiEndpointsInner) SetHttpVerbNil(b bool)`

 SetHttpVerbNil sets the value for HttpVerb to be an explicit nil

### UnsetHttpVerb
`func (o *ApiEndpointsInner) UnsetHttpVerb()`

UnsetHttpVerb ensures that no value is present for HttpVerb, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


