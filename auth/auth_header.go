package auth

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/iota-uz/click/clickapi"
	"time"
)

func GenerateAuthHeader(t time.Time, merchantUserID int32, secretKey string) string {
	timestamp := t.Unix()
	data := fmt.Sprintf("%d%s", timestamp, secretKey)

	hash := sha1.New()
	hash.Write([]byte(data))
	digest := hex.EncodeToString(hash.Sum(nil))

	return fmt.Sprintf("%d:%s:%d", merchantUserID, digest, timestamp)
}

func WithAuthContext(ctx context.Context, merchantUserID int32, secretKey string, t time.Time) context.Context {
	authHeader := GenerateAuthHeader(t, merchantUserID, secretKey)

	return context.WithValue(
		ctx,
		clickapi.ContextAPIKeys,
		map[string]clickapi.APIKey{
			"AuthHeader": {Key: authHeader},
		},
	)
}
