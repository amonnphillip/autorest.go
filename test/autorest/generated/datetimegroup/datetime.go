// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package datetimegroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"time"
)

// DatetimeOperations contains the methods for the Datetime group.
type DatetimeOperations interface {
	// GetInvalid - Get invalid datetime value
	GetInvalid(ctx context.Context) (*TimeResponse, error)
	// GetLocalNegativeOffsetLowercaseMaxDateTime - Get max datetime value with positive num offset 9999-12-31t23:59:59.999-14:00
	GetLocalNegativeOffsetLowercaseMaxDateTime(ctx context.Context) (*TimeResponse, error)
	// GetLocalNegativeOffsetMinDateTime - Get min datetime value 0001-01-01T00:00:00-14:00
	GetLocalNegativeOffsetMinDateTime(ctx context.Context) (*TimeResponse, error)
	// GetLocalNegativeOffsetUppercaseMaxDateTime - Get max datetime value with positive num offset 9999-12-31T23:59:59.999-14:00
	GetLocalNegativeOffsetUppercaseMaxDateTime(ctx context.Context) (*TimeResponse, error)
	// GetLocalPositiveOffsetLowercaseMaxDateTime - Get max datetime value with positive num offset 9999-12-31t23:59:59.999+14:00
	GetLocalPositiveOffsetLowercaseMaxDateTime(ctx context.Context) (*TimeResponse, error)
	// GetLocalPositiveOffsetMinDateTime - Get min datetime value 0001-01-01T00:00:00+14:00
	GetLocalPositiveOffsetMinDateTime(ctx context.Context) (*TimeResponse, error)
	// GetLocalPositiveOffsetUppercaseMaxDateTime - Get max datetime value with positive num offset 9999-12-31T23:59:59.999+14:00
	GetLocalPositiveOffsetUppercaseMaxDateTime(ctx context.Context) (*TimeResponse, error)
	// GetNull - Get null datetime value
	GetNull(ctx context.Context) (*TimeResponse, error)
	// GetOverflow - Get overflow datetime value
	GetOverflow(ctx context.Context) (*TimeResponse, error)
	// GetUTCLowercaseMaxDateTime - Get max datetime value 9999-12-31t23:59:59.999z
	GetUTCLowercaseMaxDateTime(ctx context.Context) (*TimeResponse, error)
	// GetUTCMinDateTime - Get min datetime value 0001-01-01T00:00:00Z
	GetUTCMinDateTime(ctx context.Context) (*TimeResponse, error)
	// GetUTCUppercaseMaxDateTime - Get max datetime value 9999-12-31T23:59:59.999Z
	GetUTCUppercaseMaxDateTime(ctx context.Context) (*TimeResponse, error)
	// GetUTCUppercaseMaxDateTime7Digits - Get max datetime value 9999-12-31T23:59:59.9999999Z
	GetUTCUppercaseMaxDateTime7Digits(ctx context.Context) (*TimeResponse, error)
	// GetUnderflow - Get underflow datetime value
	GetUnderflow(ctx context.Context) (*TimeResponse, error)
	// PutLocalNegativeOffsetMaxDateTime - Put max datetime value with positive numoffset 9999-12-31t23:59:59.999-14:00
	PutLocalNegativeOffsetMaxDateTime(ctx context.Context, datetimeBody time.Time) (*http.Response, error)
	// PutLocalNegativeOffsetMinDateTime - Put min datetime value 0001-01-01T00:00:00-14:00
	PutLocalNegativeOffsetMinDateTime(ctx context.Context, datetimeBody time.Time) (*http.Response, error)
	// PutLocalPositiveOffsetMaxDateTime - Put max datetime value with positive numoffset 9999-12-31t23:59:59.999+14:00
	PutLocalPositiveOffsetMaxDateTime(ctx context.Context, datetimeBody time.Time) (*http.Response, error)
	// PutLocalPositiveOffsetMinDateTime - Put min datetime value 0001-01-01T00:00:00+14:00
	PutLocalPositiveOffsetMinDateTime(ctx context.Context, datetimeBody time.Time) (*http.Response, error)
	// PutUTCMaxDateTime - Put max datetime value 9999-12-31T23:59:59.999Z
	PutUTCMaxDateTime(ctx context.Context, datetimeBody time.Time) (*http.Response, error)
	// PutUTCMaxDateTime7Digits - Put max datetime value 9999-12-31T23:59:59.9999999Z
	PutUTCMaxDateTime7Digits(ctx context.Context, datetimeBody time.Time) (*http.Response, error)
	// PutUTCMinDateTime - Put min datetime value 0001-01-01T00:00:00Z
	PutUTCMinDateTime(ctx context.Context, datetimeBody time.Time) (*http.Response, error)
}

