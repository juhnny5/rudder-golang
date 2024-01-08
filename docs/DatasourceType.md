# DatasourceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Data source type name | [optional] 
**Parameters** | Pointer to [**DatasourceTypeParameters**](DatasourceTypeParameters.md) |  | [optional] 

## Methods

### NewDatasourceType

`func NewDatasourceType() *DatasourceType`

NewDatasourceType instantiates a new DatasourceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasourceTypeWithDefaults

`func NewDatasourceTypeWithDefaults() *DatasourceType`

NewDatasourceTypeWithDefaults instantiates a new DatasourceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DatasourceType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatasourceType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatasourceType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DatasourceType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameters

`func (o *DatasourceType) GetParameters() DatasourceTypeParameters`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *DatasourceType) GetParametersOk() (*DatasourceTypeParameters, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *DatasourceType) SetParameters(v DatasourceTypeParameters)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *DatasourceType) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


