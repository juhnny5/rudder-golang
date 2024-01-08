# CveCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CveId** | Pointer to **string** | CVE id | [optional] 
**Score** | Pointer to [**CveCheckScore**](CveCheckScore.md) |  | [optional] 
**Nodes** | Pointer to **[]string** | Id of Nodes affected by this CVE | [optional] 
**Packages** | Pointer to [**[]CveCheckPackagesInner**](CveCheckPackagesInner.md) | Packages affected by this CVE | [optional] 

## Methods

### NewCveCheck

`func NewCveCheck() *CveCheck`

NewCveCheck instantiates a new CveCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCveCheckWithDefaults

`func NewCveCheckWithDefaults() *CveCheck`

NewCveCheckWithDefaults instantiates a new CveCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCveId

`func (o *CveCheck) GetCveId() string`

GetCveId returns the CveId field if non-nil, zero value otherwise.

### GetCveIdOk

`func (o *CveCheck) GetCveIdOk() (*string, bool)`

GetCveIdOk returns a tuple with the CveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveId

`func (o *CveCheck) SetCveId(v string)`

SetCveId sets CveId field to given value.

### HasCveId

`func (o *CveCheck) HasCveId() bool`

HasCveId returns a boolean if a field has been set.

### GetScore

`func (o *CveCheck) GetScore() CveCheckScore`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *CveCheck) GetScoreOk() (*CveCheckScore, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *CveCheck) SetScore(v CveCheckScore)`

SetScore sets Score field to given value.

### HasScore

`func (o *CveCheck) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetNodes

`func (o *CveCheck) GetNodes() []string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *CveCheck) GetNodesOk() (*[]string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *CveCheck) SetNodes(v []string)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *CveCheck) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetPackages

`func (o *CveCheck) GetPackages() []CveCheckPackagesInner`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *CveCheck) GetPackagesOk() (*[]CveCheckPackagesInner, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *CveCheck) SetPackages(v []CveCheckPackagesInner)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *CveCheck) HasPackages() bool`

HasPackages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


