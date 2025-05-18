package auth

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// ValidatePrepareSignString validates sign_string for the Prepare phase.
// Formula: md5(click_trans_id + service_id + SECRET_KEY + merchant_trans_id + amount + action + sign_time)
func ValidatePrepareSignString(
	actualSign string,
	clickTransId int64,
	serviceId int64,
	secretKey string,
	merchantTransId string,
	amount float64,
	action int32,
	signTime string, // format: "YYYY-MM-DD HH:mm:ss"
) bool {
	// Formatted amount without decimal part
	amountStr := fmt.Sprintf("%.0f", amount)

	// We form a line according to the CLICK specification
	signString := fmt.Sprintf(
		"%d%d%s%s%s%d%s",
		clickTransId,
		serviceId,
		secretKey,
		merchantTransId,
		amountStr,
		action,
		signTime,
	)

	// Calculate MD5 hash
	hash := md5.Sum([]byte(signString))
	expectedSign := hex.EncodeToString(hash[:])

	// Compare safely
	return hmac.Equal([]byte(expectedSign), []byte(actualSign))
}

// ValidateCompleteSignString validates sign_string for the Complete phase.
// Formula: md5(click_trans_id + service_id + SECRET_KEY + merchant_trans_id + merchant_prepare_id + amount + action + sign_time)
func ValidateCompleteSignString(
	actualSign string,
	clickTransId int64,
	serviceId int64,
	secretKey string,
	merchantTransId string,
	merchantPrepareId int64,
	amount float64,
	action int32,
	signTime string, // format: "YYYY-MM-DD HH:mm:ss"
) bool {
	// Formatted amount without decimal part
	amountStr := fmt.Sprintf("%.0f", amount)

	// We form a line according to the CLICK specification
	signString := fmt.Sprintf(
		"%d%d%s%s%d%s%d%s",
		clickTransId,
		serviceId,
		secretKey,
		merchantTransId,
		merchantPrepareId,
		amountStr,
		action,
		signTime,
	)

	// Calculate MD5 hash
	hash := md5.Sum([]byte(signString))
	expectedSign := hex.EncodeToString(hash[:])

	// Compare safely
	return hmac.Equal([]byte(expectedSign), []byte(actualSign))
}
