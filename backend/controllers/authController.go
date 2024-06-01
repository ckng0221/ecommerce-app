package controllers

import (
	"context"
	"ecommerce-app/config"
	"ecommerce-app/initializers"
	"ecommerce-app/middlewares"
	"ecommerce-app/models"
	"ecommerce-app/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"clevergo.tech/jsend"
	"github.com/coreos/go-oidc"
)

// login when having cookies
func Login(w http.ResponseWriter, r *http.Request) {
	// User login based on access token
	var bodyObj struct {
		Code  string
		State string
		Nonce string
	}

	cookieState, err := r.Cookie("state")
	if err != nil {
		jsend.Fail(w, "State not found", http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		jsend.Fail(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &bodyObj)
	if err != nil {
		jsend.Fail(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	if bodyObj.Code == "" {
		jsend.Fail(w, "No access token", http.StatusUnauthorized)
		return
	}

	tokenClaims, idToken, err := getTokenClaimJwtFromLogin(bodyObj.Code, bodyObj.State, cookieState.Value, bodyObj.Nonce)
	if err != nil {
		log.Print(err.Error())
		jsend.Fail(w, err, http.StatusUnauthorized)
		return
	}

	// register user if not exist
	var user models.User
	err = initializers.Db.Where("sub = ?", tokenClaims.Sub).Limit(1).Find(&user).Error
	log.Println(err)
	sub := tokenClaims.Sub
	profilePic := tokenClaims.Picture
	if user.ID == 0 {
		log.Printf("user sub:  %s not found, creating user...\n", tokenClaims.Sub)
		initializers.Db.Create(&models.User{
			Name:       tokenClaims.Name,
			Email:      tokenClaims.Email,
			Sub:        &sub,
			ProfilePic: &profilePic,
		})
		log.Printf("user %s created\n", tokenClaims.Name)
	}

	// Set cookies
	cookie := http.Cookie{
		Name:     "Authorization",
		Value:    idToken,
		MaxAge:   3600 * 24 * 30,
		Path:     "/",
		Domain:   "",
		Secure:   false,
		HttpOnly: false,
	}
	http.SetCookie(w, &cookie)

	// respond
	jsend.Success(w, map[string]string{"name": tokenClaims.Name, "sub": tokenClaims.Sub, "access_token": idToken})
}

// Google login first > user login
func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	state, err := utils.RandString(16)
	if err != nil {
		log.Println(err)
		jsend.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	nonce, err := utils.RandString(16)
	if err != nil {
		log.Println(err)
		jsend.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// c.SetCookie("state", state, 3600*24*30, "", "", false, false)
	// c.SetCookie("nonce", nonce, 3600*24*30, "", "", false, false)

	config, err := config.GoogleConfig()
	if err != nil {
		log.Println(err)
		jsend.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	url := config.AuthCodeURL(state, oidc.Nonce(nonce))

	jsend.Success(w, map[string]string{
		"state": state,
		"nonce": nonce,
		"url":   url,
	})
}

func getTokenClaimJwtFromLogin(code, state, cookieState, nonce string) (config.IDTokenClaims, string, error) {
	if state != cookieState {
		return config.IDTokenClaims{}, "", errors.New("state does not match")
	}
	ctx := context.Background()
	authConfig, _ := config.GoogleConfig()
	verifier := config.GetVerifier()

	oauth2Token, err := authConfig.Exchange(ctx, code)
	if err != nil {
		return config.IDTokenClaims{}, "", err
	}

	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		log.Println("No id_token field in oauth2 token.")
		return config.IDTokenClaims{}, "", err
	}

	// JWT token from identify provider
	idToken, err := verifier.Verify(ctx, rawIDToken)
	if err != nil {
		log.Println("Failed to verify id token", err)
		return config.IDTokenClaims{}, "", err
	}

	if idToken.Nonce != nonce {
		return config.IDTokenClaims{}, "", errors.New("nonce does not match")
	}

	var tokenClaims config.IDTokenClaims
	if err := idToken.Claims(&tokenClaims); err != nil {
		// handle error
		log.Println("Failed to unmarshal claim")
		return config.IDTokenClaims{}, "", err
	}

	return tokenClaims, rawIDToken, nil
}

func Validate(w http.ResponseWriter, r *http.Request) {
	user, _ := middlewares.GerUserFromContext(r)

	jsend.Success(w, user, http.StatusOK)
}

func requireOwner(r *http.Request, ownerId string) error {
	requestUser, err := middlewares.GerUserFromContext(r)
	if err != nil {
		return errors.New("failed to get user from context")
	}
	if fmt.Sprint(requestUser.ID) != ownerId && requestUser.Role != "admin" {
		return utils.ErrForbidden
	}
	return nil
}

func requireAdmin(r *http.Request) error {
	requestUser, _ := middlewares.GerUserFromContext(r)
	if requestUser.Role != "admin" {
		return utils.ErrForbidden
	}
	return nil
}
