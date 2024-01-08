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

// checks if the DirectiveNew type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DirectiveNew{}

// DirectiveNew struct for DirectiveNew
type DirectiveNew struct {
	// The id of the directive the clone will be based onto. If this parameter if provided,  the new directive will be a clone of this source. Other value will override values from the source.
	Source *string `json:"source,omitempty"`
	// Directive id
	Id *string `json:"id,omitempty"`
	// Human readable name of the directive
	DisplayName *string `json:"displayName,omitempty"`
	// One line directive description
	ShortDescription *string `json:"shortDescription,omitempty"`
	// Description of the technique (rendered as markdown)
	LongDescription *string `json:"longDescription,omitempty"`
	// Directive id
	TechniqueName *string `json:"techniqueName,omitempty"`
	// Directive id
	TechniqueVersion *string `json:"techniqueVersion,omitempty"`
	// Directive priority. `0` has highest priority.
	Priority *int32 `json:"priority,omitempty"`
	// Is the directive enabled
	Enabled *bool `json:"enabled,omitempty"`
	// If true it is an internal Rudder directive
	System *bool `json:"system,omitempty"`
	Tags []DirectiveTagsInner `json:"tags,omitempty"`
	// Directive parameters (depends on the source technique)
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// NewDirectiveNew instantiates a new DirectiveNew object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectiveNew() *DirectiveNew {
	this := DirectiveNew{}
	return &this
}

// NewDirectiveNewWithDefaults instantiates a new DirectiveNew object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectiveNewWithDefaults() *DirectiveNew {
	this := DirectiveNew{}
	return &this
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *DirectiveNew) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectiveNew) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *DirectiveNew) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *DirectiveNew) SetSource(v string) {
	o.Source = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DirectiveNew) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectiveNew) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DirectiveNew) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DirectiveNew) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *DirectiveNew) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectiveNew) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *DirectiveNew) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *DirectiveNew) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetShortDescription returns the ShortDescription field value if set, zero value otherwise.
func (o *DirectiveNew) GetShortDescription() string {
	if o == nil || IsNil(o.ShortDescription) {
		var ret string
		return ret
	}
	return *o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectiveNew) GetShortDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ShortDescription) {
		return nil, false
	}
	return o.ShortDescription, true
}

// HasShortDescription returns a boolean if a field has been set.
func (o *DirectiveNew) HasShortDescription() bool {
	if o != nil && !IsNil(o.ShortDescription) {
		return true
	}

	return false
}

// SetShortDescription gets a reference to the given string and assigns it to the ShortDescription field.
func (o *DirectiveNew) SetShortDescription(v string) {
	o.ShortDescription = &v
}

// GetLongDescription returns the LongDescription field value if set, zero value otherwise.
func (o *DirectiveNew) GetLongDescription() string {
	if o == nil || IsNil(o.LongDescription) {
		var ret string
		return ret
	}
	return *o.LongDescription
}

// GetLongDescriptionOk returns a tuple with the LongDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectiveNew) GetLongDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.LongDescription) {
		return nil, false
	}
	return o.LongDescription, true
}

// HasLongDescription returns a boolean if a field has been set.
func (o *DirectiveNew) HasLongDescription() bool {
	if o != nil && !IsNil(o.LongDescription) {
		return true
	}

	return false
}

// SetLongDescription gets a reference to the given string and assigns it to the LongDescription field.
func (o *DirectiveNew) SetLongDescription(v string) {
	o.LongDescription = &v
}

// GetTechniqueName returns the TechniqueName field value if set, zero value otherwise.
func (o *DirectiveNew) GetTechniqueName() string {
	if o == nil || IsNil(o.TechniqueName) {
		var ret string
		return ret
	}
	return *o.TechniqueName
}

// GetTechniqueNameOk returns a tuple with the TechniqueName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectiveNew) GetTechniqueNameOk() (*string, bool) {
	if o == nil || IsNil(o.TechniqueName) {
		return nil, false
	}
	return o.TechniqueName, true
}

