package auth

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

// amountCandidates returns the set of amount string representations that Click
// may have signed. Click signs the merchant URL byte-for-byte, so the same
// numeric amount can appear as "38400", "38400.00", "38400.5", "38400.50",
// etc. depending on how the merchant formatted it when generating the payment
// link. We accept any of these to stay compatible with both legacy
// integer-amount links and current N.NN links.
func amountCandidates(amount float64) []string {
	seen := make(map[string]struct{}, 4)
	out := make([]string, 0, 4)
	add := func(s string) {
		if _, ok := seen[s]; ok {
			return
		}
		seen[s] = struct{}{}
		out = append(out, s)
	}

	// N.NN — current Click button spec (https://docs.click.uz/click-button/).
	add(fmt.Sprintf("%.2f", amount))
	// Minimal representation, e.g. 38400 or 38400.5 — preserves whatever
	// fractional digits the merchant URL actually had.
	add(strconv.FormatFloat(amount, 'f', -1, 64))
	// Integer-only — legacy v1.0.4 behaviour. Some merchants still have payment
	// links generated with %.0f.
	add(fmt.Sprintf("%.0f", amount))
	// One-decimal representation, e.g. 38400.5, in case Click trims trailing
	// zeros only.
	add(fmt.Sprintf("%.1f", amount))
	return out
}

// equalSign computes md5(payload) and compares against actualSign in constant
// time.
func equalSign(payload, actualSign string) bool {
	hash := md5.Sum([]byte(payload))
	return hmac.Equal([]byte(hex.EncodeToString(hash[:])), []byte(actualSign))
}

// ValidatePrepareSignString validates sign_string for the Prepare phase.
// Formula: md5(click_trans_id + service_id + SECRET_KEY + merchant_trans_id + amount + action + sign_time)
//
// The amount segment is matched against several string representations
// (see amountCandidates) so that signatures generated against either
// "38400" or "38400.00" both validate.
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
	for _, amountStr := range amountCandidates(amount) {
		payload := fmt.Sprintf(
			"%d%d%s%s%s%d%s",
			clickTransId,
			serviceId,
			secretKey,
			merchantTransId,
			amountStr,
			action,
			signTime,
		)
		if equalSign(payload, actualSign) {
			return true
		}
	}
	return false
}

// ValidateCompleteSignString validates sign_string for the Complete phase.
// Formula: md5(click_trans_id + service_id + SECRET_KEY + merchant_trans_id + merchant_prepare_id + amount + action + sign_time)
//
// See ValidatePrepareSignString for the amount-matching rationale.
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
	for _, amountStr := range amountCandidates(amount) {
		payload := fmt.Sprintf(
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
		if equalSign(payload, actualSign) {
			return true
		}
	}
	return false
}
