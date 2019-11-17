package model

// Information about the reconciliation request.
type ATMTransaction25 struct {

	// Type of logical or physical operation on the ATM for which the counters are computed.
	TypeOfOperation *ATMOperation1Code `xml:"TpOfOpr,omitempty"`

	// Identification of the reconciliation transaction.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the reconciliation period.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId"`

	// Current totals of the ATM.
	ATMTotals []*ATMTotals1 `xml:"ATMTtls,omitempty"`

	// Information on the cassette of the ATM.
	Cassette []*ATMCassette2 `xml:"Csstt,omitempty"`

	// Transaction counters that are set to zero after a reconciliation with counter reinitialisation command.
	TransactionTotals []*ATMTotals3 `xml:"TxTtls,omitempty"`

	// Total number of retained cards.
	RetainedCard *Number `xml:"RtndCard,omitempty"`

	// Additional information about reconciliation.
	AdditionalTransactionInformation *Max140Text `xml:"AddtlTxInf,omitempty"`
}

func (a *ATMTransaction25) SetTypeOfOperation(value string) {
	a.TypeOfOperation = (*ATMOperation1Code)(&value)
}

func (a *ATMTransaction25) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction25) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction25) AddATMTotals() *ATMTotals1 {
	newValue := new(ATMTotals1)
	a.ATMTotals = append(a.ATMTotals, newValue)
	return newValue
}

func (a *ATMTransaction25) AddCassette() *ATMCassette2 {
	newValue := new(ATMCassette2)
	a.Cassette = append(a.Cassette, newValue)
	return newValue
}

func (a *ATMTransaction25) AddTransactionTotals() *ATMTotals3 {
	newValue := new(ATMTotals3)
	a.TransactionTotals = append(a.TransactionTotals, newValue)
	return newValue
}

func (a *ATMTransaction25) SetRetainedCard(value string) {
	a.RetainedCard = (*Number)(&value)
}

func (a *ATMTransaction25) SetAdditionalTransactionInformation(value string) {
	a.AdditionalTransactionInformation = (*Max140Text)(&value)
}
