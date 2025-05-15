package auth

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// ValidateSignString verifies the validity of the sign_string based on all required parameters.
// All numeric values are passed in their native types and formatted precisely according to the protocol.
func ValidateSignString(
	actualSign string,
	clickTransID int64,
	serviceID int,
	secretKey string,
	merchantTransID string,
	amount float64,
	action int,
	signTime string, // format: "YYYY-MM-DD HH:mm:ss"
) bool {
	// Construct the raw data string in the exact required order
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

	// Generate the MD5 hash
	hash := md5.Sum([]byte(data))
	expectedSign := hex.EncodeToString(hash[:])

	// Securely compare the provided sign_string with the expected one
	return hmac.Equal([]byte(actualSign), []byte(expectedSign))
}
