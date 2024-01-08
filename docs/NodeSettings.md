# NodeSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Properties** | Pointer to [**[]NodeFullPropertiesInner**](NodeFullPropertiesInner.md) |  | [optional] 
**PolicyMode** | Pointer to **string** | In which mode the node will apply its configuration policy. Use &#x60;default&#x60; to use the global mode. | [optional] 
**State** | Pointer to **string** | The node life cycle state. See [dedicated doc](https://docs.rudder.io/reference/current/usage/advanced_node_management.html#node-lifecycle) for more information. | [optional] 
**AgentKey** | Pointer to [**AgentKey**](AgentKey.md) |  | [optional] 

## Methods

### NewNodeSettings

`func NewNodeSettings() *NodeSettings`

NewNodeSettings instantiates a new NodeSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeSettingsWithDefaults

`func NewNodeSettingsWithDefaults() *NodeSettings`

NewNodeSettingsWithDefaults instantiates a new NodeSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *NodeSettings) GetProperties() []NodeFullPropertiesInner`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *NodeSettings) GetPropertiesOk() (*[]NodeFullPropertiesInner, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *NodeSettings) SetProperties(v []NodeFullPropertiesInner)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *NodeSettings) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetPolicyMode

`func (o *NodeSettings) GetPolicyMode() string`

GetPolicyMode returns the PolicyMode field if non-nil, zero value otherwise.

### GetPolicyModeOk

`func (o *NodeSettings) GetPolicyModeOk() (*string, bool)`

GetPolicyModeOk returns a tuple with the PolicyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyMode

`func (o *NodeSettings) SetPolicyMode(v string)`

SetPolicyMode sets PolicyMode field to given value.

### HasPolicyMode

`func (o *NodeSettings) HasPolicyMode() bool`

HasPolicyMode returns a boolean if a field has been set.

### GetState

`func (o *NodeSettings) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NodeSettings) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NodeSettings) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NodeSettings) HasState() bool`

HasState returns a boolean if a field has been set.

### GetAgentKey

`func (o *NodeSettings) GetAgentKey() AgentKey`

GetAgentKey returns the AgentKey field if non-nil, zero value otherwise.

### GetAgentKeyOk

`func (o *NodeSettings) GetAgentKeyOk() (*AgentKey, bool)`

GetAgentKeyOk returns a tuple with the AgentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentKey

`func (o *NodeSettings) SetAgentKey(v AgentKey)`

SetAgentKey sets AgentKey field to given value.

### HasAgentKey

`func (o *NodeSettings) HasAgentKey() bool`

HasAgentKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


