# CveDetailsCvssv3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseScore** | Pointer to **float32** | CVSS V3 base score | [optional] 
**Vector** | Pointer to **string** | CVSS V3 vector | [optional] 
**BaseSeverity** | Pointer to **string** | CVSS V3 Severity | [optional] 

## Methods

### NewCveDetailsCvssv3

`func NewCveDetailsCvssv3() *CveDetailsCvssv3`

NewCveDetailsCvssv3 instantiates a new CveDetailsCvssv3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCveDetailsCvssv3WithDefaults

`func NewCveDetailsCvssv3WithDefaults() *CveDetailsCvssv3`

NewCveDetailsCvssv3WithDefaults instantiates a new CveDetailsCvssv3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseScore

`func (o *CveDetailsCvssv3) GetBaseScore() float32`

GetBaseScore returns the BaseScore field if non-nil, zero value otherwise.

### GetBaseScoreOk

`func (o *CveDetailsCvssv3) GetBaseScoreOk() (*float32, bool)`

GetBaseScoreOk returns a tuple with the BaseScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseScore

`func (o *CveDetailsCvssv3) SetBaseScore(v float32)`

SetBaseScore sets BaseScore field to given value.

### HasBaseScore

`func (o *CveDetailsCvssv3) HasBaseScore() bool`

HasBaseScore returns a boolean if a field has been set.

### GetVector

`func (o *CveDetailsCvssv3) GetVector() string`

GetVector returns the Vector field if non-nil, zero value otherwise.

### GetVectorOk

`func (o *CveDetailsCvssv3) GetVectorOk() (*string, bool)`

GetVectorOk returns a tuple with the Vector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVector

`func (o *CveDetailsCvssv3) SetVector(v string)`

SetVector sets Vector field to given value.

### HasVector

`func (o *CveDetailsCvssv3) HasVector() bool`

HasVector returns a boolean if a field has been set.

### GetBaseSeverity

`func (o *CveDetailsCvssv3) GetBaseSeverity() string`

GetBaseSeverity returns the BaseSeverity field if non-nil, zero value otherwise.

### GetBaseSeverityOk

`func (o *CveDetailsCvssv3) GetBaseSeverityOk() (*string, bool)`

GetBaseSeverityOk returns a tuple with the BaseSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSeverity

`func (o *CveDetailsCvssv3) SetBaseSeverity(v string)`

SetBaseSeverity sets BaseSeverity field to given value.

### HasBaseSeverity

`func (o *CveDetailsCvssv3) HasBaseSeverity() bool`

HasBaseSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


