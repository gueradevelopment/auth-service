package handlers

import (
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

// const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

var githubOauthConfig = &oauth2.Config{
	RedirectURL: "http://localhost:3000/auth/github/callback",
	// Scopes:      []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint: github.Endpoint,
}

func oauthGithubLogin(w http.ResponseWriter, r *http.Request) {
	// Create oauthState cookie
	oauthState := generateStateOauthCookie(w)
	githubOauthConfig.ClientID = os.Getenv("github_CLIENT_ID")
	githubOauthConfig.ClientSecret = os.Getenv("github_CLIENT_SECRET")
	u := googleOauthConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

// func generateStateOauthCookie(w http.ResponseWriter) string {
// 	var expiration = time.Now().Add(365 * 24 * time.Hour)

// 	b := make([]byte, 16)
// 	rand.Read(b)
// 	state := base64.URLEncoding.EncodeToString(b)
// 	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
// 	http.SetCookie(w, &cookie)

// 	return state
// }

// func oauthGoogleCallback(w http.ResponseWriter, r *http.Request) {
// 	// Read oauthState from Cookie
// 	oauthState, _ := r.Cookie("oauthstate")

// 	if r.FormValue("state") != oauthState.Value {
// 		log.Println("Invalid oauth Google state")
// 		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
// 		return
// 	}

// 	data, err := getUserDataFromGoogle(r.FormValue("code"))
// 	if err != nil {
// 		log.Println(err.Error())
// 		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
// 		return
// 	}

// 	// SUCCESS!!
// 	userToken := getToken(data)
// 	cookieEmail := &http.Cookie{
// 		Name:  "email",
// 		Value: userToken.Email,
// 		Path:  "/",
// 	}
// 	cookieToken := &http.Cookie{
// 		Name:  "token",
// 		Value: userToken.Token,
// 		Path:  "/",
// 	}
// 	http.SetCookie(w, cookieEmail)
// 	http.SetCookie(w, cookieToken)
// 	http.Redirect(w, r, "http://localhost:8080/#/", 301)
// 	// w.Header().Add("Content-Type", "application/json")
// 	// json.NewEncoder(w).Encode(userToken)
// 	// fmt.Fprintf(w, "UserInfo: %s\n", data)
// }

// func getUserDataFromGoogle(code string) ([]byte, error) {
// 	// Use code to get token and get user info from Google.

// 	token, err := googleOauthConfig.Exchange(context.Background(), code)
// 	if err != nil {
// 		return nil, fmt.Errorf("Wrong code exchange: %s", err.Error())
// 	}
// 	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
// 	if err != nil {
// 		return nil, fmt.Errorf("Failed to get user information: %s", err.Error())
// 	}
// 	defer response.Body.Close()
// 	contents, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		return nil, fmt.Errorf("Failed to read response: %s", err.Error())
// 	}
// 	return contents, nil
// }
