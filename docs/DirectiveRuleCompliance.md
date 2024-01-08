# DirectiveRuleCompliance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | id of the rule | 
**Name** | **string** | Name of the rule | 
**Mode** | Pointer to **string** |  | [optional] 
**Compliance** | **float32** | Directive compliance level | 
**ComplianceDetails** | [**NodeComplianceComponentComplianceDetails**](NodeComplianceComponentComplianceDetails.md) |  | 
**Components** | Pointer to [**[]DirectiveRuleComplianceComponentsInner**](DirectiveRuleComplianceComponentsInner.md) |  | [optional] 

## Methods

### NewDirectiveRuleCompliance

`func NewDirectiveRuleCompliance(id string, name string, compliance float32, complianceDetails NodeComplianceComponentComplianceDetails, ) *DirectiveRuleCompliance`

NewDirectiveRuleCompliance instantiates a new DirectiveRuleCompliance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectiveRuleComplianceWithDefaults

`func NewDirectiveRuleComplianceWithDefaults() *DirectiveRuleCompliance`

NewDirectiveRuleComplianceWithDefaults instantiates a new DirectiveRuleCompliance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DirectiveRuleCompliance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DirectiveRuleCompliance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DirectiveRuleCompliance) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DirectiveRuleCompliance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DirectiveRuleCompliance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DirectiveRuleCompliance) SetName(v string)`

SetName sets Name field to given value.


### GetMode

`func (o *DirectiveRuleCompliance) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *DirectiveRuleCompliance) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *DirectiveRuleCompliance) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *DirectiveRuleCompliance) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetCompliance

`func (o *DirectiveRuleCompliance) GetCompliance() float32`

GetCompliance returns the Compliance field if non-nil, zero value otherwise.

### GetComplianceOk

`func (o *DirectiveRuleCompliance) GetComplianceOk() (*float32, bool)`

GetComplianceOk returns a tuple with the Compliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliance

`func (o *DirectiveRuleCompliance) SetCompliance(v float32)`

SetCompliance sets Compliance field to given value.


### GetComplianceDetails

`func (o *DirectiveRuleCompliance) GetComplianceDetails() NodeComplianceComponentComplianceDetails`

GetComplianceDetails returns the ComplianceDetails field if non-nil, zero value otherwise.

### GetComplianceDetailsOk

`func (o *DirectiveRuleCompliance) GetComplianceDetailsOk() (*NodeComplianceComponentComplianceDetails, bool)`

GetComplianceDetailsOk returns a tuple with the ComplianceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceDetails

`func (o *DirectiveRuleCompliance) SetComplianceDetails(v NodeComplianceComponentComplianceDetails)`

SetComplianceDetails sets ComplianceDetails field to given value.


### GetComponents

`func (o *DirectiveRuleCompliance) GetComponents() []DirectiveRuleComplianceComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *DirectiveRuleCompliance) GetComponentsOk() (*[]DirectiveRuleComplianceComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *DirectiveRuleCompliance) SetComponents(v []DirectiveRuleComplianceComponentsInner)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *DirectiveRuleCompliance) HasComponents() bool`

HasComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


