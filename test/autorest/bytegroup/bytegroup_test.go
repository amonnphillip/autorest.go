// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package bytegrouptest

import (
	"context"
	"generatortests/autorest/generated/bytegroup"
	"generatortests/helpers"
	"net/http"
	"testing"
)

func getByteClient(t *testing.T) bytegroup.ByteOperations {
	client, err := bytegroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create byte client: %v", err)
	}
	return client.ByteOperations()
}

func TestGetEmpty(t *testing.T) {
	client := getByteClient(t)
	result, err := client.GetEmpty(context.Background())
	if err != nil {
		t.Fatalf("GetEmpty: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	helpers.DeepEqualOrFatal(t, result.Value, &[]byte{})
}

func TestGetInvalid(t *testing.T) {
	client := getByteClient(t)
	result, err := client.GetInvalid(context.Background())
	// TODO: verify error response is clear and actionable
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if result != nil {
		t.Fatal("expected nil result")
	}
}

func TestGetNonASCII(t *testing.T) {
	client := getByteClient(t)
	result, err := client.GetNonASCII(context.Background())
	if err != nil {
		t.Fatalf("GetNonASCII: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	helpers.DeepEqualOrFatal(t, result.Value, &[]byte{0xFF, 0xFE, 0xFD, 0xFC, 0xFB, 0xFA, 0xF9, 0xF8, 0xF7, 0xF6})
}

func TestGetNull(t *testing.T) {
	client := getByteClient(t)
	result, err := client.GetNull(context.Background())
	if err != nil {
		t.Fatalf("GetNull: %v", err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	helpers.DeepEqualOrFatal(t, result.Value, (*[]byte)(nil))
}

func TestPutNonASCII(t *testing.T) {
	client := getByteClient(t)
	result, err := client.PutNonASCII(context.Background(), []byte{0xFF, 0xFE, 0xFD, 0xFC, 0xFB, 0xFA, 0xF9, 0xF8, 0xF7, 0xF6})
	if err != nil {
		t.Fatalf("PutNonASCII: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}