# TechniqueBlockCallsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Method call id | [optional] 
**Component** | Pointer to **string** | Component is used in reporting to identify this method call. You can see it as a name | [optional] 
**Method** | Pointer to **string** | Id of the method called | [optional] 
**Condition** | Pointer to **string** | Condition to run this method. | [optional] 
**DisableReporting** | Pointer to **bool** | Should the reporting of this method be disabled | [optional] 
**Parameters** | Pointer to [**[]TechniqueMethodCallParametersInner**](TechniqueMethodCallParametersInner.md) | Parameters for this method call | [optional] 
**Calls** | Pointer to [**[]TechniqueBlockCallsInner**](TechniqueBlockCallsInner.md) | Method and blocks contained by this block | [optional] 
**ReportingLogic** | Pointer to [**TechniqueBlockReportingLogic**](TechniqueBlockReportingLogic.md) |  | [optional] 

## Methods

### NewTechniqueBlockCallsInner

`func NewTechniqueBlockCallsInner() *TechniqueBlockCallsInner`

NewTechniqueBlockCallsInner instantiates a new TechniqueBlockCallsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTechniqueBlockCallsInnerWithDefaults

`func NewTechniqueBlockCallsInnerWithDefaults() *TechniqueBlockCallsInner`

NewTechniqueBlockCallsInnerWithDefaults instantiates a new TechniqueBlockCallsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TechniqueBlockCallsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TechniqueBlockCallsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TechniqueBlockCallsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TechniqueBlockCallsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetComponent

`func (o *TechniqueBlockCallsInner) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *TechniqueBlockCallsInner) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *TechniqueBlockCallsInner) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *TechniqueBlockCallsInner) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetMethod

`func (o *TechniqueBlockCallsInner) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *TechniqueBlockCallsInner) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *TechniqueBlockCallsInner) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *TechniqueBlockCallsInner) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetCondition

`func (o *TechniqueBlockCallsInner) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *TechniqueBlockCallsInner) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *TechniqueBlockCallsInner) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *TechniqueBlockCallsInner) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetDisableReporting

`func (o *TechniqueBlockCallsInner) GetDisableReporting() bool`

GetDisableReporting returns the DisableReporting field if non-nil, zero value otherwise.

### GetDisableReportingOk

`func (o *TechniqueBlockCallsInner) GetDisableReportingOk() (*bool, bool)`

GetDisableReportingOk returns a tuple with the DisableReporting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableReporting

`func (o *TechniqueBlockCallsInner) SetDisableReporting(v bool)`

SetDisableReporting sets DisableReporting field to given value.

### HasDisableReporting

`func (o *TechniqueBlockCallsInner) HasDisableReporting() bool`

HasDisableReporting returns a boolean if a field has been set.

### GetParameters

`func (o *TechniqueBlockCallsInner) GetParameters() []TechniqueMethodCallParametersInner`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *TechniqueBlockCallsInner) GetParametersOk() (*[]TechniqueMethodCallParametersInner, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *TechniqueBlockCallsInner) SetParameters(v []TechniqueMethodCallParametersInner)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *TechniqueBlockCallsInner) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetCalls

`func (o *TechniqueBlockCallsInner) GetCalls() []TechniqueBlockCallsInner`

GetCalls returns the Calls field if non-nil, zero value otherwise.

### GetCallsOk

`func (o *TechniqueBlockCallsInner) GetCallsOk() (*[]TechniqueBlockCallsInner, bool)`

GetCallsOk returns a tuple with the Calls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalls

`func (o *TechniqueBlockCallsInner) SetCalls(v []TechniqueBlockCallsInner)`

SetCalls sets Calls field to given value.

### HasCalls

`func (o *TechniqueBlockCallsInner) HasCalls() bool`

HasCalls returns a boolean if a field has been set.

### GetReportingLogic

`func (o *TechniqueBlockCallsInner) GetReportingLogic() TechniqueBlockReportingLogic`

GetReportingLogic returns the ReportingLogic field if non-nil, zero value otherwise.

### GetReportingLogicOk

`func (o *TechniqueBlockCallsInner) GetReportingLogicOk() (*TechniqueBlockReportingLogic, bool)`

GetReportingLogicOk returns a tuple with the ReportingLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingLogic

`func (o *TechniqueBlockCallsInner) SetReportingLogic(v TechniqueBlockReportingLogic)`

SetReportingLogic sets ReportingLogic field to given value.

### HasReportingLogic

`func (o *TechniqueBlockCallsInner) HasReportingLogic() bool`

HasReportingLogic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


