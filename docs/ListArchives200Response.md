# ListArchives200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Result of the request | 
**Action** | **string** | The kind of the archive | 
**Data** | [**ListArchives200ResponseData**](ListArchives200ResponseData.md) |  | 

## Methods

### NewListArchives200Response

`func NewListArchives200Response(result string, action string, data ListArchives200ResponseData, ) *ListArchives200Response`

NewListArchives200Response instantiates a new ListArchives200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListArchives200ResponseWithDefaults

`func NewListArchives200ResponseWithDefaults() *ListArchives200Response`

NewListArchives200ResponseWithDefaults instantiates a new ListArchives200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ListArchives200Response) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ListArchives200Response) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ListArchives200Response) SetResult(v string)`

SetResult sets Result field to given value.


### GetAction

`func (o *ListArchives200Response) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ListArchives200Response) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ListArchives200Response) SetAction(v string)`

SetAction sets Action field to given value.


### GetData

`func (o *ListArchives200Response) GetData() ListArchives200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListArchives200Response) GetDataOk() (*ListArchives200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListArchives200Response) SetData(v ListArchives200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

