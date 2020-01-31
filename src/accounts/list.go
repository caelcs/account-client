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

type Attributes struct {
	Country                  string `json:"country"`
	Base_currency            string `json:"base_currency"`
	Account_number           string `json:"account_number"`
	Bank_id                  string `json:"bank_id"`
	Bank_id_code             string `json:"bank_id_code"`
	Bic                      string `json:"bic"`
	Iban                     string `json:"iban"`
	Account_classification   string `json:"account_classification"`
	Join_account             bool   `json:"join_account"`
	Account_matching_opt_out bool   `json:"account_matching_opt_out"`
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
