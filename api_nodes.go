/*
Rudder API

Download OpenAPI specification: [openapi.yml](openapi.yml)  **Other documentation sources**:  * [Main documentation](https://docs.rudder.io) * [Internal relay API](https://docs.rudder.io/api/relay/)  # Introduction  Rudder exposes a REST API, enabling the user to interact with Rudder without using the webapp, for example, in scripts or cron jobs.  ## Authentication  The Rudder REST API uses simple API keys for authentication. All requests must be authenticated (except from a generic status API). The tokens are 32-character strings, passed in a `X-API-Token` header, like in:   ```bash curl --header \"X-API-Token: yourToken\" https://rudder.example.com/rudder/api/latest/rules ```  The tokens are the API equivalent of a password, and must be secured just like a password would be.  ### API accounts  The accounts are managed in the Web interface. There are two possible types of accounts:  * **Global API accounts**: they are not linked to a Rudder user, and are managed by Rudder administrators in the _Administration -> API accounts_ page. You should define an expiration date whenever possible.  ![General API tokens settings](assets/api-tokens.png \"General API tokens settings\")  * **User tokens**: they are linked to a Rudder user, and give the same rights the user has. There can be only one token by user. This feature is provided by the `api-authorizatons` plugin.  ![User API token](assets/api-user.png \"User API token\")  When an action produces a change of configuration on the server, the API account that made it will be recorded in the event log, like for a Web interaction.  ### Authorization  When using Rudder without the `api-authorizatons` plugin, only global accounts are available, with two possible privilege levels, read-only or write. With the `api-authorizatons` plugin, you also get access to:  * User tokens, which have the same permissions as the user, using the Rudder roles and permissions feature. * Custom ACLs on global API accounts. They provide fine-grained permissions on every endpoint:  ![Custom API ACL](assets/custom-acl.png \"Custom API ACL\")  As a general principle, you should create dedicated tokens with the least privilege level for each different interaction you have with the API. This limits the risks of exploitation if a token is stolen, and allows tracking the activity of each token separately. Token renewal is also easier when they are only used for a limited purpose.  ## Versioning  Each time the API is extended with new features (new functions, new parameters, new responses, ...), it will be assigned a new version number. This will allow you to keep your existing scripts (based on previous behavior). Versions will always be integers (no 2.1 or 3.3, just 2, 3, 4, ...) or `latest`.  You can change the version of the API used by setting it either within the url or in a header:  * the URL: each URL is prefixed by its version id, like `/api/version/function`.  ```bash # Version 10 curl -X GET -H \"X-API-Token: yourToken\" https://rudder.example.com/rudder/api/10/rules # Latest curl -X GET -H \"X-API-Token: yourToken\" https://rudder.example.com/rudder/api/latest/rules # Wrong (not an integer) => 404 not found curl -X GET -H \"X-API-Token: yourToken\" https://rudder.example.com/rudder/api/3.14/rules ```  * the HTTP headers. You can add the **X-API-Version** header to your request. The value needs to be an integer or `latest`.  ```bash # Version 10 curl -X GET -H \"X-API-Token: yourToken\" -H \"X-API-Version: 10\" https://rudder.example.com/rudder/api/rules # Wrong => Error response indicating which versions are available curl -X GET -H \"X-API-Token: yourToken\" -H \"X-API-Version: 3.14\" https://rudder.example.com/rudder/api/rules ```  In the future, we may declare some versions as deprecated, in order to remove them in a later version of Rudder, but we will never remove any versions without warning, or without a safe period of time to allow migration from previous versions.   <h4>Existing versions</h4> <table>   <thead>     <tr>       <th style=\"width: 20%\">Version</th>       <th style=\"width: 20%\">Rudder versions it appeared in</th>       <th style=\"width: 70%\">Description</th>     </tr>   </thead>   <tbody>     <tr>       <td class=\"code\">1</td>       <td class=\"code\">Never released (for internal use only)</td>       <td>Experimental version</td>     </tr>     <tr>       <td class=\"code\">2 to 10 (deprecated)</td>       <td class=\"code\">4.3 and before</td>       <td>These versions provided the core set of API features for rules, directives, nodes global parameters, change requests and compliance, rudder settings, and system API</td>     </tr>     <tr>       <td class=\"code\">11</td>       <td class=\"code\">5.0</td>       <td>New system API (replacing old localhost v1 api): status, maintenance operations and server behavior</td>     </tr>     <tr>       <td class=\"code\">12</td>       <td class=\"code\">6.0 and 6.1</td>       <td>Node key management</td>     </tr>     <tr>       <td class=\"code\">13</td>       <td class=\"code\">6.2</td>       <td><ul>         <li>Node status endpoint</li>         <li>System health check</li>         <li>System maintenance job to purge software [that endpoint was back-ported in 6.1]</li>       </ul></td>     </tr>     <tr>       <td class=\"code\">14</td>       <td class=\"code\">7.0</td>       <td><ul>         <li>Secret management</li>         <li>Directive tree</li>         <li>Improve techniques management</li>         <li>Demote a relay</li>       </ul></td>     </tr>     <tr>       <td class=\"code\">15</td>       <td class=\"code\">7.1</td>       <td><ul>         <li>Package updates in nodes</li>       </ul></td>     </tr>     <tr>       <td class=\"code\">16</td>       <td class=\"code\">7.2</td>       <td><ul>         <li>Create node API included from plugin</li>         <li>Configuration archive import/export</li>       </ul></td>     </tr>     <tr>       <td class=\"code\">17</td>       <td class=\"code\">7.3</td>       <td><ul>         <li>Compliance by directive</li>         <li>Path campaigns API included</li>       </ul></td>     </tr>     <tr>       <td class=\"code\">18</td>       <td class=\"code\">8.0</td>       <td><ul>         <li>Allowed network </li>         <li>Improve the structure of `/settings/allowed_networks` output</li>       </ul></td>     </tr>   </tbody> </table>   ## Response format  All responses from the API are in the JSON format.  ```json {   \"action\": \"The name of the called function\",   \"id\": \"The ID of the element you want, if relevant\",   \"result\": \"The result of your action: success or error\",   \"data\": \"Only present if this is a success and depends on the function, it's usually a JSON object\",   \"errorDetails\": \"Only present if this is an error, it contains the error message\" } ```   * __Success__ responses are sent with the 200 HTTP (Success) code  * __Error__ responses are sent with a HTTP error code (mostly 5xx...)   ## HTTP method  Rudder's REST API is based on the usage of [HTTP methods](http://www.w3.org/Protocols/rfc2616/rfc2616-sec9.html). We use them to indicate what action will be done by the request. Currently, we use four of them:   * **GET**: search or retrieve information (get rule details, get a group, ...)  * **PUT**: add new objects (create a directive, clone a Rule, ...)  * **DELETE**: remove objects (delete a node, delete a parameter, ...)  * **POST**: update existing objects (update a directive, reload a group, ...)   ## Parameters  ### General parameters  Some parameters are available for almost all API functions. They will be described in this section. They must be part of the query and can't be submitted in a JSON form.  #### Available for all requests  <table>   <thead>     <tr>       <th style=\"width: 30%\">Field</th>       <th style=\"width: 10%\">Type</th>       <th style=\"width: 70%\">Description</th>     </tr>   </thead>   <tbody>     <tr>       <td class=\"code\">prettify</td>       <td><b>boolean</b><br><i>optional</i></td>       <td>         Determine if the answer should be prettified (human friendly) or not. We recommend using this for debugging purposes, but not for general script usage as this does add some unnecessary load on the server side.         <p class=\"default-value\">Default value: <code>false</code></p>       </td>     </tr>   </tbody> </table>   #### Available for modification requests (PUT/POST/DELETE)  <table>   <thead>     <tr>       <th style=\"width: 25%\">Field</th>       <th style=\"width: 12%\">Type</th>       <th style=\"width: 70%\">Description</th>     </tr>   </thead>   <tbody>     <tr>       <td class=\"code\">reason</td>       <td><b>string</b><br><i>optional</i> or <i>required</i></td>       <td>         Set a message to explain the change. If you set the reason messages to be mandatory in the web interface, failing to supply this value will lead to an error.         <p class=\"default-value\">Default value: <code>\"\"</code></p>       </td>     </tr>     <tr>       <td class=\"code\">changeRequestName</td>       <td><b>string</b><br><i>optional</i></td>       <td>         Set the change request name, is used only if workflows are enabled. The default value depends on the function called         <p class=\"default-value\">Default value: <code>A default string for each function</code></p>       </td>     </tr>     <tr>       <td class=\"code\">changeRequestDescription</td>       <td><b>string</b><br><i>optional</i></td>       <td>         Set the change request description, is used only if workflows are enabled.         <p class=\"default-value\">Default value: <code>\"\"</code></p>       </td>     </tr>   </tbody> </table>   ### Passing parameters  Parameters to the API can be sent:  * As part of the URL for resource identification  * As data for POST/PUT requests    * Directly in JSON format    * As request arguments  #### As part of the URL for resource identification  Parameters in URLs are used to indicate which resource you want to interact with. The function will not work if this resource is missing.  ```bash # Get the Rule of ID \"id\" curl -H \"X-API-Token: yourToken\" https://rudder.example.com/rudder/api/latest/rules/id ```  CAUTION: To avoid surprising behavior, do not put a '/' at the end of a URL: it would be interpreted as '/[empty string parameter]' and redirected to '/index', likely not what you wanted to do.   #### Sending data for POST/PUT requests  ##### Directly in JSON format  JSON format is the preferred way to interact with Rudder API for creating or updating resources. You'll also have to set the *Content-Type* header to **application/json** (without it the JSON content would be ignored). In a `curl` `POST` request, that header can be provided with the `-H` parameter:  ```bash curl -X POST -H \"Content-Type: application/json\" ... ```  The supplied file must contain a valid JSON: strings need quotes, booleans and integers don't, etc.  The (human-readable) format is:  ```json {   \"key1\": \"value1\",   \"key2\": false,   \"key3\": 42 } ```  Here is an example with inlined data:  ```bash # Update the Rule 'id' with a new name, disabled, and setting it one directive curl -X POST -H \"X-API-Token: yourToken\" -H  \"Content-Type: application/json\" https://rudder.example.com/rudder/api/rules/latest/{id}   -d '{ \"displayName\": \"new name\", \"enabled\": false, \"directives\": \"directiveId\"}' ```  You can also pass a supply the JSON in a file:  ```bash # Update the Rule 'id' with a new name, disabled, and setting it one directive curl -X POST -H \"X-API-Token: yourToken\" -H \"Content-Type: application/json\" https://rudder.example.com/rudder/api/rules/latest/{id} -d @jsonParam ```  Note that the general parameters view in the previous chapter cannot be passed in a JSON, and you will need to pass them a URL parameters if you want them to be taken into account (you can't mix JSON and request parameters):  ```bash # Update the Rule 'id' with a new name, disabled, and setting it one directive with reason message \"Reason used\" curl -X POST -H \"X-API-Token: yourToken\" -H \"Content-Type: application/json\" \"https://rudder.example.com/rudder/api/rules/latest/{id}?reason=Reason used\" -d @jsonParam -d \"reason=Reason ignored\" ```  ##### Request parameters  In some cases, when you have little, simple data to update, JSON can feel bloated. In such cases, you can use request parameters. You will need to pass one parameter for each data you want to change.  Parameters follow the following schema:  ``` key=value ```  You can pass parameters by two means:  * As query parameters: At the end of your url, put a **?** then your first parameter and then a **&** before next parameters. In that case, parameters need to be https://en.wikipedia.org/wiki/Percent-encoding[URL encoded]  ```bash # Update the Rule 'id' with a new name, disabled, and setting it one directive curl -X POST -H \"X-API-Token: yourToken\"  https://rudder.example.com/rudder/api/rules/latest/{id}?\"displayName=my new name\"&\"enabled=false\"&\"directives=aDirectiveId\" ```  * As request data: You can pass those parameters in the request data, they won't figure in the URL, making it lighter to read, You can pass a file that contains data.  ```bash # Update the Rule 'id' with a new name, disabled, and setting it one directive (in file directive-info.json) curl -X POST -H \"X-API-Token: yourToken\" https://rudder.example.com/rudder/api/rules/latest/{id} -d \"displayName=my new name\" -d \"enabled=false\" -d @directive-info.json ``` 

API version: 18
Contact: dev@rudder.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rudder-golang

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// NodesAPIService NodesAPI service
type NodesAPIService service

type ApiApplyPolicyRequest struct {
	ctx context.Context
	ApiService *NodesAPIService
	nodeId string
}

func (r ApiApplyPolicyRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.ApplyPolicyExecute(r)
}

/*
ApplyPolicy Trigger an agent run

This API allows to trigger an agent run on the target node. Response is a stream of the actual agent run on the node.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param nodeId Id of the target node
 @return ApiApplyPolicyRequest
*/
func (a *NodesAPIService) ApplyPolicy(ctx context.Context, nodeId string) ApiApplyPolicyRequest {
	return ApiApplyPolicyRequest{
		ApiService: a,
		ctx: ctx,
		nodeId: nodeId,
	}
}

