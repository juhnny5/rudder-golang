# DeleteParameter500Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the global property | 
**Result** | **string** | Result of the request | 
**Action** | **string** | The id of the action | 
**ErrorDetails** | Pointer to **string** |  | [optional] 

## Methods

### NewDeleteParameter500Response

`func NewDeleteParameter500Response(id string, result string, action string, ) *DeleteParameter500Response`

NewDeleteParameter500Response instantiates a new DeleteParameter500Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteParameter500ResponseWithDefaults

`func NewDeleteParameter500ResponseWithDefaults() *DeleteParameter500Response`

NewDeleteParameter500ResponseWithDefaults instantiates a new DeleteParameter500Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeleteParameter500Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeleteParameter500Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeleteParameter500Response) SetId(v string)`

SetId sets Id field to given value.


### GetResult

`func (o *DeleteParameter500Response) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DeleteParameter500Response) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DeleteParameter500Response) SetResult(v string)`

SetResult sets Result field to given value.


### GetAction

`func (o *DeleteParameter500Response) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DeleteParameter500Response) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DeleteParameter500Response) SetAction(v string)`

SetAction sets Action field to given value.


### GetErrorDetails

`func (o *DeleteParameter500Response) GetErrorDetails() string`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *DeleteParameter500Response) GetErrorDetailsOk() (*string, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *DeleteParameter500Response) SetErrorDetails(v string)`

SetErrorDetails sets ErrorDetails field to given value.

### HasErrorDetails

`func (o *DeleteParameter500Response) HasErrorDetails() bool`

HasErrorDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


