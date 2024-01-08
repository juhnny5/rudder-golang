# DirectiveNodeCompliance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | id of the node | 
**Name** | **string** | Name of the node | 
**Compliance** | **float32** | Directive compliance level | 
**ComplianceDetails** | [**NodeComplianceComponentComplianceDetails**](NodeComplianceComponentComplianceDetails.md) |  | 
**Rules** | [**RuleComplianceComponent**](RuleComplianceComponent.md) |  | 

## Methods

### NewDirectiveNodeCompliance

`func NewDirectiveNodeCompliance(id string, name string, compliance float32, complianceDetails NodeComplianceComponentComplianceDetails, rules RuleComplianceComponent, ) *DirectiveNodeCompliance`

NewDirectiveNodeCompliance instantiates a new DirectiveNodeCompliance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectiveNodeComplianceWithDefaults

`func NewDirectiveNodeComplianceWithDefaults() *DirectiveNodeCompliance`

NewDirectiveNodeComplianceWithDefaults instantiates a new DirectiveNodeCompliance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DirectiveNodeCompliance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DirectiveNodeCompliance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DirectiveNodeCompliance) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DirectiveNodeCompliance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DirectiveNodeCompliance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DirectiveNodeCompliance) SetName(v string)`

SetName sets Name field to given value.


### GetCompliance

`func (o *DirectiveNodeCompliance) GetCompliance() float32`

GetCompliance returns the Compliance field if non-nil, zero value otherwise.

### GetComplianceOk

`func (o *DirectiveNodeCompliance) GetComplianceOk() (*float32, bool)`

GetComplianceOk returns a tuple with the Compliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliance

`func (o *DirectiveNodeCompliance) SetCompliance(v float32)`

SetCompliance sets Compliance field to given value.


### GetComplianceDetails

`func (o *DirectiveNodeCompliance) GetComplianceDetails() NodeComplianceComponentComplianceDetails`

GetComplianceDetails returns the ComplianceDetails field if non-nil, zero value otherwise.

### GetComplianceDetailsOk

`func (o *DirectiveNodeCompliance) GetComplianceDetailsOk() (*NodeComplianceComponentComplianceDetails, bool)`

GetComplianceDetailsOk returns a tuple with the ComplianceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceDetails

`func (o *DirectiveNodeCompliance) SetComplianceDetails(v NodeComplianceComponentComplianceDetails)`

SetComplianceDetails sets ComplianceDetails field to given value.


### GetRules

`func (o *DirectiveNodeCompliance) GetRules() RuleComplianceComponent`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *DirectiveNodeCompliance) GetRulesOk() (*RuleComplianceComponent, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *DirectiveNodeCompliance) SetRules(v RuleComplianceComponent)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


