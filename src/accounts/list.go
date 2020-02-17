package accounts

import (
	"net/http"
	"net/url"
)

func (cli DefaultAccountsClient) List() (Account, error) {
	rel := &url.URL{Path: "/v1/organisation/accounts"}
	u := cli.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return Account{}, err
	}

	req.Header.Set("Accept", "application/vnd.api+json")

	result := Account{}
	errors := cli.executeRequest(*req, result)
	return result, errors
}

type AccountDetail struct {
	Type            string     `json:"type"`
	Id              string     `json:"id"`
	Organisation_id string     `json:"organisation_id"`
	Version         int        `json:"version"`
	Attributes      Attributes `json:"attributes"`
}

type Account struct {
	Data []AccountDetail `json:"data"`
}
