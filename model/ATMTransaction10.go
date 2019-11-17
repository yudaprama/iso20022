package model

// Response to the PIN management transaction. request.
type ATMTransaction10 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the reconciliation period assigned by the ATM.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// True if a completion advice has to be sent after the end of the transaction.
	CompletionRequired *TrueFalseIndicator `xml:"CmpltnReqrd,omitempty"`

	// Result of the PIN service.
	TransactionResponse *ResponseType3 `xml:"TxRspn"`

	// Sequence of actions to be performed by the ATM to complete the transaction.
	Action []*Action5 `xml:"Actn,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`

	// Maintenance command to perform on the ATM.
	Command []*ATMCommand1 `xml:"Cmd,omitempty"`
}

func (a *ATMTransaction10) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction10) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction10) SetCompletionRequired(value string) {
	a.CompletionRequired = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction10) AddTransactionResponse() *ResponseType3 {
	a.TransactionResponse = new(ResponseType3)
	return a.TransactionResponse
}

func (a *ATMTransaction10) AddAction() *Action5 {
	newValue := new(Action5)
	a.Action = append(a.Action, newValue)
	return newValue
}

func (a *ATMTransaction10) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}

func (a *ATMTransaction10) AddCommand() *ATMCommand1 {
	newValue := new(ATMCommand1)
	a.Command = append(a.Command, newValue)
	return newValue
}
