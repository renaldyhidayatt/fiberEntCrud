package pkg

import "github.com/golang-jwt/jwt/v4"

func ValidToken(t *jwt.Token) float64 {

	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))

	return float64(uid)
}
