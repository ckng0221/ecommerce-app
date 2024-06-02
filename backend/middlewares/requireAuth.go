package middlewares

import (
	"context"
	"ecommerce-app/config"
	"ecommerce-app/initializers"
	"ecommerce-app/models"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"clevergo.tech/jsend"
	"github.com/coreos/go-oidc/v3/oidc"
)

type CtxKey string

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerTokenArr := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		bearerToken := bearerTokenArr[len(bearerTokenArr)-1]
		apiKey := r.Header.Get("x-api-key")
		actualApiKey := os.Getenv("ADMIN_API_KEY")
		ctx := context.Background()
		var userCtxKey CtxKey = "user"
		var user models.User

		if (apiKey == actualApiKey) && (actualApiKey != "") {
			// Create a temp admin user object
			user.Role = "admin"
			ctx := context.WithValue(ctx, userCtxKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		if bearerToken == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Decode/validate it
		verifier := config.GetVerifier()

		// JWT token from identify provider
		// TODO: verify access token instead of id token
		idToken, err := verifier.Verify(ctx, bearerToken)
		if err != nil {
			if errors.Is(err, &oidc.TokenExpiredError{}) {
				log.Print(err.Error())
				log.Print("Token has expired")
				http.Error(w, "Token has expired", http.StatusUnauthorized)
				return
			} else {
				log.Print(err.Error())
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
		}

		var claims config.IDTokenClaims
		if err := idToken.Claims(&claims); err != nil {
			// handle error
			log.Printf("failed to parse claims")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Find the user with token sub
		// initializers.Db.Where("sub = ?", claims.Sub).Joins("DefaultAddress").First(&user)
		initializers.Db.Where("sub = ?", claims.Sub).First(&user)

		// log.Println("user", user.Name, user.ID)
		if user.ID == 0 {
			log.Println("User not found")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Attach user to context
		ctx = context.WithValue(ctx, userCtxKey, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GerUserFromContext(r *http.Request) (models.User, error) {
	ctx := r.Context()

	var userCtxKey CtxKey = "user"
	userI := ctx.Value(userCtxKey)

	if user, ok := userI.(models.User); ok {
		// fmt.Println("user", user)
		return user, nil
	}
	return models.User{}, errors.New("failed get user")
}

func RequireAdmin(w http.ResponseWriter, r *http.Request) {
	requestUser, err := GerUserFromContext(r)
	if err != nil {
		log.Println(err.Error())
		jsend.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if requestUser.Role != "admin" {
		jsend.Fail(w, "", http.StatusForbidden)
		return
	}
}
