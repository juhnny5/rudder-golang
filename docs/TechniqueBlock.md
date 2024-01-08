# TechniqueBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Method call id | [optional] 
**Component** | Pointer to **string** | Component is used in reporting to identify this method call. You can see it as a name | [optional] 
**Condition** | Pointer to **string** | Condition to run this method. | [optional] 
**Calls** | Pointer to [**[]TechniqueBlockCallsInner**](TechniqueBlockCallsInner.md) | Method and blocks contained by this block | [optional] 
**ReportingLogic** | Pointer to [**TechniqueBlockReportingLogic**](TechniqueBlockReportingLogic.md) |  | [optional] 

## Methods

### NewTechniqueBlock

`func NewTechniqueBlock() *TechniqueBlock`

NewTechniqueBlock instantiates a new TechniqueBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTechniqueBlockWithDefaults

`func NewTechniqueBlockWithDefaults() *TechniqueBlock`

NewTechniqueBlockWithDefaults instantiates a new TechniqueBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TechniqueBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TechniqueBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TechniqueBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TechniqueBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetComponent

`func (o *TechniqueBlock) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *TechniqueBlock) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *TechniqueBlock) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *TechniqueBlock) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetCondition

`func (o *TechniqueBlock) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *TechniqueBlock) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *TechniqueBlock) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *TechniqueBlock) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetCalls

`func (o *TechniqueBlock) GetCalls() []TechniqueBlockCallsInner`

GetCalls returns the Calls field if non-nil, zero value otherwise.

### GetCallsOk

`func (o *TechniqueBlock) GetCallsOk() (*[]TechniqueBlockCallsInner, bool)`

GetCallsOk returns a tuple with the Calls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalls

`func (o *TechniqueBlock) SetCalls(v []TechniqueBlockCallsInner)`

SetCalls sets Calls field to given value.

### HasCalls

`func (o *TechniqueBlock) HasCalls() bool`

HasCalls returns a boolean if a field has been set.

### GetReportingLogic

`func (o *TechniqueBlock) GetReportingLogic() TechniqueBlockReportingLogic`

GetReportingLogic returns the ReportingLogic field if non-nil, zero value otherwise.

### GetReportingLogicOk

`func (o *TechniqueBlock) GetReportingLogicOk() (*TechniqueBlockReportingLogic, bool)`

GetReportingLogicOk returns a tuple with the ReportingLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingLogic

`func (o *TechniqueBlock) SetReportingLogic(v TechniqueBlockReportingLogic)`

SetReportingLogic sets ReportingLogic field to given value.

### HasReportingLogic

`func (o *TechniqueBlock) HasReportingLogic() bool`

HasReportingLogic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


