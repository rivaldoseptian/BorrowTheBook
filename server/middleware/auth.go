package middleware

import (
	"context"
	"net/http"
	"server/helpers"
	"server/models"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		accessToken := r.Header.Get("accesstoken")

		if accessToken == "" {
			helpers.Response(w, 401, "Unauthorize", nil)
			return
		}

		user, err := helpers.ValidateToken(accessToken)
		if err != nil {
			helpers.Response(w, 401, err.Error(), nil)
			return
		}

		ctx := context.WithValue(r.Context(), "userinfo", user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AdminAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user := r.Context().Value("userinfo").(*helpers.MyCustomClaims)

		var borrow models.Borrow

		if borrow.UserID != user.ID && user.Role != "Admin" {
			helpers.Response(w, 401, "You Not Authorize", nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}
