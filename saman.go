package samanpayment

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
)

func (config *SamanConfig) GetTokenRequest(resnum string, transaction_key int, amount int, callback_url string, payer_phone string) (*RequestTokenResponse, error) {
	parameters := RequestToken{
		Action:              "token",
		TerminalId:          config.TerminalId,
		RedirectUrl:         callback_url,
		TxnRandomSessionKey: transaction_key,
		ResNum:              resnum,
		Amount:              amount,
		CellNumber:          payer_phone,
	}

	payload, err := json.Marshal(parameters)

	req, err := http.NewRequest(http.MethodPost, "https://sep.shaparak.ir/onlinepg/onlinepg", bytes.NewReader(payload))
	if err != nil {
		log.Fatal("Error on creating request object. ", err.Error())
		return nil, err
	}
	req.Header.Set("Content-type", "application/json")

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Error on dispatching request. ", err.Error())
		return nil, err
	}

	result := new(RequestTokenResponse)
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		log.Fatal("Error on unmarshaling json. ", err.Error())
		return nil, err
	}

	return result, nil
}
func (config *SamanConfig) VerifyTransactionRequest(transaction_key int, refnum string) (*VerifyTransactionResponse, error) {

	parameters := VerifyTransaction{
		RefNum:              refnum,
		TerminalId:          config.TerminalId,
		TxnRandomSessionKey: transaction_key,
		IgnoreNationalcode:  true,
	}
	payload, err := json.Marshal(parameters)

	req, err := http.NewRequest("POST", "https://sep.shaparak.ir/verifyTxnRandomSessionkey/ipg/VerifyTranscation", bytes.NewReader(payload))
	if err != nil {
		log.Fatal("Error on creating request object. ", err.Error())
		return nil, err
	}
	req.Header.Set("Content-type", "application/json")
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Error on dispatching request. ", err.Error())
		return nil, err
	}

	result := new(VerifyTransactionResponse)
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		log.Fatal("Error on unmarshaling json. ", err.Error())
		return nil, err
	}
	return result, nil
}
func (config *SamanConfig) ReverseTransactionRequest(transaction_key int, refnum string) (*ReverseTransactionResponse, error) {
	parameters := VerifyTransaction{
		RefNum:              refnum,
		TerminalId:          config.TerminalId,
		TxnRandomSessionKey: transaction_key,
		IgnoreNationalcode:  true,
	}
	payload, err := json.Marshal(parameters)

	req, err := http.NewRequest("POST", "https://sep.shaparak.ir/verifyTxnRandomSessionkey/ipg/ReverseTranscation", bytes.NewReader(payload))
	if err != nil {
		log.Fatal("Error on creating request object. ", err.Error())
		return nil, err
	}
	req.Header.Set("Content-type", "application/json")
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Error on dispatching request. ", err.Error())
		return nil, err
	}

	result := new(ReverseTransactionResponse)
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		log.Fatal("Error on unmarshaling json. ", err.Error())
		return nil, err
	}

	return result, nil
}

func SetValue(statuscode int) string {
	return codes[statuscode]
}
