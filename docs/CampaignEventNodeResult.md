# CampaignEventNodeResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Campaign event id | [optional] 
**Nodes** | Pointer to [**[]CampaignEventNodeResultNodesInner**](CampaignEventNodeResultNodesInner.md) |  | [optional] 

## Methods

### NewCampaignEventNodeResult

`func NewCampaignEventNodeResult() *CampaignEventNodeResult`

NewCampaignEventNodeResult instantiates a new CampaignEventNodeResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignEventNodeResultWithDefaults

`func NewCampaignEventNodeResultWithDefaults() *CampaignEventNodeResult`

NewCampaignEventNodeResultWithDefaults instantiates a new CampaignEventNodeResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CampaignEventNodeResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignEventNodeResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignEventNodeResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CampaignEventNodeResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNodes

`func (o *CampaignEventNodeResult) GetNodes() []CampaignEventNodeResultNodesInner`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *CampaignEventNodeResult) GetNodesOk() (*[]CampaignEventNodeResultNodesInner, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *CampaignEventNodeResult) SetNodes(v []CampaignEventNodeResultNodesInner)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *CampaignEventNodeResult) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


