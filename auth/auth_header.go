package auth

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/iota-uz/click"
	"time"
)

func GenerateAuthHeader(t time.Time, merchantUserId int64, secretKey string) string {
	timestamp := t.Unix()
	data := fmt.Sprintf("%d%s", timestamp, secretKey)

	hash := sha1.New()
	hash.Write([]byte(data))
	digest := hex.EncodeToString(hash.Sum(nil))

	return fmt.Sprintf("%d:%s:%d", merchantUserId, digest, timestamp)
}

func WithAuthContext(ctx context.Context, merchantUserId int64, secretKey string, t time.Time) context.Context {
	authHeader := GenerateAuthHeader(t, merchantUserId, secretKey)

	return context.WithValue(
		ctx,
		clickapi.ContextAPIKeys,
		map[string]clickapi.APIKey{
			"AuthHeader": {Key: authHeader},
		},
	)
}
