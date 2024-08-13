package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"net/http"

	"egoist/internal/database"
	"egoist/internal/database/queries"
	"egoist/internal/structs"

	"github.com/google/uuid"
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
	// get the authorization token
	authCode := r.Header.Get("Authorization")
	fmt.Println("inside", authCode)
	if authCode == "" {
		fmt.Println("Auth code is empty.")
		return
	}


	// get the access token
	oauthToken, err := googleOauthConfig.Exchange(r.Context(), authCode)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// get the google user info data
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + oauthToken.AccessToken)
	if err != nil {
	 fmt.Errorf("failed getting user info: %s", err.Error())
	 return
	}

	defer response.Body.Close()

	contents, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println("failed read response: %s", err.Error())
		return
	}

	var user structs.GoogleUser
	if err = json.Unmarshal(contents, &user); err != nil {
		fmt.Println(err.Error())
		return
	}

	db := database.ConnectDB()
	queries := queries.New(db)

	id, err := uuid.NewUUID()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	
	err = queries.InsertUser(structs.User{
		ID: id.String(),
		Email: user.Email,
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	// create user or generate JWT and Refresh Token
}

func SignInWithEmail(w http.ResponseWriter, r *http.Request) {
	// get the email and password

	// validate the inputs

	// validate if it exists

	// create the jwt and refresh token

}

func SignUpWithEmail(w http.ResponseWriter, r *http.Request) {
	// get the email and password

	// validate the inputs

	// validate if it exists

	// create the jwt and refresh token
}