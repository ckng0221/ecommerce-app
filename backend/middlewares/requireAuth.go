package middlewares

import (
	"context"
	"ecommerce-app/config"
	"ecommerce-app/initializers"
	"ecommerce-app/models"
	"fmt"
	"net/http"
	"os"
	"strings"

	"clevergo.tech/jsend"
)

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerTokenArr := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		bearerToken := bearerTokenArr[len(bearerTokenArr)-1]
		apiKey := r.Header.Get("x-api-key")
		if apiKey == os.Getenv("ADMIN_API_KEY") {
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
		initializers.Db.Where("sub = ?", claims.Sub).First(&user)

		if user.ID == 0 {
			fmt.Println("User not found")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Attach user to context
		type CtxKey string
		var userCtxKey CtxKey = "user"
		ctx = context.WithValue(r.Context(), userCtxKey, user)
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
