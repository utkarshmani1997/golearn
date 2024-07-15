package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
)

var (
	clientID             = "53f28acb35e6be15ec1a469806a021df"
	clientSecret         = "411ccdc049098daf83c1eb2b61dd3bfa29005adcf8ff9c1bfe4eb41a1415f6a5"
	redirectURI          = "http://localhost:9080/callback"                        // Callback URL registered with Jira
	jiraURL              = "https://utkarshmani.atlassian.net/rest/api/2/projects" // Endpoint to test connection
	jiraAuthURL          = "https://9909-167-103-24-113.ngrok-free.app/rest/oauth2/latest/authorize"
	jiraTokenURL         = "https://9909-167-103-24-113.ngrok-free.app/rest/oauth2/latest/token"
	jiraPermissionScopes = []string{"READ", "WRITE", "ADMIN"} //[]string{"read:issue-type:jira", "read:project:jira", "read:project.property:jira", "read:user:jira", "read:application-role:jira", "read:avatar:jira", "read:group:jira", "read:issue-type-hierarchy:jira", "read:project-category:jira", "read:project-version:jira", "read:project.component:jira", "read:field:jira", "read:field-configuration:jira", "read:issue-meta:jira", "write:issue:jira", "write:comment:jira", "write:comment.property:jira", "write:attachment:jira", "read:issue:jira", "read:label:jira", "offline_access", "read:issue-security-level:jira", "read:issue.vote:jira", "read:issue.changelog:jira", "read:status:jira", "read:comment:jira", "read:comment.property:jira", "read:project-role:jira"}
	jiraCloudIDURL       = "https://api.atlassian.com/oauth/token/accessible-resources"
)

var token *oauth2.Token

func main() {
	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURI,
		Endpoint: oauth2.Endpoint{
			AuthURL:  jiraAuthURL,
			TokenURL: jiraTokenURL,
		},
		Scopes: jiraPermissionScopes,
	}

	http.HandleFunc("/", handleRoot(conf))
	http.HandleFunc("/callback", handleCallback(conf))
	http.HandleFunc("/projects", handleProjects())
	port := os.Getenv("PORT")
	if port == "" {
		port = "9080"
	}

	log.Printf("Server started on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleRoot(conf *oauth2.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := conf.AuthCodeURL("qwertyuiuo")
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}

func handleProjects() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handleToken(token)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to handle token: %s", err.Error()), http.StatusInternalServerError)
			return
		}
		err = testJiraConnection(token)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to test connection: %s", err.Error()), http.StatusInternalServerError)
			return
		}
	}
}

func handleCallback(conf *oauth2.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		var err error
		token, err = conf.Exchange(context.Background(), code)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to exchange code: %s", err.Error()), http.StatusInternalServerError)
			return
		}
	}
}

func handleToken(token *oauth2.Token) error {
	// Check if the token is expired
	if token == nil || token.Expiry.Before(time.Now()) {
		return fmt.Errorf("Token is invalid or expired")
	}

	// If the token is about to expire in the next 30 seconds, refresh it
	if time.Until(token.Expiry) < 30*time.Second {
		conf := &oauth2.Config{
			//		ClientID:     clientID,
			//		ClientSecret: clientSecret,
			RedirectURL: redirectURI,
			Endpoint: oauth2.Endpoint{
				AuthURL:  jiraAuthURL,
				TokenURL: jiraTokenURL,
			},
			Scopes: jiraPermissionScopes,
		}

		newToken, err := conf.TokenSource(context.Background(), token).Token()
		if err != nil {
			return fmt.Errorf("Failed to refresh token: %s", err.Error())
		}

		// Handle the refreshed token - store securely
		fmt.Printf("Refreshed Access Token: %s\n", newToken.AccessToken)

		// Update the original token with the refreshed token
		*token = *newToken
	}

	return nil
}

type JiraCloudIdDetail struct {
	Id        string   `json:"id,omitempty"`
	Url       string   `json:"url,omitempty"`
	Name      string   `json:"name,omitempty"`
	AvatarUrl string   `json:"avatarUrl,omitempty"`
	Scopes    []string `json:"scopes,omitempty"`
}

func testJiraConnection(token *oauth2.Token) error {
	client := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token))
	/*
		resp, err := client.Get(jiraCloudIDURL)
		if err != nil {
			return fmt.Errorf("Failed to get jira cloud id: %s", err.Error())
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("Failed to get jira cloud id with status: %s", resp.Status)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("Failed to read response: %v", err)
		}
		var det []JiraCloudIdDetail
		err = json.Unmarshal(body, &det)
		if err != nil {
			return fmt.Errorf("Failed to unmarshal body: %s, %v", body, err)
		}
		resp.Body.Close()
	*/
	//getProjectsURL := `https://9909-167-103-24-113.ngrok-free.app/` + det[0].Id + "/rest/api/2/project/search"
	getProjectsURL := `https://9909-167-103-24-113.ngrok-free.app/rest/api/2/project/SCM`
	resp, err := client.Get(getProjectsURL)
	if err != nil {
		return fmt.Errorf("Failed to test Jira connection: %s", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Jira connection test failed with status: %s", resp.Status)
	}

	return nil
}
