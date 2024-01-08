# GetDirectivesCompliance200ResponseDataDirectivesCompliance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | id of the directive | 
**Name** | **string** | Name of the directive | 
**Mode** | **string** |  | 
**Compliance** | **float32** | Directive compliance level | 
**ComplianceDetails** | [**GetDirectivesCompliance200ResponseDataDirectivesComplianceComplianceDetails**](GetDirectivesCompliance200ResponseDataDirectivesComplianceComplianceDetails.md) |  | 
**Rules** | [**DirectiveRuleCompliance**](DirectiveRuleCompliance.md) |  | 
**Nodes** | [**DirectiveNodeCompliance**](DirectiveNodeCompliance.md) |  | 

## Methods

### NewGetDirectivesCompliance200ResponseDataDirectivesCompliance

`func NewGetDirectivesCompliance200ResponseDataDirectivesCompliance(id string, name string, mode string, compliance float32, complianceDetails GetDirectivesCompliance200ResponseDataDirectivesComplianceComplianceDetails, rules DirectiveRuleCompliance, nodes DirectiveNodeCompliance, ) *GetDirectivesCompliance200ResponseDataDirectivesCompliance`

NewGetDirectivesCompliance200ResponseDataDirectivesCompliance instantiates a new GetDirectivesCompliance200ResponseDataDirectivesCompliance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDirectivesCompliance200ResponseDataDirectivesComplianceWithDefaults

`func NewGetDirectivesCompliance200ResponseDataDirectivesComplianceWithDefaults() *GetDirectivesCompliance200ResponseDataDirectivesCompliance`

NewGetDirectivesCompliance200ResponseDataDirectivesComplianceWithDefaults instantiates a new GetDirectivesCompliance200ResponseDataDirectivesCompliance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetDirectivesCompliance200ResponseDataDirectivesCompliance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetDirectivesCompliance200ResponseDataDirectivesCompliance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetDirectivesCompliance200ResponseDataDirectivesCompliance) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GetDirectivesCompliance200ResponseDataDirectivesCompliance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDirectivesCompliance200ResponseDataDirectivesCompliance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDirectivesCompliance200ResponseDataDirectivesCompliance) SetName(v string)`

SetName sets Name field to given value.


### GetMode

`func (o *GetDirectivesCompliance200ResponseDataDirectivesCompliance) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GetDirectivesCompliance200ResponseDataDirectivesCompliance) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GetDirectivesCompliance200ResponseDataDirectivesCompliance) SetMode(v string)`

SetMode sets Mode field to given value.


### GetCompliance

`func (o *GetDirectivesCompliance200ResponseDataDirectivesCompliance) GetCompliance() float32`

GetCompliance returns the Compliance field if non-nil, zero value otherwise.

### GetComplianceOk

`func (o *GetDirectivesCompliance200ResponseDataDirectivesCompliance) GetComplianceOk() (*float32, bool)`

GetComplianceOk returns a tuple with the Compliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliance

`func (o *GetDirectivesCompliance200ResponseDataDirectivesCompliance) SetCompliance(v float32)`

SetCompliance sets Compliance field to given value.


### GetComplianceDetails

`func (o *GetDirectivesCompliance200ResponseDataDirectivesCompliance) GetComplianceDetails() GetDirectivesCompliance200ResponseDataDirectivesComplianceComplianceDetails`

GetComplianceDetails returns the ComplianceDetails field if non-nil, zero value otherwise.

### GetComplianceDetailsOk

`func (o *GetDirectivesCompliance200ResponseDataDirectivesCompliance) GetComplianceDetailsOk() (*GetDirectivesCompliance200ResponseDataDirectivesComplianceComplianceDetails, bool)`

GetComplianceDetailsOk returns a tuple with the ComplianceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceDetails

`func (o *GetDirectivesCompliance200ResponseDataDirectivesCompliance) SetComplianceDetails(v GetDirectivesCompliance200ResponseDataDirectivesComplianceComplianceDetails)`

SetComplianceDetails sets ComplianceDetails field to given value.


### GetRules

`func (o *GetDirectivesCompliance200ResponseDataDirectivesCompliance) GetRules() DirectiveRuleCompliance`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *GetDirectivesCompliance200ResponseDataDirectivesCompliance) GetRulesOk() (*DirectiveRuleCompliance, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *GetDirectivesCompliance200ResponseDataDirectivesCompliance) SetRules(v DirectiveRuleCompliance)`

SetRules sets Rules field to given value.


### GetNodes

`func (o *GetDirectivesCompliance200ResponseDataDirectivesCompliance) GetNodes() DirectiveNodeCompliance`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *GetDirectivesCompliance200ResponseDataDirectivesCompliance) GetNodesOk() (*DirectiveNodeCompliance, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *GetDirectivesCompliance200ResponseDataDirectivesCompliance) SetNodes(v DirectiveNodeCompliance)`

SetNodes sets Nodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


