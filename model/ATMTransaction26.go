package model

// Information about the reconciliation response.
type ATMTransaction26 struct {

	// Type of logical or physical operation on the ATM for which the counters are computed.
	TypeOfOperation *ATMOperation1Code `xml:"TpOfOpr,omitempty"`

	// Identification of the reconciliation transaction.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the reconciliation period assigned by the ATM.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId"`

	// Result of the reconciliation.
	TransactionResponse *ResponseType7 `xml:"TxRspn"`

	// Current totals of the ATM.
	ATMTotals []*ATMTotals1 `xml:"ATMTtls,omitempty"`

	// Information on the cassettes of the ATM.
	Cassette []*ATMCassette2 `xml:"Csstt,omitempty"`

	// Maintenance command to perform on the ATM.
	Command []*ATMCommand7 `xml:"Cmd,omitempty"`
}

func (a *ATMTransaction26) SetTypeOfOperation(value string) {
	a.TypeOfOperation = (*ATMOperation1Code)(&value)
}

func (a *ATMTransaction26) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction26) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction26) AddTransactionResponse() *ResponseType7 {
	a.TransactionResponse = new(ResponseType7)
	return a.TransactionResponse
}

func (a *ATMTransaction26) AddATMTotals() *ATMTotals1 {
	newValue := new(ATMTotals1)
	a.ATMTotals = append(a.ATMTotals, newValue)
	return newValue
}

func (a *ATMTransaction26) AddCassette() *ATMCassette2 {
	newValue := new(ATMCassette2)
	a.Cassette = append(a.Cassette, newValue)
	return newValue
}

func (a *ATMTransaction26) AddCommand() *ATMCommand7 {
	newValue := new(ATMCommand7)
	a.Command = append(a.Command, newValue)
	return newValue
}
