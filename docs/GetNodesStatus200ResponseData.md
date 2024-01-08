# GetNodesStatus200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | [**[]GetNodesStatus200ResponseDataNodesInner**](GetNodesStatus200ResponseDataNodesInner.md) | List of nodeId and associated status | 

## Methods

### NewGetNodesStatus200ResponseData

`func NewGetNodesStatus200ResponseData(nodes []GetNodesStatus200ResponseDataNodesInner, ) *GetNodesStatus200ResponseData`

NewGetNodesStatus200ResponseData instantiates a new GetNodesStatus200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNodesStatus200ResponseDataWithDefaults

`func NewGetNodesStatus200ResponseDataWithDefaults() *GetNodesStatus200ResponseData`

NewGetNodesStatus200ResponseDataWithDefaults instantiates a new GetNodesStatus200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *GetNodesStatus200ResponseData) GetNodes() []GetNodesStatus200ResponseDataNodesInner`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *GetNodesStatus200ResponseData) GetNodesOk() (*[]GetNodesStatus200ResponseDataNodesInner, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *GetNodesStatus200ResponseData) SetNodes(v []GetNodesStatus200ResponseDataNodesInner)`

SetNodes sets Nodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


