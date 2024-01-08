/*
Rudder API

Download OpenAPI specification: [openapi.yml](openapi.yml)  **Other documentation sources**:  * [Main documentation](https://docs.rudder.io) * [Internal relay API](https://docs.rudder.io/api/relay/)  # Introduction  Rudder exposes a REST API, enabling the user to interact with Rudder without using the webapp, for example, in scripts or cron jobs.  ## Authentication  The Rudder REST API uses simple API keys for authentication. All requests must be authenticated (except from a generic status API). The tokens are 32-character strings, passed in a `X-API-Token` header, like in:   ```bash curl --header \"X-API-Token: yourToken\" https://rudder.example.com/rudder/api/latest/rules ```  The tokens are the API equivalent of a password, and must be secured just like a password would be.  ### API accounts  The accounts are managed in the Web interface. There are two possible types of accounts:  * **Global API accounts**: they are not linked to a Rudder user, and are managed by Rudder administrators in the _Administration -> API accounts_ page. You should define an expiration date whenever possible.  ![General API tokens settings](assets/api-tokens.png \"General API tokens settings\")  * **User tokens**: they are linked to a Rudder user, and give the same rights the user has. There can be only one token by user. This feature is provided by the `api-authorizatons` plugin.  ![User API token](assets/api-user.png \"User API token\")  When an action produces a change of configuration on the server, the API account that made it will be recorded in the event log, like for a Web interaction.  ### Authorization  When using Rudder without the `api-authorizatons` plugin, only global accounts are available, with two possible privilege levels, read-only or write. With the `api-authorizatons` plugin, you also get access to:  * User tokens, which have the same permissions as the user, using the Rudder roles and permissions feature. * Custom ACLs on global API accounts. They provide fine-grained permissions on every endpoint:  ![Custom API ACL](assets/custom-acl.png \"Custom API ACL\")  As a general principle, you should create dedicated tokens with the least privilege level for each different interaction you have with the API. This limits the risks of exploitation if a token is stolen, and allows tracking the activity of each token separately. Token renewal is also easier when they are only used for a limited purpose.  ## Versioning  Each time the API is extended with new features (new functions, new parameters, new responses, ...), it will be assigned a new version number. This will allow you to keep your existing scripts (based on previous behavior). Versions will always be integers (no 2.1 or 3.3, just 2, 3, 4, ...) or `latest`.  You can change the version of the API used by setting it either within the url or in a header:  * the URL: each URL is prefixed by its version id, like `/api/version/function`.  ```bash # Version 10 curl -X GET -H \"X-API-Token: yourToken\" https://rudder.example.com/rudder/api/10/rules # Latest curl -X GET -H \"X-API-Token: yourToken\" https://rudder.example.com/rudder/api/latest/rules # Wrong (not an integer) => 404 not found curl -X GET -H \"X-API-Token: yourToken\" https://rudder.example.com/rudder/api/3.14/rules ```  * the HTTP headers. You can add the **X-API-Version** header to your request. The value needs to be an integer or `latest`.  ```bash # Version 10 curl -X GET -H \"X-API-Token: yourToken\" -H \"X-API-Version: 10\" https://rudder.example.com/rudder/api/rules # Wrong => Error response indicating which versions are available curl -X GET -H \"X-API-Token: yourToken\" -H \"X-API-Version: 3.14\" https://rudder.example.com/rudder/api/rules ```  In the future, we may declare some versions as deprecated, in order to remove them in a later version of Rudder, but we will never remove any versions without warning, or without a safe period of time to allow migration from previous versions.   <h4>Existing versions</h4> <table>   <thead>     <tr>       <th style=\"width: 20%\">Version</th>       <th style=\"width: 20%\">Rudder versions it appeared in</th>       <th style=\"width: 70%\">Description</th>     </tr>   </thead>   <tbody>     <tr>       <td class=\"code\">1</td>       <td class=\"code\">Never released (for internal use only)</td>       <td>Experimental version</td>     </tr>     <tr>       <td class=\"code\">2 to 10 (deprecated)</td>       <td class=\"code\">4.3 and before</td>       <td>These versions provided the core set of API features for rules, directives, nodes global parameters, change requests and compliance, rudder settings, and system API</td>     </tr>     <tr>       <td class=\"code\">11</td>       <td class=\"code\">5.0</td>       <td>New system API (replacing old localhost v1 api): status, maintenance operations and server behavior</td>     </tr>     <tr>       <td class=\"code\">12</td>       <td class=\"code\">6.0 and 6.1</td>       <td>Node key management</td>     </tr>     <tr>       <td class=\"code\">13</td>       <td class=\"code\">6.2</td>       <td><ul>         <li>Node status endpoint</li>         <li>System health check</li>         <li>System maintenance job to purge software [that endpoint was back-ported in 6.1]</li>       </ul></td>     </tr>     <tr>       <td class=\"code\">14</td>       <td class=\"code\">7.0</td>       <td><ul>         <li>Secret management</li>         <li>Directive tree</li>         <li>Improve techniques management</li>         <li>Demote a relay</li>       </ul></td>     </tr>     <tr>       <td class=\"code\">15</td>       <td class=\"code\">7.1</td>       <td><ul>         <li>Package updates in nodes</li>       </ul></td>     </tr>     <tr>       <td class=\"code\">16</td>       <td class=\"code\">7.2</td>       <td><ul>         <li>Create node API included from plugin</li>         <li>Configuration archive import/export</li>       </ul></td>     </tr>     <tr>       <td class=\"code\">17</td>       <td class=\"code\">7.3</td>       <td><ul>         <li>Compliance by directive</li>         <li>Path campaigns API included</li>       </ul></td>     </tr>     <tr>       <td class=\"code\">18</td>       <td class=\"code\">8.0</td>       <td><ul>         <li>Allowed network </li>         <li>Improve the structure of `/settings/allowed_networks` output</li>       </ul></td>     </tr>   </tbody> </table>   ## Response format  All responses from the API are in the JSON format.  ```json {   \"action\": \"The name of the called function\",   \"id\": \"The ID of the element you want, if relevant\",   \"result\": \"The result of your action: success or error\",   \"data\": \"Only present if this is a success and depends on the function, it's usually a JSON object\",   \"errorDetails\": \"Only present if this is an error, it contains the error message\" } ```   * __Success__ responses are sent with the 200 HTTP (Success) code  * __Error__ responses are sent with a HTTP error code (mostly 5xx...)   ## HTTP method  Rudder's REST API is based on the usage of [HTTP methods](http://www.w3.org/Protocols/rfc2616/rfc2616-sec9.html). We use them to indicate what action will be done by the request. Currently, we use four of them:   * **GET**: search or retrieve information (get rule details, get a group, ...)  * **PUT**: add new objects (create a directive, clone a Rule, ...)  * **DELETE**: remove objects (delete a node, delete a parameter, ...)  * **POST**: update existing objects (update a directive, reload a group, ...)   ## Parameters  ### General parameters  Some parameters are available for almost all API functions. They will be described in this section. They must be part of the query and can't be submitted in a JSON form.  #### Available for all requests  <table>   <thead>     <tr>       <th style=\"width: 30%\">Field</th>       <th style=\"width: 10%\">Type</th>       <th style=\"width: 70%\">Description</th>     </tr>   </thead>   <tbody>     <tr>       <td class=\"code\">prettify</td>       <td><b>boolean</b><br><i>optional</i></td>       <td>         Determine if the answer should be prettified (human friendly) or not. We recommend using this for debugging purposes, but not for general script usage as this does add some unnecessary load on the server side.         <p class=\"default-value\">Default value: <code>false</code></p>       </td>     </tr>   </tbody> </table>   #### Available for modification requests (PUT/POST/DELETE)  <table>   <thead>     <tr>       <th style=\"width: 25%\">Field</th>       <th style=\"width: 12%\">Type</th>       <th style=\"width: 70%\">Description</th>     </tr>   </thead>   <tbody>     <tr>       <td class=\"code\">reason</td>       <td><b>string</b><br><i>optional</i> or <i>required</i></td>       <td>         Set a message to explain the change. If you set the reason messages to be mandatory in the web interface, failing to supply this value will lead to an error.         <p class=\"default-value\">Default value: <code>\"\"</code></p>       </td>     </tr>     <tr>       <td class=\"code\">changeRequestName</td>       <td><b>string</b><br><i>optional</i></td>       <td>         Set the change request name, is used only if workflows are enabled. The default value depends on the function called         <p class=\"default-value\">Default value: <code>A default string for each function</code></p>       </td>     </tr>     <tr>       <td class=\"code\">changeRequestDescription</td>       <td><b>string</b><br><i>optional</i></td>       <td>         Set the change request description, is used only if workflows are enabled.         <p class=\"default-value\">Default value: <code>\"\"</code></p>       </td>     </tr>   </tbody> </table>   ### Passing parameters  Parameters to the API can be sent:  * As part of the URL for resource identification  * As data for POST/PUT requests    * Directly in JSON format    * As request arguments  #### As part of the URL for resource identification  Parameters in URLs are used to indicate which resource you want to interact with. The function will not work if this resource is missing.  ```bash # Get the Rule of ID \"id\" curl -H \"X-API-Token: yourToken\" https://rudder.example.com/rudder/api/latest/rules/id ```  CAUTION: To avoid surprising behavior, do not put a '/' at the end of a URL: it would be interpreted as '/[empty string parameter]' and redirected to '/index', likely not what you wanted to do.   #### Sending data for POST/PUT requests  ##### Directly in JSON format  JSON format is the preferred way to interact with Rudder API for creating or updating resources. You'll also have to set the *Content-Type* header to **application/json** (without it the JSON content would be ignored). In a `curl` `POST` request, that header can be provided with the `-H` parameter:  ```bash curl -X POST -H \"Content-Type: application/json\" ... ```  The supplied file must contain a valid JSON: strings need quotes, booleans and integers don't, etc.  The (human-readable) format is:  ```json {   \"key1\": \"value1\",   \"key2\": false,   \"key3\": 42 } ```  Here is an example with inlined data:  ```bash # Update the Rule 'id' with a new name, disabled, and setting it one directive curl -X POST -H \"X-API-Token: yourToken\" -H  \"Content-Type: application/json\" https://rudder.example.com/rudder/api/rules/latest/{id}   -d '{ \"displayName\": \"new name\", \"enabled\": false, \"directives\": \"directiveId\"}' ```  You can also pass a supply the JSON in a file:  ```bash # Update the Rule 'id' with a new name, disabled, and setting it one directive curl -X POST -H \"X-API-Token: yourToken\" -H \"Content-Type: application/json\" https://rudder.example.com/rudder/api/rules/latest/{id} -d @jsonParam ```  Note that the general parameters view in the previous chapter cannot be passed in a JSON, and you will need to pass them a URL parameters if you want them to be taken into account (you can't mix JSON and request parameters):  ```bash # Update the Rule 'id' with a new name, disabled, and setting it one directive with reason message \"Reason used\" curl -X POST -H \"X-API-Token: yourToken\" -H \"Content-Type: application/json\" \"https://rudder.example.com/rudder/api/rules/latest/{id}?reason=Reason used\" -d @jsonParam -d \"reason=Reason ignored\" ```  ##### Request parameters  In some cases, when you have little, simple data to update, JSON can feel bloated. In such cases, you can use request parameters. You will need to pass one parameter for each data you want to change.  Parameters follow the following schema:  ``` key=value ```  You can pass parameters by two means:  * As query parameters: At the end of your url, put a **?** then your first parameter and then a **&** before next parameters. In that case, parameters need to be https://en.wikipedia.org/wiki/Percent-encoding[URL encoded]  ```bash # Update the Rule 'id' with a new name, disabled, and setting it one directive curl -X POST -H \"X-API-Token: yourToken\"  https://rudder.example.com/rudder/api/rules/latest/{id}?\"displayName=my new name\"&\"enabled=false\"&\"directives=aDirectiveId\" ```  * As request data: You can pass those parameters in the request data, they won't figure in the URL, making it lighter to read, You can pass a file that contains data.  ```bash # Update the Rule 'id' with a new name, disabled, and setting it one directive (in file directive-info.json) curl -X POST -H \"X-API-Token: yourToken\" https://rudder.example.com/rudder/api/rules/latest/{id} -d \"displayName=my new name\" -d \"enabled=false\" -d @directive-info.json ``` 

API version: 18
Contact: dev@rudder.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rudder-golang

import (
	"encoding/json"
)

// checks if the GetAllSettings200ResponseDataSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAllSettings200ResponseDataSettings{}

// GetAllSettings200ResponseDataSettings struct for GetAllSettings200ResponseDataSettings
type GetAllSettings200ResponseDataSettings struct {
	// List of allowed networks for each policy server (root and relays)
	AllowedNetworks []GetAllSettings200ResponseDataSettingsAllowedNetworksInner `json:"allowed_networks,omitempty"`
	// Define the default setting for global policy mode
	GlobalPolicyMode *string `json:"global_policy_mode,omitempty"`
	// Allow overrides on this default setting
	GlobalPolicyModeOverridable *bool `json:"global_policy_mode_overridable,omitempty"`
	// Agent run schedule - time between agent runs (in minutes)
	RunFrequency *int32 `json:"run_frequency,omitempty"`
	// First agent run time - hour
	FirstRunHour *int32 `json:"first_run_hour,omitempty"`
	// First agent run time - minute
	FirstRunMinute *int32 `json:"first_run_minute,omitempty"`
	// Maximum delay after scheduled run time (random interval)
	SplayTime *int32 `json:"splay_time,omitempty"`
	// Number of days to retain modified files
	ModifiedFileTtl *int32 `json:"modified_file_ttl,omitempty"`
	// Number of days to retain agent output files
	OutputFileTtl *int32 `json:"output_file_ttl,omitempty"`
	// Require time synchronization between nodes and policy server
	RequireTimeSynchronization *bool `json:"require_time_synchronization,omitempty"`
	// Method used to synchronize data between server and relays, either \"classic\" (agent protocol, default), \"rsync\" (use rsync to synchronize, that you'll need to be manually set up), or \"disabled\" (use third party system to transmit data)
	RelayServerSynchronizationMethod *string `json:"relay_server_synchronization_method,omitempty"`
	// If **rsync** is set as a synchronization method, use rsync to synchronize policies between Rudder server and relays. If false, you'll have to synchronize policies yourself.
	RelayServerSynchronizePolicies *bool `json:"relay_server_synchronize_policies,omitempty"`
	// If **rsync** is set as a synchronization method, use rsync to synchronize shared files between Rudder server and relays. If false, you'll have to synchronize shared files yourself.
	RelayServerSynchronizeSharedFiles *bool `json:"relay_server_synchronize_shared_files,omitempty"`
	// Default reporting protocol used
	RudderReportProtocolDefault *string `json:"rudder_report_protocol_default,omitempty"`
	// Compliance reporting mode
	ReportingMode *string `json:"reporting_mode,omitempty"`
	// Send heartbeat every heartbeat_frequency runs (only on **changes-only** compliance mode)
	HeartbeatFrequency *int32 `json:"heartbeat_frequency,omitempty"`
	// Enable change audit logs
	EnableChangeMessage *bool `json:"enable_change_message,omitempty"`
	// Make message mandatory
	MandatoryChangeMessage *bool `json:"mandatory_change_message,omitempty"`
	// Explanation to display
	ChangeMessagePrompt *string `json:"change_message_prompt,omitempty"`
	// Enable Change Requests
	EnableChangeRequest *bool `json:"enable_change_request,omitempty"`
	// Allow self validation
	EnableSelfValidation *bool `json:"enable_self_validation,omitempty"`
	// Allow self deployment
	EnableSelfDeployment *bool `json:"enable_self_deployment,omitempty"`
	// Display changes graphs
	DisplayRecentChangesGraphs *bool `json:"display_recent_changes_graphs,omitempty"`
	// Enable script evaluation in Directives
	EnableJavascriptDirectives *string `json:"enable_javascript_directives,omitempty"`
	// Send anonymous usage statistics
	SendMetrics *string `json:"send_metrics,omitempty"`
	// Allow acceptation of a pending node when another one with the same hostname is already accepted
	NodeAcceptDuplicatedHostname *bool `json:"node_accept_duplicated_hostname,omitempty"`
	// Set default state for node when they are accepted within Rudder
	NodeOnacceptDefaultState *string `json:"node_onaccept_default_state,omitempty"`
	// Default policy mode for accepted node
	NodeOnacceptDefaultPolicyMode *string `json:"node_onaccept_default_policyMode,omitempty"`
	// Allows multiple reports for configuration based on a multivalued variable
	UnexpectedUnboundVarValues *bool `json:"unexpected_unbound_var_values,omitempty"`
	// Compute list of changes (repaired reports) per rule
	RudderComputeChanges *bool `json:"rudder_compute_changes,omitempty"`
	// Recompute all dynamic groups at the start of policy generation
	RudderGenerationComputeDyngroups *bool `json:"rudder_generation_compute_dyngroups,omitempty"`
	// Set the parallelism to compute dynamic group, as a number of thread (i.e. 4), or a multiplicative of the number of core (x0.5)
	RudderComputeDyngroupsMaxParallelism *string `json:"rudder_compute_dyngroups_max_parallelism,omitempty"`
	// Store all compliance levels in database
	RudderSaveDbComplianceLevels *bool `json:"rudder_save_db_compliance_levels,omitempty"`
	// Store all compliance details in database
	RudderSaveDbComplianceDetails *bool `json:"rudder_save_db_compliance_details,omitempty"`
	// Set the policy generation parallelism, either as an number of thread (i.e. 4), or a multiplicative of the number of core (x0.5)
	RudderGenerationMaxParallelism *string `json:"rudder_generation_max_parallelism,omitempty"`
	// Policy generation JS evaluation of directive parameter timeout in seconds
	RudderGenerationJsTimeout *int32 `json:"rudder_generation_js_timeout,omitempty"`
	// Policy generation continues on error during NodeConfiguration evaluation
	RudderGenerationContinueOnError *bool `json:"rudder_generation_continue_on_error,omitempty"`
	// Set a delay before starting policy generation, this will allow you to accumulate changes before they are deployed to Nodes, and can also lessen webapp resources needs. The value is a number followed by the time unit needed (seconds/s, minutes/m, hours/h ...), ie \"5m\" for 5 minutes
	RudderGenerationDelay *string `json:"rudder_generation_delay,omitempty"`
	// Should policy generation be triggered automatically after a change (value 'all'), or should we wait until a manual launch (through api or UI, value 'onlyManual') or even no policy generation at all (value \"none\")
	RudderGenerationPolicy *string `json:"rudder_generation_policy,omitempty"`
}

// NewGetAllSettings200ResponseDataSettings instantiates a new GetAllSettings200ResponseDataSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllSettings200ResponseDataSettings() *GetAllSettings200ResponseDataSettings {
	this := GetAllSettings200ResponseDataSettings{}
	var nodeAcceptDuplicatedHostname bool = false
	this.NodeAcceptDuplicatedHostname = &nodeAcceptDuplicatedHostname
	var unexpectedUnboundVarValues bool = true
	this.UnexpectedUnboundVarValues = &unexpectedUnboundVarValues
	var rudderComputeChanges bool = true
	this.RudderComputeChanges = &rudderComputeChanges
	var rudderGenerationComputeDyngroups bool = true
	this.RudderGenerationComputeDyngroups = &rudderGenerationComputeDyngroups
	var rudderComputeDyngroupsMaxParallelism string = "1"
	this.RudderComputeDyngroupsMaxParallelism = &rudderComputeDyngroupsMaxParallelism
	var rudderSaveDbComplianceLevels bool = true
	this.RudderSaveDbComplianceLevels = &rudderSaveDbComplianceLevels
	var rudderSaveDbComplianceDetails bool = false
	this.RudderSaveDbComplianceDetails = &rudderSaveDbComplianceDetails
	var rudderGenerationMaxParallelism string = "x0.5"
	this.RudderGenerationMaxParallelism = &rudderGenerationMaxParallelism
	var rudderGenerationJsTimeout int32 = 30
	this.RudderGenerationJsTimeout = &rudderGenerationJsTimeout
	var rudderGenerationContinueOnError bool = false
	this.RudderGenerationContinueOnError = &rudderGenerationContinueOnError
	var rudderGenerationDelay string = "0 seconds"
	this.RudderGenerationDelay = &rudderGenerationDelay
	var rudderGenerationPolicy string = "all"
	this.RudderGenerationPolicy = &rudderGenerationPolicy
	return &this
}

// NewGetAllSettings200ResponseDataSettingsWithDefaults instantiates a new GetAllSettings200ResponseDataSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllSettings200ResponseDataSettingsWithDefaults() *GetAllSettings200ResponseDataSettings {
	this := GetAllSettings200ResponseDataSettings{}
	var nodeAcceptDuplicatedHostname bool = false
	this.NodeAcceptDuplicatedHostname = &nodeAcceptDuplicatedHostname
	var unexpectedUnboundVarValues bool = true
	this.UnexpectedUnboundVarValues = &unexpectedUnboundVarValues
	var rudderComputeChanges bool = true
	this.RudderComputeChanges = &rudderComputeChanges
	var rudderGenerationComputeDyngroups bool = true
	this.RudderGenerationComputeDyngroups = &rudderGenerationComputeDyngroups
	var rudderComputeDyngroupsMaxParallelism string = "1"
	this.RudderComputeDyngroupsMaxParallelism = &rudderComputeDyngroupsMaxParallelism
	var rudderSaveDbComplianceLevels bool = true
	this.RudderSaveDbComplianceLevels = &rudderSaveDbComplianceLevels
	var rudderSaveDbComplianceDetails bool = false
	this.RudderSaveDbComplianceDetails = &rudderSaveDbComplianceDetails
	var rudderGenerationMaxParallelism string = "x0.5"
	this.RudderGenerationMaxParallelism = &rudderGenerationMaxParallelism
	var rudderGenerationJsTimeout int32 = 30
	this.RudderGenerationJsTimeout = &rudderGenerationJsTimeout
	var rudderGenerationContinueOnError bool = false
	this.RudderGenerationContinueOnError = &rudderGenerationContinueOnError
	var rudderGenerationDelay string = "0 seconds"
	this.RudderGenerationDelay = &rudderGenerationDelay
	var rudderGenerationPolicy string = "all"
	this.RudderGenerationPolicy = &rudderGenerationPolicy
	return &this
}

// GetAllowedNetworks returns the AllowedNetworks field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetAllowedNetworks() []GetAllSettings200ResponseDataSettingsAllowedNetworksInner {
	if o == nil || IsNil(o.AllowedNetworks) {
		var ret []GetAllSettings200ResponseDataSettingsAllowedNetworksInner
		return ret
	}
	return o.AllowedNetworks
}

// GetAllowedNetworksOk returns a tuple with the AllowedNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetAllowedNetworksOk() ([]GetAllSettings200ResponseDataSettingsAllowedNetworksInner, bool) {
	if o == nil || IsNil(o.AllowedNetworks) {
		return nil, false
	}
	return o.AllowedNetworks, true
}

// HasAllowedNetworks returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasAllowedNetworks() bool {
	if o != nil && !IsNil(o.AllowedNetworks) {
		return true
	}

	return false
}

// SetAllowedNetworks gets a reference to the given []GetAllSettings200ResponseDataSettingsAllowedNetworksInner and assigns it to the AllowedNetworks field.
func (o *GetAllSettings200ResponseDataSettings) SetAllowedNetworks(v []GetAllSettings200ResponseDataSettingsAllowedNetworksInner) {
	o.AllowedNetworks = v
}

// GetGlobalPolicyMode returns the GlobalPolicyMode field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetGlobalPolicyMode() string {
	if o == nil || IsNil(o.GlobalPolicyMode) {
		var ret string
		return ret
	}
	return *o.GlobalPolicyMode
}

// GetGlobalPolicyModeOk returns a tuple with the GlobalPolicyMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetGlobalPolicyModeOk() (*string, bool) {
	if o == nil || IsNil(o.GlobalPolicyMode) {
		return nil, false
	}
	return o.GlobalPolicyMode, true
}

// HasGlobalPolicyMode returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasGlobalPolicyMode() bool {
	if o != nil && !IsNil(o.GlobalPolicyMode) {
		return true
	}

	return false
}

// SetGlobalPolicyMode gets a reference to the given string and assigns it to the GlobalPolicyMode field.
func (o *GetAllSettings200ResponseDataSettings) SetGlobalPolicyMode(v string) {
	o.GlobalPolicyMode = &v
}

// GetGlobalPolicyModeOverridable returns the GlobalPolicyModeOverridable field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetGlobalPolicyModeOverridable() bool {
	if o == nil || IsNil(o.GlobalPolicyModeOverridable) {
		var ret bool
		return ret
	}
	return *o.GlobalPolicyModeOverridable
}

// GetGlobalPolicyModeOverridableOk returns a tuple with the GlobalPolicyModeOverridable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetGlobalPolicyModeOverridableOk() (*bool, bool) {
	if o == nil || IsNil(o.GlobalPolicyModeOverridable) {
		return nil, false
	}
	return o.GlobalPolicyModeOverridable, true
}

// HasGlobalPolicyModeOverridable returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasGlobalPolicyModeOverridable() bool {
	if o != nil && !IsNil(o.GlobalPolicyModeOverridable) {
		return true
	}

	return false
}

// SetGlobalPolicyModeOverridable gets a reference to the given bool and assigns it to the GlobalPolicyModeOverridable field.
func (o *GetAllSettings200ResponseDataSettings) SetGlobalPolicyModeOverridable(v bool) {
	o.GlobalPolicyModeOverridable = &v
}

// GetRunFrequency returns the RunFrequency field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetRunFrequency() int32 {
	if o == nil || IsNil(o.RunFrequency) {
		var ret int32
		return ret
	}
	return *o.RunFrequency
}

// GetRunFrequencyOk returns a tuple with the RunFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetRunFrequencyOk() (*int32, bool) {
	if o == nil || IsNil(o.RunFrequency) {
		return nil, false
	}
	return o.RunFrequency, true
}

// HasRunFrequency returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasRunFrequency() bool {
	if o != nil && !IsNil(o.RunFrequency) {
		return true
	}

	return false
}

// SetRunFrequency gets a reference to the given int32 and assigns it to the RunFrequency field.
func (o *GetAllSettings200ResponseDataSettings) SetRunFrequency(v int32) {
	o.RunFrequency = &v
}

// GetFirstRunHour returns the FirstRunHour field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetFirstRunHour() int32 {
	if o == nil || IsNil(o.FirstRunHour) {
		var ret int32
		return ret
	}
	return *o.FirstRunHour
}

// GetFirstRunHourOk returns a tuple with the FirstRunHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetFirstRunHourOk() (*int32, bool) {
	if o == nil || IsNil(o.FirstRunHour) {
		return nil, false
	}
	return o.FirstRunHour, true
}

// HasFirstRunHour returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasFirstRunHour() bool {
	if o != nil && !IsNil(o.FirstRunHour) {
		return true
	}

	return false
}

// SetFirstRunHour gets a reference to the given int32 and assigns it to the FirstRunHour field.
func (o *GetAllSettings200ResponseDataSettings) SetFirstRunHour(v int32) {
	o.FirstRunHour = &v
}

// GetFirstRunMinute returns the FirstRunMinute field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetFirstRunMinute() int32 {
	if o == nil || IsNil(o.FirstRunMinute) {
		var ret int32
		return ret
	}
	return *o.FirstRunMinute
}

// GetFirstRunMinuteOk returns a tuple with the FirstRunMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetFirstRunMinuteOk() (*int32, bool) {
	if o == nil || IsNil(o.FirstRunMinute) {
		return nil, false
	}
	return o.FirstRunMinute, true
}

// HasFirstRunMinute returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasFirstRunMinute() bool {
	if o != nil && !IsNil(o.FirstRunMinute) {
		return true
	}

	return false
}

// SetFirstRunMinute gets a reference to the given int32 and assigns it to the FirstRunMinute field.
func (o *GetAllSettings200ResponseDataSettings) SetFirstRunMinute(v int32) {
	o.FirstRunMinute = &v
}

// GetSplayTime returns the SplayTime field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetSplayTime() int32 {
	if o == nil || IsNil(o.SplayTime) {
		var ret int32
		return ret
	}
	return *o.SplayTime
}

// GetSplayTimeOk returns a tuple with the SplayTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetSplayTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.SplayTime) {
		return nil, false
	}
	return o.SplayTime, true
}

// HasSplayTime returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasSplayTime() bool {
	if o != nil && !IsNil(o.SplayTime) {
		return true
	}

	return false
}

// SetSplayTime gets a reference to the given int32 and assigns it to the SplayTime field.
func (o *GetAllSettings200ResponseDataSettings) SetSplayTime(v int32) {
	o.SplayTime = &v
}

// GetModifiedFileTtl returns the ModifiedFileTtl field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetModifiedFileTtl() int32 {
	if o == nil || IsNil(o.ModifiedFileTtl) {
		var ret int32
		return ret
	}
	return *o.ModifiedFileTtl
}

// GetModifiedFileTtlOk returns a tuple with the ModifiedFileTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetModifiedFileTtlOk() (*int32, bool) {
	if o == nil || IsNil(o.ModifiedFileTtl) {
		return nil, false
	}
	return o.ModifiedFileTtl, true
}

// HasModifiedFileTtl returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasModifiedFileTtl() bool {
	if o != nil && !IsNil(o.ModifiedFileTtl) {
		return true
	}

	return false
}

// SetModifiedFileTtl gets a reference to the given int32 and assigns it to the ModifiedFileTtl field.
func (o *GetAllSettings200ResponseDataSettings) SetModifiedFileTtl(v int32) {
	o.ModifiedFileTtl = &v
}

// GetOutputFileTtl returns the OutputFileTtl field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetOutputFileTtl() int32 {
	if o == nil || IsNil(o.OutputFileTtl) {
		var ret int32
		return ret
	}
	return *o.OutputFileTtl
}

// GetOutputFileTtlOk returns a tuple with the OutputFileTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetOutputFileTtlOk() (*int32, bool) {
	if o == nil || IsNil(o.OutputFileTtl) {
		return nil, false
	}
	return o.OutputFileTtl, true
}

// HasOutputFileTtl returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasOutputFileTtl() bool {
	if o != nil && !IsNil(o.OutputFileTtl) {
		return true
	}

	return false
}

// SetOutputFileTtl gets a reference to the given int32 and assigns it to the OutputFileTtl field.
func (o *GetAllSettings200ResponseDataSettings) SetOutputFileTtl(v int32) {
	o.OutputFileTtl = &v
}

// GetRequireTimeSynchronization returns the RequireTimeSynchronization field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetRequireTimeSynchronization() bool {
	if o == nil || IsNil(o.RequireTimeSynchronization) {
		var ret bool
		return ret
	}
	return *o.RequireTimeSynchronization
}

// GetRequireTimeSynchronizationOk returns a tuple with the RequireTimeSynchronization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetRequireTimeSynchronizationOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireTimeSynchronization) {
		return nil, false
	}
	return o.RequireTimeSynchronization, true
}

// HasRequireTimeSynchronization returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasRequireTimeSynchronization() bool {
	if o != nil && !IsNil(o.RequireTimeSynchronization) {
		return true
	}

	return false
}

// SetRequireTimeSynchronization gets a reference to the given bool and assigns it to the RequireTimeSynchronization field.
func (o *GetAllSettings200ResponseDataSettings) SetRequireTimeSynchronization(v bool) {
	o.RequireTimeSynchronization = &v
}

// GetRelayServerSynchronizationMethod returns the RelayServerSynchronizationMethod field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetRelayServerSynchronizationMethod() string {
	if o == nil || IsNil(o.RelayServerSynchronizationMethod) {
		var ret string
		return ret
	}
	return *o.RelayServerSynchronizationMethod
}

// GetRelayServerSynchronizationMethodOk returns a tuple with the RelayServerSynchronizationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetRelayServerSynchronizationMethodOk() (*string, bool) {
	if o == nil || IsNil(o.RelayServerSynchronizationMethod) {
		return nil, false
	}
	return o.RelayServerSynchronizationMethod, true
}

// HasRelayServerSynchronizationMethod returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasRelayServerSynchronizationMethod() bool {
	if o != nil && !IsNil(o.RelayServerSynchronizationMethod) {
		return true
	}

	return false
}

// SetRelayServerSynchronizationMethod gets a reference to the given string and assigns it to the RelayServerSynchronizationMethod field.
func (o *GetAllSettings200ResponseDataSettings) SetRelayServerSynchronizationMethod(v string) {
	o.RelayServerSynchronizationMethod = &v
}

// GetRelayServerSynchronizePolicies returns the RelayServerSynchronizePolicies field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetRelayServerSynchronizePolicies() bool {
	if o == nil || IsNil(o.RelayServerSynchronizePolicies) {
		var ret bool
		return ret
	}
	return *o.RelayServerSynchronizePolicies
}

// GetRelayServerSynchronizePoliciesOk returns a tuple with the RelayServerSynchronizePolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetRelayServerSynchronizePoliciesOk() (*bool, bool) {
	if o == nil || IsNil(o.RelayServerSynchronizePolicies) {
		return nil, false
	}
	return o.RelayServerSynchronizePolicies, true
}

// HasRelayServerSynchronizePolicies returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasRelayServerSynchronizePolicies() bool {
	if o != nil && !IsNil(o.RelayServerSynchronizePolicies) {
		return true
	}

	return false
}

// SetRelayServerSynchronizePolicies gets a reference to the given bool and assigns it to the RelayServerSynchronizePolicies field.
func (o *GetAllSettings200ResponseDataSettings) SetRelayServerSynchronizePolicies(v bool) {
	o.RelayServerSynchronizePolicies = &v
}

// GetRelayServerSynchronizeSharedFiles returns the RelayServerSynchronizeSharedFiles field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetRelayServerSynchronizeSharedFiles() bool {
	if o == nil || IsNil(o.RelayServerSynchronizeSharedFiles) {
		var ret bool
		return ret
	}
	return *o.RelayServerSynchronizeSharedFiles
}

// GetRelayServerSynchronizeSharedFilesOk returns a tuple with the RelayServerSynchronizeSharedFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetRelayServerSynchronizeSharedFilesOk() (*bool, bool) {
	if o == nil || IsNil(o.RelayServerSynchronizeSharedFiles) {
		return nil, false
	}
	return o.RelayServerSynchronizeSharedFiles, true
}

// HasRelayServerSynchronizeSharedFiles returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasRelayServerSynchronizeSharedFiles() bool {
	if o != nil && !IsNil(o.RelayServerSynchronizeSharedFiles) {
		return true
	}

	return false
}

// SetRelayServerSynchronizeSharedFiles gets a reference to the given bool and assigns it to the RelayServerSynchronizeSharedFiles field.
func (o *GetAllSettings200ResponseDataSettings) SetRelayServerSynchronizeSharedFiles(v bool) {
	o.RelayServerSynchronizeSharedFiles = &v
}

// GetRudderReportProtocolDefault returns the RudderReportProtocolDefault field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetRudderReportProtocolDefault() string {
	if o == nil || IsNil(o.RudderReportProtocolDefault) {
		var ret string
		return ret
	}
	return *o.RudderReportProtocolDefault
}

// GetRudderReportProtocolDefaultOk returns a tuple with the RudderReportProtocolDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetRudderReportProtocolDefaultOk() (*string, bool) {
	if o == nil || IsNil(o.RudderReportProtocolDefault) {
		return nil, false
	}
	return o.RudderReportProtocolDefault, true
}

// HasRudderReportProtocolDefault returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasRudderReportProtocolDefault() bool {
	if o != nil && !IsNil(o.RudderReportProtocolDefault) {
		return true
	}

	return false
}

// SetRudderReportProtocolDefault gets a reference to the given string and assigns it to the RudderReportProtocolDefault field.
func (o *GetAllSettings200ResponseDataSettings) SetRudderReportProtocolDefault(v string) {
	o.RudderReportProtocolDefault = &v
}

// GetReportingMode returns the ReportingMode field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetReportingMode() string {
	if o == nil || IsNil(o.ReportingMode) {
		var ret string
		return ret
	}
	return *o.ReportingMode
}

// GetReportingModeOk returns a tuple with the ReportingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetReportingModeOk() (*string, bool) {
	if o == nil || IsNil(o.ReportingMode) {
		return nil, false
	}
	return o.ReportingMode, true
}

// HasReportingMode returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasReportingMode() bool {
	if o != nil && !IsNil(o.ReportingMode) {
		return true
	}

	return false
}

// SetReportingMode gets a reference to the given string and assigns it to the ReportingMode field.
func (o *GetAllSettings200ResponseDataSettings) SetReportingMode(v string) {
	o.ReportingMode = &v
}

// GetHeartbeatFrequency returns the HeartbeatFrequency field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetHeartbeatFrequency() int32 {
	if o == nil || IsNil(o.HeartbeatFrequency) {
		var ret int32
		return ret
	}
	return *o.HeartbeatFrequency
}

// GetHeartbeatFrequencyOk returns a tuple with the HeartbeatFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetHeartbeatFrequencyOk() (*int32, bool) {
	if o == nil || IsNil(o.HeartbeatFrequency) {
		return nil, false
	}
	return o.HeartbeatFrequency, true
}

// HasHeartbeatFrequency returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasHeartbeatFrequency() bool {
	if o != nil && !IsNil(o.HeartbeatFrequency) {
		return true
	}

	return false
}

// SetHeartbeatFrequency gets a reference to the given int32 and assigns it to the HeartbeatFrequency field.
func (o *GetAllSettings200ResponseDataSettings) SetHeartbeatFrequency(v int32) {
	o.HeartbeatFrequency = &v
}

// GetEnableChangeMessage returns the EnableChangeMessage field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetEnableChangeMessage() bool {
	if o == nil || IsNil(o.EnableChangeMessage) {
		var ret bool
		return ret
	}
	return *o.EnableChangeMessage
}

// GetEnableChangeMessageOk returns a tuple with the EnableChangeMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetEnableChangeMessageOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableChangeMessage) {
		return nil, false
	}
	return o.EnableChangeMessage, true
}

// HasEnableChangeMessage returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasEnableChangeMessage() bool {
	if o != nil && !IsNil(o.EnableChangeMessage) {
		return true
	}

	return false
}

// SetEnableChangeMessage gets a reference to the given bool and assigns it to the EnableChangeMessage field.
func (o *GetAllSettings200ResponseDataSettings) SetEnableChangeMessage(v bool) {
	o.EnableChangeMessage = &v
}

// GetMandatoryChangeMessage returns the MandatoryChangeMessage field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetMandatoryChangeMessage() bool {
	if o == nil || IsNil(o.MandatoryChangeMessage) {
		var ret bool
		return ret
	}
	return *o.MandatoryChangeMessage
}

// GetMandatoryChangeMessageOk returns a tuple with the MandatoryChangeMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetMandatoryChangeMessageOk() (*bool, bool) {
	if o == nil || IsNil(o.MandatoryChangeMessage) {
		return nil, false
	}
	return o.MandatoryChangeMessage, true
}

// HasMandatoryChangeMessage returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasMandatoryChangeMessage() bool {
	if o != nil && !IsNil(o.MandatoryChangeMessage) {
		return true
	}

	return false
}

// SetMandatoryChangeMessage gets a reference to the given bool and assigns it to the MandatoryChangeMessage field.
func (o *GetAllSettings200ResponseDataSettings) SetMandatoryChangeMessage(v bool) {
	o.MandatoryChangeMessage = &v
}

// GetChangeMessagePrompt returns the ChangeMessagePrompt field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetChangeMessagePrompt() string {
	if o == nil || IsNil(o.ChangeMessagePrompt) {
		var ret string
		return ret
	}
	return *o.ChangeMessagePrompt
}

// GetChangeMessagePromptOk returns a tuple with the ChangeMessagePrompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetChangeMessagePromptOk() (*string, bool) {
	if o == nil || IsNil(o.ChangeMessagePrompt) {
		return nil, false
	}
	return o.ChangeMessagePrompt, true
}

// HasChangeMessagePrompt returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasChangeMessagePrompt() bool {
	if o != nil && !IsNil(o.ChangeMessagePrompt) {
		return true
	}

	return false
}

// SetChangeMessagePrompt gets a reference to the given string and assigns it to the ChangeMessagePrompt field.
func (o *GetAllSettings200ResponseDataSettings) SetChangeMessagePrompt(v string) {
	o.ChangeMessagePrompt = &v
}

// GetEnableChangeRequest returns the EnableChangeRequest field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetEnableChangeRequest() bool {
	if o == nil || IsNil(o.EnableChangeRequest) {
		var ret bool
		return ret
	}
	return *o.EnableChangeRequest
}

// GetEnableChangeRequestOk returns a tuple with the EnableChangeRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetEnableChangeRequestOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableChangeRequest) {
		return nil, false
	}
	return o.EnableChangeRequest, true
}

// HasEnableChangeRequest returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasEnableChangeRequest() bool {
	if o != nil && !IsNil(o.EnableChangeRequest) {
		return true
	}

	return false
}

// SetEnableChangeRequest gets a reference to the given bool and assigns it to the EnableChangeRequest field.
func (o *GetAllSettings200ResponseDataSettings) SetEnableChangeRequest(v bool) {
	o.EnableChangeRequest = &v
}

// GetEnableSelfValidation returns the EnableSelfValidation field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetEnableSelfValidation() bool {
	if o == nil || IsNil(o.EnableSelfValidation) {
		var ret bool
		return ret
	}
	return *o.EnableSelfValidation
}

// GetEnableSelfValidationOk returns a tuple with the EnableSelfValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetEnableSelfValidationOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableSelfValidation) {
		return nil, false
	}
	return o.EnableSelfValidation, true
}

// HasEnableSelfValidation returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasEnableSelfValidation() bool {
	if o != nil && !IsNil(o.EnableSelfValidation) {
		return true
	}

	return false
}

// SetEnableSelfValidation gets a reference to the given bool and assigns it to the EnableSelfValidation field.
func (o *GetAllSettings200ResponseDataSettings) SetEnableSelfValidation(v bool) {
	o.EnableSelfValidation = &v
}

// GetEnableSelfDeployment returns the EnableSelfDeployment field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetEnableSelfDeployment() bool {
	if o == nil || IsNil(o.EnableSelfDeployment) {
		var ret bool
		return ret
	}
	return *o.EnableSelfDeployment
}

// GetEnableSelfDeploymentOk returns a tuple with the EnableSelfDeployment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetEnableSelfDeploymentOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableSelfDeployment) {
		return nil, false
	}
	return o.EnableSelfDeployment, true
}

// HasEnableSelfDeployment returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasEnableSelfDeployment() bool {
	if o != nil && !IsNil(o.EnableSelfDeployment) {
		return true
	}

	return false
}

// SetEnableSelfDeployment gets a reference to the given bool and assigns it to the EnableSelfDeployment field.
func (o *GetAllSettings200ResponseDataSettings) SetEnableSelfDeployment(v bool) {
	o.EnableSelfDeployment = &v
}

// GetDisplayRecentChangesGraphs returns the DisplayRecentChangesGraphs field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetDisplayRecentChangesGraphs() bool {
	if o == nil || IsNil(o.DisplayRecentChangesGraphs) {
		var ret bool
		return ret
	}
	return *o.DisplayRecentChangesGraphs
}

// GetDisplayRecentChangesGraphsOk returns a tuple with the DisplayRecentChangesGraphs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetDisplayRecentChangesGraphsOk() (*bool, bool) {
	if o == nil || IsNil(o.DisplayRecentChangesGraphs) {
		return nil, false
	}
	return o.DisplayRecentChangesGraphs, true
}

// HasDisplayRecentChangesGraphs returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasDisplayRecentChangesGraphs() bool {
	if o != nil && !IsNil(o.DisplayRecentChangesGraphs) {
		return true
	}

	return false
}

// SetDisplayRecentChangesGraphs gets a reference to the given bool and assigns it to the DisplayRecentChangesGraphs field.
func (o *GetAllSettings200ResponseDataSettings) SetDisplayRecentChangesGraphs(v bool) {
	o.DisplayRecentChangesGraphs = &v
}

// GetEnableJavascriptDirectives returns the EnableJavascriptDirectives field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetEnableJavascriptDirectives() string {
	if o == nil || IsNil(o.EnableJavascriptDirectives) {
		var ret string
		return ret
	}
	return *o.EnableJavascriptDirectives
}

// GetEnableJavascriptDirectivesOk returns a tuple with the EnableJavascriptDirectives field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetEnableJavascriptDirectivesOk() (*string, bool) {
	if o == nil || IsNil(o.EnableJavascriptDirectives) {
		return nil, false
	}
	return o.EnableJavascriptDirectives, true
}

// HasEnableJavascriptDirectives returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasEnableJavascriptDirectives() bool {
	if o != nil && !IsNil(o.EnableJavascriptDirectives) {
		return true
	}

	return false
}

// SetEnableJavascriptDirectives gets a reference to the given string and assigns it to the EnableJavascriptDirectives field.
func (o *GetAllSettings200ResponseDataSettings) SetEnableJavascriptDirectives(v string) {
	o.EnableJavascriptDirectives = &v
}

// GetSendMetrics returns the SendMetrics field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetSendMetrics() string {
	if o == nil || IsNil(o.SendMetrics) {
		var ret string
		return ret
	}
	return *o.SendMetrics
}

// GetSendMetricsOk returns a tuple with the SendMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetSendMetricsOk() (*string, bool) {
	if o == nil || IsNil(o.SendMetrics) {
		return nil, false
	}
	return o.SendMetrics, true
}

// HasSendMetrics returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasSendMetrics() bool {
	if o != nil && !IsNil(o.SendMetrics) {
		return true
	}

	return false
}

// SetSendMetrics gets a reference to the given string and assigns it to the SendMetrics field.
func (o *GetAllSettings200ResponseDataSettings) SetSendMetrics(v string) {
	o.SendMetrics = &v
}

// GetNodeAcceptDuplicatedHostname returns the NodeAcceptDuplicatedHostname field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetNodeAcceptDuplicatedHostname() bool {
	if o == nil || IsNil(o.NodeAcceptDuplicatedHostname) {
		var ret bool
		return ret
	}
	return *o.NodeAcceptDuplicatedHostname
}

// GetNodeAcceptDuplicatedHostnameOk returns a tuple with the NodeAcceptDuplicatedHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetNodeAcceptDuplicatedHostnameOk() (*bool, bool) {
	if o == nil || IsNil(o.NodeAcceptDuplicatedHostname) {
		return nil, false
	}
	return o.NodeAcceptDuplicatedHostname, true
}

// HasNodeAcceptDuplicatedHostname returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasNodeAcceptDuplicatedHostname() bool {
	if o != nil && !IsNil(o.NodeAcceptDuplicatedHostname) {
		return true
	}

	return false
}

// SetNodeAcceptDuplicatedHostname gets a reference to the given bool and assigns it to the NodeAcceptDuplicatedHostname field.
func (o *GetAllSettings200ResponseDataSettings) SetNodeAcceptDuplicatedHostname(v bool) {
	o.NodeAcceptDuplicatedHostname = &v
}

// GetNodeOnacceptDefaultState returns the NodeOnacceptDefaultState field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetNodeOnacceptDefaultState() string {
	if o == nil || IsNil(o.NodeOnacceptDefaultState) {
		var ret string
		return ret
	}
	return *o.NodeOnacceptDefaultState
}

// GetNodeOnacceptDefaultStateOk returns a tuple with the NodeOnacceptDefaultState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetNodeOnacceptDefaultStateOk() (*string, bool) {
	if o == nil || IsNil(o.NodeOnacceptDefaultState) {
		return nil, false
	}
	return o.NodeOnacceptDefaultState, true
}

// HasNodeOnacceptDefaultState returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasNodeOnacceptDefaultState() bool {
	if o != nil && !IsNil(o.NodeOnacceptDefaultState) {
		return true
	}

	return false
}

// SetNodeOnacceptDefaultState gets a reference to the given string and assigns it to the NodeOnacceptDefaultState field.
func (o *GetAllSettings200ResponseDataSettings) SetNodeOnacceptDefaultState(v string) {
	o.NodeOnacceptDefaultState = &v
}

// GetNodeOnacceptDefaultPolicyMode returns the NodeOnacceptDefaultPolicyMode field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetNodeOnacceptDefaultPolicyMode() string {
	if o == nil || IsNil(o.NodeOnacceptDefaultPolicyMode) {
		var ret string
		return ret
	}
	return *o.NodeOnacceptDefaultPolicyMode
}

// GetNodeOnacceptDefaultPolicyModeOk returns a tuple with the NodeOnacceptDefaultPolicyMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetNodeOnacceptDefaultPolicyModeOk() (*string, bool) {
	if o == nil || IsNil(o.NodeOnacceptDefaultPolicyMode) {
		return nil, false
	}
	return o.NodeOnacceptDefaultPolicyMode, true
}

// HasNodeOnacceptDefaultPolicyMode returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasNodeOnacceptDefaultPolicyMode() bool {
	if o != nil && !IsNil(o.NodeOnacceptDefaultPolicyMode) {
		return true
	}

	return false
}

// SetNodeOnacceptDefaultPolicyMode gets a reference to the given string and assigns it to the NodeOnacceptDefaultPolicyMode field.
func (o *GetAllSettings200ResponseDataSettings) SetNodeOnacceptDefaultPolicyMode(v string) {
	o.NodeOnacceptDefaultPolicyMode = &v
}

// GetUnexpectedUnboundVarValues returns the UnexpectedUnboundVarValues field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetUnexpectedUnboundVarValues() bool {
	if o == nil || IsNil(o.UnexpectedUnboundVarValues) {
		var ret bool
		return ret
	}
	return *o.UnexpectedUnboundVarValues
}

// GetUnexpectedUnboundVarValuesOk returns a tuple with the UnexpectedUnboundVarValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetUnexpectedUnboundVarValuesOk() (*bool, bool) {
	if o == nil || IsNil(o.UnexpectedUnboundVarValues) {
		return nil, false
	}
	return o.UnexpectedUnboundVarValues, true
}

// HasUnexpectedUnboundVarValues returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasUnexpectedUnboundVarValues() bool {
	if o != nil && !IsNil(o.UnexpectedUnboundVarValues) {
		return true
	}

	return false
}

// SetUnexpectedUnboundVarValues gets a reference to the given bool and assigns it to the UnexpectedUnboundVarValues field.
func (o *GetAllSettings200ResponseDataSettings) SetUnexpectedUnboundVarValues(v bool) {
	o.UnexpectedUnboundVarValues = &v
}

// GetRudderComputeChanges returns the RudderComputeChanges field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetRudderComputeChanges() bool {
	if o == nil || IsNil(o.RudderComputeChanges) {
		var ret bool
		return ret
	}
	return *o.RudderComputeChanges
}

// GetRudderComputeChangesOk returns a tuple with the RudderComputeChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetRudderComputeChangesOk() (*bool, bool) {
	if o == nil || IsNil(o.RudderComputeChanges) {
		return nil, false
	}
	return o.RudderComputeChanges, true
}

// HasRudderComputeChanges returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasRudderComputeChanges() bool {
	if o != nil && !IsNil(o.RudderComputeChanges) {
		return true
	}

	return false
}

// SetRudderComputeChanges gets a reference to the given bool and assigns it to the RudderComputeChanges field.
func (o *GetAllSettings200ResponseDataSettings) SetRudderComputeChanges(v bool) {
	o.RudderComputeChanges = &v
}

// GetRudderGenerationComputeDyngroups returns the RudderGenerationComputeDyngroups field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetRudderGenerationComputeDyngroups() bool {
	if o == nil || IsNil(o.RudderGenerationComputeDyngroups) {
		var ret bool
		return ret
	}
	return *o.RudderGenerationComputeDyngroups
}

// GetRudderGenerationComputeDyngroupsOk returns a tuple with the RudderGenerationComputeDyngroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetRudderGenerationComputeDyngroupsOk() (*bool, bool) {
	if o == nil || IsNil(o.RudderGenerationComputeDyngroups) {
		return nil, false
	}
	return o.RudderGenerationComputeDyngroups, true
}

// HasRudderGenerationComputeDyngroups returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasRudderGenerationComputeDyngroups() bool {
	if o != nil && !IsNil(o.RudderGenerationComputeDyngroups) {
		return true
	}

	return false
}

// SetRudderGenerationComputeDyngroups gets a reference to the given bool and assigns it to the RudderGenerationComputeDyngroups field.
func (o *GetAllSettings200ResponseDataSettings) SetRudderGenerationComputeDyngroups(v bool) {
	o.RudderGenerationComputeDyngroups = &v
}

// GetRudderComputeDyngroupsMaxParallelism returns the RudderComputeDyngroupsMaxParallelism field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetRudderComputeDyngroupsMaxParallelism() string {
	if o == nil || IsNil(o.RudderComputeDyngroupsMaxParallelism) {
		var ret string
		return ret
	}
	return *o.RudderComputeDyngroupsMaxParallelism
}

// GetRudderComputeDyngroupsMaxParallelismOk returns a tuple with the RudderComputeDyngroupsMaxParallelism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetRudderComputeDyngroupsMaxParallelismOk() (*string, bool) {
	if o == nil || IsNil(o.RudderComputeDyngroupsMaxParallelism) {
		return nil, false
	}
	return o.RudderComputeDyngroupsMaxParallelism, true
}

// HasRudderComputeDyngroupsMaxParallelism returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasRudderComputeDyngroupsMaxParallelism() bool {
	if o != nil && !IsNil(o.RudderComputeDyngroupsMaxParallelism) {
		return true
	}

	return false
}

// SetRudderComputeDyngroupsMaxParallelism gets a reference to the given string and assigns it to the RudderComputeDyngroupsMaxParallelism field.
func (o *GetAllSettings200ResponseDataSettings) SetRudderComputeDyngroupsMaxParallelism(v string) {
	o.RudderComputeDyngroupsMaxParallelism = &v
}

// GetRudderSaveDbComplianceLevels returns the RudderSaveDbComplianceLevels field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetRudderSaveDbComplianceLevels() bool {
	if o == nil || IsNil(o.RudderSaveDbComplianceLevels) {
		var ret bool
		return ret
	}
	return *o.RudderSaveDbComplianceLevels
}

// GetRudderSaveDbComplianceLevelsOk returns a tuple with the RudderSaveDbComplianceLevels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetRudderSaveDbComplianceLevelsOk() (*bool, bool) {
	if o == nil || IsNil(o.RudderSaveDbComplianceLevels) {
		return nil, false
	}
	return o.RudderSaveDbComplianceLevels, true
}

// HasRudderSaveDbComplianceLevels returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasRudderSaveDbComplianceLevels() bool {
	if o != nil && !IsNil(o.RudderSaveDbComplianceLevels) {
		return true
	}

	return false
}

// SetRudderSaveDbComplianceLevels gets a reference to the given bool and assigns it to the RudderSaveDbComplianceLevels field.
func (o *GetAllSettings200ResponseDataSettings) SetRudderSaveDbComplianceLevels(v bool) {
	o.RudderSaveDbComplianceLevels = &v
}

// GetRudderSaveDbComplianceDetails returns the RudderSaveDbComplianceDetails field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetRudderSaveDbComplianceDetails() bool {
	if o == nil || IsNil(o.RudderSaveDbComplianceDetails) {
		var ret bool
		return ret
	}
	return *o.RudderSaveDbComplianceDetails
}

// GetRudderSaveDbComplianceDetailsOk returns a tuple with the RudderSaveDbComplianceDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetRudderSaveDbComplianceDetailsOk() (*bool, bool) {
	if o == nil || IsNil(o.RudderSaveDbComplianceDetails) {
		return nil, false
	}
	return o.RudderSaveDbComplianceDetails, true
}

// HasRudderSaveDbComplianceDetails returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasRudderSaveDbComplianceDetails() bool {
	if o != nil && !IsNil(o.RudderSaveDbComplianceDetails) {
		return true
	}

	return false
}

// SetRudderSaveDbComplianceDetails gets a reference to the given bool and assigns it to the RudderSaveDbComplianceDetails field.
func (o *GetAllSettings200ResponseDataSettings) SetRudderSaveDbComplianceDetails(v bool) {
	o.RudderSaveDbComplianceDetails = &v
}

// GetRudderGenerationMaxParallelism returns the RudderGenerationMaxParallelism field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetRudderGenerationMaxParallelism() string {
	if o == nil || IsNil(o.RudderGenerationMaxParallelism) {
		var ret string
		return ret
	}
	return *o.RudderGenerationMaxParallelism
}

// GetRudderGenerationMaxParallelismOk returns a tuple with the RudderGenerationMaxParallelism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetRudderGenerationMaxParallelismOk() (*string, bool) {
	if o == nil || IsNil(o.RudderGenerationMaxParallelism) {
		return nil, false
	}
	return o.RudderGenerationMaxParallelism, true
}

// HasRudderGenerationMaxParallelism returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasRudderGenerationMaxParallelism() bool {
	if o != nil && !IsNil(o.RudderGenerationMaxParallelism) {
		return true
	}

	return false
}

// SetRudderGenerationMaxParallelism gets a reference to the given string and assigns it to the RudderGenerationMaxParallelism field.
func (o *GetAllSettings200ResponseDataSettings) SetRudderGenerationMaxParallelism(v string) {
	o.RudderGenerationMaxParallelism = &v
}

// GetRudderGenerationJsTimeout returns the RudderGenerationJsTimeout field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetRudderGenerationJsTimeout() int32 {
	if o == nil || IsNil(o.RudderGenerationJsTimeout) {
		var ret int32
		return ret
	}
	return *o.RudderGenerationJsTimeout
}

// GetRudderGenerationJsTimeoutOk returns a tuple with the RudderGenerationJsTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetRudderGenerationJsTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.RudderGenerationJsTimeout) {
		return nil, false
	}
	return o.RudderGenerationJsTimeout, true
}

// HasRudderGenerationJsTimeout returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasRudderGenerationJsTimeout() bool {
	if o != nil && !IsNil(o.RudderGenerationJsTimeout) {
		return true
	}

	return false
}

// SetRudderGenerationJsTimeout gets a reference to the given int32 and assigns it to the RudderGenerationJsTimeout field.
func (o *GetAllSettings200ResponseDataSettings) SetRudderGenerationJsTimeout(v int32) {
	o.RudderGenerationJsTimeout = &v
}

// GetRudderGenerationContinueOnError returns the RudderGenerationContinueOnError field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetRudderGenerationContinueOnError() bool {
	if o == nil || IsNil(o.RudderGenerationContinueOnError) {
		var ret bool
		return ret
	}
	return *o.RudderGenerationContinueOnError
}

// GetRudderGenerationContinueOnErrorOk returns a tuple with the RudderGenerationContinueOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetRudderGenerationContinueOnErrorOk() (*bool, bool) {
	if o == nil || IsNil(o.RudderGenerationContinueOnError) {
		return nil, false
	}
	return o.RudderGenerationContinueOnError, true
}

// HasRudderGenerationContinueOnError returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasRudderGenerationContinueOnError() bool {
	if o != nil && !IsNil(o.RudderGenerationContinueOnError) {
		return true
	}

	return false
}

// SetRudderGenerationContinueOnError gets a reference to the given bool and assigns it to the RudderGenerationContinueOnError field.
func (o *GetAllSettings200ResponseDataSettings) SetRudderGenerationContinueOnError(v bool) {
	o.RudderGenerationContinueOnError = &v
}

// GetRudderGenerationDelay returns the RudderGenerationDelay field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetRudderGenerationDelay() string {
	if o == nil || IsNil(o.RudderGenerationDelay) {
		var ret string
		return ret
	}
	return *o.RudderGenerationDelay
}

// GetRudderGenerationDelayOk returns a tuple with the RudderGenerationDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetRudderGenerationDelayOk() (*string, bool) {
	if o == nil || IsNil(o.RudderGenerationDelay) {
		return nil, false
	}
	return o.RudderGenerationDelay, true
}

// HasRudderGenerationDelay returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasRudderGenerationDelay() bool {
	if o != nil && !IsNil(o.RudderGenerationDelay) {
		return true
	}

	return false
}

// SetRudderGenerationDelay gets a reference to the given string and assigns it to the RudderGenerationDelay field.
func (o *GetAllSettings200ResponseDataSettings) SetRudderGenerationDelay(v string) {
	o.RudderGenerationDelay = &v
}

// GetRudderGenerationPolicy returns the RudderGenerationPolicy field value if set, zero value otherwise.
func (o *GetAllSettings200ResponseDataSettings) GetRudderGenerationPolicy() string {
	if o == nil || IsNil(o.RudderGenerationPolicy) {
		var ret string
		return ret
	}
	return *o.RudderGenerationPolicy
}

// GetRudderGenerationPolicyOk returns a tuple with the RudderGenerationPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllSettings200ResponseDataSettings) GetRudderGenerationPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.RudderGenerationPolicy) {
		return nil, false
	}
	return o.RudderGenerationPolicy, true
}

// HasRudderGenerationPolicy returns a boolean if a field has been set.
func (o *GetAllSettings200ResponseDataSettings) HasRudderGenerationPolicy() bool {
	if o != nil && !IsNil(o.RudderGenerationPolicy) {
		return true
	}

	return false
}

// SetRudderGenerationPolicy gets a reference to the given string and assigns it to the RudderGenerationPolicy field.
func (o *GetAllSettings200ResponseDataSettings) SetRudderGenerationPolicy(v string) {
	o.RudderGenerationPolicy = &v
}

func (o GetAllSettings200ResponseDataSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllSettings200ResponseDataSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowedNetworks) {
		toSerialize["allowed_networks"] = o.AllowedNetworks
	}
	if !IsNil(o.GlobalPolicyMode) {
		toSerialize["global_policy_mode"] = o.GlobalPolicyMode
	}
	if !IsNil(o.GlobalPolicyModeOverridable) {
		toSerialize["global_policy_mode_overridable"] = o.GlobalPolicyModeOverridable
	}
	if !IsNil(o.RunFrequency) {
		toSerialize["run_frequency"] = o.RunFrequency
	}
	if !IsNil(o.FirstRunHour) {
		toSerialize["first_run_hour"] = o.FirstRunHour
	}
	if !IsNil(o.FirstRunMinute) {
		toSerialize["first_run_minute"] = o.FirstRunMinute
	}
	if !IsNil(o.SplayTime) {
		toSerialize["splay_time"] = o.SplayTime
	}
	if !IsNil(o.ModifiedFileTtl) {
		toSerialize["modified_file_ttl"] = o.ModifiedFileTtl
	}
	if !IsNil(o.OutputFileTtl) {
		toSerialize["output_file_ttl"] = o.OutputFileTtl
	}
	if !IsNil(o.RequireTimeSynchronization) {
		toSerialize["require_time_synchronization"] = o.RequireTimeSynchronization
	}
	if !IsNil(o.RelayServerSynchronizationMethod) {
		toSerialize["relay_server_synchronization_method"] = o.RelayServerSynchronizationMethod
	}
	if !IsNil(o.RelayServerSynchronizePolicies) {
		toSerialize["relay_server_synchronize_policies"] = o.RelayServerSynchronizePolicies
	}
	if !IsNil(o.RelayServerSynchronizeSharedFiles) {
		toSerialize["relay_server_synchronize_shared_files"] = o.RelayServerSynchronizeSharedFiles
	}
	if !IsNil(o.RudderReportProtocolDefault) {
		toSerialize["rudder_report_protocol_default"] = o.RudderReportProtocolDefault
	}
	if !IsNil(o.ReportingMode) {
		toSerialize["reporting_mode"] = o.ReportingMode
	}
	if !IsNil(o.HeartbeatFrequency) {
		toSerialize["heartbeat_frequency"] = o.HeartbeatFrequency
	}
	if !IsNil(o.EnableChangeMessage) {
		toSerialize["enable_change_message"] = o.EnableChangeMessage
	}
	if !IsNil(o.MandatoryChangeMessage) {
		toSerialize["mandatory_change_message"] = o.MandatoryChangeMessage
	}
	if !IsNil(o.ChangeMessagePrompt) {
		toSerialize["change_message_prompt"] = o.ChangeMessagePrompt
	}
	if !IsNil(o.EnableChangeRequest) {
		toSerialize["enable_change_request"] = o.EnableChangeRequest
	}
	if !IsNil(o.EnableSelfValidation) {
		toSerialize["enable_self_validation"] = o.EnableSelfValidation
	}
	if !IsNil(o.EnableSelfDeployment) {
		toSerialize["enable_self_deployment"] = o.EnableSelfDeployment
	}
	if !IsNil(o.DisplayRecentChangesGraphs) {
		toSerialize["display_recent_changes_graphs"] = o.DisplayRecentChangesGraphs
	}
	if !IsNil(o.EnableJavascriptDirectives) {
		toSerialize["enable_javascript_directives"] = o.EnableJavascriptDirectives
	}
	if !IsNil(o.SendMetrics) {
		toSerialize["send_metrics"] = o.SendMetrics
	}
	if !IsNil(o.NodeAcceptDuplicatedHostname) {
		toSerialize["node_accept_duplicated_hostname"] = o.NodeAcceptDuplicatedHostname
	}
	if !IsNil(o.NodeOnacceptDefaultState) {
		toSerialize["node_onaccept_default_state"] = o.NodeOnacceptDefaultState
	}
	if !IsNil(o.NodeOnacceptDefaultPolicyMode) {
		toSerialize["node_onaccept_default_policyMode"] = o.NodeOnacceptDefaultPolicyMode
	}
	if !IsNil(o.UnexpectedUnboundVarValues) {
		toSerialize["unexpected_unbound_var_values"] = o.UnexpectedUnboundVarValues
	}
	if !IsNil(o.RudderComputeChanges) {
		toSerialize["rudder_compute_changes"] = o.RudderComputeChanges
	}
	if !IsNil(o.RudderGenerationComputeDyngroups) {
		toSerialize["rudder_generation_compute_dyngroups"] = o.RudderGenerationComputeDyngroups
	}
	if !IsNil(o.RudderComputeDyngroupsMaxParallelism) {
		toSerialize["rudder_compute_dyngroups_max_parallelism"] = o.RudderComputeDyngroupsMaxParallelism
	}
	if !IsNil(o.RudderSaveDbComplianceLevels) {
		toSerialize["rudder_save_db_compliance_levels"] = o.RudderSaveDbComplianceLevels
	}
	if !IsNil(o.RudderSaveDbComplianceDetails) {
		toSerialize["rudder_save_db_compliance_details"] = o.RudderSaveDbComplianceDetails
	}
	if !IsNil(o.RudderGenerationMaxParallelism) {
		toSerialize["rudder_generation_max_parallelism"] = o.RudderGenerationMaxParallelism
	}
	if !IsNil(o.RudderGenerationJsTimeout) {
		toSerialize["rudder_generation_js_timeout"] = o.RudderGenerationJsTimeout
	}
	if !IsNil(o.RudderGenerationContinueOnError) {
		toSerialize["rudder_generation_continue_on_error"] = o.RudderGenerationContinueOnError
	}
	if !IsNil(o.RudderGenerationDelay) {
		toSerialize["rudder_generation_delay"] = o.RudderGenerationDelay
	}
	if !IsNil(o.RudderGenerationPolicy) {
		toSerialize["rudder_generation_policy"] = o.RudderGenerationPolicy
	}
	return toSerialize, nil
}

type NullableGetAllSettings200ResponseDataSettings struct {
	value *GetAllSettings200ResponseDataSettings
	isSet bool
}

func (v NullableGetAllSettings200ResponseDataSettings) Get() *GetAllSettings200ResponseDataSettings {
	return v.value
}

func (v *NullableGetAllSettings200ResponseDataSettings) Set(val *GetAllSettings200ResponseDataSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllSettings200ResponseDataSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllSettings200ResponseDataSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllSettings200ResponseDataSettings(val *GetAllSettings200ResponseDataSettings) *NullableGetAllSettings200ResponseDataSettings {
	return &NullableGetAllSettings200ResponseDataSettings{value: val, isSet: true}
}

func (v NullableGetAllSettings200ResponseDataSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllSettings200ResponseDataSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


