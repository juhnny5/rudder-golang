# CveDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | CVE id | [optional] 
**PublishedDate** | Pointer to **string** | Date the CVE was published | [optional] 
**LastModifiedDate** | Pointer to **string** | last Date the CVE was modified | [optional] 
**Description** | Pointer to **string** | CVE Description | [optional] 
**CweIds** | Pointer to **[]string** | List of CWE (Common Weakness Enumeration) of the CVE | [optional] 
**Cvssv3** | Pointer to [**CveDetailsCvssv3**](CveDetailsCvssv3.md) |  | [optional] 
**Cvssv2** | Pointer to [**CveDetailsCvssv2**](CveDetailsCvssv2.md) |  | [optional] 

## Methods

### NewCveDetails

`func NewCveDetails() *CveDetails`

NewCveDetails instantiates a new CveDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCveDetailsWithDefaults

`func NewCveDetailsWithDefaults() *CveDetails`

NewCveDetailsWithDefaults instantiates a new CveDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CveDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CveDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CveDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CveDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPublishedDate

`func (o *CveDetails) GetPublishedDate() string`

GetPublishedDate returns the PublishedDate field if non-nil, zero value otherwise.

### GetPublishedDateOk

`func (o *CveDetails) GetPublishedDateOk() (*string, bool)`

GetPublishedDateOk returns a tuple with the PublishedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedDate

`func (o *CveDetails) SetPublishedDate(v string)`

SetPublishedDate sets PublishedDate field to given value.

### HasPublishedDate

`func (o *CveDetails) HasPublishedDate() bool`

HasPublishedDate returns a boolean if a field has been set.

### GetLastModifiedDate

`func (o *CveDetails) GetLastModifiedDate() string`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *CveDetails) GetLastModifiedDateOk() (*string, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *CveDetails) SetLastModifiedDate(v string)`

SetLastModifiedDate sets LastModifiedDate field to given value.

### HasLastModifiedDate

`func (o *CveDetails) HasLastModifiedDate() bool`

HasLastModifiedDate returns a boolean if a field has been set.

### GetDescription

`func (o *CveDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CveDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CveDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CveDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCweIds

`func (o *CveDetails) GetCweIds() []string`

GetCweIds returns the CweIds field if non-nil, zero value otherwise.

### GetCweIdsOk

`func (o *CveDetails) GetCweIdsOk() (*[]string, bool)`

GetCweIdsOk returns a tuple with the CweIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCweIds

`func (o *CveDetails) SetCweIds(v []string)`

SetCweIds sets CweIds field to given value.

### HasCweIds

`func (o *CveDetails) HasCweIds() bool`

HasCweIds returns a boolean if a field has been set.

### GetCvssv3

`func (o *CveDetails) GetCvssv3() CveDetailsCvssv3`

GetCvssv3 returns the Cvssv3 field if non-nil, zero value otherwise.

### GetCvssv3Ok

`func (o *CveDetails) GetCvssv3Ok() (*CveDetailsCvssv3, bool)`

GetCvssv3Ok returns a tuple with the Cvssv3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssv3

`func (o *CveDetails) SetCvssv3(v CveDetailsCvssv3)`

SetCvssv3 sets Cvssv3 field to given value.

### HasCvssv3

`func (o *CveDetails) HasCvssv3() bool`

HasCvssv3 returns a boolean if a field has been set.

### GetCvssv2

`func (o *CveDetails) GetCvssv2() CveDetailsCvssv2`

GetCvssv2 returns the Cvssv2 field if non-nil, zero value otherwise.

### GetCvssv2Ok

`func (o *CveDetails) GetCvssv2Ok() (*CveDetailsCvssv2, bool)`

GetCvssv2Ok returns a tuple with the Cvssv2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssv2

`func (o *CveDetails) SetCvssv2(v CveDetailsCvssv2)`

SetCvssv2 sets Cvssv2 field to given value.

### HasCvssv2

`func (o *CveDetails) HasCvssv2() bool`

HasCvssv2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


