package controllers

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"

	"net/http"

	"egoist/internal/database"
	"egoist/internal/database/queries"
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

func SignInWithGoogle(w http.ResponseWriter, r *http.Request) {
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

	db := database.ConnectDB()
	queries := queries.New(db)

	user, err := queries.GetUserByEmail(googleUser.Email)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("user %s already exist.", googleUser.Email)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var id string
	isOnboarded := false
	if err == sql.ErrNoRows {
		createdUserId, err := queries.CreateUser(googleUser.Email, nil)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id = createdUserId
	}else{
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
	tokens.IsOnboarded = isOnboarded

	utils.ReturnJson(w, tokens, http.StatusOK)
}

func SignInWithEmail(w http.ResponseWriter, r *http.Request) {
	
	var requestBody structs.AuthRequest
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestBody)
	
	err := requestBody.ValidateAuthRequest()

	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	db := database.ConnectDB()
	queries := queries.New(db)

	user, err := queries.GetUserByEmail(requestBody.Email)
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

	utils.ReturnJson(w, tokens, http.StatusOK)
}

func SignUpWithEmail(w http.ResponseWriter, r *http.Request) {
	var requestBody structs.AuthRequest
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestBody)

	err := requestBody.ValidateAuthRequest()

	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	
	db := database.ConnectDB()
	queries := queries.New(db)

	if _, err := queries.GetUserByEmail(requestBody.Email); err != sql.ErrNoRows {
		log.Println(err.Error())
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
	userId, err := queries.CreateUser(requestBody.Email, &passwordAsStr)
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

	utils.ReturnJson(w, tokens, http.StatusOK)
}