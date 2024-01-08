# DatasourceRunParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnGeneration** | Pointer to **bool** | Trigger a fetch at the beginning of a policy generation | [optional] 
**OnNewNode** | Pointer to **bool** | Trigger a fetch when a new node is accepted, for that node | [optional] 
**Schedule** | Pointer to [**DatasourceRunParametersSchedule**](DatasourceRunParametersSchedule.md) |  | [optional] 

## Methods

### NewDatasourceRunParameters

`func NewDatasourceRunParameters() *DatasourceRunParameters`

NewDatasourceRunParameters instantiates a new DatasourceRunParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasourceRunParametersWithDefaults

`func NewDatasourceRunParametersWithDefaults() *DatasourceRunParameters`

NewDatasourceRunParametersWithDefaults instantiates a new DatasourceRunParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnGeneration

`func (o *DatasourceRunParameters) GetOnGeneration() bool`

GetOnGeneration returns the OnGeneration field if non-nil, zero value otherwise.

### GetOnGenerationOk

`func (o *DatasourceRunParameters) GetOnGenerationOk() (*bool, bool)`

GetOnGenerationOk returns a tuple with the OnGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnGeneration

`func (o *DatasourceRunParameters) SetOnGeneration(v bool)`

SetOnGeneration sets OnGeneration field to given value.

### HasOnGeneration

`func (o *DatasourceRunParameters) HasOnGeneration() bool`

HasOnGeneration returns a boolean if a field has been set.

### GetOnNewNode

`func (o *DatasourceRunParameters) GetOnNewNode() bool`

GetOnNewNode returns the OnNewNode field if non-nil, zero value otherwise.

### GetOnNewNodeOk

`func (o *DatasourceRunParameters) GetOnNewNodeOk() (*bool, bool)`

GetOnNewNodeOk returns a tuple with the OnNewNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnNewNode

`func (o *DatasourceRunParameters) SetOnNewNode(v bool)`

SetOnNewNode sets OnNewNode field to given value.

### HasOnNewNode

`func (o *DatasourceRunParameters) HasOnNewNode() bool`

HasOnNewNode returns a boolean if a field has been set.

### GetSchedule

`func (o *DatasourceRunParameters) GetSchedule() DatasourceRunParametersSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *DatasourceRunParameters) GetScheduleOk() (*DatasourceRunParametersSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *DatasourceRunParameters) SetSchedule(v DatasourceRunParametersSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *DatasourceRunParameters) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


