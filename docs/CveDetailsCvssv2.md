# CveDetailsCvssv2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseScore** | Pointer to **int32** | CVSS V2 base score | [optional] 
**Vector** | Pointer to **string** | CVSS V2 vector | [optional] 

## Methods

### NewCveDetailsCvssv2

`func NewCveDetailsCvssv2() *CveDetailsCvssv2`

NewCveDetailsCvssv2 instantiates a new CveDetailsCvssv2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCveDetailsCvssv2WithDefaults

`func NewCveDetailsCvssv2WithDefaults() *CveDetailsCvssv2`

NewCveDetailsCvssv2WithDefaults instantiates a new CveDetailsCvssv2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseScore

`func (o *CveDetailsCvssv2) GetBaseScore() int32`

GetBaseScore returns the BaseScore field if non-nil, zero value otherwise.

### GetBaseScoreOk

`func (o *CveDetailsCvssv2) GetBaseScoreOk() (*int32, bool)`

GetBaseScoreOk returns a tuple with the BaseScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseScore

`func (o *CveDetailsCvssv2) SetBaseScore(v int32)`

SetBaseScore sets BaseScore field to given value.

### HasBaseScore

`func (o *CveDetailsCvssv2) HasBaseScore() bool`

HasBaseScore returns a boolean if a field has been set.

### GetVector

`func (o *CveDetailsCvssv2) GetVector() string`

GetVector returns the Vector field if non-nil, zero value otherwise.

### GetVectorOk

`func (o *CveDetailsCvssv2) GetVectorOk() (*string, bool)`

GetVectorOk returns a tuple with the Vector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVector

`func (o *CveDetailsCvssv2) SetVector(v string)`

SetVector sets Vector field to given value.

### HasVector

`func (o *CveDetailsCvssv2) HasVector() bool`

HasVector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


