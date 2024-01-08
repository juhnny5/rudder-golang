# TechniqueMethodCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Method call id | [optional] 
**Component** | Pointer to **string** | Component is used in reporting to identify this method call. You can see it as a name | [optional] 
**Method** | Pointer to **string** | Id of the method called | [optional] 
**Condition** | Pointer to **string** | Condition to run this method. | [optional] 
**DisableReporting** | Pointer to **bool** | Should the reporting of this method be disabled | [optional] 
**Parameters** | Pointer to [**[]TechniqueMethodCallParametersInner**](TechniqueMethodCallParametersInner.md) | Parameters for this method call | [optional] 

## Methods

### NewTechniqueMethodCall

`func NewTechniqueMethodCall() *TechniqueMethodCall`

NewTechniqueMethodCall instantiates a new TechniqueMethodCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTechniqueMethodCallWithDefaults

`func NewTechniqueMethodCallWithDefaults() *TechniqueMethodCall`

NewTechniqueMethodCallWithDefaults instantiates a new TechniqueMethodCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TechniqueMethodCall) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TechniqueMethodCall) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TechniqueMethodCall) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TechniqueMethodCall) HasId() bool`

HasId returns a boolean if a field has been set.

### GetComponent

`func (o *TechniqueMethodCall) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *TechniqueMethodCall) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *TechniqueMethodCall) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *TechniqueMethodCall) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetMethod

`func (o *TechniqueMethodCall) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *TechniqueMethodCall) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *TechniqueMethodCall) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *TechniqueMethodCall) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetCondition

`func (o *TechniqueMethodCall) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *TechniqueMethodCall) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *TechniqueMethodCall) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *TechniqueMethodCall) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetDisableReporting

`func (o *TechniqueMethodCall) GetDisableReporting() bool`

GetDisableReporting returns the DisableReporting field if non-nil, zero value otherwise.

### GetDisableReportingOk

`func (o *TechniqueMethodCall) GetDisableReportingOk() (*bool, bool)`

GetDisableReportingOk returns a tuple with the DisableReporting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableReporting

`func (o *TechniqueMethodCall) SetDisableReporting(v bool)`

SetDisableReporting sets DisableReporting field to given value.

### HasDisableReporting

`func (o *TechniqueMethodCall) HasDisableReporting() bool`

HasDisableReporting returns a boolean if a field has been set.

### GetParameters

`func (o *TechniqueMethodCall) GetParameters() []TechniqueMethodCallParametersInner`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *TechniqueMethodCall) GetParametersOk() (*[]TechniqueMethodCallParametersInner, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *TechniqueMethodCall) SetParameters(v []TechniqueMethodCallParametersInner)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *TechniqueMethodCall) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


