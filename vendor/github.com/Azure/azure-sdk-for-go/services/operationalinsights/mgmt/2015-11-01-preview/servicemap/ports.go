package servicemap

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// PortsClient is the service Map API Reference
type PortsClient struct {
	ManagementClient
}

// NewPortsClient creates an instance of the PortsClient client.
func NewPortsClient(subscriptionID string) PortsClient {
	return NewPortsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPortsClientWithBaseURI creates an instance of the PortsClient client.
func NewPortsClientWithBaseURI(baseURI string, subscriptionID string) PortsClient {
	return PortsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get returns the specified port. The port must be live during the specified time interval. If the port is not live
// during the interval, status 404 (Not Found) is returned.
//
// resourceGroupName is resource group name within the specified subscriptionId. workspaceName is OMS workspace
// containing the resources of interest. machineName is machine resource name. portName is port resource name.
// startTime is UTC date and time specifying the start time of an interval. When not specified the service uses
// DateTime.UtcNow - 10m endTime is UTC date and time specifying the end time of an interval. When not specified the
// service uses DateTime.UtcNow
func (client PortsClient) Get(resourceGroupName string, workspaceName string, machineName string, portName string, startTime *date.Time, endTime *date.Time) (result Port, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_-]+`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_][a-zA-Z0-9_-]+[a-zA-Z0-9_]`, Chain: nil}}},
		{TargetValue: machineName,
			Constraints: []validation.Constraint{{Target: "machineName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "machineName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: portName,
			Constraints: []validation.Constraint{{Target: "portName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "portName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "servicemap.PortsClient", "Get")
	}

	req, err := client.GetPreparer(resourceGroupName, workspaceName, machineName, portName, startTime, endTime)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.PortsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicemap.PortsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.PortsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client PortsClient) GetPreparer(resourceGroupName string, workspaceName string, machineName string, portName string, startTime *date.Time, endTime *date.Time) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"machineName":       autorest.Encode("path", machineName),
		"portName":          autorest.Encode("path", portName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2015-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if startTime != nil {
		queryParameters["startTime"] = autorest.Encode("query", *startTime)
	}
	if endTime != nil {
		queryParameters["endTime"] = autorest.Encode("query", *endTime)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machines/{machineName}/ports/{portName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PortsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PortsClient) GetResponder(resp *http.Response) (result Port, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetLiveness obtains the liveness status of the port during the specified time interval.
//
// resourceGroupName is resource group name within the specified subscriptionId. workspaceName is OMS workspace
// containing the resources of interest. machineName is machine resource name. portName is port resource name.
// startTime is UTC date and time specifying the start time of an interval. When not specified the service uses
// DateTime.UtcNow - 10m endTime is UTC date and time specifying the end time of an interval. When not specified the
// service uses DateTime.UtcNow
func (client PortsClient) GetLiveness(resourceGroupName string, workspaceName string, machineName string, portName string, startTime *date.Time, endTime *date.Time) (result Liveness, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_-]+`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_][a-zA-Z0-9_-]+[a-zA-Z0-9_]`, Chain: nil}}},
		{TargetValue: machineName,
			Constraints: []validation.Constraint{{Target: "machineName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "machineName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: portName,
			Constraints: []validation.Constraint{{Target: "portName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "portName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "servicemap.PortsClient", "GetLiveness")
	}

	req, err := client.GetLivenessPreparer(resourceGroupName, workspaceName, machineName, portName, startTime, endTime)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.PortsClient", "GetLiveness", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetLivenessSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicemap.PortsClient", "GetLiveness", resp, "Failure sending request")
		return
	}

	result, err = client.GetLivenessResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.PortsClient", "GetLiveness", resp, "Failure responding to request")
	}

	return
}

// GetLivenessPreparer prepares the GetLiveness request.
func (client PortsClient) GetLivenessPreparer(resourceGroupName string, workspaceName string, machineName string, portName string, startTime *date.Time, endTime *date.Time) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"machineName":       autorest.Encode("path", machineName),
		"portName":          autorest.Encode("path", portName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2015-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if startTime != nil {
		queryParameters["startTime"] = autorest.Encode("query", *startTime)
	}
	if endTime != nil {
		queryParameters["endTime"] = autorest.Encode("query", *endTime)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machines/{machineName}/ports/{portName}/liveness", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetLivenessSender sends the GetLiveness request. The method will close the
// http.Response Body if it receives an error.
func (client PortsClient) GetLivenessSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetLivenessResponder handles the response to the GetLiveness request. The method always
// closes the http.Response Body.
func (client PortsClient) GetLivenessResponder(resp *http.Response) (result Liveness, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListAcceptingProcesses returns a collection of processes accepting on the specified port
//
// resourceGroupName is resource group name within the specified subscriptionId. workspaceName is OMS workspace
// containing the resources of interest. machineName is machine resource name. portName is port resource name.
// startTime is UTC date and time specifying the start time of an interval. When not specified the service uses
// DateTime.UtcNow - 10m endTime is UTC date and time specifying the end time of an interval. When not specified the
// service uses DateTime.UtcNow
func (client PortsClient) ListAcceptingProcesses(resourceGroupName string, workspaceName string, machineName string, portName string, startTime *date.Time, endTime *date.Time) (result ProcessCollection, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_-]+`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_][a-zA-Z0-9_-]+[a-zA-Z0-9_]`, Chain: nil}}},
		{TargetValue: machineName,
			Constraints: []validation.Constraint{{Target: "machineName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "machineName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: portName,
			Constraints: []validation.Constraint{{Target: "portName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "portName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "servicemap.PortsClient", "ListAcceptingProcesses")
	}

	req, err := client.ListAcceptingProcessesPreparer(resourceGroupName, workspaceName, machineName, portName, startTime, endTime)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.PortsClient", "ListAcceptingProcesses", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAcceptingProcessesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicemap.PortsClient", "ListAcceptingProcesses", resp, "Failure sending request")
		return
	}

	result, err = client.ListAcceptingProcessesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.PortsClient", "ListAcceptingProcesses", resp, "Failure responding to request")
	}

	return
}

