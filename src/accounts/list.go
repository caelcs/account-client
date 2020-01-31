package accounts

import (
	"encoding/json"
	"io/ioutil"
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

	resp, err := cli.httpClient.Do(req)

	if err != nil {
		return Account{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var accounts Account
	err = json.Unmarshal(body, &accounts)
	return accounts, err
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
