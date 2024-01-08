# RuleComplianceComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | id of the rule | 
**Name** | **string** | Name of the rule | 
**Compliance** | **float32** | Rule compliance level | 
**ComplianceDetails** | [**RuleComplianceComponentComplianceDetails**](RuleComplianceComponentComplianceDetails.md) |  | 
**Components** | [**[]RuleComplianceComponentComponentsInner**](RuleComplianceComponentComponentsInner.md) |  | 

## Methods

### NewRuleComplianceComponent

`func NewRuleComplianceComponent(id string, name string, compliance float32, complianceDetails RuleComplianceComponentComplianceDetails, components []RuleComplianceComponentComponentsInner, ) *RuleComplianceComponent`

NewRuleComplianceComponent instantiates a new RuleComplianceComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleComplianceComponentWithDefaults

`func NewRuleComplianceComponentWithDefaults() *RuleComplianceComponent`

NewRuleComplianceComponentWithDefaults instantiates a new RuleComplianceComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RuleComplianceComponent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuleComplianceComponent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuleComplianceComponent) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *RuleComplianceComponent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleComplianceComponent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleComplianceComponent) SetName(v string)`

SetName sets Name field to given value.


### GetCompliance

`func (o *RuleComplianceComponent) GetCompliance() float32`

GetCompliance returns the Compliance field if non-nil, zero value otherwise.

### GetComplianceOk

`func (o *RuleComplianceComponent) GetComplianceOk() (*float32, bool)`

GetComplianceOk returns a tuple with the Compliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliance

`func (o *RuleComplianceComponent) SetCompliance(v float32)`

SetCompliance sets Compliance field to given value.


### GetComplianceDetails

`func (o *RuleComplianceComponent) GetComplianceDetails() RuleComplianceComponentComplianceDetails`

GetComplianceDetails returns the ComplianceDetails field if non-nil, zero value otherwise.

### GetComplianceDetailsOk

`func (o *RuleComplianceComponent) GetComplianceDetailsOk() (*RuleComplianceComponentComplianceDetails, bool)`

GetComplianceDetailsOk returns a tuple with the ComplianceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceDetails

`func (o *RuleComplianceComponent) SetComplianceDetails(v RuleComplianceComponentComplianceDetails)`

SetComplianceDetails sets ComplianceDetails field to given value.


### GetComponents

`func (o *RuleComplianceComponent) GetComponents() []RuleComplianceComponentComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *RuleComplianceComponent) GetComponentsOk() (*[]RuleComplianceComponentComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *RuleComplianceComponent) SetComponents(v []RuleComplianceComponentComponentsInner)`

SetComponents sets Components field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


