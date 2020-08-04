package samanpayment

type SamanConfig struct {
	TerminalId int `json:"TerminalId"`
}

type RequestToken struct {
	Action              string `json:"action"`
	TerminalId          int    `json:"TerminalId"`
	RedirectUrl         string `json:"RedirectUrl"`
	TxnRandomSessionKey int    `json:"TxnRandomSessionKey"`
	ResNum              string `json:"ResNum"`
	Amount              int    `json:"Amount"`
	CellNumber          string `json:"CellNumber"`
}
type RequestTokenResponse struct {
	Status           int    `json:"status"`
	ErrorCode        string `json:"errorCode"`
	ErrorDescription string `json:"errorDesc"`
	Token            string `json:"token"`
}
type CallbackResponse struct {
	MID        string `json:"MID"`
	State      string `json:"State"`
	Status     int    `json:"Status"`
	RRN        string `json:"RRN"`
	RefNum     string `json:"RefNum"`
	ResNum     string `json:"ResNum"`
	TerminalId int    `json:"TerminalId"`
	TraceNo    string `json:"TraceNo"`
	Amount     int    `json:"Amount"`
	Wage       string `json:"Wage"`
	SecurePan  string `json:"SecurePan"`
}
type VerifyTransaction struct {
	RefNum              string `json:"RefNum"`
	TerminalId          int    `json:"TerminalNumber"`
	TxnRandomSessionKey int    `json:"TxnRandomSessionKey"`
	IgnoreNationalcode  bool   `json:"IgnoreNationalcode"`
}
type VerifyInfo struct {
	ReferenceNumber string `json:"RRN"`
	RefNum          string `json:"RefNum"`
	MaskedPan       string `json:"MaskedPan"`
	HashedPan       string `json:"HashedPan"`
	TerminalId      int    `json:"TerminalNumber"`
	Amount          int    `json:"OrginalAmount"`
	AffectiveAmount int    `json:"AffectiveAmount"`
	StraceDate      string `json:"StraceDate"`
	TraceNo         string `json:"StraceNo"`
}
type VerifyTransactionResponse struct {
	VerifyInfo        VerifyInfo `json:"TransactionDetail"`
	ResultCode        int        `json:"ResultCode"`
	ResultDescription string     `json:"ResultDescription"`
	Success           bool       `json:"Success"`
}
type ReverseTransaction struct {
	RefNum              string `json:"RefNum"`
	TerminalId          int    `json:"TerminalNumber"`
	TxnRandomSessionKey int    `json:"TxnRandomSessionKey"`
	IgnoreNationalcode  bool   `json:"IgnoreNationalcode"`
}
type ReverseTransactionResponse struct {
	VerifyInfo        VerifyInfo `json:"TransactionDetail"`
	ResultCode        int        `json:"ResultCode"`
	ResultDescription string     `json:"ResultDescription"`
	Success           bool       `json:"Success"`
}
