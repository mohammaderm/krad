package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, _ := jwt.ParseWithClaims(r.Header["Token"][0], &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("invalid signing method")
				}
				return []byte(Secretkey), nil
			})
			if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
				ctx := context.WithValue(r.Context(), "UserId", claims.Id)
				next.ServeHTTP(w, r.WithContext(ctx))

			} else {
				// fmt.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Expired token"))
			}

		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
		}
	})
}

func (claims JwtClaims) Valid() error {
	var now = time.Now().UTC().Unix()
	if claims.VerifyExpiresAt(now, true) && claims.VerifyIssuer(Issue, true) {
		return nil
	}
	return fmt.Errorf("token expired")
}
