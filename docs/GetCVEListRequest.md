# GetCVEListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CveIds** | Pointer to **[]string** |  | [optional] 
**OnlyScore** | Pointer to **bool** | Only send score of the cve, and not the whole detailed list | [optional] [default to false]
**MinScore** | Pointer to **string** | Only send CVE with a score higher than the value | [optional] 
**MaxScore** | Pointer to **string** | Only send CVE with a score lower than the value | [optional] 
**PublishedDate** | Pointer to **string** | Only send CVE with a publication date more recent than the value | [optional] 

## Methods

### NewGetCVEListRequest

`func NewGetCVEListRequest() *GetCVEListRequest`

NewGetCVEListRequest instantiates a new GetCVEListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCVEListRequestWithDefaults

`func NewGetCVEListRequestWithDefaults() *GetCVEListRequest`

NewGetCVEListRequestWithDefaults instantiates a new GetCVEListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCveIds

`func (o *GetCVEListRequest) GetCveIds() []string`

GetCveIds returns the CveIds field if non-nil, zero value otherwise.

### GetCveIdsOk

`func (o *GetCVEListRequest) GetCveIdsOk() (*[]string, bool)`

GetCveIdsOk returns a tuple with the CveIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveIds

`func (o *GetCVEListRequest) SetCveIds(v []string)`

SetCveIds sets CveIds field to given value.

### HasCveIds

`func (o *GetCVEListRequest) HasCveIds() bool`

HasCveIds returns a boolean if a field has been set.

### GetOnlyScore

`func (o *GetCVEListRequest) GetOnlyScore() bool`

GetOnlyScore returns the OnlyScore field if non-nil, zero value otherwise.

### GetOnlyScoreOk

`func (o *GetCVEListRequest) GetOnlyScoreOk() (*bool, bool)`

GetOnlyScoreOk returns a tuple with the OnlyScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyScore

`func (o *GetCVEListRequest) SetOnlyScore(v bool)`

SetOnlyScore sets OnlyScore field to given value.

### HasOnlyScore

`func (o *GetCVEListRequest) HasOnlyScore() bool`

HasOnlyScore returns a boolean if a field has been set.

### GetMinScore

`func (o *GetCVEListRequest) GetMinScore() string`

GetMinScore returns the MinScore field if non-nil, zero value otherwise.

### GetMinScoreOk

`func (o *GetCVEListRequest) GetMinScoreOk() (*string, bool)`

GetMinScoreOk returns a tuple with the MinScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinScore

`func (o *GetCVEListRequest) SetMinScore(v string)`

SetMinScore sets MinScore field to given value.

### HasMinScore

`func (o *GetCVEListRequest) HasMinScore() bool`

HasMinScore returns a boolean if a field has been set.

### GetMaxScore

`func (o *GetCVEListRequest) GetMaxScore() string`

GetMaxScore returns the MaxScore field if non-nil, zero value otherwise.

### GetMaxScoreOk

`func (o *GetCVEListRequest) GetMaxScoreOk() (*string, bool)`

GetMaxScoreOk returns a tuple with the MaxScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxScore

`func (o *GetCVEListRequest) SetMaxScore(v string)`

SetMaxScore sets MaxScore field to given value.

### HasMaxScore

`func (o *GetCVEListRequest) HasMaxScore() bool`

HasMaxScore returns a boolean if a field has been set.

### GetPublishedDate

`func (o *GetCVEListRequest) GetPublishedDate() string`

GetPublishedDate returns the PublishedDate field if non-nil, zero value otherwise.

### GetPublishedDateOk

`func (o *GetCVEListRequest) GetPublishedDateOk() (*string, bool)`

GetPublishedDateOk returns a tuple with the PublishedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedDate

`func (o *GetCVEListRequest) SetPublishedDate(v string)`

SetPublishedDate sets PublishedDate field to given value.

### HasPublishedDate

`func (o *GetCVEListRequest) HasPublishedDate() bool`

HasPublishedDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


