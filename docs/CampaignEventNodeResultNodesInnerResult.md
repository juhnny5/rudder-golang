# CampaignEventNodeResultNodesInnerResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Campaign result | [optional] 
**SoftwareUpdated** | Pointer to [**[]CampaignEventNodeResultNodesInnerResultSoftwareUpdatedInner**](CampaignEventNodeResultNodesInnerResultSoftwareUpdatedInner.md) | List of updated software | [optional] 
**Output** | Pointer to **string** | campaign standard output | [optional] 
**Errors** | Pointer to **string** | campaign standard errors | [optional] 

## Methods

### NewCampaignEventNodeResultNodesInnerResult

`func NewCampaignEventNodeResultNodesInnerResult() *CampaignEventNodeResultNodesInnerResult`

NewCampaignEventNodeResultNodesInnerResult instantiates a new CampaignEventNodeResultNodesInnerResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignEventNodeResultNodesInnerResultWithDefaults

`func NewCampaignEventNodeResultNodesInnerResultWithDefaults() *CampaignEventNodeResultNodesInnerResult`

NewCampaignEventNodeResultNodesInnerResultWithDefaults instantiates a new CampaignEventNodeResultNodesInnerResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CampaignEventNodeResultNodesInnerResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CampaignEventNodeResultNodesInnerResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CampaignEventNodeResultNodesInnerResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CampaignEventNodeResultNodesInnerResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSoftwareUpdated

`func (o *CampaignEventNodeResultNodesInnerResult) GetSoftwareUpdated() []CampaignEventNodeResultNodesInnerResultSoftwareUpdatedInner`

GetSoftwareUpdated returns the SoftwareUpdated field if non-nil, zero value otherwise.

### GetSoftwareUpdatedOk

`func (o *CampaignEventNodeResultNodesInnerResult) GetSoftwareUpdatedOk() (*[]CampaignEventNodeResultNodesInnerResultSoftwareUpdatedInner, bool)`

GetSoftwareUpdatedOk returns a tuple with the SoftwareUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareUpdated

`func (o *CampaignEventNodeResultNodesInnerResult) SetSoftwareUpdated(v []CampaignEventNodeResultNodesInnerResultSoftwareUpdatedInner)`

SetSoftwareUpdated sets SoftwareUpdated field to given value.

### HasSoftwareUpdated

`func (o *CampaignEventNodeResultNodesInnerResult) HasSoftwareUpdated() bool`

HasSoftwareUpdated returns a boolean if a field has been set.

### GetOutput

`func (o *CampaignEventNodeResultNodesInnerResult) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *CampaignEventNodeResultNodesInnerResult) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *CampaignEventNodeResultNodesInnerResult) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *CampaignEventNodeResultNodesInnerResult) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetErrors

`func (o *CampaignEventNodeResultNodesInnerResult) GetErrors() string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CampaignEventNodeResultNodesInnerResult) GetErrorsOk() (*string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CampaignEventNodeResultNodesInnerResult) SetErrors(v string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CampaignEventNodeResultNodesInnerResult) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