// Execute executes the request
//  @return string
func (a *NodesAPIService) ApplyPolicyExecute(r ApiApplyPolicyRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NodesAPIService.ApplyPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/nodes/{nodeId}/applyPolicy"
	localVarPath = strings.Replace(localVarPath, "{"+"nodeId"+"}", url.PathEscape(parameterValueToString(r.nodeId, "nodeId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["API-Tokens"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApplyPolicyAllNodesRequest struct {
	ctx context.Context
	ApiService *NodesAPIService
}

func (r ApiApplyPolicyAllNodesRequest) Execute() (*ApplyPolicyAllNodes200Response, *http.Response, error) {
	return r.ApiService.ApplyPolicyAllNodesExecute(r)
}

/*
ApplyPolicyAllNodes Trigger an agent run on all nodes

This API allows to trigger an agent run on all nodes. Response contains a json stating if agent could be started on each node, but not if the run went fine and do not display any output from it. You can see the result of the run in Rudder web interface or in the each agent logs.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApplyPolicyAllNodesRequest
*/
func (a *NodesAPIService) ApplyPolicyAllNodes(ctx context.Context) ApiApplyPolicyAllNodesRequest {
	return ApiApplyPolicyAllNodesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApplyPolicyAllNodes200Response
func (a *NodesAPIService) ApplyPolicyAllNodesExecute(r ApiApplyPolicyAllNodesRequest) (*ApplyPolicyAllNodes200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApplyPolicyAllNodes200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NodesAPIService.ApplyPolicyAllNodes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/nodes/applyPolicy"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["API-Tokens"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiChangePendingNodeStatusRequest struct {
	ctx context.Context
	ApiService *NodesAPIService
	nodeId string
	changePendingNodeStatusRequest *ChangePendingNodeStatusRequest
}

func (r ApiChangePendingNodeStatusRequest) ChangePendingNodeStatusRequest(changePendingNodeStatusRequest ChangePendingNodeStatusRequest) ApiChangePendingNodeStatusRequest {
	r.changePendingNodeStatusRequest = &changePendingNodeStatusRequest
	return r
}

func (r ApiChangePendingNodeStatusRequest) Execute() (*ChangePendingNodeStatus200Response, *http.Response, error) {
	return r.ApiService.ChangePendingNodeStatusExecute(r)
}

/*
ChangePendingNodeStatus Update pending Node status

Accept or refuse a pending node

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param nodeId Id of the target node
 @return ApiChangePendingNodeStatusRequest
*/
func (a *NodesAPIService) ChangePendingNodeStatus(ctx context.Context, nodeId string) ApiChangePendingNodeStatusRequest {
	return ApiChangePendingNodeStatusRequest{
		ApiService: a,
		ctx: ctx,
		nodeId: nodeId,
	}
}

// Execute executes the request
//  @return ChangePendingNodeStatus200Response
func (a *NodesAPIService) ChangePendingNodeStatusExecute(r ApiChangePendingNodeStatusRequest) (*ChangePendingNodeStatus200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ChangePendingNodeStatus200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NodesAPIService.ChangePendingNodeStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/nodes/pending/{nodeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"nodeId"+"}", url.PathEscape(parameterValueToString(r.nodeId, "nodeId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.changePendingNodeStatusRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["API-Tokens"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreateNodesRequest struct {
	ctx context.Context
	ApiService *NodesAPIService
	nodeAddInner *[]NodeAddInner
}

func (r ApiCreateNodesRequest) NodeAddInner(nodeAddInner []NodeAddInner) ApiCreateNodesRequest {
	r.nodeAddInner = &nodeAddInner
	return r
}

func (r ApiCreateNodesRequest) Execute() (*CreateNodes200Response, *http.Response, error) {
	return r.ApiService.CreateNodesExecute(r)
}

/*
CreateNodes Create one or several new nodes

Use the provided array of node information to create new nodes

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateNodesRequest
*/
func (a *NodesAPIService) CreateNodes(ctx context.Context) ApiCreateNodesRequest {
	return ApiCreateNodesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateNodes200Response
func (a *NodesAPIService) CreateNodesExecute(r ApiCreateNodesRequest) (*CreateNodes200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateNodes200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NodesAPIService.CreateNodes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/nodes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.nodeAddInner == nil {
		return localVarReturnValue, nil, reportError("nodeAddInner is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.nodeAddInner
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["API-Tokens"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteNodeRequest struct {
	ctx context.Context
	ApiService *NodesAPIService
	nodeId string
	mode *string
}

// Deletion mode to use, either move to trash (&#39;move&#39;, default) or erase (&#39;erase&#39;)
func (r ApiDeleteNodeRequest) Mode(mode string) ApiDeleteNodeRequest {
	r.mode = &mode
	return r
}

func (r ApiDeleteNodeRequest) Execute() (*DeleteNode200Response, *http.Response, error) {
	return r.ApiService.DeleteNodeExecute(r)
}

/*
DeleteNode Delete a node

Remove a node from the Rudder server. It won't be managed anymore.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param nodeId Id of the target node
 @return ApiDeleteNodeRequest
*/
func (a *NodesAPIService) DeleteNode(ctx context.Context, nodeId string) ApiDeleteNodeRequest {
	return ApiDeleteNodeRequest{
		ApiService: a,
		ctx: ctx,
		nodeId: nodeId,
	}
}

// Execute executes the request
//  @return DeleteNode200Response
func (a *NodesAPIService) DeleteNodeExecute(r ApiDeleteNodeRequest) (*DeleteNode200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DeleteNode200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NodesAPIService.DeleteNode")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/nodes/{nodeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"nodeId"+"}", url.PathEscape(parameterValueToString(r.nodeId, "nodeId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.mode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mode", r.mode, "")
	} else {
		var defaultValue string = "move"
		r.mode = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["API-Tokens"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetNodesStatusRequest struct {
	ctx context.Context
	ApiService *NodesAPIService
	ids *string
}

// Comma separated list of node Ids
func (r ApiGetNodesStatusRequest) Ids(ids string) ApiGetNodesStatusRequest {
	r.ids = &ids
	return r
}

func (r ApiGetNodesStatusRequest) Execute() (*GetNodesStatus200Response, *http.Response, error) {
	return r.ApiService.GetNodesStatusExecute(r)
}

/*
GetNodesStatus Get nodes acceptation status

Get acceptation status (pending, accepted, deleted, unknown) of a list of nodes

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetNodesStatusRequest
*/
func (a *NodesAPIService) GetNodesStatus(ctx context.Context) ApiGetNodesStatusRequest {
	return ApiGetNodesStatusRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetNodesStatus200Response
func (a *NodesAPIService) GetNodesStatusExecute(r ApiGetNodesStatusRequest) (*GetNodesStatus200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNodesStatus200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NodesAPIService.GetNodesStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/nodes/status"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ids == nil {
		return localVarReturnValue, nil, reportError("ids is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "ids", r.ids, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["API-Tokens"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListAcceptedNodesRequest struct {
	ctx context.Context
	ApiService *NodesAPIService
	include *string
	query *map[string]interface{}
	where *[]map[string]interface{}
	composition *string
	select_ *string
}

// Level of information to include from the node inventory. Some base levels are defined (**minimal**, **default**, **full**). You can add fields you want to a base level by adding them to the list, possible values are keys from json answer. If you don&#39;t provide a base level, they will be added to &#x60;default&#x60; level, so if you only want os details, use &#x60;minimal,os&#x60; as the value for this parameter. * **minimal** includes: &#x60;id&#x60;, &#x60;hostname&#x60; and &#x60;status&#x60; * **default** includes **minimal** plus &#x60;architectureDescription&#x60;, &#x60;description&#x60;, &#x60;ipAddresses&#x60;, &#x60;lastRunDate&#x60;, &#x60;lastInventoryDate&#x60;, &#x60;machine&#x60;, &#x60;os&#x60;, &#x60;managementTechnology&#x60;, &#x60;policyServerId&#x60;, &#x60;properties&#x60; (be careful! Only node own properties, if you also need inherited properties, look at the dedicated &#x60;/nodes/{id}/inheritedProperties&#x60; endpoint), &#x60;policyMode &#x60;, &#x60;ram&#x60; and &#x60;timezone&#x60; * **full** includes: **default** plus &#x60;accounts&#x60;, &#x60;bios&#x60;, &#x60;controllers&#x60;, &#x60;environmentVariables&#x60;, &#x60;fileSystems&#x60;, &#x60;managementTechnologyDetails&#x60;, &#x60;memories&#x60;, &#x60;networkInterfaces&#x60;, &#x60;ports&#x60;, &#x60;processes&#x60;, &#x60;processors&#x60;, &#x60;slots&#x60;, &#x60;software&#x60;, &#x60;softwareUpdate&#x60;, &#x60;sound&#x60;, &#x60;storage&#x60;, &#x60;videos&#x60; and &#x60;virtualMachines&#x60;
func (r ApiListAcceptedNodesRequest) Include(include string) ApiListAcceptedNodesRequest {
	r.include = &include
	return r
}

// The criterion you want to find for your nodes. Replaces the &#x60;where&#x60;, &#x60;composition&#x60; and &#x60;select&#x60; parameters in a single parameter.
func (r ApiListAcceptedNodesRequest) Query(query map[string]interface{}) ApiListAcceptedNodesRequest {
	r.query = &query
	return r
}

// The criterion you want to find for your nodes
func (r ApiListAcceptedNodesRequest) Where(where []map[string]interface{}) ApiListAcceptedNodesRequest {
	r.where = &where
	return r
}

// Boolean operator to use between each  &#x60;where&#x60; criteria.
func (r ApiListAcceptedNodesRequest) Composition(composition string) ApiListAcceptedNodesRequest {
	r.composition = &composition
	return r
}

// What kind of data we want to include. Here we can get policy servers/relay by setting &#x60;nodeAndPolicyServer&#x60;. Only used if &#x60;where&#x60; is defined.
func (r ApiListAcceptedNodesRequest) Select_(select_ string) ApiListAcceptedNodesRequest {
	r.select_ = &select_
	return r
}

func (r ApiListAcceptedNodesRequest) Execute() (*ListAcceptedNodes200Response, *http.Response, error) {
	return r.ApiService.ListAcceptedNodesExecute(r)
}

/*
ListAcceptedNodes List managed nodes

Get information about the nodes managed by the target server

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListAcceptedNodesRequest
*/
func (a *NodesAPIService) ListAcceptedNodes(ctx context.Context) ApiListAcceptedNodesRequest {
	return ApiListAcceptedNodesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListAcceptedNodes200Response
func (a *NodesAPIService) ListAcceptedNodesExecute(r ApiListAcceptedNodesRequest) (*ListAcceptedNodes200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListAcceptedNodes200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NodesAPIService.ListAcceptedNodes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/nodes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "")
	} else {
		var defaultValue string = "default"
		r.include = &defaultValue
	}
	if r.query != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "query", r.query, "")
	}
	if r.where != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "where", r.where, "csv")
	}
	if r.composition != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "composition", r.composition, "")
	} else {
		var defaultValue string = "and"
		r.composition = &defaultValue
	}
	if r.select_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "select", r.select_, "")
	} else {
		var defaultValue string = "node"
		r.select_ = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["API-Tokens"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListPendingNodesRequest struct {
	ctx context.Context
	ApiService *NodesAPIService
	include *string
	query *map[string]interface{}
	where *[]map[string]interface{}
	composition *string
	select_ *string
}

// Level of information to include from the node inventory. Some base levels are defined (**minimal**, **default**, **full**). You can add fields you want to a base level by adding them to the list, possible values are keys from json answer. If you don&#39;t provide a base level, they will be added to &#x60;default&#x60; level, so if you only want os details, use &#x60;minimal,os&#x60; as the value for this parameter. * **minimal** includes: &#x60;id&#x60;, &#x60;hostname&#x60; and &#x60;status&#x60; * **default** includes **minimal** plus &#x60;architectureDescription&#x60;, &#x60;description&#x60;, &#x60;ipAddresses&#x60;, &#x60;lastRunDate&#x60;, &#x60;lastInventoryDate&#x60;, &#x60;machine&#x60;, &#x60;os&#x60;, &#x60;managementTechnology&#x60;, &#x60;policyServerId&#x60;, &#x60;properties&#x60; (be careful! Only node own properties, if you also need inherited properties, look at the dedicated &#x60;/nodes/{id}/inheritedProperties&#x60; endpoint), &#x60;policyMode &#x60;, &#x60;ram&#x60; and &#x60;timezone&#x60; * **full** includes: **default** plus &#x60;accounts&#x60;, &#x60;bios&#x60;, &#x60;controllers&#x60;, &#x60;environmentVariables&#x60;, &#x60;fileSystems&#x60;, &#x60;managementTechnologyDetails&#x60;, &#x60;memories&#x60;, &#x60;networkInterfaces&#x60;, &#x60;ports&#x60;, &#x60;processes&#x60;, &#x60;processors&#x60;, &#x60;slots&#x60;, &#x60;software&#x60;, &#x60;softwareUpdate&#x60;, &#x60;sound&#x60;, &#x60;storage&#x60;, &#x60;videos&#x60; and &#x60;virtualMachines&#x60;
func (r ApiListPendingNodesRequest) Include(include string) ApiListPendingNodesRequest {
	r.include = &include
	return r
}

// The criterion you want to find for your nodes. Replaces the &#x60;where&#x60;, &#x60;composition&#x60; and &#x60;select&#x60; parameters in a single parameter.
func (r ApiListPendingNodesRequest) Query(query map[string]interface{}) ApiListPendingNodesRequest {
	r.query = &query
	return r
}

// The criterion you want to find for your nodes
func (r ApiListPendingNodesRequest) Where(where []map[string]interface{}) ApiListPendingNodesRequest {
	r.where = &where
	return r
}

// Boolean operator to use between each  &#x60;where&#x60; criteria.
func (r ApiListPendingNodesRequest) Composition(composition string) ApiListPendingNodesRequest {
	r.composition = &composition
	return r
}

// What kind of data we want to include. Here we can get policy servers/relay by setting &#x60;nodeAndPolicyServer&#x60;. Only used if &#x60;where&#x60; is defined.
func (r ApiListPendingNodesRequest) Select_(select_ string) ApiListPendingNodesRequest {
	r.select_ = &select_
	return r
}

func (r ApiListPendingNodesRequest) Execute() (*ListPendingNodes200Response, *http.Response, error) {
	return r.ApiService.ListPendingNodesExecute(r)
}

/*
ListPendingNodes List pending nodes

Get information about the nodes pending acceptation

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListPendingNodesRequest
*/
func (a *NodesAPIService) ListPendingNodes(ctx context.Context) ApiListPendingNodesRequest {
	return ApiListPendingNodesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListPendingNodes200Response
func (a *NodesAPIService) ListPendingNodesExecute(r ApiListPendingNodesRequest) (*ListPendingNodes200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListPendingNodes200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NodesAPIService.ListPendingNodes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/nodes/pending"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "")
	} else {
		var defaultValue string = "default"
		r.include = &defaultValue
	}
	if r.query != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "query", r.query, "")
	}
	if r.where != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "where", r.where, "csv")
	}
	if r.composition != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "composition", r.composition, "")
	} else {
		var defaultValue string = "and"
		r.composition = &defaultValue
	}
	if r.select_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "select", r.select_, "")
	} else {
		var defaultValue string = "node"
		r.select_ = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["API-Tokens"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiNodeDetailsRequest struct {
	ctx context.Context
	ApiService *NodesAPIService
	nodeId string
	include *string
}

// Level of information to include from the node inventory. Some base levels are defined (**minimal**, **default**, **full**). You can add fields you want to a base level by adding them to the list, possible values are keys from json answer. If you don&#39;t provide a base level, they will be added to &#x60;default&#x60; level, so if you only want os details, use &#x60;minimal,os&#x60; as the value for this parameter. * **minimal** includes: &#x60;id&#x60;, &#x60;hostname&#x60; and &#x60;status&#x60; * **default** includes **minimal** plus &#x60;architectureDescription&#x60;, &#x60;description&#x60;, &#x60;ipAddresses&#x60;, &#x60;lastRunDate&#x60;, &#x60;lastInventoryDate&#x60;, &#x60;machine&#x60;, &#x60;os&#x60;, &#x60;managementTechnology&#x60;, &#x60;policyServerId&#x60;, &#x60;properties&#x60; (be careful! Only node own properties, if you also need inherited properties, look at the dedicated &#x60;/nodes/{id}/inheritedProperties&#x60; endpoint), &#x60;policyMode &#x60;, &#x60;ram&#x60; and &#x60;timezone&#x60; * **full** includes: **default** plus &#x60;accounts&#x60;, &#x60;bios&#x60;, &#x60;controllers&#x60;, &#x60;environmentVariables&#x60;, &#x60;fileSystems&#x60;, &#x60;managementTechnologyDetails&#x60;, &#x60;memories&#x60;, &#x60;networkInterfaces&#x60;, &#x60;ports&#x60;, &#x60;processes&#x60;, &#x60;processors&#x60;, &#x60;slots&#x60;, &#x60;software&#x60;, &#x60;softwareUpdate&#x60;, &#x60;sound&#x60;, &#x60;storage&#x60;, &#x60;videos&#x60; and &#x60;virtualMachines&#x60;
func (r ApiNodeDetailsRequest) Include(include string) ApiNodeDetailsRequest {
	r.include = &include
	return r
}

func (r ApiNodeDetailsRequest) Execute() (*NodeDetails200Response, *http.Response, error) {
	return r.ApiService.NodeDetailsExecute(r)
}

/*
NodeDetails Get information about a node

Get details about a given node

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param nodeId Id of the target node
 @return ApiNodeDetailsRequest
*/
func (a *NodesAPIService) NodeDetails(ctx context.Context, nodeId string) ApiNodeDetailsRequest {
	return ApiNodeDetailsRequest{
		ApiService: a,
		ctx: ctx,
		nodeId: nodeId,
	}
}

// Execute executes the request
//  @return NodeDetails200Response
func (a *NodesAPIService) NodeDetailsExecute(r ApiNodeDetailsRequest) (*NodeDetails200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NodeDetails200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NodesAPIService.NodeDetails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/nodes/{nodeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"nodeId"+"}", url.PathEscape(parameterValueToString(r.nodeId, "nodeId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "")
	} else {
		var defaultValue string = "default"
		r.include = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["API-Tokens"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiNodeInheritedPropertiesRequest struct {
	ctx context.Context
	ApiService *NodesAPIService
	nodeId string
}

func (r ApiNodeInheritedPropertiesRequest) Execute() (*NodeInheritedProperties200Response, *http.Response, error) {
	return r.ApiService.NodeInheritedPropertiesExecute(r)
}

/*
NodeInheritedProperties Get inherited node properties for a node

This API returns all node properties for a node, including group inherited ones.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param nodeId Id of the target node
 @return ApiNodeInheritedPropertiesRequest
*/
func (a *NodesAPIService) NodeInheritedProperties(ctx context.Context, nodeId string) ApiNodeInheritedPropertiesRequest {
	return ApiNodeInheritedPropertiesRequest{
		ApiService: a,
		ctx: ctx,
		nodeId: nodeId,
	}
}

// Execute executes the request
//  @return NodeInheritedProperties200Response
func (a *NodesAPIService) NodeInheritedPropertiesExecute(r ApiNodeInheritedPropertiesRequest) (*NodeInheritedProperties200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NodeInheritedProperties200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NodesAPIService.NodeInheritedProperties")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/nodes/{nodeId}/inheritedProperties"
	localVarPath = strings.Replace(localVarPath, "{"+"nodeId"+"}", url.PathEscape(parameterValueToString(r.nodeId, "nodeId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["API-Tokens"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateNodeRequest struct {
	ctx context.Context
	ApiService *NodesAPIService
	nodeId string
	nodeSettings *NodeSettings
}

func (r ApiUpdateNodeRequest) NodeSettings(nodeSettings NodeSettings) ApiUpdateNodeRequest {
	r.nodeSettings = &nodeSettings
	return r
}

func (r ApiUpdateNodeRequest) Execute() (*UpdateNode200Response, *http.Response, error) {
	return r.ApiService.UpdateNodeExecute(r)
}

/*
UpdateNode Update node settings and properties

Update node settings and properties

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param nodeId Id of the target node
 @return ApiUpdateNodeRequest
*/
func (a *NodesAPIService) UpdateNode(ctx context.Context, nodeId string) ApiUpdateNodeRequest {
	return ApiUpdateNodeRequest{
		ApiService: a,
		ctx: ctx,
		nodeId: nodeId,
	}
}

// Execute executes the request
//  @return UpdateNode200Response
func (a *NodesAPIService) UpdateNodeExecute(r ApiUpdateNodeRequest) (*UpdateNode200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UpdateNode200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NodesAPIService.UpdateNode")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/nodes/{nodeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"nodeId"+"}", url.PathEscape(parameterValueToString(r.nodeId, "nodeId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.nodeSettings
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["API-Tokens"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
