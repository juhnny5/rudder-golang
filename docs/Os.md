# Os

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | **string** | For Linux, a distribution, for Windows, the commercial name | 
**Version** | **string** | A string representation of the version | 
**FullName** | **string** | The long description name of the os | 
**ServicePack** | Pointer to **string** | a service pack informationnal string | [optional] 

## Methods

### NewOs

`func NewOs(type_ string, name string, version string, fullName string, ) *Os`

NewOs instantiates a new Os object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsWithDefaults

`func NewOsWithDefaults() *Os`

NewOsWithDefaults instantiates a new Os object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Os) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Os) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Os) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *Os) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Os) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Os) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *Os) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Os) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Os) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetFullName

`func (o *Os) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *Os) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *Os) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetServicePack

`func (o *Os) GetServicePack() string`

GetServicePack returns the ServicePack field if non-nil, zero value otherwise.

### GetServicePackOk

`func (o *Os) GetServicePackOk() (*string, bool)`

GetServicePackOk returns a tuple with the ServicePack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePack

`func (o *Os) SetServicePack(v string)`

SetServicePack sets ServicePack field to given value.

### HasServicePack

`func (o *Os) HasServicePack() bool`

HasServicePack returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


