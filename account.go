package goiex

import (
	"net/http"
	"net/url"
)

// Account struct to interface with /account endpoints
type Account struct {
	iex
}

// Metadata struct
type Metadata struct {
	PayAsYouGoEnabled    bool   `json:"payAsYouGoEnabled"`
	EffectiveDate        int64  `json:"effectiveDate"`
	EndDateEffective     int64  `json:"endDateEffective"`
	SubscriptionTermType string `json:"subscriptionTermType"`
	TierName             string `json:"tierName"`
	MessageLimit         int    `json:"messageLimit"`
	MessagesUsed         int    `json:"messagesUsed"`
}

// NewAccount return new Account
func NewAccount(token, version string, base *url.URL, httpClient *http.Client) *Account {
	apiurl, err := url.Parse("account/")
	if err != nil {
		panic(err)
	}

	return &Account{
		iex{
			token:   token,
			version: version,
			url:     base,
			apiurl:  apiurl,
			client:  httpClient,
		},
	}
}

// Token return token string
func (a *Account) Token() string {
	return a.token
}

// Version return version string
func (a *Account) Version() string {
	return a.version
}

// URL return URL base
func (a *Account) URL() *url.URL {
	return a.url
}

// APIURL return APIURL
func (a *Account) APIURL() *url.URL {
	return a.apiurl
}

// Client return HTTP client
func (a *Account) Client() *http.Client {
	return a.client
}

// Metadata GET /account/metadata
func (a *Account) Metadata() (metadata *Metadata, err error) {
	err = get(a, &metadata, "metadata", nil)
	return
}