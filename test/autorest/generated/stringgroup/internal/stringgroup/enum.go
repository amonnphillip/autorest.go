// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package stringgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
)

// EnumOperations contains the methods for the Enum group.
type EnumOperations struct{}

// GetNotExpandableCreateRequest creates the GetNotExpandable request.
func (EnumOperations) GetNotExpandableCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/string/enum/notExpandable"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// GetNotExpandableHandleResponse handles the GetNotExpandable response.
func (EnumOperations) GetNotExpandableHandleResponse(resp *azcore.Response) (*EnumGetNotExpandableResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := EnumGetNotExpandableResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetReferencedCreateRequest creates the GetReferenced request.
func (EnumOperations) GetReferencedCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/string/enum/Referenced"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// GetReferencedHandleResponse handles the GetReferenced response.
func (EnumOperations) GetReferencedHandleResponse(resp *azcore.Response) (*EnumGetReferencedResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := EnumGetReferencedResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetReferencedConstantCreateRequest creates the GetReferencedConstant request.
func (EnumOperations) GetReferencedConstantCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/string/enum/ReferencedConstant"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// GetReferencedConstantHandleResponse handles the GetReferencedConstant response.
func (EnumOperations) GetReferencedConstantHandleResponse(resp *azcore.Response) (*EnumGetReferencedConstantResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := EnumGetReferencedConstantResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.RefColorConstant)
}

// PutNotExpandableCreateRequest creates the PutNotExpandable request.
func (EnumOperations) PutNotExpandableCreateRequest(u url.URL, stringBody Colors) (*azcore.Request, error) {
	urlPath := "/string/enum/notExpandable"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(stringBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutNotExpandableHandleResponse handles the PutNotExpandable response.
func (EnumOperations) PutNotExpandableHandleResponse(resp *azcore.Response) (*EnumPutNotExpandableResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &EnumPutNotExpandableResponse{RawResponse: resp.Response}, nil
}

// PutReferencedCreateRequest creates the PutReferenced request.
func (EnumOperations) PutReferencedCreateRequest(u url.URL, enumStringBody Colors) (*azcore.Request, error) {
	urlPath := "/string/enum/Referenced"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(enumStringBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutReferencedHandleResponse handles the PutReferenced response.
func (EnumOperations) PutReferencedHandleResponse(resp *azcore.Response) (*EnumPutReferencedResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &EnumPutReferencedResponse{RawResponse: resp.Response}, nil
}

// PutReferencedConstantCreateRequest creates the PutReferencedConstant request.
func (EnumOperations) PutReferencedConstantCreateRequest(u url.URL, enumStringBody RefColorConstant) (*azcore.Request, error) {
	urlPath := "/string/enum/ReferencedConstant"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(enumStringBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutReferencedConstantHandleResponse handles the PutReferencedConstant response.
func (EnumOperations) PutReferencedConstantHandleResponse(resp *azcore.Response) (*EnumPutReferencedConstantResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &EnumPutReferencedConstantResponse{RawResponse: resp.Response}, nil
}
