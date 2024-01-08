# GetAllSettings200ResponseDataSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedNetworks** | Pointer to [**[]GetAllSettings200ResponseDataSettingsAllowedNetworksInner**](GetAllSettings200ResponseDataSettingsAllowedNetworksInner.md) | List of allowed networks for each policy server (root and relays) | [optional] 
**GlobalPolicyMode** | Pointer to **string** | Define the default setting for global policy mode | [optional] 
**GlobalPolicyModeOverridable** | Pointer to **bool** | Allow overrides on this default setting | [optional] 
**RunFrequency** | Pointer to **int32** | Agent run schedule - time between agent runs (in minutes) | [optional] 
**FirstRunHour** | Pointer to **int32** | First agent run time - hour | [optional] 
**FirstRunMinute** | Pointer to **int32** | First agent run time - minute | [optional] 
**SplayTime** | Pointer to **int32** | Maximum delay after scheduled run time (random interval) | [optional] 
**ModifiedFileTtl** | Pointer to **int32** | Number of days to retain modified files | [optional] 
**OutputFileTtl** | Pointer to **int32** | Number of days to retain agent output files | [optional] 
**RequireTimeSynchronization** | Pointer to **bool** | Require time synchronization between nodes and policy server | [optional] 
**RelayServerSynchronizationMethod** | Pointer to **string** | Method used to synchronize data between server and relays, either \&quot;classic\&quot; (agent protocol, default), \&quot;rsync\&quot; (use rsync to synchronize, that you&#39;ll need to be manually set up), or \&quot;disabled\&quot; (use third party system to transmit data) | [optional] 
**RelayServerSynchronizePolicies** | Pointer to **bool** | If **rsync** is set as a synchronization method, use rsync to synchronize policies between Rudder server and relays. If false, you&#39;ll have to synchronize policies yourself. | [optional] 
**RelayServerSynchronizeSharedFiles** | Pointer to **bool** | If **rsync** is set as a synchronization method, use rsync to synchronize shared files between Rudder server and relays. If false, you&#39;ll have to synchronize shared files yourself. | [optional] 
**RudderReportProtocolDefault** | Pointer to **string** | Default reporting protocol used | [optional] 
**ReportingMode** | Pointer to **string** | Compliance reporting mode | [optional] 
**HeartbeatFrequency** | Pointer to **int32** | Send heartbeat every heartbeat_frequency runs (only on **changes-only** compliance mode) | [optional] 
**EnableChangeMessage** | Pointer to **bool** | Enable change audit logs | [optional] 
**MandatoryChangeMessage** | Pointer to **bool** | Make message mandatory | [optional] 
**ChangeMessagePrompt** | Pointer to **string** | Explanation to display | [optional] 
**EnableChangeRequest** | Pointer to **bool** | Enable Change Requests | [optional] 
**EnableSelfValidation** | Pointer to **bool** | Allow self validation | [optional] 
**EnableSelfDeployment** | Pointer to **bool** | Allow self deployment | [optional] 
**DisplayRecentChangesGraphs** | Pointer to **bool** | Display changes graphs | [optional] 
**EnableJavascriptDirectives** | Pointer to **string** | Enable script evaluation in Directives | [optional] 
**SendMetrics** | Pointer to **string** | Send anonymous usage statistics | [optional] 
**NodeAcceptDuplicatedHostname** | Pointer to **bool** | Allow acceptation of a pending node when another one with the same hostname is already accepted | [optional] [default to false]
**NodeOnacceptDefaultState** | Pointer to **string** | Set default state for node when they are accepted within Rudder | [optional] 
**NodeOnacceptDefaultPolicyMode** | Pointer to **string** | Default policy mode for accepted node | [optional] 
**UnexpectedUnboundVarValues** | Pointer to **bool** | Allows multiple reports for configuration based on a multivalued variable | [optional] [default to true]
**RudderComputeChanges** | Pointer to **bool** | Compute list of changes (repaired reports) per rule | [optional] [default to true]
**RudderGenerationComputeDyngroups** | Pointer to **bool** | Recompute all dynamic groups at the start of policy generation | [optional] [default to true]
**RudderComputeDyngroupsMaxParallelism** | Pointer to **string** | Set the parallelism to compute dynamic group, as a number of thread (i.e. 4), or a multiplicative of the number of core (x0.5) | [optional] [default to "1"]
**RudderSaveDbComplianceLevels** | Pointer to **bool** | Store all compliance levels in database | [optional] [default to true]
**RudderSaveDbComplianceDetails** | Pointer to **bool** | Store all compliance details in database | [optional] [default to false]
**RudderGenerationMaxParallelism** | Pointer to **string** | Set the policy generation parallelism, either as an number of thread (i.e. 4), or a multiplicative of the number of core (x0.5) | [optional] [default to "x0.5"]
**RudderGenerationJsTimeout** | Pointer to **int32** | Policy generation JS evaluation of directive parameter timeout in seconds | [optional] [default to 30]
**RudderGenerationContinueOnError** | Pointer to **bool** | Policy generation continues on error during NodeConfiguration evaluation | [optional] [default to false]
**RudderGenerationDelay** | Pointer to **string** | Set a delay before starting policy generation, this will allow you to accumulate changes before they are deployed to Nodes, and can also lessen webapp resources needs. The value is a number followed by the time unit needed (seconds/s, minutes/m, hours/h ...), ie \&quot;5m\&quot; for 5 minutes | [optional] [default to "0 seconds"]
**RudderGenerationPolicy** | Pointer to **string** | Should policy generation be triggered automatically after a change (value &#39;all&#39;), or should we wait until a manual launch (through api or UI, value &#39;onlyManual&#39;) or even no policy generation at all (value \&quot;none\&quot;) | [optional] [default to "all"]

## Methods

### NewGetAllSettings200ResponseDataSettings

`func NewGetAllSettings200ResponseDataSettings() *GetAllSettings200ResponseDataSettings`

NewGetAllSettings200ResponseDataSettings instantiates a new GetAllSettings200ResponseDataSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllSettings200ResponseDataSettingsWithDefaults

`func NewGetAllSettings200ResponseDataSettingsWithDefaults() *GetAllSettings200ResponseDataSettings`

NewGetAllSettings200ResponseDataSettingsWithDefaults instantiates a new GetAllSettings200ResponseDataSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedNetworks

`func (o *GetAllSettings200ResponseDataSettings) GetAllowedNetworks() []GetAllSettings200ResponseDataSettingsAllowedNetworksInner`

GetAllowedNetworks returns the AllowedNetworks field if non-nil, zero value otherwise.

### GetAllowedNetworksOk

`func (o *GetAllSettings200ResponseDataSettings) GetAllowedNetworksOk() (*[]GetAllSettings200ResponseDataSettingsAllowedNetworksInner, bool)`

GetAllowedNetworksOk returns a tuple with the AllowedNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNetworks

`func (o *GetAllSettings200ResponseDataSettings) SetAllowedNetworks(v []GetAllSettings200ResponseDataSettingsAllowedNetworksInner)`

SetAllowedNetworks sets AllowedNetworks field to given value.

### HasAllowedNetworks

`func (o *GetAllSettings200ResponseDataSettings) HasAllowedNetworks() bool`

HasAllowedNetworks returns a boolean if a field has been set.

### GetGlobalPolicyMode

`func (o *GetAllSettings200ResponseDataSettings) GetGlobalPolicyMode() string`

GetGlobalPolicyMode returns the GlobalPolicyMode field if non-nil, zero value otherwise.

### GetGlobalPolicyModeOk

`func (o *GetAllSettings200ResponseDataSettings) GetGlobalPolicyModeOk() (*string, bool)`

GetGlobalPolicyModeOk returns a tuple with the GlobalPolicyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalPolicyMode

`func (o *GetAllSettings200ResponseDataSettings) SetGlobalPolicyMode(v string)`

SetGlobalPolicyMode sets GlobalPolicyMode field to given value.

### HasGlobalPolicyMode

`func (o *GetAllSettings200ResponseDataSettings) HasGlobalPolicyMode() bool`

HasGlobalPolicyMode returns a boolean if a field has been set.

### GetGlobalPolicyModeOverridable

`func (o *GetAllSettings200ResponseDataSettings) GetGlobalPolicyModeOverridable() bool`

GetGlobalPolicyModeOverridable returns the GlobalPolicyModeOverridable field if non-nil, zero value otherwise.

### GetGlobalPolicyModeOverridableOk

`func (o *GetAllSettings200ResponseDataSettings) GetGlobalPolicyModeOverridableOk() (*bool, bool)`

GetGlobalPolicyModeOverridableOk returns a tuple with the GlobalPolicyModeOverridable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalPolicyModeOverridable

`func (o *GetAllSettings200ResponseDataSettings) SetGlobalPolicyModeOverridable(v bool)`

SetGlobalPolicyModeOverridable sets GlobalPolicyModeOverridable field to given value.

### HasGlobalPolicyModeOverridable

`func (o *GetAllSettings200ResponseDataSettings) HasGlobalPolicyModeOverridable() bool`

HasGlobalPolicyModeOverridable returns a boolean if a field has been set.

### GetRunFrequency

`func (o *GetAllSettings200ResponseDataSettings) GetRunFrequency() int32`

GetRunFrequency returns the RunFrequency field if non-nil, zero value otherwise.

### GetRunFrequencyOk

`func (o *GetAllSettings200ResponseDataSettings) GetRunFrequencyOk() (*int32, bool)`

GetRunFrequencyOk returns a tuple with the RunFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunFrequency

`func (o *GetAllSettings200ResponseDataSettings) SetRunFrequency(v int32)`

SetRunFrequency sets RunFrequency field to given value.

### HasRunFrequency

`func (o *GetAllSettings200ResponseDataSettings) HasRunFrequency() bool`

HasRunFrequency returns a boolean if a field has been set.

### GetFirstRunHour

`func (o *GetAllSettings200ResponseDataSettings) GetFirstRunHour() int32`

GetFirstRunHour returns the FirstRunHour field if non-nil, zero value otherwise.

### GetFirstRunHourOk

`func (o *GetAllSettings200ResponseDataSettings) GetFirstRunHourOk() (*int32, bool)`

GetFirstRunHourOk returns a tuple with the FirstRunHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstRunHour

`func (o *GetAllSettings200ResponseDataSettings) SetFirstRunHour(v int32)`

SetFirstRunHour sets FirstRunHour field to given value.

### HasFirstRunHour

`func (o *GetAllSettings200ResponseDataSettings) HasFirstRunHour() bool`

HasFirstRunHour returns a boolean if a field has been set.

### GetFirstRunMinute

`func (o *GetAllSettings200ResponseDataSettings) GetFirstRunMinute() int32`

GetFirstRunMinute returns the FirstRunMinute field if non-nil, zero value otherwise.

### GetFirstRunMinuteOk

`func (o *GetAllSettings200ResponseDataSettings) GetFirstRunMinuteOk() (*int32, bool)`

GetFirstRunMinuteOk returns a tuple with the FirstRunMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstRunMinute

`func (o *GetAllSettings200ResponseDataSettings) SetFirstRunMinute(v int32)`

SetFirstRunMinute sets FirstRunMinute field to given value.

### HasFirstRunMinute

`func (o *GetAllSettings200ResponseDataSettings) HasFirstRunMinute() bool`

HasFirstRunMinute returns a boolean if a field has been set.

### GetSplayTime

`func (o *GetAllSettings200ResponseDataSettings) GetSplayTime() int32`

GetSplayTime returns the SplayTime field if non-nil, zero value otherwise.

### GetSplayTimeOk

`func (o *GetAllSettings200ResponseDataSettings) GetSplayTimeOk() (*int32, bool)`

GetSplayTimeOk returns a tuple with the SplayTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplayTime

`func (o *GetAllSettings200ResponseDataSettings) SetSplayTime(v int32)`

SetSplayTime sets SplayTime field to given value.

### HasSplayTime

`func (o *GetAllSettings200ResponseDataSettings) HasSplayTime() bool`

HasSplayTime returns a boolean if a field has been set.

### GetModifiedFileTtl

`func (o *GetAllSettings200ResponseDataSettings) GetModifiedFileTtl() int32`

GetModifiedFileTtl returns the ModifiedFileTtl field if non-nil, zero value otherwise.

### GetModifiedFileTtlOk

`func (o *GetAllSettings200ResponseDataSettings) GetModifiedFileTtlOk() (*int32, bool)`

GetModifiedFileTtlOk returns a tuple with the ModifiedFileTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedFileTtl

`func (o *GetAllSettings200ResponseDataSettings) SetModifiedFileTtl(v int32)`

SetModifiedFileTtl sets ModifiedFileTtl field to given value.

### HasModifiedFileTtl

`func (o *GetAllSettings200ResponseDataSettings) HasModifiedFileTtl() bool`

HasModifiedFileTtl returns a boolean if a field has been set.

### GetOutputFileTtl

`func (o *GetAllSettings200ResponseDataSettings) GetOutputFileTtl() int32`

GetOutputFileTtl returns the OutputFileTtl field if non-nil, zero value otherwise.

### GetOutputFileTtlOk

`func (o *GetAllSettings200ResponseDataSettings) GetOutputFileTtlOk() (*int32, bool)`

GetOutputFileTtlOk returns a tuple with the OutputFileTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFileTtl

`func (o *GetAllSettings200ResponseDataSettings) SetOutputFileTtl(v int32)`

SetOutputFileTtl sets OutputFileTtl field to given value.

### HasOutputFileTtl

`func (o *GetAllSettings200ResponseDataSettings) HasOutputFileTtl() bool`

HasOutputFileTtl returns a boolean if a field has been set.

### GetRequireTimeSynchronization

`func (o *GetAllSettings200ResponseDataSettings) GetRequireTimeSynchronization() bool`

GetRequireTimeSynchronization returns the RequireTimeSynchronization field if non-nil, zero value otherwise.

### GetRequireTimeSynchronizationOk

`func (o *GetAllSettings200ResponseDataSettings) GetRequireTimeSynchronizationOk() (*bool, bool)`

GetRequireTimeSynchronizationOk returns a tuple with the RequireTimeSynchronization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireTimeSynchronization

`func (o *GetAllSettings200ResponseDataSettings) SetRequireTimeSynchronization(v bool)`

SetRequireTimeSynchronization sets RequireTimeSynchronization field to given value.

### HasRequireTimeSynchronization

`func (o *GetAllSettings200ResponseDataSettings) HasRequireTimeSynchronization() bool`

HasRequireTimeSynchronization returns a boolean if a field has been set.

### GetRelayServerSynchronizationMethod

`func (o *GetAllSettings200ResponseDataSettings) GetRelayServerSynchronizationMethod() string`

GetRelayServerSynchronizationMethod returns the RelayServerSynchronizationMethod field if non-nil, zero value otherwise.

### GetRelayServerSynchronizationMethodOk

`func (o *GetAllSettings200ResponseDataSettings) GetRelayServerSynchronizationMethodOk() (*string, bool)`

GetRelayServerSynchronizationMethodOk returns a tuple with the RelayServerSynchronizationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayServerSynchronizationMethod

`func (o *GetAllSettings200ResponseDataSettings) SetRelayServerSynchronizationMethod(v string)`

SetRelayServerSynchronizationMethod sets RelayServerSynchronizationMethod field to given value.

### HasRelayServerSynchronizationMethod

`func (o *GetAllSettings200ResponseDataSettings) HasRelayServerSynchronizationMethod() bool`

HasRelayServerSynchronizationMethod returns a boolean if a field has been set.

### GetRelayServerSynchronizePolicies

`func (o *GetAllSettings200ResponseDataSettings) GetRelayServerSynchronizePolicies() bool`

GetRelayServerSynchronizePolicies returns the RelayServerSynchronizePolicies field if non-nil, zero value otherwise.

### GetRelayServerSynchronizePoliciesOk

`func (o *GetAllSettings200ResponseDataSettings) GetRelayServerSynchronizePoliciesOk() (*bool, bool)`

GetRelayServerSynchronizePoliciesOk returns a tuple with the RelayServerSynchronizePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayServerSynchronizePolicies

`func (o *GetAllSettings200ResponseDataSettings) SetRelayServerSynchronizePolicies(v bool)`

SetRelayServerSynchronizePolicies sets RelayServerSynchronizePolicies field to given value.

### HasRelayServerSynchronizePolicies

`func (o *GetAllSettings200ResponseDataSettings) HasRelayServerSynchronizePolicies() bool`

HasRelayServerSynchronizePolicies returns a boolean if a field has been set.

### GetRelayServerSynchronizeSharedFiles

`func (o *GetAllSettings200ResponseDataSettings) GetRelayServerSynchronizeSharedFiles() bool`

GetRelayServerSynchronizeSharedFiles returns the RelayServerSynchronizeSharedFiles field if non-nil, zero value otherwise.

### GetRelayServerSynchronizeSharedFilesOk

`func (o *GetAllSettings200ResponseDataSettings) GetRelayServerSynchronizeSharedFilesOk() (*bool, bool)`

GetRelayServerSynchronizeSharedFilesOk returns a tuple with the RelayServerSynchronizeSharedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayServerSynchronizeSharedFiles

`func (o *GetAllSettings200ResponseDataSettings) SetRelayServerSynchronizeSharedFiles(v bool)`

SetRelayServerSynchronizeSharedFiles sets RelayServerSynchronizeSharedFiles field to given value.

### HasRelayServerSynchronizeSharedFiles

`func (o *GetAllSettings200ResponseDataSettings) HasRelayServerSynchronizeSharedFiles() bool`

HasRelayServerSynchronizeSharedFiles returns a boolean if a field has been set.

### GetRudderReportProtocolDefault

`func (o *GetAllSettings200ResponseDataSettings) GetRudderReportProtocolDefault() string`

GetRudderReportProtocolDefault returns the RudderReportProtocolDefault field if non-nil, zero value otherwise.

### GetRudderReportProtocolDefaultOk

`func (o *GetAllSettings200ResponseDataSettings) GetRudderReportProtocolDefaultOk() (*string, bool)`

GetRudderReportProtocolDefaultOk returns a tuple with the RudderReportProtocolDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRudderReportProtocolDefault

`func (o *GetAllSettings200ResponseDataSettings) SetRudderReportProtocolDefault(v string)`

SetRudderReportProtocolDefault sets RudderReportProtocolDefault field to given value.

### HasRudderReportProtocolDefault

`func (o *GetAllSettings200ResponseDataSettings) HasRudderReportProtocolDefault() bool`

HasRudderReportProtocolDefault returns a boolean if a field has been set.

### GetReportingMode

`func (o *GetAllSettings200ResponseDataSettings) GetReportingMode() string`

GetReportingMode returns the ReportingMode field if non-nil, zero value otherwise.

### GetReportingModeOk

`func (o *GetAllSettings200ResponseDataSettings) GetReportingModeOk() (*string, bool)`

GetReportingModeOk returns a tuple with the ReportingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingMode

`func (o *GetAllSettings200ResponseDataSettings) SetReportingMode(v string)`

SetReportingMode sets ReportingMode field to given value.

### HasReportingMode

`func (o *GetAllSettings200ResponseDataSettings) HasReportingMode() bool`

HasReportingMode returns a boolean if a field has been set.

### GetHeartbeatFrequency

`func (o *GetAllSettings200ResponseDataSettings) GetHeartbeatFrequency() int32`

GetHeartbeatFrequency returns the HeartbeatFrequency field if non-nil, zero value otherwise.

### GetHeartbeatFrequencyOk

`func (o *GetAllSettings200ResponseDataSettings) GetHeartbeatFrequencyOk() (*int32, bool)`

GetHeartbeatFrequencyOk returns a tuple with the HeartbeatFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartbeatFrequency

`func (o *GetAllSettings200ResponseDataSettings) SetHeartbeatFrequency(v int32)`

SetHeartbeatFrequency sets HeartbeatFrequency field to given value.

### HasHeartbeatFrequency

`func (o *GetAllSettings200ResponseDataSettings) HasHeartbeatFrequency() bool`

HasHeartbeatFrequency returns a boolean if a field has been set.

### GetEnableChangeMessage

`func (o *GetAllSettings200ResponseDataSettings) GetEnableChangeMessage() bool`

GetEnableChangeMessage returns the EnableChangeMessage field if non-nil, zero value otherwise.

### GetEnableChangeMessageOk

`func (o *GetAllSettings200ResponseDataSettings) GetEnableChangeMessageOk() (*bool, bool)`

GetEnableChangeMessageOk returns a tuple with the EnableChangeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableChangeMessage

`func (o *GetAllSettings200ResponseDataSettings) SetEnableChangeMessage(v bool)`

SetEnableChangeMessage sets EnableChangeMessage field to given value.

### HasEnableChangeMessage

`func (o *GetAllSettings200ResponseDataSettings) HasEnableChangeMessage() bool`

HasEnableChangeMessage returns a boolean if a field has been set.

### GetMandatoryChangeMessage

`func (o *GetAllSettings200ResponseDataSettings) GetMandatoryChangeMessage() bool`

GetMandatoryChangeMessage returns the MandatoryChangeMessage field if non-nil, zero value otherwise.

### GetMandatoryChangeMessageOk

`func (o *GetAllSettings200ResponseDataSettings) GetMandatoryChangeMessageOk() (*bool, bool)`

GetMandatoryChangeMessageOk returns a tuple with the MandatoryChangeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatoryChangeMessage

`func (o *GetAllSettings200ResponseDataSettings) SetMandatoryChangeMessage(v bool)`

SetMandatoryChangeMessage sets MandatoryChangeMessage field to given value.

### HasMandatoryChangeMessage

`func (o *GetAllSettings200ResponseDataSettings) HasMandatoryChangeMessage() bool`

HasMandatoryChangeMessage returns a boolean if a field has been set.

### GetChangeMessagePrompt

`func (o *GetAllSettings200ResponseDataSettings) GetChangeMessagePrompt() string`

GetChangeMessagePrompt returns the ChangeMessagePrompt field if non-nil, zero value otherwise.

### GetChangeMessagePromptOk

`func (o *GetAllSettings200ResponseDataSettings) GetChangeMessagePromptOk() (*string, bool)`

GetChangeMessagePromptOk returns a tuple with the ChangeMessagePrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeMessagePrompt

`func (o *GetAllSettings200ResponseDataSettings) SetChangeMessagePrompt(v string)`

SetChangeMessagePrompt sets ChangeMessagePrompt field to given value.

### HasChangeMessagePrompt

`func (o *GetAllSettings200ResponseDataSettings) HasChangeMessagePrompt() bool`

HasChangeMessagePrompt returns a boolean if a field has been set.

### GetEnableChangeRequest

`func (o *GetAllSettings200ResponseDataSettings) GetEnableChangeRequest() bool`

GetEnableChangeRequest returns the EnableChangeRequest field if non-nil, zero value otherwise.

### GetEnableChangeRequestOk

`func (o *GetAllSettings200ResponseDataSettings) GetEnableChangeRequestOk() (*bool, bool)`

GetEnableChangeRequestOk returns a tuple with the EnableChangeRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableChangeRequest

`func (o *GetAllSettings200ResponseDataSettings) SetEnableChangeRequest(v bool)`

SetEnableChangeRequest sets EnableChangeRequest field to given value.

### HasEnableChangeRequest

`func (o *GetAllSettings200ResponseDataSettings) HasEnableChangeRequest() bool`

HasEnableChangeRequest returns a boolean if a field has been set.

### GetEnableSelfValidation

`func (o *GetAllSettings200ResponseDataSettings) GetEnableSelfValidation() bool`

GetEnableSelfValidation returns the EnableSelfValidation field if non-nil, zero value otherwise.

### GetEnableSelfValidationOk

`func (o *GetAllSettings200ResponseDataSettings) GetEnableSelfValidationOk() (*bool, bool)`

GetEnableSelfValidationOk returns a tuple with the EnableSelfValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSelfValidation

`func (o *GetAllSettings200ResponseDataSettings) SetEnableSelfValidation(v bool)`

SetEnableSelfValidation sets EnableSelfValidation field to given value.

### HasEnableSelfValidation

`func (o *GetAllSettings200ResponseDataSettings) HasEnableSelfValidation() bool`

HasEnableSelfValidation returns a boolean if a field has been set.

### GetEnableSelfDeployment

`func (o *GetAllSettings200ResponseDataSettings) GetEnableSelfDeployment() bool`

GetEnableSelfDeployment returns the EnableSelfDeployment field if non-nil, zero value otherwise.

### GetEnableSelfDeploymentOk

`func (o *GetAllSettings200ResponseDataSettings) GetEnableSelfDeploymentOk() (*bool, bool)`

GetEnableSelfDeploymentOk returns a tuple with the EnableSelfDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSelfDeployment

`func (o *GetAllSettings200ResponseDataSettings) SetEnableSelfDeployment(v bool)`

SetEnableSelfDeployment sets EnableSelfDeployment field to given value.

### HasEnableSelfDeployment

`func (o *GetAllSettings200ResponseDataSettings) HasEnableSelfDeployment() bool`

HasEnableSelfDeployment returns a boolean if a field has been set.

### GetDisplayRecentChangesGraphs

`func (o *GetAllSettings200ResponseDataSettings) GetDisplayRecentChangesGraphs() bool`

GetDisplayRecentChangesGraphs returns the DisplayRecentChangesGraphs field if non-nil, zero value otherwise.

### GetDisplayRecentChangesGraphsOk

`func (o *GetAllSettings200ResponseDataSettings) GetDisplayRecentChangesGraphsOk() (*bool, bool)`

GetDisplayRecentChangesGraphsOk returns a tuple with the DisplayRecentChangesGraphs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayRecentChangesGraphs

`func (o *GetAllSettings200ResponseDataSettings) SetDisplayRecentChangesGraphs(v bool)`

SetDisplayRecentChangesGraphs sets DisplayRecentChangesGraphs field to given value.

### HasDisplayRecentChangesGraphs

`func (o *GetAllSettings200ResponseDataSettings) HasDisplayRecentChangesGraphs() bool`

HasDisplayRecentChangesGraphs returns a boolean if a field has been set.

### GetEnableJavascriptDirectives

`func (o *GetAllSettings200ResponseDataSettings) GetEnableJavascriptDirectives() string`

GetEnableJavascriptDirectives returns the EnableJavascriptDirectives field if non-nil, zero value otherwise.

### GetEnableJavascriptDirectivesOk

`func (o *GetAllSettings200ResponseDataSettings) GetEnableJavascriptDirectivesOk() (*string, bool)`

GetEnableJavascriptDirectivesOk returns a tuple with the EnableJavascriptDirectives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableJavascriptDirectives

`func (o *GetAllSettings200ResponseDataSettings) SetEnableJavascriptDirectives(v string)`

SetEnableJavascriptDirectives sets EnableJavascriptDirectives field to given value.

### HasEnableJavascriptDirectives

`func (o *GetAllSettings200ResponseDataSettings) HasEnableJavascriptDirectives() bool`

HasEnableJavascriptDirectives returns a boolean if a field has been set.

### GetSendMetrics

`func (o *GetAllSettings200ResponseDataSettings) GetSendMetrics() string`

GetSendMetrics returns the SendMetrics field if non-nil, zero value otherwise.

### GetSendMetricsOk

`func (o *GetAllSettings200ResponseDataSettings) GetSendMetricsOk() (*string, bool)`

GetSendMetricsOk returns a tuple with the SendMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendMetrics

`func (o *GetAllSettings200ResponseDataSettings) SetSendMetrics(v string)`

SetSendMetrics sets SendMetrics field to given value.

### HasSendMetrics

`func (o *GetAllSettings200ResponseDataSettings) HasSendMetrics() bool`

HasSendMetrics returns a boolean if a field has been set.

### GetNodeAcceptDuplicatedHostname

`func (o *GetAllSettings200ResponseDataSettings) GetNodeAcceptDuplicatedHostname() bool`

GetNodeAcceptDuplicatedHostname returns the NodeAcceptDuplicatedHostname field if non-nil, zero value otherwise.

### GetNodeAcceptDuplicatedHostnameOk

`func (o *GetAllSettings200ResponseDataSettings) GetNodeAcceptDuplicatedHostnameOk() (*bool, bool)`

GetNodeAcceptDuplicatedHostnameOk returns a tuple with the NodeAcceptDuplicatedHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeAcceptDuplicatedHostname

`func (o *GetAllSettings200ResponseDataSettings) SetNodeAcceptDuplicatedHostname(v bool)`

SetNodeAcceptDuplicatedHostname sets NodeAcceptDuplicatedHostname field to given value.

### HasNodeAcceptDuplicatedHostname

`func (o *GetAllSettings200ResponseDataSettings) HasNodeAcceptDuplicatedHostname() bool`

HasNodeAcceptDuplicatedHostname returns a boolean if a field has been set.

### GetNodeOnacceptDefaultState

`func (o *GetAllSettings200ResponseDataSettings) GetNodeOnacceptDefaultState() string`

GetNodeOnacceptDefaultState returns the NodeOnacceptDefaultState field if non-nil, zero value otherwise.

### GetNodeOnacceptDefaultStateOk

`func (o *GetAllSettings200ResponseDataSettings) GetNodeOnacceptDefaultStateOk() (*string, bool)`

GetNodeOnacceptDefaultStateOk returns a tuple with the NodeOnacceptDefaultState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeOnacceptDefaultState

`func (o *GetAllSettings200ResponseDataSettings) SetNodeOnacceptDefaultState(v string)`

SetNodeOnacceptDefaultState sets NodeOnacceptDefaultState field to given value.

### HasNodeOnacceptDefaultState

`func (o *GetAllSettings200ResponseDataSettings) HasNodeOnacceptDefaultState() bool`

HasNodeOnacceptDefaultState returns a boolean if a field has been set.

### GetNodeOnacceptDefaultPolicyMode

`func (o *GetAllSettings200ResponseDataSettings) GetNodeOnacceptDefaultPolicyMode() string`

GetNodeOnacceptDefaultPolicyMode returns the NodeOnacceptDefaultPolicyMode field if non-nil, zero value otherwise.

### GetNodeOnacceptDefaultPolicyModeOk

`func (o *GetAllSettings200ResponseDataSettings) GetNodeOnacceptDefaultPolicyModeOk() (*string, bool)`

GetNodeOnacceptDefaultPolicyModeOk returns a tuple with the NodeOnacceptDefaultPolicyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeOnacceptDefaultPolicyMode

`func (o *GetAllSettings200ResponseDataSettings) SetNodeOnacceptDefaultPolicyMode(v string)`

SetNodeOnacceptDefaultPolicyMode sets NodeOnacceptDefaultPolicyMode field to given value.

### HasNodeOnacceptDefaultPolicyMode

`func (o *GetAllSettings200ResponseDataSettings) HasNodeOnacceptDefaultPolicyMode() bool`

HasNodeOnacceptDefaultPolicyMode returns a boolean if a field has been set.

### GetUnexpectedUnboundVarValues

`func (o *GetAllSettings200ResponseDataSettings) GetUnexpectedUnboundVarValues() bool`

GetUnexpectedUnboundVarValues returns the UnexpectedUnboundVarValues field if non-nil, zero value otherwise.

### GetUnexpectedUnboundVarValuesOk

`func (o *GetAllSettings200ResponseDataSettings) GetUnexpectedUnboundVarValuesOk() (*bool, bool)`

GetUnexpectedUnboundVarValuesOk returns a tuple with the UnexpectedUnboundVarValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnexpectedUnboundVarValues

`func (o *GetAllSettings200ResponseDataSettings) SetUnexpectedUnboundVarValues(v bool)`

SetUnexpectedUnboundVarValues sets UnexpectedUnboundVarValues field to given value.

### HasUnexpectedUnboundVarValues

`func (o *GetAllSettings200ResponseDataSettings) HasUnexpectedUnboundVarValues() bool`

HasUnexpectedUnboundVarValues returns a boolean if a field has been set.

### GetRudderComputeChanges

`func (o *GetAllSettings200ResponseDataSettings) GetRudderComputeChanges() bool`

GetRudderComputeChanges returns the RudderComputeChanges field if non-nil, zero value otherwise.

### GetRudderComputeChangesOk

`func (o *GetAllSettings200ResponseDataSettings) GetRudderComputeChangesOk() (*bool, bool)`

GetRudderComputeChangesOk returns a tuple with the RudderComputeChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRudderComputeChanges

`func (o *GetAllSettings200ResponseDataSettings) SetRudderComputeChanges(v bool)`

SetRudderComputeChanges sets RudderComputeChanges field to given value.

### HasRudderComputeChanges

`func (o *GetAllSettings200ResponseDataSettings) HasRudderComputeChanges() bool`

HasRudderComputeChanges returns a boolean if a field has been set.

### GetRudderGenerationComputeDyngroups

`func (o *GetAllSettings200ResponseDataSettings) GetRudderGenerationComputeDyngroups() bool`

GetRudderGenerationComputeDyngroups returns the RudderGenerationComputeDyngroups field if non-nil, zero value otherwise.

### GetRudderGenerationComputeDyngroupsOk

`func (o *GetAllSettings200ResponseDataSettings) GetRudderGenerationComputeDyngroupsOk() (*bool, bool)`

GetRudderGenerationComputeDyngroupsOk returns a tuple with the RudderGenerationComputeDyngroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRudderGenerationComputeDyngroups

`func (o *GetAllSettings200ResponseDataSettings) SetRudderGenerationComputeDyngroups(v bool)`

SetRudderGenerationComputeDyngroups sets RudderGenerationComputeDyngroups field to given value.

### HasRudderGenerationComputeDyngroups

`func (o *GetAllSettings200ResponseDataSettings) HasRudderGenerationComputeDyngroups() bool`

HasRudderGenerationComputeDyngroups returns a boolean if a field has been set.

### GetRudderComputeDyngroupsMaxParallelism

`func (o *GetAllSettings200ResponseDataSettings) GetRudderComputeDyngroupsMaxParallelism() string`

GetRudderComputeDyngroupsMaxParallelism returns the RudderComputeDyngroupsMaxParallelism field if non-nil, zero value otherwise.

### GetRudderComputeDyngroupsMaxParallelismOk

`func (o *GetAllSettings200ResponseDataSettings) GetRudderComputeDyngroupsMaxParallelismOk() (*string, bool)`

GetRudderComputeDyngroupsMaxParallelismOk returns a tuple with the RudderComputeDyngroupsMaxParallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRudderComputeDyngroupsMaxParallelism

`func (o *GetAllSettings200ResponseDataSettings) SetRudderComputeDyngroupsMaxParallelism(v string)`

SetRudderComputeDyngroupsMaxParallelism sets RudderComputeDyngroupsMaxParallelism field to given value.

### HasRudderComputeDyngroupsMaxParallelism

`func (o *GetAllSettings200ResponseDataSettings) HasRudderComputeDyngroupsMaxParallelism() bool`

HasRudderComputeDyngroupsMaxParallelism returns a boolean if a field has been set.

### GetRudderSaveDbComplianceLevels

`func (o *GetAllSettings200ResponseDataSettings) GetRudderSaveDbComplianceLevels() bool`

GetRudderSaveDbComplianceLevels returns the RudderSaveDbComplianceLevels field if non-nil, zero value otherwise.

### GetRudderSaveDbComplianceLevelsOk

`func (o *GetAllSettings200ResponseDataSettings) GetRudderSaveDbComplianceLevelsOk() (*bool, bool)`

GetRudderSaveDbComplianceLevelsOk returns a tuple with the RudderSaveDbComplianceLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRudderSaveDbComplianceLevels

`func (o *GetAllSettings200ResponseDataSettings) SetRudderSaveDbComplianceLevels(v bool)`

SetRudderSaveDbComplianceLevels sets RudderSaveDbComplianceLevels field to given value.

### HasRudderSaveDbComplianceLevels

`func (o *GetAllSettings200ResponseDataSettings) HasRudderSaveDbComplianceLevels() bool`

HasRudderSaveDbComplianceLevels returns a boolean if a field has been set.

### GetRudderSaveDbComplianceDetails

`func (o *GetAllSettings200ResponseDataSettings) GetRudderSaveDbComplianceDetails() bool`

GetRudderSaveDbComplianceDetails returns the RudderSaveDbComplianceDetails field if non-nil, zero value otherwise.

### GetRudderSaveDbComplianceDetailsOk

`func (o *GetAllSettings200ResponseDataSettings) GetRudderSaveDbComplianceDetailsOk() (*bool, bool)`

GetRudderSaveDbComplianceDetailsOk returns a tuple with the RudderSaveDbComplianceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRudderSaveDbComplianceDetails

`func (o *GetAllSettings200ResponseDataSettings) SetRudderSaveDbComplianceDetails(v bool)`

SetRudderSaveDbComplianceDetails sets RudderSaveDbComplianceDetails field to given value.

### HasRudderSaveDbComplianceDetails

`func (o *GetAllSettings200ResponseDataSettings) HasRudderSaveDbComplianceDetails() bool`

HasRudderSaveDbComplianceDetails returns a boolean if a field has been set.

### GetRudderGenerationMaxParallelism

`func (o *GetAllSettings200ResponseDataSettings) GetRudderGenerationMaxParallelism() string`

GetRudderGenerationMaxParallelism returns the RudderGenerationMaxParallelism field if non-nil, zero value otherwise.

### GetRudderGenerationMaxParallelismOk

`func (o *GetAllSettings200ResponseDataSettings) GetRudderGenerationMaxParallelismOk() (*string, bool)`

GetRudderGenerationMaxParallelismOk returns a tuple with the RudderGenerationMaxParallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRudderGenerationMaxParallelism

`func (o *GetAllSettings200ResponseDataSettings) SetRudderGenerationMaxParallelism(v string)`

SetRudderGenerationMaxParallelism sets RudderGenerationMaxParallelism field to given value.

### HasRudderGenerationMaxParallelism

`func (o *GetAllSettings200ResponseDataSettings) HasRudderGenerationMaxParallelism() bool`

HasRudderGenerationMaxParallelism returns a boolean if a field has been set.

### GetRudderGenerationJsTimeout

`func (o *GetAllSettings200ResponseDataSettings) GetRudderGenerationJsTimeout() int32`

GetRudderGenerationJsTimeout returns the RudderGenerationJsTimeout field if non-nil, zero value otherwise.

### GetRudderGenerationJsTimeoutOk

`func (o *GetAllSettings200ResponseDataSettings) GetRudderGenerationJsTimeoutOk() (*int32, bool)`

GetRudderGenerationJsTimeoutOk returns a tuple with the RudderGenerationJsTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRudderGenerationJsTimeout

`func (o *GetAllSettings200ResponseDataSettings) SetRudderGenerationJsTimeout(v int32)`

SetRudderGenerationJsTimeout sets RudderGenerationJsTimeout field to given value.

### HasRudderGenerationJsTimeout

`func (o *GetAllSettings200ResponseDataSettings) HasRudderGenerationJsTimeout() bool`

HasRudderGenerationJsTimeout returns a boolean if a field has been set.

### GetRudderGenerationContinueOnError

`func (o *GetAllSettings200ResponseDataSettings) GetRudderGenerationContinueOnError() bool`

GetRudderGenerationContinueOnError returns the RudderGenerationContinueOnError field if non-nil, zero value otherwise.

### GetRudderGenerationContinueOnErrorOk

`func (o *GetAllSettings200ResponseDataSettings) GetRudderGenerationContinueOnErrorOk() (*bool, bool)`

GetRudderGenerationContinueOnErrorOk returns a tuple with the RudderGenerationContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRudderGenerationContinueOnError

`func (o *GetAllSettings200ResponseDataSettings) SetRudderGenerationContinueOnError(v bool)`

SetRudderGenerationContinueOnError sets RudderGenerationContinueOnError field to given value.

### HasRudderGenerationContinueOnError

`func (o *GetAllSettings200ResponseDataSettings) HasRudderGenerationContinueOnError() bool`

HasRudderGenerationContinueOnError returns a boolean if a field has been set.

### GetRudderGenerationDelay

`func (o *GetAllSettings200ResponseDataSettings) GetRudderGenerationDelay() string`

GetRudderGenerationDelay returns the RudderGenerationDelay field if non-nil, zero value otherwise.

### GetRudderGenerationDelayOk

`func (o *GetAllSettings200ResponseDataSettings) GetRudderGenerationDelayOk() (*string, bool)`

GetRudderGenerationDelayOk returns a tuple with the RudderGenerationDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRudderGenerationDelay

`func (o *GetAllSettings200ResponseDataSettings) SetRudderGenerationDelay(v string)`

SetRudderGenerationDelay sets RudderGenerationDelay field to given value.

### HasRudderGenerationDelay

`func (o *GetAllSettings200ResponseDataSettings) HasRudderGenerationDelay() bool`

HasRudderGenerationDelay returns a boolean if a field has been set.

### GetRudderGenerationPolicy

`func (o *GetAllSettings200ResponseDataSettings) GetRudderGenerationPolicy() string`

GetRudderGenerationPolicy returns the RudderGenerationPolicy field if non-nil, zero value otherwise.

### GetRudderGenerationPolicyOk

`func (o *GetAllSettings200ResponseDataSettings) GetRudderGenerationPolicyOk() (*string, bool)`

GetRudderGenerationPolicyOk returns a tuple with the RudderGenerationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRudderGenerationPolicy

`func (o *GetAllSettings200ResponseDataSettings) SetRudderGenerationPolicy(v string)`

SetRudderGenerationPolicy sets RudderGenerationPolicy field to given value.

### HasRudderGenerationPolicy

`func (o *GetAllSettings200ResponseDataSettings) HasRudderGenerationPolicy() bool`

HasRudderGenerationPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


