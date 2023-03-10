// https://forum.duplicacy.com/t/onedrive-using-own-credentials/6507/8

package duplicacy

import (
	"encoding/json"
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/oauth2"
	"net/http"
)

const URLBASE string = "http://localhost"
const SERVERPORT string = "53682"

const STARTPAGE string = `
<html lang="en">
  <head>
    <title>OneDrive for Duplicacy</title>
    <link href="//maxcdn.bootstrapcdn.com/bootstrap/3.3.4/css/bootstrap.min.css" rel="stylesheet">
  </head>
  <body>
    <div class="container">
      <br>
      <p class="lead">
        This web page allows you to download Microsoft OneDrive credentials
        (access token and refresh token) to be used with Duplicacy.
      </p>
      <p class="lead">
        When you're ready to proceed, please click on "Download my credentials"
        below. It will take you to the OneDrive website and ask you to log in to OneDrive and grant
        Duplicacy the permission to back up files to your OneDrive.
      </p>
      <p class="lead">
        You will then receive a file named
        <code>%v</code> which can be supplied to Duplicacy when prompted.
        This web page never logs or saves any information, or in
       anyway interacts with OneDrive on your behalf beyond providing you with a
       token and refreshing that token for you.
      </p>
      <br><br>
      <p class="text-center">
        <a href="%v" type="button" class="btn btn-lg btn-info btn-fill">
          Download my credentials
        </a>
      </p>
    </div>
  </body>
</html>
`

var (
	oneOauthConfig oauth2.Config
	odbFileName    string
)

// Start page to authorize and download initial token...
func odbStartHandler(w http.ResponseWriter, r *http.Request) {
	url := oneOauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, fmt.Sprintf(STARTPAGE+"\n", odbFileName, url))
}

// OauthHandler ...
func odbOauthHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("code") == "" {
		http.Redirect(w, r, "/odb_start", 302)
		return
	}

	token, err := oneOauthConfig.Exchange(r.Context(), r.URL.Query().Get("code"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error exchanging the code for an access token: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Content-Disposition", "attachment; filename="+odbFileName)
	if err := json.NewEncoder(w).Encode(token); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding the token in JSON: %v", err), http.StatusInternalServerError)
		return
	}
}

// RefreshHandler ...
func odbRefreshHandler(w http.ResponseWriter, r *http.Request) {
	var token oauth2.Token
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&token); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding the token from the request: %v", err), http.StatusBadRequest)
		return
	}

	tokenSource := oneOauthConfig.TokenSource(r.Context(), &token)
	newToken, err := tokenSource.Token()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching a new token: %v", err), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(newToken); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding the token in JSON: %v", err), http.StatusInternalServerError)
		return
	}
}

func Authorize(preference Preference) {
	http.HandleFunc("/odb_start", odbStartHandler)
	http.HandleFunc("/", odbOauthHandler)
	http.HandleFunc("/odb_refresh", odbRefreshHandler)

	clientIdPrompt := fmt.Sprintf("Enter client_id for custom Azure app: ")
	clientId := GetPassword(preference, "odb_client_id", clientIdPrompt, true, true)

	clientSecretPrompt := fmt.Sprintf("Enter client_secret for custom Azure app: ")
	clientSecret := GetPassword(preference, "odb_client_secret", clientSecretPrompt, true, false)

	odbFileName = "odb-" + clientId + "-token.json"

	authURL := URLBASE + ":" + SERVERPORT + "/odb_start"

	oneOauthConfig = oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  URLBASE + ":" + SERVERPORT + "/",
		Scopes:       []string{"Files.ReadWrite", "offline_access"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://login.microsoftonline.com/common/oauth2/v2.0/authorize",
			TokenURL: "https://login.microsoftonline.com/common/oauth2/v2.0/token",
		},
	}

	_ = open.Start(authURL)
	LOG_INFO("AUTHORIZE_PROMPT", "If your browser doesn't open automatically go to the following link: %s", authURL)

	var err error = nil

	err = http.ListenAndServe("127.0.0.1:"+SERVERPORT, nil)

	if err != nil {
		fmt.Printf("Failed to start the server: %v\n", err)
	}
}
