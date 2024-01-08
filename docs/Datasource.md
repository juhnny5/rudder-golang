# Datasource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier of the data source to create. | [optional] 
**Name** | Pointer to **string** | The human readable name of the data source to create. | [optional] 
**Description** | Pointer to **string** | Description of the goal of the data source to create. | [optional] 
**Enabled** | Pointer to **bool** | Enable or disable data source. | [optional] 
**UpdateTimeout** | Pointer to **int32** | Duration in seconds before aborting data source update. The main goal is to prevent never ending requests. If a periodicity if configured, you should set that timeout at a lower value. | [optional] 
**RunParameters** | Pointer to [**DatasourceRunParameters**](DatasourceRunParameters.md) |  | [optional] 
**Type** | Pointer to [**DatasourceType**](DatasourceType.md) |  | [optional] 

## Methods

### NewDatasource

`func NewDatasource() *Datasource`

NewDatasource instantiates a new Datasource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasourceWithDefaults

`func NewDatasourceWithDefaults() *Datasource`

NewDatasourceWithDefaults instantiates a new Datasource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Datasource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Datasource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Datasource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Datasource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Datasource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Datasource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Datasource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Datasource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Datasource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Datasource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Datasource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Datasource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *Datasource) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Datasource) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Datasource) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Datasource) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUpdateTimeout

`func (o *Datasource) GetUpdateTimeout() int32`

GetUpdateTimeout returns the UpdateTimeout field if non-nil, zero value otherwise.

### GetUpdateTimeoutOk

`func (o *Datasource) GetUpdateTimeoutOk() (*int32, bool)`

GetUpdateTimeoutOk returns a tuple with the UpdateTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTimeout

`func (o *Datasource) SetUpdateTimeout(v int32)`

SetUpdateTimeout sets UpdateTimeout field to given value.

### HasUpdateTimeout

`func (o *Datasource) HasUpdateTimeout() bool`

HasUpdateTimeout returns a boolean if a field has been set.

### GetRunParameters

`func (o *Datasource) GetRunParameters() DatasourceRunParameters`

GetRunParameters returns the RunParameters field if non-nil, zero value otherwise.

### GetRunParametersOk

`func (o *Datasource) GetRunParametersOk() (*DatasourceRunParameters, bool)`

GetRunParametersOk returns a tuple with the RunParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunParameters

`func (o *Datasource) SetRunParameters(v DatasourceRunParameters)`

SetRunParameters sets RunParameters field to given value.

### HasRunParameters

`func (o *Datasource) HasRunParameters() bool`

HasRunParameters returns a boolean if a field has been set.

### GetType

`func (o *Datasource) GetType() DatasourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Datasource) GetTypeOk() (*DatasourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Datasource) SetType(v DatasourceType)`

SetType sets Type field to given value.

### HasType

`func (o *Datasource) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


