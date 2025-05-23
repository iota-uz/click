/*
CLICK-PASS-API

Testing FiscalizationAPIService

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

func Test_clickapi_FiscalizationAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FiscalizationAPIService GetFiscalData", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serviceId int64
		var paymentId int64

		resp, httpRes, err := apiClient.FiscalizationAPI.GetFiscalData(context.Background(), serviceId, paymentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FiscalizationAPIService SubmitFiscalItems", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FiscalizationAPI.SubmitFiscalItems(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FiscalizationAPIService SubmitFiscalQRCode", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FiscalizationAPI.SubmitFiscalQRCode(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
