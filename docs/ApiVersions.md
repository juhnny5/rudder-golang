# ApiVersions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Latest** | Pointer to **int32** | Latest API version available | [optional] 
**All** | Pointer to [**[]ApiVersion**](ApiVersion.md) | List of API version and status | [optional] 

## Methods

### NewApiVersions

`func NewApiVersions() *ApiVersions`

NewApiVersions instantiates a new ApiVersions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiVersionsWithDefaults

`func NewApiVersionsWithDefaults() *ApiVersions`

NewApiVersionsWithDefaults instantiates a new ApiVersions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatest

`func (o *ApiVersions) GetLatest() int32`

GetLatest returns the Latest field if non-nil, zero value otherwise.

### GetLatestOk

`func (o *ApiVersions) GetLatestOk() (*int32, bool)`

GetLatestOk returns a tuple with the Latest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatest

`func (o *ApiVersions) SetLatest(v int32)`

SetLatest sets Latest field to given value.

### HasLatest

`func (o *ApiVersions) HasLatest() bool`

HasLatest returns a boolean if a field has been set.

### GetAll

`func (o *ApiVersions) GetAll() []ApiVersion`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *ApiVersions) GetAllOk() (*[]ApiVersion, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *ApiVersions) SetAll(v []ApiVersion)`

SetAll sets All field to given value.

### HasAll

`func (o *ApiVersions) HasAll() bool`

HasAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


