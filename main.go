package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// run -addr https://vault.hafslundnett.io -token krgeg -role role

	args, err := getCommandLineArguments(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	token := &gitToken{Token: args.githubToken}
	js, err := json.Marshal(token)
	if err != nil {
		log.Fatal(err)
	}
	body := bytes.NewBuffer(js)

	url := args.addr + "/v1/auth/github/login"

	req, err := http.NewRequest(http.MethodPost, url, body)

	client, err := NewClient()
	if err != nil {
		log.Fatal(err)
	}

	var clientToken Token
	if err = client.Do(req, &clientToken); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("login succesfull, token is %s", clientToken.Auth.ClientToken)
}



// gitToken holds github authentication information to be formatted to a bytes buffer
type gitToken struct {
	Token string `json:"token"`
}

// Token is used for authenticating Vault requests
type Token struct {
	RequestID     string `json:"request_id"`
	LeaseID       string `json:"lease_id"`
	LeaseDuration int    `json:"lease_duration"`
	Renewable     bool   `json:"renewable"`
	Auth          Auth   `json:"auth"`
	//CreatedAt     time.Time
}

// Auth contains the token information for authenticating Vault requests
type Auth struct {
	ClientToken   string                 `json:"client_token"`
	TokenType     string                 `json:"token_type"`
	Accessor      string                 `json:"accessor"`
	EntityID      string                 `json:"entity_id"`
	LeaseDuration int                    `json:"lease_duration"`
	Renewable     bool                   `json:"renewable"`
	Orphan        bool                   `json:"orphan"`
	Policies      []string               `json:"policies"`
	TokenPolicies []string               `json:"token_policies"`
	Metadata      map[string]interface{} `json:"metadata"`
}