# DirectiveRuleComplianceComponentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the component | 
**Compliance** | **float32** | Directive compliance level | 
**ComplianceDetails** | [**NodeComplianceComponentComplianceDetails**](NodeComplianceComponentComplianceDetails.md) |  | 
**Nodes** | [**NodeComplianceComponent**](NodeComplianceComponent.md) |  | 

## Methods

### NewDirectiveRuleComplianceComponentsInner

`func NewDirectiveRuleComplianceComponentsInner(name string, compliance float32, complianceDetails NodeComplianceComponentComplianceDetails, nodes NodeComplianceComponent, ) *DirectiveRuleComplianceComponentsInner`

NewDirectiveRuleComplianceComponentsInner instantiates a new DirectiveRuleComplianceComponentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectiveRuleComplianceComponentsInnerWithDefaults

`func NewDirectiveRuleComplianceComponentsInnerWithDefaults() *DirectiveRuleComplianceComponentsInner`

NewDirectiveRuleComplianceComponentsInnerWithDefaults instantiates a new DirectiveRuleComplianceComponentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DirectiveRuleComplianceComponentsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DirectiveRuleComplianceComponentsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DirectiveRuleComplianceComponentsInner) SetName(v string)`

SetName sets Name field to given value.


### GetCompliance

`func (o *DirectiveRuleComplianceComponentsInner) GetCompliance() float32`

GetCompliance returns the Compliance field if non-nil, zero value otherwise.

### GetComplianceOk

`func (o *DirectiveRuleComplianceComponentsInner) GetComplianceOk() (*float32, bool)`

GetComplianceOk returns a tuple with the Compliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliance

`func (o *DirectiveRuleComplianceComponentsInner) SetCompliance(v float32)`

SetCompliance sets Compliance field to given value.


### GetComplianceDetails

`func (o *DirectiveRuleComplianceComponentsInner) GetComplianceDetails() NodeComplianceComponentComplianceDetails`

GetComplianceDetails returns the ComplianceDetails field if non-nil, zero value otherwise.

### GetComplianceDetailsOk

`func (o *DirectiveRuleComplianceComponentsInner) GetComplianceDetailsOk() (*NodeComplianceComponentComplianceDetails, bool)`

GetComplianceDetailsOk returns a tuple with the ComplianceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceDetails

`func (o *DirectiveRuleComplianceComponentsInner) SetComplianceDetails(v NodeComplianceComponentComplianceDetails)`

SetComplianceDetails sets ComplianceDetails field to given value.


### GetNodes

`func (o *DirectiveRuleComplianceComponentsInner) GetNodes() NodeComplianceComponent`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *DirectiveRuleComplianceComponentsInner) GetNodesOk() (*NodeComplianceComponent, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *DirectiveRuleComplianceComponentsInner) SetNodes(v NodeComplianceComponent)`

SetNodes sets Nodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


