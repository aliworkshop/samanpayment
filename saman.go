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
	codes := map[int]string{
		1:    "تراکنش توسط خریدار کنسل شده است.",
		2:    "پرداخت با موفقیت انجام شده است.",
		3:    "پرداخت انجام نشد.",
		4:    "کاربر در بازه زمانی تعیین شده پاسخی ارسال نکرده است.",
		5:    "پارامتر های ارسالی نامعتبر است.",
		8:    "آدرس سرور پذیرنده نامعتبر است.",
		10:   "توکن ارسال شده یافت نشد.",
		11:   "با این شماره ترمینال فقط تراکنش های توکنی قابل پرداخت هستند.",
		12:   "شماره ترمینال ارسال شده یافت نشد.",
		-100: "لطفا ورودی متد را کنترل نمائید.",
		-101: "مدل ورودی دارای پارامترهای نامعتبر است.",
		-102: "پردازش درخواست با خطا مواجه گردید.",
		-103: "پاسخی از سرور دریافت نشد.",
		-104: "ترمینال ارسالی غیرفعال می باشد.",
		-105: "ترمینال ارسالی در سیستم موجود نمی باشد.",
		-106: "آدرس آیپی درخواستی غیرمجار می باشد.",
		-107: "امکان وریفای کردن تراکنش موردنظر وجود ندارد.",
		-108: "امکان وریفای سریع برای این ترمینال وجود ندارد.",
		-111: "امکان تایید تراکنش وجود ندارد.",
		-112: "امکان برگشت تراکنش وجود ندارد.",
	}
	return codes[statuscode]
}
