package accounts

import (
	"encoding/json"
	"io/ioutil"
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

func (cli DefaultAccountsClient) executeRequest(request http.Request, result interface{}) error {
	resp, err := cli.httpClient.Do(&request)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &result)
	return err
}
