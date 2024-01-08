# NodeFullSoftwareUpdateInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | name of software that can be updated | [optional] 
**Version** | Pointer to **string** | available version for software | [optional] 
**Arch** | Pointer to **string** | CPU architecture of the update | [optional] 
**From** | Pointer to **string** | tool that discovered that update | [optional] 
**Kind** | Pointer to **string** | if available, kind of patch provided by that update, else none | [optional] 
**Source** | Pointer to **string** | information about the source providing that update | [optional] 
**Description** | Pointer to **string** | details about the content of the update, if available | [optional] 
**Severity** | Pointer to **string** | if available, the severity of the update | [optional] 
**Ids** | Pointer to **[]string** |  | [optional] 

## Methods

### NewNodeFullSoftwareUpdateInner

`func NewNodeFullSoftwareUpdateInner() *NodeFullSoftwareUpdateInner`

NewNodeFullSoftwareUpdateInner instantiates a new NodeFullSoftwareUpdateInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeFullSoftwareUpdateInnerWithDefaults

`func NewNodeFullSoftwareUpdateInnerWithDefaults() *NodeFullSoftwareUpdateInner`

NewNodeFullSoftwareUpdateInnerWithDefaults instantiates a new NodeFullSoftwareUpdateInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NodeFullSoftwareUpdateInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeFullSoftwareUpdateInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeFullSoftwareUpdateInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NodeFullSoftwareUpdateInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *NodeFullSoftwareUpdateInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NodeFullSoftwareUpdateInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NodeFullSoftwareUpdateInner) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NodeFullSoftwareUpdateInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetArch

`func (o *NodeFullSoftwareUpdateInner) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *NodeFullSoftwareUpdateInner) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *NodeFullSoftwareUpdateInner) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *NodeFullSoftwareUpdateInner) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetFrom

`func (o *NodeFullSoftwareUpdateInner) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *NodeFullSoftwareUpdateInner) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *NodeFullSoftwareUpdateInner) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *NodeFullSoftwareUpdateInner) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetKind

`func (o *NodeFullSoftwareUpdateInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NodeFullSoftwareUpdateInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NodeFullSoftwareUpdateInner) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NodeFullSoftwareUpdateInner) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetSource

`func (o *NodeFullSoftwareUpdateInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *NodeFullSoftwareUpdateInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *NodeFullSoftwareUpdateInner) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *NodeFullSoftwareUpdateInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDescription

`func (o *NodeFullSoftwareUpdateInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NodeFullSoftwareUpdateInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NodeFullSoftwareUpdateInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NodeFullSoftwareUpdateInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSeverity

`func (o *NodeFullSoftwareUpdateInner) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *NodeFullSoftwareUpdateInner) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *NodeFullSoftwareUpdateInner) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *NodeFullSoftwareUpdateInner) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetIds

`func (o *NodeFullSoftwareUpdateInner) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *NodeFullSoftwareUpdateInner) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *NodeFullSoftwareUpdateInner) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *NodeFullSoftwareUpdateInner) HasIds() bool`

HasIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


