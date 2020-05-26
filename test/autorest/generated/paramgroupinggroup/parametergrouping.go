// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package paramgroupinggroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ParameterGroupingOperations contains the methods for the ParameterGrouping group.
type ParameterGroupingOperations interface {
	// PostMultiParamGroups - Post parameters from multiple different parameter groups
	PostMultiParamGroups(ctx context.Context, firstParameterGroup *FirstParameterGroup, parameterGroupingPostMultiParamGroupsSecondParamGroup *ParameterGroupingPostMultiParamGroupsSecondParamGroup) (*http.Response, error)
	// PostOptional - Post a bunch of optional parameters grouped
	PostOptional(ctx context.Context, parameterGroupingPostOptionalParameters *ParameterGroupingPostOptionalParameters) (*http.Response, error)
	// PostRequired - Post a bunch of required parameters grouped
	PostRequired(ctx context.Context, parameterGroupingPostRequiredParameters ParameterGroupingPostRequiredParameters) (*http.Response, error)
	// PostSharedParameterGroupObject - Post parameters with a shared parameter group object
	PostSharedParameterGroupObject(ctx context.Context, firstParameterGroup *FirstParameterGroup) (*http.Response, error)
}

// parameterGroupingOperations implements the ParameterGroupingOperations interface.
type parameterGroupingOperations struct {
	*Client
}

// PostMultiParamGroups - Post parameters from multiple different parameter groups
func (client *parameterGroupingOperations) PostMultiParamGroups(ctx context.Context, firstParameterGroup *FirstParameterGroup, parameterGroupingPostMultiParamGroupsSecondParamGroup *ParameterGroupingPostMultiParamGroupsSecondParamGroup) (*http.Response, error) {
	req, err := client.postMultiParamGroupsCreateRequest(firstParameterGroup, parameterGroupingPostMultiParamGroupsSecondParamGroup)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.postMultiParamGroupsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// postMultiParamGroupsCreateRequest creates the PostMultiParamGroups request.
func (client *parameterGroupingOperations) postMultiParamGroupsCreateRequest(firstParameterGroup *FirstParameterGroup, parameterGroupingPostMultiParamGroupsSecondParamGroup *ParameterGroupingPostMultiParamGroupsSecondParamGroup) (*azcore.Request, error) {
	urlPath := "/parameterGrouping/postMultipleParameterGroups"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	if firstParameterGroup != nil && firstParameterGroup.QueryOne != nil {
		query.Set("query-one", strconv.FormatInt(int64(*firstParameterGroup.QueryOne), 10))
	}
	if parameterGroupingPostMultiParamGroupsSecondParamGroup != nil && parameterGroupingPostMultiParamGroupsSecondParamGroup.QueryTwo != nil {
		query.Set("query-two", strconv.FormatInt(int64(*parameterGroupingPostMultiParamGroupsSecondParamGroup.QueryTwo), 10))
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPost, *u)
	if firstParameterGroup != nil && firstParameterGroup.HeaderOne != nil {
		req.Header.Set("header-one", *firstParameterGroup.HeaderOne)
	}
	if parameterGroupingPostMultiParamGroupsSecondParamGroup != nil && parameterGroupingPostMultiParamGroupsSecondParamGroup.HeaderTwo != nil {
		req.Header.Set("header-two", *parameterGroupingPostMultiParamGroupsSecondParamGroup.HeaderTwo)
	}
	return req, nil
}

// postMultiParamGroupsHandleResponse handles the PostMultiParamGroups response.
func (client *parameterGroupingOperations) postMultiParamGroupsHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postMultiParamGroupsHandleError(resp)
	}
	return resp.Response, nil
}

