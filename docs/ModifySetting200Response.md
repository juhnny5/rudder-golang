# ModifySetting200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the setting | 
**Result** | **string** | Result of the request | 
**Action** | **string** | The id of the action | 
**Data** | [**GetSetting200ResponseData**](GetSetting200ResponseData.md) |  | 

## Methods

### NewModifySetting200Response

`func NewModifySetting200Response(id string, result string, action string, data GetSetting200ResponseData, ) *ModifySetting200Response`

NewModifySetting200Response instantiates a new ModifySetting200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifySetting200ResponseWithDefaults

`func NewModifySetting200ResponseWithDefaults() *ModifySetting200Response`

NewModifySetting200ResponseWithDefaults instantiates a new ModifySetting200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModifySetting200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModifySetting200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModifySetting200Response) SetId(v string)`

SetId sets Id field to given value.


### GetResult

`func (o *ModifySetting200Response) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ModifySetting200Response) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ModifySetting200Response) SetResult(v string)`

SetResult sets Result field to given value.


### GetAction

`func (o *ModifySetting200Response) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ModifySetting200Response) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ModifySetting200Response) SetAction(v string)`

SetAction sets Action field to given value.


### GetData

`func (o *ModifySetting200Response) GetData() GetSetting200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ModifySetting200Response) GetDataOk() (*GetSetting200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ModifySetting200Response) SetData(v GetSetting200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


