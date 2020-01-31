package accounts

import (
	"net/http"
	"net/url"
)

type AccountsClient interface {
	List() (Account, error)
	Create(account Account2) (Account2, error)
}

type DefaultAccountsClient struct {
	BaseURL    *url.URL
	httpClient *http.Client
}
