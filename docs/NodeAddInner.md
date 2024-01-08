# NodeAddInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Rudder node unique identifier in /opt/rudder/etc/uuid.hive | 
**Hostname** | **string** | The fully qualified name of the node | 
**Status** | **string** | Target status of the node | 
**Os** | [**Os**](Os.md) |  | 
**PolicyServerId** | Pointer to **string** | The policy server ID for that node. By default, \&quot;root\&quot; | [optional] 
**MachineType** | **string** | The kind of machine for the node (use vm for a generic VM) | 
**State** | Pointer to **string** | Node lifecycle state. Can only be specified when status&#x3D;accepted. If not specified, enable is used | [optional] 
**PolicyMode** | Pointer to **string** | The policy mode for the node. Can only be specified when status&#x3D;accepted. If not specified, the default (global) mode will be used | [optional] 
**AgentKey** | Pointer to [**AgentKey**](AgentKey.md) |  | [optional] 
**Properties** | [**[]NodeFullPropertiesInner**](NodeFullPropertiesInner.md) | Node properties (either set by user or filled by third party sources) | 
**IpAddresses** | **[]string** | an array of IPs. | 
**Timezone** | Pointer to [**Timezone**](Timezone.md) |  | [optional] 

## Methods

### NewNodeAddInner

`func NewNodeAddInner(id string, hostname string, status string, os Os, machineType string, properties []NodeFullPropertiesInner, ipAddresses []string, ) *NodeAddInner`

NewNodeAddInner instantiates a new NodeAddInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeAddInnerWithDefaults

`func NewNodeAddInnerWithDefaults() *NodeAddInner`

NewNodeAddInnerWithDefaults instantiates a new NodeAddInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NodeAddInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeAddInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeAddInner) SetId(v string)`

SetId sets Id field to given value.


### GetHostname

`func (o *NodeAddInner) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *NodeAddInner) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *NodeAddInner) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetStatus

`func (o *NodeAddInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NodeAddInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NodeAddInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetOs

`func (o *NodeAddInner) GetOs() Os`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *NodeAddInner) GetOsOk() (*Os, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *NodeAddInner) SetOs(v Os)`

SetOs sets Os field to given value.


### GetPolicyServerId

`func (o *NodeAddInner) GetPolicyServerId() string`

GetPolicyServerId returns the PolicyServerId field if non-nil, zero value otherwise.

### GetPolicyServerIdOk

`func (o *NodeAddInner) GetPolicyServerIdOk() (*string, bool)`

GetPolicyServerIdOk returns a tuple with the PolicyServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyServerId

`func (o *NodeAddInner) SetPolicyServerId(v string)`

SetPolicyServerId sets PolicyServerId field to given value.

### HasPolicyServerId

`func (o *NodeAddInner) HasPolicyServerId() bool`

HasPolicyServerId returns a boolean if a field has been set.

### GetMachineType

`func (o *NodeAddInner) GetMachineType() string`

GetMachineType returns the MachineType field if non-nil, zero value otherwise.

### GetMachineTypeOk

`func (o *NodeAddInner) GetMachineTypeOk() (*string, bool)`

GetMachineTypeOk returns a tuple with the MachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineType

`func (o *NodeAddInner) SetMachineType(v string)`

SetMachineType sets MachineType field to given value.


### GetState

`func (o *NodeAddInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NodeAddInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NodeAddInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NodeAddInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPolicyMode

`func (o *NodeAddInner) GetPolicyMode() string`

GetPolicyMode returns the PolicyMode field if non-nil, zero value otherwise.

### GetPolicyModeOk

`func (o *NodeAddInner) GetPolicyModeOk() (*string, bool)`

GetPolicyModeOk returns a tuple with the PolicyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyMode

`func (o *NodeAddInner) SetPolicyMode(v string)`

SetPolicyMode sets PolicyMode field to given value.

### HasPolicyMode

`func (o *NodeAddInner) HasPolicyMode() bool`

HasPolicyMode returns a boolean if a field has been set.

### GetAgentKey

`func (o *NodeAddInner) GetAgentKey() AgentKey`

GetAgentKey returns the AgentKey field if non-nil, zero value otherwise.

### GetAgentKeyOk

`func (o *NodeAddInner) GetAgentKeyOk() (*AgentKey, bool)`

GetAgentKeyOk returns a tuple with the AgentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentKey

`func (o *NodeAddInner) SetAgentKey(v AgentKey)`

SetAgentKey sets AgentKey field to given value.

### HasAgentKey

`func (o *NodeAddInner) HasAgentKey() bool`

HasAgentKey returns a boolean if a field has been set.

### GetProperties

`func (o *NodeAddInner) GetProperties() []NodeFullPropertiesInner`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *NodeAddInner) GetPropertiesOk() (*[]NodeFullPropertiesInner, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *NodeAddInner) SetProperties(v []NodeFullPropertiesInner)`

SetProperties sets Properties field to given value.


### GetIpAddresses

`func (o *NodeAddInner) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *NodeAddInner) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *NodeAddInner) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.


### GetTimezone

`func (o *NodeAddInner) GetTimezone() Timezone`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *NodeAddInner) GetTimezoneOk() (*Timezone, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *NodeAddInner) SetTimezone(v Timezone)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *NodeAddInner) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


