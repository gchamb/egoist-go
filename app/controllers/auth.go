package controllers

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"

	"net/http"

	"egoist/app"
	"egoist/internal/structs"
	"egoist/internal/utils"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig = &oauth2.Config{
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",
	},
	Endpoint: google.Endpoint,
}

func SignInWithGoogle(global *app.Globals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authCode := r.Header.Get("Authorization")
		if authCode == "" {
			log.Println("Authorization code was empty")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		oauthToken, err := googleOauthConfig.Exchange(r.Context(), authCode)
		if err != nil {
			log.Fatal(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + oauthToken.AccessToken)
		if err != nil {
			log.Println("failed getting user info:", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer response.Body.Close()
		contents, err := io.ReadAll(response.Body)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var googleUser structs.GoogleUser
		if err = json.Unmarshal(contents, &googleUser); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// happens no matter the provider
		user, err := global.Queries.GetUserByEmail(googleUser.Email)
		if err != nil && err != sql.ErrNoRows {
			log.Printf(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var id string
		isOnboarded := false
		if err == sql.ErrNoRows {
			createdUserId, err := global.Queries.CreateUser(&googleUser.Email, nil, nil)
			if err != nil {
				log.Println(err.Error())
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			id = createdUserId
		} else {
			id = user.ID
			if user.GoalWeight != nil {
				isOnboarded = true
			}
		}
	
		tokens, err := utils.GenerateTokens(id)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		tokens.Uid = id
		tokens.IsOnboarded = isOnboarded

		utils.ReturnJson(w, tokens, http.StatusOK)
	}
}

func SignInWithApple(global *app.Globals) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		idToken := r.Header.Get("Authorization")

		if idToken == "" {
			log.Println("Authorization code was empty")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		claims, err := utils.VerifyToken(idToken, true)
		if err != nil {
			log.Fatal(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		
		appleId, err := claims.GetSubject()
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		user, err := global.Queries.GetUserByAppleID(appleId)
		if err != nil && err != sql.ErrNoRows {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var id string
		isOnboarded := false	
		if err == sql.ErrNoRows {
			createdId, err := global.Queries.CreateUser(nil, nil, &appleId)

			if err != nil {
				log.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			id = createdId
		}else {
			id = user.ID
			if user.GoalWeight != nil {
				isOnboarded = true
			}
		}

		tokens, err := utils.GenerateTokens(id)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		tokens.Uid = id
		tokens.IsOnboarded = isOnboarded

		utils.ReturnJson(w, tokens, http.StatusOK)
	}
}

func SignInWithEmail(global *app.Globals) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody structs.AuthRequest
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&requestBody)

		err := requestBody.ValidateAuthRequest()

		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		user, err := global.Queries.GetUserByEmail(requestBody.Email)
		isOnboarded := false
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(requestBody.Password)); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if user.GoalWeight != nil {
			isOnboarded = true
		}

		tokens, err := utils.GenerateTokens(user.ID)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		tokens.IsOnboarded = isOnboarded
		tokens.Uid = user.ID

		utils.ReturnJson(w, tokens, http.StatusOK)
	}

}

func SignUpWithEmail(global *app.Globals) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody structs.AuthRequest
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&requestBody)

		err := requestBody.ValidateAuthRequest()

		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		user, err := global.Queries.GetUserByEmail(requestBody.Email)
		if err != nil && err != sql.ErrNoRows{
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if user.ID != "" {
			log.Println("user already exists")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		password, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), 12)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		passwordAsStr := string(password)
		userId, err := global.Queries.CreateUser(&requestBody.Email, &passwordAsStr, nil)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		tokens, err := utils.GenerateTokens(userId)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		tokens.Uid = userId

		utils.ReturnJson(w, tokens, http.StatusOK)
	}

}
