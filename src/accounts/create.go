package accounts

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (cli DefaultAccountsClient) Create(account Account2) (Account2, error) {
	reqBody, err := json.Marshal(account)
	rel := &url.URL{Path: "/v1/organisation/accounts"}
	u := cli.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(reqBody))
	if err != nil {
		return Account2{}, err
	}

	req.Header.Set("Accept", "application/vnd.api+json")

	resp, err := cli.httpClient.Do(req)

	if err != nil {
		return Account2{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var accounts Account2
	err = json.Unmarshal(body, &accounts)
	return accounts, err
}

type Attributes2 struct {
	Attributes
	Title                          string   `json:"title"`
	First_name                     string   `json:"first_name"`
	Bank_account_name              string   `json:"bank_account_name"`
	Alternative_bank_account_names []string `json:"alternative_bank_account_names"`
}

type AccountDetail2 struct {
	Type            string      `json:"type"`
	Id              string      `json:"id"`
	Organisation_id string      `json:"organisation_id"`
	Version         int         `json:"version"`
	Attributes      Attributes2 `json:"attributes"`
}

type Account2 struct {
	Data []AccountDetail2 `json:"data"`
}
