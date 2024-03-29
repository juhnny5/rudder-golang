/*
Rudder API

Testing ChangeRequestsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package rudder-golang

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/juhnny5/rudder-golang"
)

func Test_rudder-golang_ChangeRequestsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ChangeRequestsAPIService AcceptChangeRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var changeRequestId int32

		resp, httpRes, err := apiClient.ChangeRequestsAPI.AcceptChangeRequest(context.Background(), changeRequestId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChangeRequestsAPIService ChangeRequestDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var changeRequestId int32

		resp, httpRes, err := apiClient.ChangeRequestsAPI.ChangeRequestDetails(context.Background(), changeRequestId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChangeRequestsAPIService DeclineChangeRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var changeRequestId int32

		resp, httpRes, err := apiClient.ChangeRequestsAPI.DeclineChangeRequest(context.Background(), changeRequestId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChangeRequestsAPIService ListChangeRequests", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChangeRequestsAPI.ListChangeRequests(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChangeRequestsAPIService ListUsers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChangeRequestsAPI.ListUsers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChangeRequestsAPIService RemoveValidatedUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var username string

		resp, httpRes, err := apiClient.ChangeRequestsAPI.RemoveValidatedUser(context.Background(), username).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChangeRequestsAPIService SaveWorkflowUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChangeRequestsAPI.SaveWorkflowUser(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChangeRequestsAPIService UpdateChangeRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var changeRequestId int32

		resp, httpRes, err := apiClient.ChangeRequestsAPI.UpdateChangeRequest(context.Background(), changeRequestId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
