package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aliworkshop/samanpayment"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const BaseUrl = "http://localhost:8080"

func main() {

	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)

	http.HandleFunc("/connect", Connect)
	http.HandleFunc("/callback", Callback)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Connect(w http.ResponseWriter, r *http.Request) {

	timestamp := strconv.Itoa(int(time.Now().Unix()))
	randNum := strconv.Itoa(rand.Int())[:6]
	resnum := strconv.Itoa(rand.Int())[:10]
	transaction_key := rand.Int()
	amount := 100000
	mobile := "09123456789"

	callback_url := BaseUrl + "/ipg/callback?time=" + timestamp + "&rand=" + randNum + "&transaction_key=" + strconv.Itoa(transaction_key)
	saman := samanpayment.SamanConfig{
		TerminalId: 123456789,
	}
	result, err := saman.GetTokenRequest(resnum, transaction_key, amount, callback_url, mobile)
	if err != nil {
		response(w, http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	if result.Status != 1 {
		errorCode, _ := strconv.Atoi(result.ErrorCode)
		response(w, http.StatusInternalServerError, map[string]interface{}{
			"message": samanpayment.SetValue(errorCode),
		})
		return
	}
	token := result.Token

	temp, _ := ioutil.ReadFile("connect.html")
	data := map[string]interface{}{
		"base_url": BaseUrl,
		"token":    token,
	}
	t := template.Must(template.New("connect").Parse(string(temp)))
	buf := &bytes.Buffer{}
	if err := t.Execute(buf, data); err != nil {
		response(w, http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	_, _ = w.Write(buf.Bytes())
}

func Callback(w http.ResponseWriter, r *http.Request) {
	var resp samanpayment.CallbackResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusBadRequest)
		response(w, http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	//timestamp := r.URL.Query().Get("time")
	//randStr := r.URL.Query().Get("rand")
	transaction_key, _ := strconv.Atoi(r.URL.Query().Get("transaction_key"))

	if resp.Status != 2 {
		http.Error(w, "پرداخت ناموفق بود! شرح خطا : "+samanpayment.SetValue(resp.Status), http.StatusBadRequest)
		return
	} else if resp.RefNum == "" {
		http.Error(w, "رسید دیجیتال ست نشده است!", http.StatusBadRequest)
		return
	} else {
		// check resp.RefNum is exists in database
		//if true{
		//	http.Error(w, "رسید دیجیتال استفاده شده است!", http.StatusBadRequest)
		//	return
		//}
		saman := samanpayment.SamanConfig{
			TerminalId: 123456789,
		}
		res, err := saman.VerifyTransactionRequest(transaction_key, resp.RefNum)
		if err != nil {
			response(w, http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
			return
		}
		payed_amount := 100000
		if res.ResultCode != 0 {
			response(w, http.StatusBadRequest, map[string]interface{}{
				"message": "خطا در تایید پرداخت! شرح خطا : " + samanpayment.SetValue(res.ResultCode),
			})
			return
		} else if res.VerifyInfo.Amount != payed_amount {
			result, reverse_err := saman.ReverseTransactionRequest(transaction_key, resp.RefNum)
			if reverse_err != nil {
				http.Error(w, reverse_err.Error(), http.StatusBadRequest)
			}
			if result.ResultCode != 0 {
				http.Error(w, "خطا در بازگشت پول: "+result.ResultDescription, http.StatusBadRequest)
			}
			response(w, http.StatusBadRequest, map[string]interface{}{
				"message": "مبلغ پرداختی نادرست است !",
			})
			return
		} else {
			// payment is success
			response(w, http.StatusOK, map[string]interface{}{
				"message": "افزایش اعتبار با موفقیت انجام شد. \n\n مبلغ پرداخت شده :  " + strconv.Itoa(payed_amount) + " ریال \n\n کد رهگیری بانک  :  " + res.VerifyInfo.TraceNo,
			})
			return
		}
	}
}

func response(w http.ResponseWriter, statusCode int, i interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	enc := json.NewEncoder(w)
	enc.Encode(i)
}
