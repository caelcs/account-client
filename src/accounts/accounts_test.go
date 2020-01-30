package accounts

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListAccounts(t *testing.T) {
	// When
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"data": [
			{
				"type": "accounts",
				"id": "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
				"organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
				"version": 0,
				"attributes": {
				  "country": "GB",
				  "base_currency": "GBP",
				  "account_number": "41426819",
				  "bank_id": "400300",
				  "bank_id_code": "GBDSC",
				  "bic": "NWBKGB22",
				  "iban": "GB11NWBK40030041426819",
				  "account_classification": "Personal",
				  "joint_account": false,
				  "account_matching_opt_out": false
				}
			  }
		]}`)
	}))
	defer ts.Close()

	BaseURL, _ := url.Parse(ts.URL)

	//When
	var client AccountsClient = DefaultAccountsClient{BaseURL, ts.Client()}
	body, error := client.List()

	//Then
	assert.NotNil(t, client)
	assert.Nil(t, error)
	assert.NotNil(t, body)
	assert.NotNil(t, body.Data)
	assert.NotNil(t, body.Data[0].Attributes)
	assert.Equal(t, "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc", body.Data[0].Id)
}

func TestCreateAccounts(t *testing.T) {
	// When
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"data": [
			{
				"type": "accounts",
				"id": "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
				"organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
				"version": 0,
				"attributes": {
				  "country": "GB",
				  "base_currency": "GBP",
				  "account_number": "41426819",
				  "bank_id": "400300",
				  "bank_id_code": "GBDSC",
				  "bic": "NWBKGB22",
				  "iban": "GB11NWBK40030041426819",
				  "account_classification": "Personal",
				  "joint_account": false,
				  "account_matching_opt_out": false
				}
			  }
		]}`)
	}))
	defer ts.Close()

	BaseURL, _ := url.Parse(ts.URL)

	attributes := Attributes2{}
	attributes.Attributes.Country = "GB"
	attributes.Attributes.Base_currency = "GBP"
	attributes.Attributes.Account_number = "41426819"
	attributes.Attributes.Bank_id = "400300"
	attributes.Attributes.Bank_id_code = "GBDSC"
	attributes.Attributes.Bic = "NWBKGB22"
	attributes.Attributes.Iban = "GB11NWBK40030041426819"
	attributes.Account_classification = "Personal"
	attributes.Attributes.Join_account = false
	attributes.Attributes.Account_matching_opt_out = false
	attributes.Title = "sdfsdf"
	attributes.First_name = "Aasdasd"
	attributes.Bank_account_name = "sdsadfsdf"
	attributes.Alternative_bank_account_names = nil

	accountDetail2 := AccountDetail2{Type: "accounts", Id: "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc", Organisation_id: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c", Version: 0, Attributes: attributes}
	account := Account2{Data: []AccountDetail2{accountDetail2}}

	//When
	var client AccountsClient = DefaultAccountsClient{BaseURL, ts.Client()}
	body, error := client.Create(account)

	assert.NotNil(t, client)
	assert.Nil(t, error)
	assert.NotNil(t, body)
	assert.NotNil(t, body.Data)
	assert.NotNil(t, body.Data[0].Attributes)
	assert.Equal(t, "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc", body.Data[0].Id)
}
