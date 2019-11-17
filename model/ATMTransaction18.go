package model

// Acknowledgement of a completion advice.
type ATMTransaction18 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Response to the advice.
	Response *Response4Code `xml:"Rspn"`

	// Detail of the response.
	ResponseReason *ResultDetail4Code `xml:"RspnRsn,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`

	// Maintenance command to perform on the ATM.
	Command []*ATMCommand7 `xml:"Cmd,omitempty"`
}

func (a *ATMTransaction18) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction18) SetResponse(value string) {
	a.Response = (*Response4Code)(&value)
}

func (a *ATMTransaction18) SetResponseReason(value string) {
	a.ResponseReason = (*ResultDetail4Code)(&value)
}

func (a *ATMTransaction18) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}

func (a *ATMTransaction18) AddCommand() *ATMCommand7 {
	newValue := new(ATMCommand7)
	a.Command = append(a.Command, newValue)
	return newValue
}
