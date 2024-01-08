# GetRulesCompliance200ResponseDataRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | id of the rule | 
**Mode** | **string** |  | 
**Compliance** | **float32** | Rule compliance level | 
**ComplianceDetails** | [**GetGlobalCompliance200ResponseDataGlobalComplianceComplianceDetails**](GetGlobalCompliance200ResponseDataGlobalComplianceComplianceDetails.md) |  | 

## Methods

### NewGetRulesCompliance200ResponseDataRulesInner

`func NewGetRulesCompliance200ResponseDataRulesInner(id string, mode string, compliance float32, complianceDetails GetGlobalCompliance200ResponseDataGlobalComplianceComplianceDetails, ) *GetRulesCompliance200ResponseDataRulesInner`

NewGetRulesCompliance200ResponseDataRulesInner instantiates a new GetRulesCompliance200ResponseDataRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRulesCompliance200ResponseDataRulesInnerWithDefaults

`func NewGetRulesCompliance200ResponseDataRulesInnerWithDefaults() *GetRulesCompliance200ResponseDataRulesInner`

NewGetRulesCompliance200ResponseDataRulesInnerWithDefaults instantiates a new GetRulesCompliance200ResponseDataRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetRulesCompliance200ResponseDataRulesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetRulesCompliance200ResponseDataRulesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetRulesCompliance200ResponseDataRulesInner) SetId(v string)`

SetId sets Id field to given value.


### GetMode

`func (o *GetRulesCompliance200ResponseDataRulesInner) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GetRulesCompliance200ResponseDataRulesInner) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GetRulesCompliance200ResponseDataRulesInner) SetMode(v string)`

SetMode sets Mode field to given value.


### GetCompliance

`func (o *GetRulesCompliance200ResponseDataRulesInner) GetCompliance() float32`

GetCompliance returns the Compliance field if non-nil, zero value otherwise.

### GetComplianceOk

`func (o *GetRulesCompliance200ResponseDataRulesInner) GetComplianceOk() (*float32, bool)`

GetComplianceOk returns a tuple with the Compliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliance

`func (o *GetRulesCompliance200ResponseDataRulesInner) SetCompliance(v float32)`

SetCompliance sets Compliance field to given value.


### GetComplianceDetails

`func (o *GetRulesCompliance200ResponseDataRulesInner) GetComplianceDetails() GetGlobalCompliance200ResponseDataGlobalComplianceComplianceDetails`

GetComplianceDetails returns the ComplianceDetails field if non-nil, zero value otherwise.

### GetComplianceDetailsOk

`func (o *GetRulesCompliance200ResponseDataRulesInner) GetComplianceDetailsOk() (*GetGlobalCompliance200ResponseDataGlobalComplianceComplianceDetails, bool)`

GetComplianceDetailsOk returns a tuple with the ComplianceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceDetails

`func (o *GetRulesCompliance200ResponseDataRulesInner) SetComplianceDetails(v GetGlobalCompliance200ResponseDataGlobalComplianceComplianceDetails)`

SetComplianceDetails sets ComplianceDetails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


