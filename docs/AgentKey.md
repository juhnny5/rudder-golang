# AgentKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** | Certificate (or public key for &lt;6.0 agents) used by the agent. Be careful write a \&quot;\\n\&quot; after header line and before footer line, JSON does not keep formatting in string. | 
**Status** | Pointer to **string** | Certification status of the security token (reset to &#x60;undefined&#x60; to trust a new certificate). If &#x60;certified&#x60;, inventory signature check will be enforced | [optional] 

## Methods

### NewAgentKey

`func NewAgentKey(value string, ) *AgentKey`

NewAgentKey instantiates a new AgentKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentKeyWithDefaults

`func NewAgentKeyWithDefaults() *AgentKey`

NewAgentKeyWithDefaults instantiates a new AgentKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *AgentKey) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AgentKey) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AgentKey) SetValue(v string)`

SetValue sets Value field to given value.


### GetStatus

`func (o *AgentKey) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentKey) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentKey) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentKey) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


