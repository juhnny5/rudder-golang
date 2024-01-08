# GetGlobalCompliance200ResponseDataGlobalCompliance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compliance** | **float32** | Global compliance level (&#x60;-1&#x60; when no policies are defined) | 
**ComplianceDetails** | Pointer to [**GetGlobalCompliance200ResponseDataGlobalComplianceComplianceDetails**](GetGlobalCompliance200ResponseDataGlobalComplianceComplianceDetails.md) |  | [optional] 

## Methods

### NewGetGlobalCompliance200ResponseDataGlobalCompliance

`func NewGetGlobalCompliance200ResponseDataGlobalCompliance(compliance float32, ) *GetGlobalCompliance200ResponseDataGlobalCompliance`

NewGetGlobalCompliance200ResponseDataGlobalCompliance instantiates a new GetGlobalCompliance200ResponseDataGlobalCompliance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGlobalCompliance200ResponseDataGlobalComplianceWithDefaults

`func NewGetGlobalCompliance200ResponseDataGlobalComplianceWithDefaults() *GetGlobalCompliance200ResponseDataGlobalCompliance`

NewGetGlobalCompliance200ResponseDataGlobalComplianceWithDefaults instantiates a new GetGlobalCompliance200ResponseDataGlobalCompliance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompliance

`func (o *GetGlobalCompliance200ResponseDataGlobalCompliance) GetCompliance() float32`

GetCompliance returns the Compliance field if non-nil, zero value otherwise.

### GetComplianceOk

`func (o *GetGlobalCompliance200ResponseDataGlobalCompliance) GetComplianceOk() (*float32, bool)`

GetComplianceOk returns a tuple with the Compliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliance

`func (o *GetGlobalCompliance200ResponseDataGlobalCompliance) SetCompliance(v float32)`

SetCompliance sets Compliance field to given value.


### GetComplianceDetails

`func (o *GetGlobalCompliance200ResponseDataGlobalCompliance) GetComplianceDetails() GetGlobalCompliance200ResponseDataGlobalComplianceComplianceDetails`

GetComplianceDetails returns the ComplianceDetails field if non-nil, zero value otherwise.

### GetComplianceDetailsOk

`func (o *GetGlobalCompliance200ResponseDataGlobalCompliance) GetComplianceDetailsOk() (*GetGlobalCompliance200ResponseDataGlobalComplianceComplianceDetails, bool)`

GetComplianceDetailsOk returns a tuple with the ComplianceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceDetails

`func (o *GetGlobalCompliance200ResponseDataGlobalCompliance) SetComplianceDetails(v GetGlobalCompliance200ResponseDataGlobalComplianceComplianceDetails)`

SetComplianceDetails sets ComplianceDetails field to given value.

### HasComplianceDetails

`func (o *GetGlobalCompliance200ResponseDataGlobalCompliance) HasComplianceDetails() bool`

HasComplianceDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


