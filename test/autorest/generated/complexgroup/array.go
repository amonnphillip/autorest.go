// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// ArrayOperations contains the methods for the Array group.
type ArrayOperations interface {
	// GetEmpty - Get complex types with array property which is empty
	GetEmpty(ctx context.Context) (*ArrayWrapperResponse, error)
	// GetNotProvided - Get complex types with array property while server doesn't provide a response payload
	GetNotProvided(ctx context.Context) (*ArrayWrapperResponse, error)
	// GetValid - Get complex types with array property
	GetValid(ctx context.Context) (*ArrayWrapperResponse, error)
	// PutEmpty - Put complex types with array property which is empty
	PutEmpty(ctx context.Context, complexBody ArrayWrapper) (*http.Response, error)
	// PutValid - Put complex types with array property
	PutValid(ctx context.Context, complexBody ArrayWrapper) (*http.Response, error)
}

// arrayOperations implements the ArrayOperations interface.
type arrayOperations struct {
	*Client
}

// GetEmpty - Get complex types with array property which is empty
func (client *arrayOperations) GetEmpty(ctx context.Context) (*ArrayWrapperResponse, error) {
	req, err := client.getEmptyCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getEmptyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getEmptyCreateRequest creates the GetEmpty request.
func (client *arrayOperations) getEmptyCreateRequest() (*azcore.Request, error) {
	urlPath := "/complex/array/empty"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getEmptyHandleResponse handles the GetEmpty response.
func (client *arrayOperations) getEmptyHandleResponse(resp *azcore.Response) (*ArrayWrapperResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getEmptyHandleError(resp)
	}
	result := ArrayWrapperResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ArrayWrapper)
}

// getEmptyHandleError handles the GetEmpty error response.
func (client *arrayOperations) getEmptyHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetNotProvided - Get complex types with array property while server doesn't provide a response payload
func (client *arrayOperations) GetNotProvided(ctx context.Context) (*ArrayWrapperResponse, error) {
	req, err := client.getNotProvidedCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getNotProvidedHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getNotProvidedCreateRequest creates the GetNotProvided request.
func (client *arrayOperations) getNotProvidedCreateRequest() (*azcore.Request, error) {
	urlPath := "/complex/array/notprovided"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getNotProvidedHandleResponse handles the GetNotProvided response.
func (client *arrayOperations) getNotProvidedHandleResponse(resp *azcore.Response) (*ArrayWrapperResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getNotProvidedHandleError(resp)
	}
	result := ArrayWrapperResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ArrayWrapper)
}

// getNotProvidedHandleError handles the GetNotProvided error response.
func (client *arrayOperations) getNotProvidedHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetValid - Get complex types with array property
func (client *arrayOperations) GetValid(ctx context.Context) (*ArrayWrapperResponse, error) {
	req, err := client.getValidCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getValidCreateRequest creates the GetValid request.
func (client *arrayOperations) getValidCreateRequest() (*azcore.Request, error) {
	urlPath := "/complex/array/valid"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getValidHandleResponse handles the GetValid response.
func (client *arrayOperations) getValidHandleResponse(resp *azcore.Response) (*ArrayWrapperResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getValidHandleError(resp)
	}
	result := ArrayWrapperResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ArrayWrapper)
}

// getValidHandleError handles the GetValid error response.
func (client *arrayOperations) getValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutEmpty - Put complex types with array property which is empty
func (client *arrayOperations) PutEmpty(ctx context.Context, complexBody ArrayWrapper) (*http.Response, error) {
	req, err := client.putEmptyCreateRequest(complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putEmptyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putEmptyCreateRequest creates the PutEmpty request.
func (client *arrayOperations) putEmptyCreateRequest(complexBody ArrayWrapper) (*azcore.Request, error) {
	urlPath := "/complex/array/empty"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(complexBody)
}

// putEmptyHandleResponse handles the PutEmpty response.
func (client *arrayOperations) putEmptyHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putEmptyHandleError(resp)
	}
	return resp.Response, nil
}

// putEmptyHandleError handles the PutEmpty error response.
func (client *arrayOperations) putEmptyHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutValid - Put complex types with array property
func (client *arrayOperations) PutValid(ctx context.Context, complexBody ArrayWrapper) (*http.Response, error) {
	req, err := client.putValidCreateRequest(complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putValidCreateRequest creates the PutValid request.
func (client *arrayOperations) putValidCreateRequest(complexBody ArrayWrapper) (*azcore.Request, error) {
	urlPath := "/complex/array/valid"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(complexBody)
}

// putValidHandleResponse handles the PutValid response.
func (client *arrayOperations) putValidHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putValidHandleError(resp)
	}
	return resp.Response, nil
}

// putValidHandleError handles the PutValid error response.
func (client *arrayOperations) putValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
