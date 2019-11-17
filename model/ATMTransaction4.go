package model

// Acknowledgement of a completion advice.
type ATMTransaction4 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Response to the withdrawal advice.
	Response *Response4Code `xml:"Rspn"`

	// Detail of the response.
	ResponseReason *ResultDetail2Code `xml:"RspnRsn,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`

	// Maintenance command to perform on the ATM.
	Command []*ATMCommand1 `xml:"Cmd,omitempty"`
}

func (a *ATMTransaction4) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction4) SetResponse(value string) {
	a.Response = (*Response4Code)(&value)
}

func (a *ATMTransaction4) SetResponseReason(value string) {
	a.ResponseReason = (*ResultDetail2Code)(&value)
}

func (a *ATMTransaction4) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}

func (a *ATMTransaction4) AddCommand() *ATMCommand1 {
	newValue := new(ATMCommand1)
	a.Command = append(a.Command, newValue)
	return newValue
}
