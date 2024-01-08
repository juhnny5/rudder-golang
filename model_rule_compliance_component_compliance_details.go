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

// checks if the RuleComplianceComponentComplianceDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleComplianceComponentComplianceDetails{}

// RuleComplianceComponentComplianceDetails struct for RuleComplianceComponentComplianceDetails
type RuleComplianceComponentComplianceDetails struct {
	SuccessAlreadyOK *float32 `json:"successAlreadyOK,omitempty"`
	NoReport *float32 `json:"noReport,omitempty"`
	SuccessNotApplicable *float32 `json:"successNotApplicable,omitempty"`
	UnexpectedMissingComponent *float32 `json:"unexpectedMissingComponent,omitempty"`
	Error *float32 `json:"error,omitempty"`
	UnexpectedUnknownComponent *float32 `json:"unexpectedUnknownComponent,omitempty"`
	SuccessRepaired *float32 `json:"successRepaired,omitempty"`
}

// NewRuleComplianceComponentComplianceDetails instantiates a new RuleComplianceComponentComplianceDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleComplianceComponentComplianceDetails() *RuleComplianceComponentComplianceDetails {
	this := RuleComplianceComponentComplianceDetails{}
	return &this
}

// NewRuleComplianceComponentComplianceDetailsWithDefaults instantiates a new RuleComplianceComponentComplianceDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleComplianceComponentComplianceDetailsWithDefaults() *RuleComplianceComponentComplianceDetails {
	this := RuleComplianceComponentComplianceDetails{}
	return &this
}

// GetSuccessAlreadyOK returns the SuccessAlreadyOK field value if set, zero value otherwise.
func (o *RuleComplianceComponentComplianceDetails) GetSuccessAlreadyOK() float32 {
	if o == nil || IsNil(o.SuccessAlreadyOK) {
		var ret float32
		return ret
	}
	return *o.SuccessAlreadyOK
}

// GetSuccessAlreadyOKOk returns a tuple with the SuccessAlreadyOK field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleComplianceComponentComplianceDetails) GetSuccessAlreadyOKOk() (*float32, bool) {
	if o == nil || IsNil(o.SuccessAlreadyOK) {
		return nil, false
	}
	return o.SuccessAlreadyOK, true
}

// HasSuccessAlreadyOK returns a boolean if a field has been set.
func (o *RuleComplianceComponentComplianceDetails) HasSuccessAlreadyOK() bool {
	if o != nil && !IsNil(o.SuccessAlreadyOK) {
		return true
	}

	return false
}

// SetSuccessAlreadyOK gets a reference to the given float32 and assigns it to the SuccessAlreadyOK field.
func (o *RuleComplianceComponentComplianceDetails) SetSuccessAlreadyOK(v float32) {
	o.SuccessAlreadyOK = &v
}

// GetNoReport returns the NoReport field value if set, zero value otherwise.
func (o *RuleComplianceComponentComplianceDetails) GetNoReport() float32 {
	if o == nil || IsNil(o.NoReport) {
		var ret float32
		return ret
	}
	return *o.NoReport
}

// GetNoReportOk returns a tuple with the NoReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleComplianceComponentComplianceDetails) GetNoReportOk() (*float32, bool) {
	if o == nil || IsNil(o.NoReport) {
		return nil, false
	}
	return o.NoReport, true
}

// HasNoReport returns a boolean if a field has been set.
func (o *RuleComplianceComponentComplianceDetails) HasNoReport() bool {
	if o != nil && !IsNil(o.NoReport) {
		return true
	}

	return false
}

// SetNoReport gets a reference to the given float32 and assigns it to the NoReport field.
func (o *RuleComplianceComponentComplianceDetails) SetNoReport(v float32) {
	o.NoReport = &v
}

// GetSuccessNotApplicable returns the SuccessNotApplicable field value if set, zero value otherwise.
func (o *RuleComplianceComponentComplianceDetails) GetSuccessNotApplicable() float32 {
	if o == nil || IsNil(o.SuccessNotApplicable) {
		var ret float32
		return ret
	}
	return *o.SuccessNotApplicable
}

// GetSuccessNotApplicableOk returns a tuple with the SuccessNotApplicable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleComplianceComponentComplianceDetails) GetSuccessNotApplicableOk() (*float32, bool) {
	if o == nil || IsNil(o.SuccessNotApplicable) {
		return nil, false
	}
	return o.SuccessNotApplicable, true
}

// HasSuccessNotApplicable returns a boolean if a field has been set.
func (o *RuleComplianceComponentComplianceDetails) HasSuccessNotApplicable() bool {
	if o != nil && !IsNil(o.SuccessNotApplicable) {
		return true
	}

	return false
}

// SetSuccessNotApplicable gets a reference to the given float32 and assigns it to the SuccessNotApplicable field.
func (o *RuleComplianceComponentComplianceDetails) SetSuccessNotApplicable(v float32) {
	o.SuccessNotApplicable = &v
}

// GetUnexpectedMissingComponent returns the UnexpectedMissingComponent field value if set, zero value otherwise.
func (o *RuleComplianceComponentComplianceDetails) GetUnexpectedMissingComponent() float32 {
	if o == nil || IsNil(o.UnexpectedMissingComponent) {
		var ret float32
		return ret
	}
	return *o.UnexpectedMissingComponent
}

