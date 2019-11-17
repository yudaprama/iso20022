package model

// Acknowledgement by the acquirer, of the completion advice of the card payment at the acceptor.
type AcceptorCompletionAdviceResponse4 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment33 `xml:"Envt"`

	// Card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransactionAdviceResponse4 `xml:"Tx"`

	// Instructions for contacting the terminal management host.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`

	// Additional information incorporated as an extension to the message.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AcceptorCompletionAdviceResponse4) AddEnvironment() *CardPaymentEnvironment33 {
	a.Environment = new(CardPaymentEnvironment33)
	return a.Environment
}

func (a *AcceptorCompletionAdviceResponse4) AddTransaction() *CardPaymentTransactionAdviceResponse4 {
	a.Transaction = new(CardPaymentTransactionAdviceResponse4)
	return a.Transaction
}

func (a *AcceptorCompletionAdviceResponse4) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}

func (a *AcceptorCompletionAdviceResponse4) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}
