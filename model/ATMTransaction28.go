package model

// Acknowledgement of the exception advice.
type ATMTransaction28 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId,omitempty"`

	// Response to the advice.
	Response *Response2Code `xml:"Rspn"`

	// Maintenance command to perform on the ATM.
	Command []*ATMCommand7 `xml:"Cmd,omitempty"`
}

func (a *ATMTransaction28) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction28) SetResponse(value string) {
	a.Response = (*Response2Code)(&value)
}

func (a *ATMTransaction28) AddCommand() *ATMCommand7 {
	newValue := new(ATMCommand7)
	a.Command = append(a.Command, newValue)
	return newValue
}
