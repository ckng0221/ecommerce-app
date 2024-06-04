package middlewares

import (
	"context"
	"ecommerce-app/initializers"
	"ecommerce-app/models"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"clevergo.tech/jsend"
	"github.com/golang-jwt/jwt"
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
		// verifier := config.GetVerifier()

		// JWT token from identify provider

		// idToken, err := verifier.Verify(ctx, bearerToken)
		// if err != nil {
		// 	if errors.Is(err, &oidc.TokenExpiredError{}) {
		// 		log.Print(err.Error())
		// 		log.Print("Token has expired")
		// 		http.Error(w, "Token has expired", http.StatusUnauthorized)
		// 		return
		// 	} else {
		// 		log.Print(err.Error())
		// 		http.Error(w, "Invalid token", http.StatusUnauthorized)
		// 		return
		// 	}
		// }

		// var claims config.IDTokenClaims
		// if err := idToken.Claims(&claims); err != nil {
		// 	// handle error
		// 	log.Printf("failed to parse claims")
		// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		// 	return
		// }

		// Decode/validate it
		token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			if err.Error() == "Token is expired" {
				jsend.Fail(w, "token expired", http.StatusUnauthorized)
				return
			}
			// return
			log.Println(err.Error())
			jsend.Fail(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// Check the exp
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				fmt.Println("token expired")
				jsend.Fail(w, "token expired", http.StatusUnauthorized)
				return
			}

			// Find the user with token sub
			var user models.User
			// initializers.Db.First(&user, claims["sub"])
			initializers.Db.Where("sub = ?", claims["sub"]).First(&user)

			if user.ID == 0 {
				fmt.Println("User not found")
				jsend.Fail(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			// fmt.Println(user.ID)
			// Attach to req
			ctx = context.WithValue(ctx, userCtxKey, user)

			// Continue
			next.ServeHTTP(w, r.WithContext(ctx))

			// fmt.Println(claims["foo"], claims["nbf"])
		} else {
			jsend.Fail(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
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
