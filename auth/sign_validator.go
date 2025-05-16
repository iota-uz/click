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
	clickTransID int64,
	serviceID int,
	secretKey string,
	merchantTransID string,
	amount float64,
	action int,
	signTime string, // format: "YYYY-MM-DD HH:mm:ss"
) bool {
	data := fmt.Sprintf(
		"%d%d%s%s%.2f%d%s",
		clickTransID,
		serviceID,
		secretKey,
		merchantTransID,
		amount,
		action,
		signTime,
	)

	hash := md5.Sum([]byte(data))
	expectedSign := hex.EncodeToString(hash[:])
	return hmac.Equal([]byte(actualSign), []byte(expectedSign))
}

// ValidateCompleteSignString validates sign_string for the Complete phase.
// Formula: md5(click_trans_id + service_id + SECRET_KEY + merchant_trans_id + merchant_prepare_id + amount + action + sign_time)
func ValidateCompleteSignString(
	actualSign string,
	clickTransID int64,
	serviceID int,
	secretKey string,
	merchantTransID string,
	merchantPrepareID int64,
	amount float64,
	action int,
	signTime string, // format: "YYYY-MM-DD HH:mm:ss"
) bool {
	data := fmt.Sprintf(
		"%d%d%s%s%d%.2f%d%s",
		clickTransID,
		serviceID,
		secretKey,
		merchantTransID,
		merchantPrepareID,
		amount,
		action,
		signTime,
	)

	hash := md5.Sum([]byte(data))
	expectedSign := hex.EncodeToString(hash[:])
	return hmac.Equal([]byte(actualSign), []byte(expectedSign))
}