// GetUnexpectedMissingComponentOk returns a tuple with the UnexpectedMissingComponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleComplianceComponentComplianceDetails) GetUnexpectedMissingComponentOk() (*float32, bool) {
	if o == nil || IsNil(o.UnexpectedMissingComponent) {
		return nil, false
	}
	return o.UnexpectedMissingComponent, true
}

// HasUnexpectedMissingComponent returns a boolean if a field has been set.
func (o *RuleComplianceComponentComplianceDetails) HasUnexpectedMissingComponent() bool {
	if o != nil && !IsNil(o.UnexpectedMissingComponent) {
		return true
	}

	return false
}

// SetUnexpectedMissingComponent gets a reference to the given float32 and assigns it to the UnexpectedMissingComponent field.
func (o *RuleComplianceComponentComplianceDetails) SetUnexpectedMissingComponent(v float32) {
	o.UnexpectedMissingComponent = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *RuleComplianceComponentComplianceDetails) GetError() float32 {
	if o == nil || IsNil(o.Error) {
		var ret float32
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleComplianceComponentComplianceDetails) GetErrorOk() (*float32, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *RuleComplianceComponentComplianceDetails) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given float32 and assigns it to the Error field.
func (o *RuleComplianceComponentComplianceDetails) SetError(v float32) {
	o.Error = &v
}

// GetUnexpectedUnknownComponent returns the UnexpectedUnknownComponent field value if set, zero value otherwise.
func (o *RuleComplianceComponentComplianceDetails) GetUnexpectedUnknownComponent() float32 {
	if o == nil || IsNil(o.UnexpectedUnknownComponent) {
		var ret float32
		return ret
	}
	return *o.UnexpectedUnknownComponent
}

// GetUnexpectedUnknownComponentOk returns a tuple with the UnexpectedUnknownComponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleComplianceComponentComplianceDetails) GetUnexpectedUnknownComponentOk() (*float32, bool) {
	if o == nil || IsNil(o.UnexpectedUnknownComponent) {
		return nil, false
	}
	return o.UnexpectedUnknownComponent, true
}

// HasUnexpectedUnknownComponent returns a boolean if a field has been set.
func (o *RuleComplianceComponentComplianceDetails) HasUnexpectedUnknownComponent() bool {
	if o != nil && !IsNil(o.UnexpectedUnknownComponent) {
		return true
	}

	return false
}

// SetUnexpectedUnknownComponent gets a reference to the given float32 and assigns it to the UnexpectedUnknownComponent field.
func (o *RuleComplianceComponentComplianceDetails) SetUnexpectedUnknownComponent(v float32) {
	o.UnexpectedUnknownComponent = &v
}

// GetSuccessRepaired returns the SuccessRepaired field value if set, zero value otherwise.
func (o *RuleComplianceComponentComplianceDetails) GetSuccessRepaired() float32 {
	if o == nil || IsNil(o.SuccessRepaired) {
		var ret float32
		return ret
	}
	return *o.SuccessRepaired
}

// GetSuccessRepairedOk returns a tuple with the SuccessRepaired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleComplianceComponentComplianceDetails) GetSuccessRepairedOk() (*float32, bool) {
	if o == nil || IsNil(o.SuccessRepaired) {
		return nil, false
	}
	return o.SuccessRepaired, true
}

// HasSuccessRepaired returns a boolean if a field has been set.
func (o *RuleComplianceComponentComplianceDetails) HasSuccessRepaired() bool {
	if o != nil && !IsNil(o.SuccessRepaired) {
		return true
	}

	return false
}

// SetSuccessRepaired gets a reference to the given float32 and assigns it to the SuccessRepaired field.
func (o *RuleComplianceComponentComplianceDetails) SetSuccessRepaired(v float32) {
	o.SuccessRepaired = &v
}

func (o RuleComplianceComponentComplianceDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleComplianceComponentComplianceDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SuccessAlreadyOK) {
		toSerialize["successAlreadyOK"] = o.SuccessAlreadyOK
	}
	if !IsNil(o.NoReport) {
		toSerialize["noReport"] = o.NoReport
	}
	if !IsNil(o.SuccessNotApplicable) {
		toSerialize["successNotApplicable"] = o.SuccessNotApplicable
	}
	if !IsNil(o.UnexpectedMissingComponent) {
		toSerialize["unexpectedMissingComponent"] = o.UnexpectedMissingComponent
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.UnexpectedUnknownComponent) {
		toSerialize["unexpectedUnknownComponent"] = o.UnexpectedUnknownComponent
	}
	if !IsNil(o.SuccessRepaired) {
		toSerialize["successRepaired"] = o.SuccessRepaired
	}
	return toSerialize, nil
}

type NullableRuleComplianceComponentComplianceDetails struct {
	value *RuleComplianceComponentComplianceDetails
	isSet bool
}

func (v NullableRuleComplianceComponentComplianceDetails) Get() *RuleComplianceComponentComplianceDetails {
	return v.value
}

func (v *NullableRuleComplianceComponentComplianceDetails) Set(val *RuleComplianceComponentComplianceDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleComplianceComponentComplianceDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleComplianceComponentComplianceDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleComplianceComponentComplianceDetails(val *RuleComplianceComponentComplianceDetails) *NullableRuleComplianceComponentComplianceDetails {
	return &NullableRuleComplianceComponentComplianceDetails{value: val, isSet: true}
}

func (v NullableRuleComplianceComponentComplianceDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleComplianceComponentComplianceDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


