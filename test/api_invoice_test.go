/*
CLICK-PASS-API

Testing InvoiceAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package clickapi

import (
	"context"
	openapiclient "github.com/iota-uz/click/clickapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_clickapi_InvoiceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test InvoiceAPIService CheckInvoiceStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serviceId int64
		var invoiceId int64

		resp, httpRes, err := apiClient.InvoiceAPI.CheckInvoiceStatus(context.Background(), serviceId, invoiceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InvoiceAPIService CreateInvoice", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.InvoiceAPI.CreateInvoice(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
