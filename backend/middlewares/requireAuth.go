package middlewares

import (
	"context"
	"ecommerce-app/config"
	"ecommerce-app/initializers"
	"ecommerce-app/models"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"clevergo.tech/jsend"
)

type CtxKey string

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerTokenArr := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		bearerToken := bearerTokenArr[len(bearerTokenArr)-1]
		apiKey := r.Header.Get("x-api-key")
		actualApiKey := os.Getenv("ADMIN_API_KEY")
		if (apiKey == actualApiKey) && (actualApiKey != "") {
			// Create a temp admin user object
			type CtxKey string
			var userCtxKey CtxKey = "user"
			ctx := context.WithValue(r.Context(), userCtxKey, models.User{Role: "admin"})
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		if bearerToken == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Decode/validate it
		verifier := config.GetVerifier()
		ctx := context.Background()

		// JWT token from identify provider
		idToken, err := verifier.Verify(ctx, bearerToken)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		var claims config.IDTokenClaims
		if err := idToken.Claims(&claims); err != nil {
			// handle error
			fmt.Printf("Sub not found")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Find the user with token sub
		var user models.User
		initializers.Db.Where("sub = ?", claims.Sub).Joins("DefaultAddress").First(&user)

		fmt.Println("user", user.Name, user.ID)
		if user.ID == 0 {
			log.Println("User not found")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Attach user to context
		var userCtxKey CtxKey = "user"
		ctx = context.WithValue(ctx, userCtxKey, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RequireAdmin(w http.ResponseWriter, r *http.Request) {
	requestUser := r.Context().Value("user")
	if requestUser.(models.User).Role != "admin" {
		jsend.Fail(w, "", http.StatusForbidden)
		return
	}
}

func GerUserFromContext(r *http.Request) (models.User, error) {
	ctx := r.Context()

	var userCtxKey CtxKey = "user"
	userI := ctx.Value(userCtxKey)
	if user, ok := userI.(models.User); ok {
		return user, nil
	}
	return models.User{}, errors.New("cannot get user")
}