// ListAcceptingProcessesPreparer prepares the ListAcceptingProcesses request.
func (client PortsClient) ListAcceptingProcessesPreparer(resourceGroupName string, workspaceName string, machineName string, portName string, startTime *date.Time, endTime *date.Time) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"machineName":       autorest.Encode("path", machineName),
		"portName":          autorest.Encode("path", portName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2015-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if startTime != nil {
		queryParameters["startTime"] = autorest.Encode("query", *startTime)
	}
	if endTime != nil {
		queryParameters["endTime"] = autorest.Encode("query", *endTime)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machines/{machineName}/ports/{portName}/acceptingProcesses", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListAcceptingProcessesSender sends the ListAcceptingProcesses request. The method will close the
// http.Response Body if it receives an error.
func (client PortsClient) ListAcceptingProcessesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListAcceptingProcessesResponder handles the response to the ListAcceptingProcesses request. The method always
// closes the http.Response Body.
func (client PortsClient) ListAcceptingProcessesResponder(resp *http.Response) (result ProcessCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListAcceptingProcessesNextResults retrieves the next set of results, if any.
func (client PortsClient) ListAcceptingProcessesNextResults(lastResults ProcessCollection) (result ProcessCollection, err error) {
	req, err := lastResults.ProcessCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "servicemap.PortsClient", "ListAcceptingProcesses", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListAcceptingProcessesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "servicemap.PortsClient", "ListAcceptingProcesses", resp, "Failure sending next results request")
	}

	result, err = client.ListAcceptingProcessesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.PortsClient", "ListAcceptingProcesses", resp, "Failure responding to next results request")
	}

	return
}

// ListAcceptingProcessesComplete gets all elements from the list without paging.
func (client PortsClient) ListAcceptingProcessesComplete(resourceGroupName string, workspaceName string, machineName string, portName string, startTime *date.Time, endTime *date.Time, cancel <-chan struct{}) (<-chan Process, <-chan error) {
	resultChan := make(chan Process)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListAcceptingProcesses(resourceGroupName, workspaceName, machineName, portName, startTime, endTime)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListAcceptingProcessesNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}

// ListConnections returns a collection of connections established via the specified port.
//
// resourceGroupName is resource group name within the specified subscriptionId. workspaceName is OMS workspace
// containing the resources of interest. machineName is machine resource name. portName is port resource name.
// startTime is UTC date and time specifying the start time of an interval. When not specified the service uses
// DateTime.UtcNow - 10m endTime is UTC date and time specifying the end time of an interval. When not specified the
// service uses DateTime.UtcNow
func (client PortsClient) ListConnections(resourceGroupName string, workspaceName string, machineName string, portName string, startTime *date.Time, endTime *date.Time) (result ConnectionCollection, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_-]+`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_][a-zA-Z0-9_-]+[a-zA-Z0-9_]`, Chain: nil}}},
		{TargetValue: machineName,
			Constraints: []validation.Constraint{{Target: "machineName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "machineName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: portName,
			Constraints: []validation.Constraint{{Target: "portName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "portName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "servicemap.PortsClient", "ListConnections")
	}

	req, err := client.ListConnectionsPreparer(resourceGroupName, workspaceName, machineName, portName, startTime, endTime)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.PortsClient", "ListConnections", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListConnectionsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicemap.PortsClient", "ListConnections", resp, "Failure sending request")
		return
	}

	result, err = client.ListConnectionsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.PortsClient", "ListConnections", resp, "Failure responding to request")
	}

	return
}

// ListConnectionsPreparer prepares the ListConnections request.
func (client PortsClient) ListConnectionsPreparer(resourceGroupName string, workspaceName string, machineName string, portName string, startTime *date.Time, endTime *date.Time) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"machineName":       autorest.Encode("path", machineName),
		"portName":          autorest.Encode("path", portName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2015-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if startTime != nil {
		queryParameters["startTime"] = autorest.Encode("query", *startTime)
	}
	if endTime != nil {
		queryParameters["endTime"] = autorest.Encode("query", *endTime)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machines/{machineName}/ports/{portName}/connections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListConnectionsSender sends the ListConnections request. The method will close the
// http.Response Body if it receives an error.
func (client PortsClient) ListConnectionsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListConnectionsResponder handles the response to the ListConnections request. The method always
// closes the http.Response Body.
func (client PortsClient) ListConnectionsResponder(resp *http.Response) (result ConnectionCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListConnectionsNextResults retrieves the next set of results, if any.
func (client PortsClient) ListConnectionsNextResults(lastResults ConnectionCollection) (result ConnectionCollection, err error) {
	req, err := lastResults.ConnectionCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "servicemap.PortsClient", "ListConnections", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListConnectionsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "servicemap.PortsClient", "ListConnections", resp, "Failure sending next results request")
	}

	result, err = client.ListConnectionsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.PortsClient", "ListConnections", resp, "Failure responding to next results request")
	}

	return
}

// ListConnectionsComplete gets all elements from the list without paging.
func (client PortsClient) ListConnectionsComplete(resourceGroupName string, workspaceName string, machineName string, portName string, startTime *date.Time, endTime *date.Time, cancel <-chan struct{}) (<-chan Connection, <-chan error) {
	resultChan := make(chan Connection)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListConnections(resourceGroupName, workspaceName, machineName, portName, startTime, endTime)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListConnectionsNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}
