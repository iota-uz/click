package auth

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"
)

func md5Hex(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

func TestValidatePrepareSignString_AmountWithDecimals(t *testing.T) {
	const (
		clickTransId    int64 = 3647147957
		serviceId       int64 = 73438
		secret                = "TEST_SECRET"
		merchantTransId       = "policy:ec197639-9b85-404c-a6ac-2752983055ed"
		amount                = 33139.73
		action          int32 = 0
		signTime              = "2026-04-29 10:38:52"
	)

	expected := md5Hex(fmt.Sprintf("%d%d%s%s%.2f%d%s",
		clickTransId, serviceId, secret, merchantTransId, amount, action, signTime,
	))

	if !ValidatePrepareSignString(
		expected, clickTransId, serviceId, secret, merchantTransId, amount, action, signTime,
	) {
		t.Fatalf("expected sign with %%.2f amount to validate")
	}
}

func TestValidatePrepareSignString_AmountWithoutDecimals(t *testing.T) {
	const (
		clickTransId    int64 = 1
		serviceId       int64 = 2
		secret                = "S"
		merchantTransId       = "M"
		amount                = 1000.0
		action          int32 = 0
		signTime              = "2025-01-01 00:00:00"
	)

	expected := md5Hex(fmt.Sprintf("%d%d%s%s%.2f%d%s",
		clickTransId, serviceId, secret, merchantTransId, amount, action, signTime,
	))

	if !ValidatePrepareSignString(
		expected, clickTransId, serviceId, secret, merchantTransId, amount, action, signTime,
	) {
		t.Fatalf("expected sign with %%.2f amount (1000.00) to validate")
	}
}

func TestValidatePrepareSignString_RejectsBadSign(t *testing.T) {
	if ValidatePrepareSignString(
		"deadbeef", 1, 2, "S", "M", 100.0, 0, "2025-01-01 00:00:00",
	) {
		t.Fatalf("expected invalid signature to be rejected")
	}
}

func TestValidateCompleteSignString_AmountWithDecimals(t *testing.T) {
	const (
		clickTransId      int64 = 9
		serviceId         int64 = 73438
		secret                  = "S"
		merchantTransId         = "policy:abc"
		merchantPrepareId int64 = 1700000000
		amount                  = 33139.73
		action            int32 = 0
		signTime                = "2026-04-29 10:38:52"
	)

	expected := md5Hex(fmt.Sprintf("%d%d%s%s%d%.2f%d%s",
		clickTransId, serviceId, secret, merchantTransId, merchantPrepareId, amount, action, signTime,
	))

	if !ValidateCompleteSignString(
		expected, clickTransId, serviceId, secret, merchantTransId, merchantPrepareId, amount, action, signTime,
	) {
		t.Fatalf("expected complete sign with %%.2f amount to validate")
	}
}

func TestValidateCompleteSignString_RejectsBadSign(t *testing.T) {
	if ValidateCompleteSignString(
		"deadbeef", 1, 2, "S", "M", 100, 100.0, 0, "2025-01-01 00:00:00",
	) {
		t.Fatalf("expected invalid signature to be rejected")
	}
}

// Regression: legacy merchant URLs were generated with %.0f, so the integer
// amount "38400" reaches Click and Click signs against that exact string.
// The validator must accept it even when the controller passes amount as a
// float64.
func TestValidatePrepareSignString_AcceptsIntegerAmountString(t *testing.T) {
	const (
		clickTransId    int64 = 3647911201
		serviceId       int64 = 73438
		secret                = "TEST_SECRET"
		merchantTransId       = "policy:c44a6289-e610-4442-9b7c-fafeabc5e9ba"
		amount                = 38400.0
		action          int32 = 0
		signTime              = "2026-04-29 18:02:18"
	)

	expected := md5Hex(fmt.Sprintf("%d%d%s%s%d%d%s",
		clickTransId, serviceId, secret, merchantTransId, int64(amount), action, signTime,
	))

	if !ValidatePrepareSignString(
		expected, clickTransId, serviceId, secret, merchantTransId, amount, action, signTime,
	) {
		t.Fatalf("expected sign with integer amount string (38400) to validate")
	}
}

func TestValidateCompleteSignString_AcceptsIntegerAmountString(t *testing.T) {
	const (
		clickTransId      int64 = 3647911201
		serviceId         int64 = 73438
		secret                  = "TEST_SECRET"
		merchantTransId         = "policy:c44a6289-e610-4442-9b7c-fafeabc5e9ba"
		merchantPrepareId int64 = 1714389738
		amount                  = 38400.0
		action            int32 = 1
		signTime                = "2026-04-29 18:05:00"
	)

	expected := md5Hex(fmt.Sprintf("%d%d%s%s%d%d%d%s",
		clickTransId, serviceId, secret, merchantTransId, merchantPrepareId, int64(amount), action, signTime,
	))

	if !ValidateCompleteSignString(
		expected, clickTransId, serviceId, secret, merchantTransId, merchantPrepareId, amount, action, signTime,
	) {
		t.Fatalf("expected complete sign with integer amount string (38400) to validate")
	}
}

// Regression: amounts with a single fractional digit (e.g. "38400.5") show
// up when Click strips trailing zeros. The validator must accept that too.
func TestValidatePrepareSignString_AcceptsSingleDecimalAmount(t *testing.T) {
	const (
		clickTransId    int64 = 1
		serviceId       int64 = 2
		secret                = "S"
		merchantTransId       = "M"
		amount                = 38400.5
		action          int32 = 0
		signTime              = "2025-01-01 00:00:00"
	)

	expected := md5Hex(fmt.Sprintf("%d%d%s%s%s%d%s",
		clickTransId, serviceId, secret, merchantTransId, "38400.5", action, signTime,
	))

	if !ValidatePrepareSignString(
		expected, clickTransId, serviceId, secret, merchantTransId, amount, action, signTime,
	) {
		t.Fatalf("expected sign with single-decimal amount (38400.5) to validate")
	}
}
