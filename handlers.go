package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func (conf *Config) GetCurrencyPriceInfo(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	symbol := vars[SYMBOL]

	if symbol != "all" {
		currencyURL := conf.CryptoURL + "/currency/" + symbol
		resp1, err := call("GET", currencyURL, nil)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "error occured in getting currency info")
			return
		}
		defer resp1.Body.Close()

		if resp1.StatusCode != http.StatusOK {
			writeError(w, http.StatusBadRequest, "not a valid currency")
			return
		}
	}

	arg := "all"
	if symbol != "all" {
		arg = symbol
	}
	resp, err := conf.GetAllCurrencies(w, r, arg)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error occured in getting currency information")
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)

	if symbol == "all" {
		json.NewEncoder(w).Encode(resp)
	} else {
		json.NewEncoder(w).Encode(resp[0])
	}
}

//GetAllCurrencies - to get currency details
func (conf *Config) GetAllCurrencies(w http.ResponseWriter, r *http.Request, symbol string) ([]Response, error) {
	var response []Response

	currencyPricesURL := conf.CryptoURL + "/ticker"

	q := make(map[string]string)
	if symbol == "all" {
		q["symbols"] = ALL
	} else {
		symbolX := "BTCUSD"

		if symbol == "ETH" {
			symbolX = "ETHBTC"
		}
		q["symbols"] = symbolX
	}

	resp, err1 := call("GET", currencyPricesURL, q)
	if err1 != nil {
		return response, err1
	}

	if resp.StatusCode != http.StatusOK {
		return response, err1
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(contents, &response); err != nil {
		return response, err
	}

	finalResponse := []Response{}

	for _, v := range response {

		var currencyRes Response
		symbolX := "BTC"
		symbolCurr := "USD"

		if v.Symbol == "ETHBTC" {
			symbolX = "ETH"
			symbolCurr = "BTC"

		}

		currencyURL := conf.CryptoURL + "/currency/" + symbolX
		resp1, err := call("GET", currencyURL, nil)
		if err != nil {
			return response, err
		}

		curContents, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			return response, err
		}

		if err := json.Unmarshal(curContents, &currencyRes); err != nil {
			return response, err
		}

		v.ID = symbolX
		v.FullName = currencyRes.FullName
		v.FeeCurrency = symbolCurr
		v.Symbol = ""

		finalResponse = append(finalResponse, v)
	}

	return finalResponse, nil

}

func call(method string, url string, values map[string]string) (*http.Response, error) {

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	if values != nil {
		q := req.URL.Query()
		for key, value := range values {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}

	// do an http get
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func writeError(w http.ResponseWriter, errorCode int, errorMessage string) {
	err := APIError{
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(errorCode)
	json.NewEncoder(w).Encode(err)
}

//error object
type APIError struct {
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

type Response struct {
	ID          string `json:"id,omitempty"`
	FullName    string `json:"fullName,omitempty"`
	Ask         string `json:"ask,omitempty"`
	Bid         string `json:"bid,omitempty"`
	Last        string `json:"last,omitempty"`
	Open        string `json:"open,omitempty"`
	Low         string `json:"low,omitempty"`
	High        string `json:"high,omitempty"`
	FeeCurrency string `json:"feeCurrency,omitempty"`
	Symbol      string `json:"symbol,omitempty"`
}

type CurrencyInfo struct {
	ID       string `json:"id,omitempty"`
	FullName string `json:"fullName,omitempty"`
}
