# RuleComplianceComponentComponentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the component | 
**Compliance** | **float32** | directive compliance level | 
**ComplianceDetails** | [**RuleComplianceComponentComplianceDetails**](RuleComplianceComponentComplianceDetails.md) |  | 
**Values** | [**[]RuleComplianceComponentComponentsInnerValuesInner**](RuleComplianceComponentComponentsInnerValuesInner.md) |  | 

## Methods

### NewRuleComplianceComponentComponentsInner

`func NewRuleComplianceComponentComponentsInner(name string, compliance float32, complianceDetails RuleComplianceComponentComplianceDetails, values []RuleComplianceComponentComponentsInnerValuesInner, ) *RuleComplianceComponentComponentsInner`

NewRuleComplianceComponentComponentsInner instantiates a new RuleComplianceComponentComponentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleComplianceComponentComponentsInnerWithDefaults

`func NewRuleComplianceComponentComponentsInnerWithDefaults() *RuleComplianceComponentComponentsInner`

NewRuleComplianceComponentComponentsInnerWithDefaults instantiates a new RuleComplianceComponentComponentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RuleComplianceComponentComponentsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleComplianceComponentComponentsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleComplianceComponentComponentsInner) SetName(v string)`

SetName sets Name field to given value.


### GetCompliance

`func (o *RuleComplianceComponentComponentsInner) GetCompliance() float32`

GetCompliance returns the Compliance field if non-nil, zero value otherwise.

### GetComplianceOk

`func (o *RuleComplianceComponentComponentsInner) GetComplianceOk() (*float32, bool)`

GetComplianceOk returns a tuple with the Compliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliance

`func (o *RuleComplianceComponentComponentsInner) SetCompliance(v float32)`

SetCompliance sets Compliance field to given value.


### GetComplianceDetails

`func (o *RuleComplianceComponentComponentsInner) GetComplianceDetails() RuleComplianceComponentComplianceDetails`

GetComplianceDetails returns the ComplianceDetails field if non-nil, zero value otherwise.

### GetComplianceDetailsOk

`func (o *RuleComplianceComponentComponentsInner) GetComplianceDetailsOk() (*RuleComplianceComponentComplianceDetails, bool)`

GetComplianceDetailsOk returns a tuple with the ComplianceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceDetails

`func (o *RuleComplianceComponentComponentsInner) SetComplianceDetails(v RuleComplianceComponentComplianceDetails)`

SetComplianceDetails sets ComplianceDetails field to given value.


### GetValues

`func (o *RuleComplianceComponentComponentsInner) GetValues() []RuleComplianceComponentComponentsInnerValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *RuleComplianceComponentComponentsInner) GetValuesOk() (*[]RuleComplianceComponentComponentsInnerValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *RuleComplianceComponentComponentsInner) SetValues(v []RuleComplianceComponentComponentsInnerValuesInner)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


