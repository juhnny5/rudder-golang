# NodeFullFileSystemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Type of file system (&#x60;ext4&#x60;, &#x60;swap&#x60;, etc.) | [optional] 
**MountPoint** | Pointer to **string** | Mount point | [optional] 
**Description** | Pointer to **string** | Description of the file system | [optional] 
**FileCount** | Pointer to **int32** | Number of files | [optional] 
**FreeSpace** | Pointer to **int32** | Free space remaining | [optional] 
**TotalSpace** | Pointer to **int32** | Total space | [optional] 

## Methods

### NewNodeFullFileSystemsInner

`func NewNodeFullFileSystemsInner() *NodeFullFileSystemsInner`

NewNodeFullFileSystemsInner instantiates a new NodeFullFileSystemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeFullFileSystemsInnerWithDefaults

`func NewNodeFullFileSystemsInnerWithDefaults() *NodeFullFileSystemsInner`

NewNodeFullFileSystemsInnerWithDefaults instantiates a new NodeFullFileSystemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NodeFullFileSystemsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeFullFileSystemsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeFullFileSystemsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NodeFullFileSystemsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMountPoint

`func (o *NodeFullFileSystemsInner) GetMountPoint() string`

GetMountPoint returns the MountPoint field if non-nil, zero value otherwise.

### GetMountPointOk

`func (o *NodeFullFileSystemsInner) GetMountPointOk() (*string, bool)`

GetMountPointOk returns a tuple with the MountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoint

`func (o *NodeFullFileSystemsInner) SetMountPoint(v string)`

SetMountPoint sets MountPoint field to given value.

### HasMountPoint

`func (o *NodeFullFileSystemsInner) HasMountPoint() bool`

HasMountPoint returns a boolean if a field has been set.

### GetDescription

`func (o *NodeFullFileSystemsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NodeFullFileSystemsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NodeFullFileSystemsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NodeFullFileSystemsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFileCount

`func (o *NodeFullFileSystemsInner) GetFileCount() int32`

GetFileCount returns the FileCount field if non-nil, zero value otherwise.

### GetFileCountOk

`func (o *NodeFullFileSystemsInner) GetFileCountOk() (*int32, bool)`

GetFileCountOk returns a tuple with the FileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCount

`func (o *NodeFullFileSystemsInner) SetFileCount(v int32)`

SetFileCount sets FileCount field to given value.

### HasFileCount

`func (o *NodeFullFileSystemsInner) HasFileCount() bool`

HasFileCount returns a boolean if a field has been set.

### GetFreeSpace

`func (o *NodeFullFileSystemsInner) GetFreeSpace() int32`

GetFreeSpace returns the FreeSpace field if non-nil, zero value otherwise.

### GetFreeSpaceOk

`func (o *NodeFullFileSystemsInner) GetFreeSpaceOk() (*int32, bool)`

GetFreeSpaceOk returns a tuple with the FreeSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpace

`func (o *NodeFullFileSystemsInner) SetFreeSpace(v int32)`

SetFreeSpace sets FreeSpace field to given value.

### HasFreeSpace

`func (o *NodeFullFileSystemsInner) HasFreeSpace() bool`

HasFreeSpace returns a boolean if a field has been set.

### GetTotalSpace

`func (o *NodeFullFileSystemsInner) GetTotalSpace() int32`

GetTotalSpace returns the TotalSpace field if non-nil, zero value otherwise.

### GetTotalSpaceOk

`func (o *NodeFullFileSystemsInner) GetTotalSpaceOk() (*int32, bool)`

GetTotalSpaceOk returns a tuple with the TotalSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSpace

`func (o *NodeFullFileSystemsInner) SetTotalSpace(v int32)`

SetTotalSpace sets TotalSpace field to given value.

### HasTotalSpace

`func (o *NodeFullFileSystemsInner) HasTotalSpace() bool`

HasTotalSpace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