// postMultiParamGroupsHandleError handles the PostMultiParamGroups error response.
func (client *parameterGroupingOperations) postMultiParamGroupsHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PostOptional - Post a bunch of optional parameters grouped
func (client *parameterGroupingOperations) PostOptional(ctx context.Context, parameterGroupingPostOptionalParameters *ParameterGroupingPostOptionalParameters) (*http.Response, error) {
	req, err := client.postOptionalCreateRequest(parameterGroupingPostOptionalParameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.postOptionalHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// postOptionalCreateRequest creates the PostOptional request.
func (client *parameterGroupingOperations) postOptionalCreateRequest(parameterGroupingPostOptionalParameters *ParameterGroupingPostOptionalParameters) (*azcore.Request, error) {
	urlPath := "/parameterGrouping/postOptional"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	if parameterGroupingPostOptionalParameters != nil && parameterGroupingPostOptionalParameters.Query != nil {
		query.Set("query", strconv.FormatInt(int64(*parameterGroupingPostOptionalParameters.Query), 10))
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPost, *u)
	if parameterGroupingPostOptionalParameters != nil && parameterGroupingPostOptionalParameters.CustomHeader != nil {
		req.Header.Set("customHeader", *parameterGroupingPostOptionalParameters.CustomHeader)
	}
	return req, nil
}

// postOptionalHandleResponse handles the PostOptional response.
func (client *parameterGroupingOperations) postOptionalHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postOptionalHandleError(resp)
	}
	return resp.Response, nil
}

// postOptionalHandleError handles the PostOptional error response.
func (client *parameterGroupingOperations) postOptionalHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PostRequired - Post a bunch of required parameters grouped
func (client *parameterGroupingOperations) PostRequired(ctx context.Context, parameterGroupingPostRequiredParameters ParameterGroupingPostRequiredParameters) (*http.Response, error) {
	req, err := client.postRequiredCreateRequest(parameterGroupingPostRequiredParameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.postRequiredHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// postRequiredCreateRequest creates the PostRequired request.
func (client *parameterGroupingOperations) postRequiredCreateRequest(parameterGroupingPostRequiredParameters ParameterGroupingPostRequiredParameters) (*azcore.Request, error) {
	urlPath := "/parameterGrouping/postRequired/{path}"
	urlPath = strings.ReplaceAll(urlPath, "{path}", url.PathEscape(parameterGroupingPostRequiredParameters.PathParameter))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	if parameterGroupingPostRequiredParameters.Query != nil {
		query.Set("query", strconv.FormatInt(int64(*parameterGroupingPostRequiredParameters.Query), 10))
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPost, *u)
	if parameterGroupingPostRequiredParameters.CustomHeader != nil {
		req.Header.Set("customHeader", *parameterGroupingPostRequiredParameters.CustomHeader)
	}
	return req, req.MarshalAsJSON(parameterGroupingPostRequiredParameters.Body)
}

// postRequiredHandleResponse handles the PostRequired response.
func (client *parameterGroupingOperations) postRequiredHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postRequiredHandleError(resp)
	}
	return resp.Response, nil
}

// postRequiredHandleError handles the PostRequired error response.
func (client *parameterGroupingOperations) postRequiredHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PostSharedParameterGroupObject - Post parameters with a shared parameter group object
func (client *parameterGroupingOperations) PostSharedParameterGroupObject(ctx context.Context, firstParameterGroup *FirstParameterGroup) (*http.Response, error) {
	req, err := client.postSharedParameterGroupObjectCreateRequest(firstParameterGroup)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.postSharedParameterGroupObjectHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// postSharedParameterGroupObjectCreateRequest creates the PostSharedParameterGroupObject request.
func (client *parameterGroupingOperations) postSharedParameterGroupObjectCreateRequest(firstParameterGroup *FirstParameterGroup) (*azcore.Request, error) {
	urlPath := "/parameterGrouping/sharedParameterGroupObject"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	if firstParameterGroup != nil && firstParameterGroup.QueryOne != nil {
		query.Set("query-one", strconv.FormatInt(int64(*firstParameterGroup.QueryOne), 10))
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPost, *u)
	if firstParameterGroup != nil && firstParameterGroup.HeaderOne != nil {
		req.Header.Set("header-one", *firstParameterGroup.HeaderOne)
	}
	return req, nil
}

// postSharedParameterGroupObjectHandleResponse handles the PostSharedParameterGroupObject response.
func (client *parameterGroupingOperations) postSharedParameterGroupObjectHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postSharedParameterGroupObjectHandleError(resp)
	}
	return resp.Response, nil
}

// postSharedParameterGroupObjectHandleError handles the PostSharedParameterGroupObject error response.
func (client *parameterGroupingOperations) postSharedParameterGroupObjectHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
