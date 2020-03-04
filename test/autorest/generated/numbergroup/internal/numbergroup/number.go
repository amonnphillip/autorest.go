// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package numbergroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
)

// NumberOperations contains the methods for the Number group.
type NumberOperations struct{}

// GetBigDecimalCreateRequest creates the GetBigDecimal request.
func (NumberOperations) GetBigDecimalCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/number/big/decimal/2.5976931e+101"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// GetBigDecimalHandleResponse handles the GetBigDecimal response.
func (NumberOperations) GetBigDecimalHandleResponse(resp *azcore.Response) (*NumberGetBigDecimalResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := NumberGetBigDecimalResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetBigDecimalNegativeDecimalCreateRequest creates the GetBigDecimalNegativeDecimal request.
func (NumberOperations) GetBigDecimalNegativeDecimalCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/number/big/decimal/-99999999.99"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// GetBigDecimalNegativeDecimalHandleResponse handles the GetBigDecimalNegativeDecimal response.
func (NumberOperations) GetBigDecimalNegativeDecimalHandleResponse(resp *azcore.Response) (*NumberGetBigDecimalNegativeDecimalResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := NumberGetBigDecimalNegativeDecimalResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetBigDecimalPositiveDecimalCreateRequest creates the GetBigDecimalPositiveDecimal request.
func (NumberOperations) GetBigDecimalPositiveDecimalCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/number/big/decimal/99999999.99"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// GetBigDecimalPositiveDecimalHandleResponse handles the GetBigDecimalPositiveDecimal response.
func (NumberOperations) GetBigDecimalPositiveDecimalHandleResponse(resp *azcore.Response) (*NumberGetBigDecimalPositiveDecimalResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := NumberGetBigDecimalPositiveDecimalResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetBigDoubleCreateRequest creates the GetBigDouble request.
func (NumberOperations) GetBigDoubleCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/number/big/double/2.5976931e+101"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// GetBigDoubleHandleResponse handles the GetBigDouble response.
func (NumberOperations) GetBigDoubleHandleResponse(resp *azcore.Response) (*NumberGetBigDoubleResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := NumberGetBigDoubleResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetBigDoubleNegativeDecimalCreateRequest creates the GetBigDoubleNegativeDecimal request.
func (NumberOperations) GetBigDoubleNegativeDecimalCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/number/big/double/-99999999.99"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// GetBigDoubleNegativeDecimalHandleResponse handles the GetBigDoubleNegativeDecimal response.
func (NumberOperations) GetBigDoubleNegativeDecimalHandleResponse(resp *azcore.Response) (*NumberGetBigDoubleNegativeDecimalResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := NumberGetBigDoubleNegativeDecimalResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetBigDoublePositiveDecimalCreateRequest creates the GetBigDoublePositiveDecimal request.
func (NumberOperations) GetBigDoublePositiveDecimalCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/number/big/double/99999999.99"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// GetBigDoublePositiveDecimalHandleResponse handles the GetBigDoublePositiveDecimal response.
func (NumberOperations) GetBigDoublePositiveDecimalHandleResponse(resp *azcore.Response) (*NumberGetBigDoublePositiveDecimalResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := NumberGetBigDoublePositiveDecimalResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetBigFloatCreateRequest creates the GetBigFloat request.
func (NumberOperations) GetBigFloatCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/number/big/float/3.402823e+20"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// GetBigFloatHandleResponse handles the GetBigFloat response.
func (NumberOperations) GetBigFloatHandleResponse(resp *azcore.Response) (*NumberGetBigFloatResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := NumberGetBigFloatResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetInvalidDecimalCreateRequest creates the GetInvalidDecimal request.
func (NumberOperations) GetInvalidDecimalCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/number/invaliddecimal"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// GetInvalidDecimalHandleResponse handles the GetInvalidDecimal response.
func (NumberOperations) GetInvalidDecimalHandleResponse(resp *azcore.Response) (*NumberGetInvalidDecimalResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := NumberGetInvalidDecimalResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetInvalidDoubleCreateRequest creates the GetInvalidDouble request.
func (NumberOperations) GetInvalidDoubleCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/number/invaliddouble"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// GetInvalidDoubleHandleResponse handles the GetInvalidDouble response.
func (NumberOperations) GetInvalidDoubleHandleResponse(resp *azcore.Response) (*NumberGetInvalidDoubleResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := NumberGetInvalidDoubleResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetInvalidFloatCreateRequest creates the GetInvalidFloat request.
func (NumberOperations) GetInvalidFloatCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/number/invalidfloat"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// GetInvalidFloatHandleResponse handles the GetInvalidFloat response.
func (NumberOperations) GetInvalidFloatHandleResponse(resp *azcore.Response) (*NumberGetInvalidFloatResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := NumberGetInvalidFloatResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetNullCreateRequest creates the GetNull request.
func (NumberOperations) GetNullCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/number/null"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// GetNullHandleResponse handles the GetNull response.
func (NumberOperations) GetNullHandleResponse(resp *azcore.Response) (*NumberGetNullResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := NumberGetNullResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetSmallDecimalCreateRequest creates the GetSmallDecimal request.
func (NumberOperations) GetSmallDecimalCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/number/small/decimal/2.5976931e-101"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// GetSmallDecimalHandleResponse handles the GetSmallDecimal response.
func (NumberOperations) GetSmallDecimalHandleResponse(resp *azcore.Response) (*NumberGetSmallDecimalResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := NumberGetSmallDecimalResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetSmallDoubleCreateRequest creates the GetSmallDouble request.
func (NumberOperations) GetSmallDoubleCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/number/small/double/2.5976931e-101"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// GetSmallDoubleHandleResponse handles the GetSmallDouble response.
func (NumberOperations) GetSmallDoubleHandleResponse(resp *azcore.Response) (*NumberGetSmallDoubleResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := NumberGetSmallDoubleResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetSmallFloatCreateRequest creates the GetSmallFloat request.
func (NumberOperations) GetSmallFloatCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/number/small/float/3.402823e-20"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// GetSmallFloatHandleResponse handles the GetSmallFloat response.
func (NumberOperations) GetSmallFloatHandleResponse(resp *azcore.Response) (*NumberGetSmallFloatResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := NumberGetSmallFloatResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// PutBigDecimalCreateRequest creates the PutBigDecimal request.
func (NumberOperations) PutBigDecimalCreateRequest(u url.URL, numberBody float64) (*azcore.Request, error) {
	urlPath := "/number/big/decimal/2.5976931e+101"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(numberBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutBigDecimalHandleResponse handles the PutBigDecimal response.
func (NumberOperations) PutBigDecimalHandleResponse(resp *azcore.Response) (*NumberPutBigDecimalResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &NumberPutBigDecimalResponse{RawResponse: resp.Response}, nil
}

// PutBigDecimalNegativeDecimalCreateRequest creates the PutBigDecimalNegativeDecimal request.
func (NumberOperations) PutBigDecimalNegativeDecimalCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/number/big/decimal/-99999999.99"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(-99999999.99)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutBigDecimalNegativeDecimalHandleResponse handles the PutBigDecimalNegativeDecimal response.
func (NumberOperations) PutBigDecimalNegativeDecimalHandleResponse(resp *azcore.Response) (*NumberPutBigDecimalNegativeDecimalResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &NumberPutBigDecimalNegativeDecimalResponse{RawResponse: resp.Response}, nil
}

// PutBigDecimalPositiveDecimalCreateRequest creates the PutBigDecimalPositiveDecimal request.
func (NumberOperations) PutBigDecimalPositiveDecimalCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/number/big/decimal/99999999.99"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(99999999.99)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutBigDecimalPositiveDecimalHandleResponse handles the PutBigDecimalPositiveDecimal response.
func (NumberOperations) PutBigDecimalPositiveDecimalHandleResponse(resp *azcore.Response) (*NumberPutBigDecimalPositiveDecimalResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &NumberPutBigDecimalPositiveDecimalResponse{RawResponse: resp.Response}, nil
}

// PutBigDoubleCreateRequest creates the PutBigDouble request.
func (NumberOperations) PutBigDoubleCreateRequest(u url.URL, numberBody float64) (*azcore.Request, error) {
	urlPath := "/number/big/double/2.5976931e+101"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(numberBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutBigDoubleHandleResponse handles the PutBigDouble response.
func (NumberOperations) PutBigDoubleHandleResponse(resp *azcore.Response) (*NumberPutBigDoubleResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &NumberPutBigDoubleResponse{RawResponse: resp.Response}, nil
}

// PutBigDoubleNegativeDecimalCreateRequest creates the PutBigDoubleNegativeDecimal request.
func (NumberOperations) PutBigDoubleNegativeDecimalCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/number/big/double/-99999999.99"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(-99999999.99)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutBigDoubleNegativeDecimalHandleResponse handles the PutBigDoubleNegativeDecimal response.
func (NumberOperations) PutBigDoubleNegativeDecimalHandleResponse(resp *azcore.Response) (*NumberPutBigDoubleNegativeDecimalResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &NumberPutBigDoubleNegativeDecimalResponse{RawResponse: resp.Response}, nil
}

// PutBigDoublePositiveDecimalCreateRequest creates the PutBigDoublePositiveDecimal request.
func (NumberOperations) PutBigDoublePositiveDecimalCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/number/big/double/99999999.99"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(99999999.99)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutBigDoublePositiveDecimalHandleResponse handles the PutBigDoublePositiveDecimal response.
func (NumberOperations) PutBigDoublePositiveDecimalHandleResponse(resp *azcore.Response) (*NumberPutBigDoublePositiveDecimalResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &NumberPutBigDoublePositiveDecimalResponse{RawResponse: resp.Response}, nil
}

// PutBigFloatCreateRequest creates the PutBigFloat request.
func (NumberOperations) PutBigFloatCreateRequest(u url.URL, numberBody float32) (*azcore.Request, error) {
	urlPath := "/number/big/float/3.402823e+20"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(numberBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutBigFloatHandleResponse handles the PutBigFloat response.
func (NumberOperations) PutBigFloatHandleResponse(resp *azcore.Response) (*NumberPutBigFloatResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &NumberPutBigFloatResponse{RawResponse: resp.Response}, nil
}

// PutSmallDecimalCreateRequest creates the PutSmallDecimal request.
func (NumberOperations) PutSmallDecimalCreateRequest(u url.URL, numberBody float64) (*azcore.Request, error) {
	urlPath := "/number/small/decimal/2.5976931e-101"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(numberBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutSmallDecimalHandleResponse handles the PutSmallDecimal response.
func (NumberOperations) PutSmallDecimalHandleResponse(resp *azcore.Response) (*NumberPutSmallDecimalResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &NumberPutSmallDecimalResponse{RawResponse: resp.Response}, nil
}

// PutSmallDoubleCreateRequest creates the PutSmallDouble request.
func (NumberOperations) PutSmallDoubleCreateRequest(u url.URL, numberBody float64) (*azcore.Request, error) {
	urlPath := "/number/small/double/2.5976931e-101"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(numberBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutSmallDoubleHandleResponse handles the PutSmallDouble response.
func (NumberOperations) PutSmallDoubleHandleResponse(resp *azcore.Response) (*NumberPutSmallDoubleResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &NumberPutSmallDoubleResponse{RawResponse: resp.Response}, nil
}

// PutSmallFloatCreateRequest creates the PutSmallFloat request.
func (NumberOperations) PutSmallFloatCreateRequest(u url.URL, numberBody float32) (*azcore.Request, error) {
	urlPath := "/number/small/float/3.402823e-20"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(numberBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutSmallFloatHandleResponse handles the PutSmallFloat response.
func (NumberOperations) PutSmallFloatHandleResponse(resp *azcore.Response) (*NumberPutSmallFloatResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &NumberPutSmallFloatResponse{RawResponse: resp.Response}, nil
}
