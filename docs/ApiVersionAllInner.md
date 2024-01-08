# ApiVersionAllInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** | API Version | 
**Status** | **interface{}** | Status of the API, either maintained or deprecated | 

## Methods

### NewApiVersionAllInner

`func NewApiVersionAllInner(version int32, status interface{}, ) *ApiVersionAllInner`

NewApiVersionAllInner instantiates a new ApiVersionAllInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiVersionAllInnerWithDefaults

`func NewApiVersionAllInnerWithDefaults() *ApiVersionAllInner`

NewApiVersionAllInnerWithDefaults instantiates a new ApiVersionAllInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ApiVersionAllInner) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApiVersionAllInner) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApiVersionAllInner) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetStatus

`func (o *ApiVersionAllInner) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiVersionAllInner) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiVersionAllInner) SetStatus(v interface{})`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *ApiVersionAllInner) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ApiVersionAllInner) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


