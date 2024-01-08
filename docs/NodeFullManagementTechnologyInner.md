# NodeFullManagementTechnologyInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Agent name | 
**Version** | Pointer to **string** | Agent version | [optional] 
**Capabilities** | Pointer to **[]string** | List of agent capabilities | [optional] 
**NodeKind** | Pointer to **string** | kind of node for the management engine, like &#x60;root&#x60;, &#x60;relay&#x60;, &#x60;node&#x60;, &#x60;root-component&#x60; | [optional] 
**RootComponents** | Pointer to **[]string** | Roles fulfilled by the agent | [optional] 

## Methods

### NewNodeFullManagementTechnologyInner

`func NewNodeFullManagementTechnologyInner(name string, ) *NodeFullManagementTechnologyInner`

NewNodeFullManagementTechnologyInner instantiates a new NodeFullManagementTechnologyInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeFullManagementTechnologyInnerWithDefaults

`func NewNodeFullManagementTechnologyInnerWithDefaults() *NodeFullManagementTechnologyInner`

NewNodeFullManagementTechnologyInnerWithDefaults instantiates a new NodeFullManagementTechnologyInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NodeFullManagementTechnologyInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeFullManagementTechnologyInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeFullManagementTechnologyInner) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *NodeFullManagementTechnologyInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NodeFullManagementTechnologyInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NodeFullManagementTechnologyInner) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NodeFullManagementTechnologyInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCapabilities

`func (o *NodeFullManagementTechnologyInner) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *NodeFullManagementTechnologyInner) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *NodeFullManagementTechnologyInner) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *NodeFullManagementTechnologyInner) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetNodeKind

`func (o *NodeFullManagementTechnologyInner) GetNodeKind() string`

GetNodeKind returns the NodeKind field if non-nil, zero value otherwise.

### GetNodeKindOk

`func (o *NodeFullManagementTechnologyInner) GetNodeKindOk() (*string, bool)`

GetNodeKindOk returns a tuple with the NodeKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeKind

`func (o *NodeFullManagementTechnologyInner) SetNodeKind(v string)`

SetNodeKind sets NodeKind field to given value.

### HasNodeKind

`func (o *NodeFullManagementTechnologyInner) HasNodeKind() bool`

HasNodeKind returns a boolean if a field has been set.

### GetRootComponents

`func (o *NodeFullManagementTechnologyInner) GetRootComponents() []string`

GetRootComponents returns the RootComponents field if non-nil, zero value otherwise.

### GetRootComponentsOk

`func (o *NodeFullManagementTechnologyInner) GetRootComponentsOk() (*[]string, bool)`

GetRootComponentsOk returns a tuple with the RootComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootComponents

`func (o *NodeFullManagementTechnologyInner) SetRootComponents(v []string)`

SetRootComponents sets RootComponents field to given value.

### HasRootComponents

`func (o *NodeFullManagementTechnologyInner) HasRootComponents() bool`

HasRootComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


