package tokenservice

import "encoding/base64"

func IsValidRefreshToken(token string) bool {
	b, err := base64.URLEncoding.DecodeString(token)

	return len(b) == 32 && err == nil
}