// HasTechniqueName returns a boolean if a field has been set.
func (o *DirectiveNew) HasTechniqueName() bool {
	if o != nil && !IsNil(o.TechniqueName) {
		return true
	}

	return false
}

// SetTechniqueName gets a reference to the given string and assigns it to the TechniqueName field.
func (o *DirectiveNew) SetTechniqueName(v string) {
	o.TechniqueName = &v
}

// GetTechniqueVersion returns the TechniqueVersion field value if set, zero value otherwise.
func (o *DirectiveNew) GetTechniqueVersion() string {
	if o == nil || IsNil(o.TechniqueVersion) {
		var ret string
		return ret
	}
	return *o.TechniqueVersion
}

// GetTechniqueVersionOk returns a tuple with the TechniqueVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectiveNew) GetTechniqueVersionOk() (*string, bool) {
	if o == nil || IsNil(o.TechniqueVersion) {
		return nil, false
	}
	return o.TechniqueVersion, true
}

// HasTechniqueVersion returns a boolean if a field has been set.
func (o *DirectiveNew) HasTechniqueVersion() bool {
	if o != nil && !IsNil(o.TechniqueVersion) {
		return true
	}

	return false
}

// SetTechniqueVersion gets a reference to the given string and assigns it to the TechniqueVersion field.
func (o *DirectiveNew) SetTechniqueVersion(v string) {
	o.TechniqueVersion = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *DirectiveNew) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectiveNew) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *DirectiveNew) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *DirectiveNew) SetPriority(v int32) {
	o.Priority = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *DirectiveNew) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectiveNew) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *DirectiveNew) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *DirectiveNew) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *DirectiveNew) GetSystem() bool {
	if o == nil || IsNil(o.System) {
		var ret bool
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectiveNew) GetSystemOk() (*bool, bool) {
	if o == nil || IsNil(o.System) {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *DirectiveNew) HasSystem() bool {
	if o != nil && !IsNil(o.System) {
		return true
	}

	return false
}

// SetSystem gets a reference to the given bool and assigns it to the System field.
func (o *DirectiveNew) SetSystem(v bool) {
	o.System = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DirectiveNew) GetTags() []DirectiveTagsInner {
	if o == nil || IsNil(o.Tags) {
		var ret []DirectiveTagsInner
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectiveNew) GetTagsOk() ([]DirectiveTagsInner, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DirectiveNew) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []DirectiveTagsInner and assigns it to the Tags field.
func (o *DirectiveNew) SetTags(v []DirectiveTagsInner) {
	o.Tags = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *DirectiveNew) GetParameters() map[string]interface{} {
	if o == nil || IsNil(o.Parameters) {
		var ret map[string]interface{}
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectiveNew) GetParametersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Parameters) {
		return map[string]interface{}{}, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *DirectiveNew) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]interface{} and assigns it to the Parameters field.
func (o *DirectiveNew) SetParameters(v map[string]interface{}) {
	o.Parameters = v
}

func (o DirectiveNew) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DirectiveNew) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.ShortDescription) {
		toSerialize["shortDescription"] = o.ShortDescription
	}
	if !IsNil(o.LongDescription) {
		toSerialize["longDescription"] = o.LongDescription
	}
	if !IsNil(o.TechniqueName) {
		toSerialize["techniqueName"] = o.TechniqueName
	}
	if !IsNil(o.TechniqueVersion) {
		toSerialize["techniqueVersion"] = o.TechniqueVersion
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.System) {
		toSerialize["system"] = o.System
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	return toSerialize, nil
}

type NullableDirectiveNew struct {
	value *DirectiveNew
	isSet bool
}

func (v NullableDirectiveNew) Get() *DirectiveNew {
	return v.value
}

func (v *NullableDirectiveNew) Set(val *DirectiveNew) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectiveNew) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectiveNew) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectiveNew(val *DirectiveNew) *NullableDirectiveNew {
	return &NullableDirectiveNew{value: val, isSet: true}
}

func (v NullableDirectiveNew) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectiveNew) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


