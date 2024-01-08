# ParameterDetails200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the property | 
**Result** | **string** | Result of the request | 
**Action** | **string** | The id of the action | 
**Data** | [**ListParameters200ResponseData**](ListParameters200ResponseData.md) |  | 

## Methods

### NewParameterDetails200Response

`func NewParameterDetails200Response(id string, result string, action string, data ListParameters200ResponseData, ) *ParameterDetails200Response`

NewParameterDetails200Response instantiates a new ParameterDetails200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterDetails200ResponseWithDefaults

`func NewParameterDetails200ResponseWithDefaults() *ParameterDetails200Response`

NewParameterDetails200ResponseWithDefaults instantiates a new ParameterDetails200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ParameterDetails200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ParameterDetails200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ParameterDetails200Response) SetId(v string)`

SetId sets Id field to given value.


### GetResult

`func (o *ParameterDetails200Response) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ParameterDetails200Response) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ParameterDetails200Response) SetResult(v string)`

SetResult sets Result field to given value.


### GetAction

`func (o *ParameterDetails200Response) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ParameterDetails200Response) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ParameterDetails200Response) SetAction(v string)`

SetAction sets Action field to given value.


### GetData

`func (o *ParameterDetails200Response) GetData() ListParameters200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ParameterDetails200Response) GetDataOk() (*ListParameters200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ParameterDetails200Response) SetData(v ListParameters200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


