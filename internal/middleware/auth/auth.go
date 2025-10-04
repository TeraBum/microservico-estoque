package middleware

import (
	"api-estoque/internal/config"
	"api-estoque/internal/utils"
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Email string `json:"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/name"`
	Role  string `json:"http://schemas.microsoft.com/ws/2008/06/identity/claims/role"`
	jwt.RegisteredClaims
}

type contextKey string

const userContextKey = contextKey("userClaims")

var logger = utils.SetupLogger()

func JWTAuthMiddleware(allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			tokenString := extractToken(r)
			if tokenString == "" {
				logger.Warn("Acesso negado por: header authorization faltando ou invalido")
				http.Error(w, "header authorization faltando ou invalido", http.StatusUnauthorized)
				return
			}

			claims := &Claims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					logger.Warn("Acesso negado por: metodo de assinatura invalido")
					return nil, errors.New("metodo de assinatura invalido")
				}
				return []byte(config.Env.JwtSecret), nil
			})

			if err != nil || !token.Valid {
				logger.Warn("Acesso negado por: token invalido ou expirado")
				http.Error(w, "token invalido ou expirado", http.StatusUnauthorized)
				return
			}

			if len(allowedRoles) > 0 && !hasAllowedRole(claims.Role, allowedRoles) {
				logger.Warn("Acesso negado por: role incompativel")
				http.Error(w, "forbidden: role incompativel", http.StatusForbidden)
				return
			}

			ctx := context.WithValue(r.Context(), userContextKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func extractToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}
	return parts[1]
}

func hasAllowedRole(userRole string, allowed []string) bool {
	for _, ar := range allowed {
		if userRole == ar {
			return true
		}
	}
	return false
}

func GetUserClaims(r *http.Request) *Claims {
	if claims, ok := r.Context().Value(userContextKey).(*Claims); ok {
		return claims
	}
	return nil
}
