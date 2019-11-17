package model

// Information about the reconciliation response.
type ATMTransaction12 struct {

	// Type of logical or physical operation on the ATM for which the counters are computed.
	TypeOfOperation *ATMOperation1Code `xml:"TpOfOpr,omitempty"`

	// Identification of the reconciliation transaction.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the reconciliation period assigned by the ATM.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId"`

	// Result of the reconciliation.
	TransactionResponse *ResponseType3 `xml:"TxRspn"`

	// Maintenance command to perform on the ATM.
	Command []*ATMCommand1 `xml:"Cmd,omitempty"`
}

func (a *ATMTransaction12) SetTypeOfOperation(value string) {
	a.TypeOfOperation = (*ATMOperation1Code)(&value)
}

func (a *ATMTransaction12) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction12) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction12) AddTransactionResponse() *ResponseType3 {
	a.TransactionResponse = new(ResponseType3)
	return a.TransactionResponse
}

func (a *ATMTransaction12) AddCommand() *ATMCommand1 {
	newValue := new(ATMCommand1)
	a.Command = append(a.Command, newValue)
	return newValue
}
