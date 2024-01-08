# GetAllowedNetworks200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Target policy server ID | 
**Result** | **string** | Result of the request | 
**Action** | **string** | The id of the action | 
**Data** | [**GetAllowedNetworks200ResponseData**](GetAllowedNetworks200ResponseData.md) |  | 

## Methods

### NewGetAllowedNetworks200Response

`func NewGetAllowedNetworks200Response(id string, result string, action string, data GetAllowedNetworks200ResponseData, ) *GetAllowedNetworks200Response`

NewGetAllowedNetworks200Response instantiates a new GetAllowedNetworks200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllowedNetworks200ResponseWithDefaults

`func NewGetAllowedNetworks200ResponseWithDefaults() *GetAllowedNetworks200Response`

NewGetAllowedNetworks200ResponseWithDefaults instantiates a new GetAllowedNetworks200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetAllowedNetworks200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAllowedNetworks200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAllowedNetworks200Response) SetId(v string)`

SetId sets Id field to given value.


### GetResult

`func (o *GetAllowedNetworks200Response) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetAllowedNetworks200Response) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetAllowedNetworks200Response) SetResult(v string)`

SetResult sets Result field to given value.


### GetAction

`func (o *GetAllowedNetworks200Response) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GetAllowedNetworks200Response) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GetAllowedNetworks200Response) SetAction(v string)`

SetAction sets Action field to given value.


### GetData

`func (o *GetAllowedNetworks200Response) GetData() GetAllowedNetworks200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAllowedNetworks200Response) GetDataOk() (*GetAllowedNetworks200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAllowedNetworks200Response) SetData(v GetAllowedNetworks200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


