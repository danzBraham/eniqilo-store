package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/danzBraham/eniqilo-store/internal/errors/autherror"
	"github.com/danzBraham/eniqilo-store/internal/helpers/httphelper"
	"github.com/danzBraham/eniqilo-store/internal/helpers/jwt"
)

type ContextKey string

const ContextUserIDKey ContextKey = "userID"

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			httphelper.ErrorResponse(w, http.StatusUnauthorized, autherror.ErrMissingAuthHeader)
			return
		}

		authFields := strings.Fields(authHeader)
		if len(authFields) < 2 || authFields[0] != "Bearer" {
			httphelper.ErrorResponse(w, http.StatusUnauthorized, autherror.ErrInvalidAuthHeader)
			return
		}

		tokenString := authFields[1]

		token, err := jwt.VerifyToken(tokenString)
		if err != nil {
			httphelper.ErrorResponse(w, http.StatusUnauthorized, err)
			return
		}

		ctx := context.WithValue(r.Context(), ContextUserIDKey, token.UserID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