// datetimeOperations implements the DatetimeOperations interface.
type datetimeOperations struct {
	*Client
}

// GetInvalid - Get invalid datetime value
func (client *datetimeOperations) GetInvalid(ctx context.Context) (*TimeResponse, error) {
	req, err := client.getInvalidCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getInvalidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getInvalidCreateRequest creates the GetInvalid request.
func (client *datetimeOperations) getInvalidCreateRequest() (*azcore.Request, error) {
	urlPath := "/datetime/invalid"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getInvalidHandleResponse handles the GetInvalid response.
func (client *datetimeOperations) getInvalidHandleResponse(resp *azcore.Response) (*TimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getInvalidHandleError(resp)
	}
	var aux *timeRFC3339
	err := resp.UnmarshalAsJSON(&aux)
	return &TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, err
}

// getInvalidHandleError handles the GetInvalid error response.
func (client *datetimeOperations) getInvalidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetLocalNegativeOffsetLowercaseMaxDateTime - Get max datetime value with positive num offset 9999-12-31t23:59:59.999-14:00
func (client *datetimeOperations) GetLocalNegativeOffsetLowercaseMaxDateTime(ctx context.Context) (*TimeResponse, error) {
	req, err := client.getLocalNegativeOffsetLowercaseMaxDateTimeCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getLocalNegativeOffsetLowercaseMaxDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getLocalNegativeOffsetLowercaseMaxDateTimeCreateRequest creates the GetLocalNegativeOffsetLowercaseMaxDateTime request.
func (client *datetimeOperations) getLocalNegativeOffsetLowercaseMaxDateTimeCreateRequest() (*azcore.Request, error) {
	urlPath := "/datetime/max/localnegativeoffset/lowercase"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getLocalNegativeOffsetLowercaseMaxDateTimeHandleResponse handles the GetLocalNegativeOffsetLowercaseMaxDateTime response.
func (client *datetimeOperations) getLocalNegativeOffsetLowercaseMaxDateTimeHandleResponse(resp *azcore.Response) (*TimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getLocalNegativeOffsetLowercaseMaxDateTimeHandleError(resp)
	}
	var aux *timeRFC3339
	err := resp.UnmarshalAsJSON(&aux)
	return &TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, err
}

// getLocalNegativeOffsetLowercaseMaxDateTimeHandleError handles the GetLocalNegativeOffsetLowercaseMaxDateTime error response.
func (client *datetimeOperations) getLocalNegativeOffsetLowercaseMaxDateTimeHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetLocalNegativeOffsetMinDateTime - Get min datetime value 0001-01-01T00:00:00-14:00
func (client *datetimeOperations) GetLocalNegativeOffsetMinDateTime(ctx context.Context) (*TimeResponse, error) {
	req, err := client.getLocalNegativeOffsetMinDateTimeCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getLocalNegativeOffsetMinDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getLocalNegativeOffsetMinDateTimeCreateRequest creates the GetLocalNegativeOffsetMinDateTime request.
func (client *datetimeOperations) getLocalNegativeOffsetMinDateTimeCreateRequest() (*azcore.Request, error) {
	urlPath := "/datetime/min/localnegativeoffset"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getLocalNegativeOffsetMinDateTimeHandleResponse handles the GetLocalNegativeOffsetMinDateTime response.
func (client *datetimeOperations) getLocalNegativeOffsetMinDateTimeHandleResponse(resp *azcore.Response) (*TimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getLocalNegativeOffsetMinDateTimeHandleError(resp)
	}
	var aux *timeRFC3339
	err := resp.UnmarshalAsJSON(&aux)
	return &TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, err
}

// getLocalNegativeOffsetMinDateTimeHandleError handles the GetLocalNegativeOffsetMinDateTime error response.
func (client *datetimeOperations) getLocalNegativeOffsetMinDateTimeHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetLocalNegativeOffsetUppercaseMaxDateTime - Get max datetime value with positive num offset 9999-12-31T23:59:59.999-14:00
func (client *datetimeOperations) GetLocalNegativeOffsetUppercaseMaxDateTime(ctx context.Context) (*TimeResponse, error) {
	req, err := client.getLocalNegativeOffsetUppercaseMaxDateTimeCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getLocalNegativeOffsetUppercaseMaxDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getLocalNegativeOffsetUppercaseMaxDateTimeCreateRequest creates the GetLocalNegativeOffsetUppercaseMaxDateTime request.
func (client *datetimeOperations) getLocalNegativeOffsetUppercaseMaxDateTimeCreateRequest() (*azcore.Request, error) {
	urlPath := "/datetime/max/localnegativeoffset/uppercase"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getLocalNegativeOffsetUppercaseMaxDateTimeHandleResponse handles the GetLocalNegativeOffsetUppercaseMaxDateTime response.
func (client *datetimeOperations) getLocalNegativeOffsetUppercaseMaxDateTimeHandleResponse(resp *azcore.Response) (*TimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getLocalNegativeOffsetUppercaseMaxDateTimeHandleError(resp)
	}
	var aux *timeRFC3339
	err := resp.UnmarshalAsJSON(&aux)
	return &TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, err
}

// getLocalNegativeOffsetUppercaseMaxDateTimeHandleError handles the GetLocalNegativeOffsetUppercaseMaxDateTime error response.
func (client *datetimeOperations) getLocalNegativeOffsetUppercaseMaxDateTimeHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetLocalPositiveOffsetLowercaseMaxDateTime - Get max datetime value with positive num offset 9999-12-31t23:59:59.999+14:00
func (client *datetimeOperations) GetLocalPositiveOffsetLowercaseMaxDateTime(ctx context.Context) (*TimeResponse, error) {
	req, err := client.getLocalPositiveOffsetLowercaseMaxDateTimeCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getLocalPositiveOffsetLowercaseMaxDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getLocalPositiveOffsetLowercaseMaxDateTimeCreateRequest creates the GetLocalPositiveOffsetLowercaseMaxDateTime request.
func (client *datetimeOperations) getLocalPositiveOffsetLowercaseMaxDateTimeCreateRequest() (*azcore.Request, error) {
	urlPath := "/datetime/max/localpositiveoffset/lowercase"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getLocalPositiveOffsetLowercaseMaxDateTimeHandleResponse handles the GetLocalPositiveOffsetLowercaseMaxDateTime response.
func (client *datetimeOperations) getLocalPositiveOffsetLowercaseMaxDateTimeHandleResponse(resp *azcore.Response) (*TimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getLocalPositiveOffsetLowercaseMaxDateTimeHandleError(resp)
	}
	var aux *timeRFC3339
	err := resp.UnmarshalAsJSON(&aux)
	return &TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, err
}

// getLocalPositiveOffsetLowercaseMaxDateTimeHandleError handles the GetLocalPositiveOffsetLowercaseMaxDateTime error response.
func (client *datetimeOperations) getLocalPositiveOffsetLowercaseMaxDateTimeHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetLocalPositiveOffsetMinDateTime - Get min datetime value 0001-01-01T00:00:00+14:00
func (client *datetimeOperations) GetLocalPositiveOffsetMinDateTime(ctx context.Context) (*TimeResponse, error) {
	req, err := client.getLocalPositiveOffsetMinDateTimeCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getLocalPositiveOffsetMinDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getLocalPositiveOffsetMinDateTimeCreateRequest creates the GetLocalPositiveOffsetMinDateTime request.
func (client *datetimeOperations) getLocalPositiveOffsetMinDateTimeCreateRequest() (*azcore.Request, error) {
	urlPath := "/datetime/min/localpositiveoffset"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getLocalPositiveOffsetMinDateTimeHandleResponse handles the GetLocalPositiveOffsetMinDateTime response.
func (client *datetimeOperations) getLocalPositiveOffsetMinDateTimeHandleResponse(resp *azcore.Response) (*TimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getLocalPositiveOffsetMinDateTimeHandleError(resp)
	}
	var aux *timeRFC3339
	err := resp.UnmarshalAsJSON(&aux)
	return &TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, err
}

// getLocalPositiveOffsetMinDateTimeHandleError handles the GetLocalPositiveOffsetMinDateTime error response.
func (client *datetimeOperations) getLocalPositiveOffsetMinDateTimeHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetLocalPositiveOffsetUppercaseMaxDateTime - Get max datetime value with positive num offset 9999-12-31T23:59:59.999+14:00
func (client *datetimeOperations) GetLocalPositiveOffsetUppercaseMaxDateTime(ctx context.Context) (*TimeResponse, error) {
	req, err := client.getLocalPositiveOffsetUppercaseMaxDateTimeCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getLocalPositiveOffsetUppercaseMaxDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getLocalPositiveOffsetUppercaseMaxDateTimeCreateRequest creates the GetLocalPositiveOffsetUppercaseMaxDateTime request.
func (client *datetimeOperations) getLocalPositiveOffsetUppercaseMaxDateTimeCreateRequest() (*azcore.Request, error) {
	urlPath := "/datetime/max/localpositiveoffset/uppercase"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getLocalPositiveOffsetUppercaseMaxDateTimeHandleResponse handles the GetLocalPositiveOffsetUppercaseMaxDateTime response.
func (client *datetimeOperations) getLocalPositiveOffsetUppercaseMaxDateTimeHandleResponse(resp *azcore.Response) (*TimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getLocalPositiveOffsetUppercaseMaxDateTimeHandleError(resp)
	}
	var aux *timeRFC3339
	err := resp.UnmarshalAsJSON(&aux)
	return &TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, err
}

// getLocalPositiveOffsetUppercaseMaxDateTimeHandleError handles the GetLocalPositiveOffsetUppercaseMaxDateTime error response.
func (client *datetimeOperations) getLocalPositiveOffsetUppercaseMaxDateTimeHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetNull - Get null datetime value
func (client *datetimeOperations) GetNull(ctx context.Context) (*TimeResponse, error) {
	req, err := client.getNullCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getNullHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getNullCreateRequest creates the GetNull request.
func (client *datetimeOperations) getNullCreateRequest() (*azcore.Request, error) {
	urlPath := "/datetime/null"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getNullHandleResponse handles the GetNull response.
func (client *datetimeOperations) getNullHandleResponse(resp *azcore.Response) (*TimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getNullHandleError(resp)
	}
	var aux *timeRFC3339
	err := resp.UnmarshalAsJSON(&aux)
	return &TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, err
}

// getNullHandleError handles the GetNull error response.
func (client *datetimeOperations) getNullHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetOverflow - Get overflow datetime value
func (client *datetimeOperations) GetOverflow(ctx context.Context) (*TimeResponse, error) {
	req, err := client.getOverflowCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getOverflowHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getOverflowCreateRequest creates the GetOverflow request.
func (client *datetimeOperations) getOverflowCreateRequest() (*azcore.Request, error) {
	urlPath := "/datetime/overflow"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getOverflowHandleResponse handles the GetOverflow response.
func (client *datetimeOperations) getOverflowHandleResponse(resp *azcore.Response) (*TimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getOverflowHandleError(resp)
	}
	var aux *timeRFC3339
	err := resp.UnmarshalAsJSON(&aux)
	return &TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, err
}

// getOverflowHandleError handles the GetOverflow error response.
func (client *datetimeOperations) getOverflowHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetUTCLowercaseMaxDateTime - Get max datetime value 9999-12-31t23:59:59.999z
func (client *datetimeOperations) GetUTCLowercaseMaxDateTime(ctx context.Context) (*TimeResponse, error) {
	req, err := client.getUtcLowercaseMaxDateTimeCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getUtcLowercaseMaxDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getUtcLowercaseMaxDateTimeCreateRequest creates the GetUTCLowercaseMaxDateTime request.
func (client *datetimeOperations) getUtcLowercaseMaxDateTimeCreateRequest() (*azcore.Request, error) {
	urlPath := "/datetime/max/utc/lowercase"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getUtcLowercaseMaxDateTimeHandleResponse handles the GetUTCLowercaseMaxDateTime response.
func (client *datetimeOperations) getUtcLowercaseMaxDateTimeHandleResponse(resp *azcore.Response) (*TimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getUtcLowercaseMaxDateTimeHandleError(resp)
	}
	var aux *timeRFC3339
	err := resp.UnmarshalAsJSON(&aux)
	return &TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, err
}

// getUtcLowercaseMaxDateTimeHandleError handles the GetUTCLowercaseMaxDateTime error response.
func (client *datetimeOperations) getUtcLowercaseMaxDateTimeHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetUTCMinDateTime - Get min datetime value 0001-01-01T00:00:00Z
func (client *datetimeOperations) GetUTCMinDateTime(ctx context.Context) (*TimeResponse, error) {
	req, err := client.getUtcMinDateTimeCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getUtcMinDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getUtcMinDateTimeCreateRequest creates the GetUTCMinDateTime request.
func (client *datetimeOperations) getUtcMinDateTimeCreateRequest() (*azcore.Request, error) {
	urlPath := "/datetime/min/utc"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getUtcMinDateTimeHandleResponse handles the GetUTCMinDateTime response.
func (client *datetimeOperations) getUtcMinDateTimeHandleResponse(resp *azcore.Response) (*TimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getUtcMinDateTimeHandleError(resp)
	}
	var aux *timeRFC3339
	err := resp.UnmarshalAsJSON(&aux)
	return &TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, err
}

// getUtcMinDateTimeHandleError handles the GetUTCMinDateTime error response.
func (client *datetimeOperations) getUtcMinDateTimeHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetUTCUppercaseMaxDateTime - Get max datetime value 9999-12-31T23:59:59.999Z
func (client *datetimeOperations) GetUTCUppercaseMaxDateTime(ctx context.Context) (*TimeResponse, error) {
	req, err := client.getUtcUppercaseMaxDateTimeCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getUtcUppercaseMaxDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getUtcUppercaseMaxDateTimeCreateRequest creates the GetUTCUppercaseMaxDateTime request.
func (client *datetimeOperations) getUtcUppercaseMaxDateTimeCreateRequest() (*azcore.Request, error) {
	urlPath := "/datetime/max/utc/uppercase"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getUtcUppercaseMaxDateTimeHandleResponse handles the GetUTCUppercaseMaxDateTime response.
func (client *datetimeOperations) getUtcUppercaseMaxDateTimeHandleResponse(resp *azcore.Response) (*TimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getUtcUppercaseMaxDateTimeHandleError(resp)
	}
	var aux *timeRFC3339
	err := resp.UnmarshalAsJSON(&aux)
	return &TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, err
}

// getUtcUppercaseMaxDateTimeHandleError handles the GetUTCUppercaseMaxDateTime error response.
func (client *datetimeOperations) getUtcUppercaseMaxDateTimeHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetUTCUppercaseMaxDateTime7Digits - Get max datetime value 9999-12-31T23:59:59.9999999Z
func (client *datetimeOperations) GetUTCUppercaseMaxDateTime7Digits(ctx context.Context) (*TimeResponse, error) {
	req, err := client.getUtcUppercaseMaxDateTime7DigitsCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getUtcUppercaseMaxDateTime7DigitsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getUtcUppercaseMaxDateTime7DigitsCreateRequest creates the GetUTCUppercaseMaxDateTime7Digits request.
func (client *datetimeOperations) getUtcUppercaseMaxDateTime7DigitsCreateRequest() (*azcore.Request, error) {
	urlPath := "/datetime/max/utc7ms/uppercase"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getUtcUppercaseMaxDateTime7DigitsHandleResponse handles the GetUTCUppercaseMaxDateTime7Digits response.
func (client *datetimeOperations) getUtcUppercaseMaxDateTime7DigitsHandleResponse(resp *azcore.Response) (*TimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getUtcUppercaseMaxDateTime7DigitsHandleError(resp)
	}
	var aux *timeRFC3339
	err := resp.UnmarshalAsJSON(&aux)
	return &TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, err
}

// getUtcUppercaseMaxDateTime7DigitsHandleError handles the GetUTCUppercaseMaxDateTime7Digits error response.
func (client *datetimeOperations) getUtcUppercaseMaxDateTime7DigitsHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetUnderflow - Get underflow datetime value
func (client *datetimeOperations) GetUnderflow(ctx context.Context) (*TimeResponse, error) {
	req, err := client.getUnderflowCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getUnderflowHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getUnderflowCreateRequest creates the GetUnderflow request.
func (client *datetimeOperations) getUnderflowCreateRequest() (*azcore.Request, error) {
	urlPath := "/datetime/underflow"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getUnderflowHandleResponse handles the GetUnderflow response.
func (client *datetimeOperations) getUnderflowHandleResponse(resp *azcore.Response) (*TimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getUnderflowHandleError(resp)
	}
	var aux *timeRFC3339
	err := resp.UnmarshalAsJSON(&aux)
	return &TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, err
}

// getUnderflowHandleError handles the GetUnderflow error response.
func (client *datetimeOperations) getUnderflowHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutLocalNegativeOffsetMaxDateTime - Put max datetime value with positive numoffset 9999-12-31t23:59:59.999-14:00
func (client *datetimeOperations) PutLocalNegativeOffsetMaxDateTime(ctx context.Context, datetimeBody time.Time) (*http.Response, error) {
	req, err := client.putLocalNegativeOffsetMaxDateTimeCreateRequest(datetimeBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putLocalNegativeOffsetMaxDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putLocalNegativeOffsetMaxDateTimeCreateRequest creates the PutLocalNegativeOffsetMaxDateTime request.
func (client *datetimeOperations) putLocalNegativeOffsetMaxDateTimeCreateRequest(datetimeBody time.Time) (*azcore.Request, error) {
	urlPath := "/datetime/max/localnegativeoffset"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(datetimeBody)
}

// putLocalNegativeOffsetMaxDateTimeHandleResponse handles the PutLocalNegativeOffsetMaxDateTime response.
func (client *datetimeOperations) putLocalNegativeOffsetMaxDateTimeHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putLocalNegativeOffsetMaxDateTimeHandleError(resp)
	}
	return resp.Response, nil
}

// putLocalNegativeOffsetMaxDateTimeHandleError handles the PutLocalNegativeOffsetMaxDateTime error response.
func (client *datetimeOperations) putLocalNegativeOffsetMaxDateTimeHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutLocalNegativeOffsetMinDateTime - Put min datetime value 0001-01-01T00:00:00-14:00
func (client *datetimeOperations) PutLocalNegativeOffsetMinDateTime(ctx context.Context, datetimeBody time.Time) (*http.Response, error) {
	req, err := client.putLocalNegativeOffsetMinDateTimeCreateRequest(datetimeBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putLocalNegativeOffsetMinDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putLocalNegativeOffsetMinDateTimeCreateRequest creates the PutLocalNegativeOffsetMinDateTime request.
func (client *datetimeOperations) putLocalNegativeOffsetMinDateTimeCreateRequest(datetimeBody time.Time) (*azcore.Request, error) {
	urlPath := "/datetime/min/localnegativeoffset"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(datetimeBody)
}

// putLocalNegativeOffsetMinDateTimeHandleResponse handles the PutLocalNegativeOffsetMinDateTime response.
func (client *datetimeOperations) putLocalNegativeOffsetMinDateTimeHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putLocalNegativeOffsetMinDateTimeHandleError(resp)
	}
	return resp.Response, nil
}

// putLocalNegativeOffsetMinDateTimeHandleError handles the PutLocalNegativeOffsetMinDateTime error response.
func (client *datetimeOperations) putLocalNegativeOffsetMinDateTimeHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutLocalPositiveOffsetMaxDateTime - Put max datetime value with positive numoffset 9999-12-31t23:59:59.999+14:00
func (client *datetimeOperations) PutLocalPositiveOffsetMaxDateTime(ctx context.Context, datetimeBody time.Time) (*http.Response, error) {
	req, err := client.putLocalPositiveOffsetMaxDateTimeCreateRequest(datetimeBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putLocalPositiveOffsetMaxDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putLocalPositiveOffsetMaxDateTimeCreateRequest creates the PutLocalPositiveOffsetMaxDateTime request.
func (client *datetimeOperations) putLocalPositiveOffsetMaxDateTimeCreateRequest(datetimeBody time.Time) (*azcore.Request, error) {
	urlPath := "/datetime/max/localpositiveoffset"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(datetimeBody)
}

// putLocalPositiveOffsetMaxDateTimeHandleResponse handles the PutLocalPositiveOffsetMaxDateTime response.
func (client *datetimeOperations) putLocalPositiveOffsetMaxDateTimeHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putLocalPositiveOffsetMaxDateTimeHandleError(resp)
	}
	return resp.Response, nil
}

// putLocalPositiveOffsetMaxDateTimeHandleError handles the PutLocalPositiveOffsetMaxDateTime error response.
func (client *datetimeOperations) putLocalPositiveOffsetMaxDateTimeHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutLocalPositiveOffsetMinDateTime - Put min datetime value 0001-01-01T00:00:00+14:00
func (client *datetimeOperations) PutLocalPositiveOffsetMinDateTime(ctx context.Context, datetimeBody time.Time) (*http.Response, error) {
	req, err := client.putLocalPositiveOffsetMinDateTimeCreateRequest(datetimeBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putLocalPositiveOffsetMinDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putLocalPositiveOffsetMinDateTimeCreateRequest creates the PutLocalPositiveOffsetMinDateTime request.
func (client *datetimeOperations) putLocalPositiveOffsetMinDateTimeCreateRequest(datetimeBody time.Time) (*azcore.Request, error) {
	urlPath := "/datetime/min/localpositiveoffset"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(datetimeBody)
}

// putLocalPositiveOffsetMinDateTimeHandleResponse handles the PutLocalPositiveOffsetMinDateTime response.
func (client *datetimeOperations) putLocalPositiveOffsetMinDateTimeHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putLocalPositiveOffsetMinDateTimeHandleError(resp)
	}
	return resp.Response, nil
}

// putLocalPositiveOffsetMinDateTimeHandleError handles the PutLocalPositiveOffsetMinDateTime error response.
func (client *datetimeOperations) putLocalPositiveOffsetMinDateTimeHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutUTCMaxDateTime - Put max datetime value 9999-12-31T23:59:59.999Z
func (client *datetimeOperations) PutUTCMaxDateTime(ctx context.Context, datetimeBody time.Time) (*http.Response, error) {
	req, err := client.putUtcMaxDateTimeCreateRequest(datetimeBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putUtcMaxDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putUtcMaxDateTimeCreateRequest creates the PutUTCMaxDateTime request.
func (client *datetimeOperations) putUtcMaxDateTimeCreateRequest(datetimeBody time.Time) (*azcore.Request, error) {
	urlPath := "/datetime/max/utc"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(datetimeBody)
}

// putUtcMaxDateTimeHandleResponse handles the PutUTCMaxDateTime response.
func (client *datetimeOperations) putUtcMaxDateTimeHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putUtcMaxDateTimeHandleError(resp)
	}
	return resp.Response, nil
}

// putUtcMaxDateTimeHandleError handles the PutUTCMaxDateTime error response.
func (client *datetimeOperations) putUtcMaxDateTimeHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutUTCMaxDateTime7Digits - Put max datetime value 9999-12-31T23:59:59.9999999Z
func (client *datetimeOperations) PutUTCMaxDateTime7Digits(ctx context.Context, datetimeBody time.Time) (*http.Response, error) {
	req, err := client.putUtcMaxDateTime7DigitsCreateRequest(datetimeBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putUtcMaxDateTime7DigitsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putUtcMaxDateTime7DigitsCreateRequest creates the PutUTCMaxDateTime7Digits request.
func (client *datetimeOperations) putUtcMaxDateTime7DigitsCreateRequest(datetimeBody time.Time) (*azcore.Request, error) {
	urlPath := "/datetime/max/utc7ms"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(datetimeBody)
}

// putUtcMaxDateTime7DigitsHandleResponse handles the PutUTCMaxDateTime7Digits response.
func (client *datetimeOperations) putUtcMaxDateTime7DigitsHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putUtcMaxDateTime7DigitsHandleError(resp)
	}
	return resp.Response, nil
}

// putUtcMaxDateTime7DigitsHandleError handles the PutUTCMaxDateTime7Digits error response.
func (client *datetimeOperations) putUtcMaxDateTime7DigitsHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutUTCMinDateTime - Put min datetime value 0001-01-01T00:00:00Z
func (client *datetimeOperations) PutUTCMinDateTime(ctx context.Context, datetimeBody time.Time) (*http.Response, error) {
	req, err := client.putUtcMinDateTimeCreateRequest(datetimeBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putUtcMinDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putUtcMinDateTimeCreateRequest creates the PutUTCMinDateTime request.
func (client *datetimeOperations) putUtcMinDateTimeCreateRequest(datetimeBody time.Time) (*azcore.Request, error) {
	urlPath := "/datetime/min/utc"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(datetimeBody)
}

// putUtcMinDateTimeHandleResponse handles the PutUTCMinDateTime response.
func (client *datetimeOperations) putUtcMinDateTimeHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putUtcMinDateTimeHandleError(resp)
	}
	return resp.Response, nil
}

// putUtcMinDateTimeHandleError handles the PutUTCMinDateTime error response.
func (client *datetimeOperations) putUtcMinDateTimeHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
