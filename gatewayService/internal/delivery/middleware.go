package delivery

import (
	"context"
	"fmt"

	"github.com/mohammaderm/todoMicroService/gatewayService/config"
	"github.com/mohammaderm/todoMicroService/gatewayService/internal/types"
	"github.com/mohammaderm/todoMicroService/gatewayService/pkg/monitoring"

	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

type key int

const accountInfoKeyCtx key = iota + 1

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

type JwtClaims struct {
	Email string `json:"email"`
	Id    uint   `json:"id"`
	jwt.StandardClaims
}

func NewResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}
func Auth(cfg *config.Token) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header["Token"] != nil {
				token, err := jwt.ParseWithClaims(r.Header["Token"][0], &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("invalid signing method")
					}
					return []byte(cfg.Secretkey), nil
				})

				if token != nil && token.Valid {
					claims, ok := token.Claims.(*JwtClaims)
					if !ok {
						w.WriteHeader(http.StatusUnauthorized)
						w.Write([]byte("Unauthorized"))
					}
					ctx := context.WithValue(r.Context(), accountInfoKeyCtx, types.AccountInfo{
						Id: uint64(claims.Id),
					})
					next.ServeHTTP(w, r.WithContext(ctx))
				} else {

					validationError, _ := err.(*jwt.ValidationError)
					if validationError.Errors&jwt.ValidationErrorExpired != 0 {
						w.WriteHeader(http.StatusUnauthorized)
						w.Write([]byte("Unauthorized/token expired"))
					} else if validationError.Errors&jwt.ValidationErrorIssuer != 0 {
						w.WriteHeader(http.StatusUnauthorized)
						w.Write([]byte("Unauthorized"))
					} else {
						w.WriteHeader(http.StatusUnauthorized)
						w.Write([]byte("Unauthorized"))
					}
				}

			} else {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			}
		})
	}
}

func MetricMiddleware(callector monitoring.MetricsCallectors) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.Path
			startedAt := time.Now()
			wr := NewResponseWriter(w)
			next.ServeHTTP(wr, r)
			duration := time.Since(startedAt)
			statusCode := wr.statusCode
			callector.HttpResponseTime(r.Method, path, statusCode, duration)
			// callector.HttpRequestCount(r.Method, path, statusCode)
		})
	}
}
