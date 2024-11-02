package middlewares

import (
	"context"
	"net/http"
	"strings"

	middlewareenum "github.com/negeek/ecommerce-api-assessment/enums"
	"github.com/negeek/ecommerce-api-assessment/repositories"
	"github.com/negeek/ecommerce-api-assessment/utils"
)

func AuthenticationMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")
		if bearerToken == "" {
			utils.JsonResponse(w, false, http.StatusUnauthorized, "Provide Auth Token", nil)
			return
		}

		bearerTokenArr := strings.Split(bearerToken, " ")
		if len(bearerTokenArr) != 2 {
			utils.JsonResponse(w, false, http.StatusUnauthorized, "Invalid Authorization Header", nil)
			return
		}

		bearer, token := bearerTokenArr[0], bearerTokenArr[1]
		if bearer != "Bearer" {
			utils.JsonResponse(w, false, http.StatusUnauthorized, "Invalid Authorization Header", nil)
			return
		}

		claim, err := utils.VerifyJwt(token)
		if err != nil {
			utils.JsonResponse(w, false, http.StatusUnauthorized, "Invalid Token", nil)
			return
		}

		user := &repositories.User{}
		err = user.FindByID(claim.ID)
		if err != nil {
			utils.JsonResponse(w, false, http.StatusUnauthorized, "Invalid User", nil)
			return
		}

		ctxWithUser := context.WithValue(r.Context(), middlewareenum.UserContextKey, claim.ID)
		rWithUser := r.WithContext(ctxWithUser)

		handler.ServeHTTP(w, rWithUser)
	})
}

func AuthAdminMiddleware(handler http.Handler) http.Handler {
	return AuthenticationMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value(middlewareenum.UserContextKey).(int)
		user := &repositories.User{}
		err := user.FindByID(userID)
		if err != nil || !user.IsAdmin() {
			utils.JsonResponse(w, false, http.StatusForbidden, "Access Denied: Admin privileges required", nil)
			return
		}
		handler.ServeHTTP(w, r)
	}))
}
