package model

// Diagnostic response from the acquirer.
type AcceptorDiagnosticResponse3 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment31 `xml:"Envt"`

	// Instructions for contacting the terminal management host.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`
}

func (a *AcceptorDiagnosticResponse3) AddEnvironment() *CardPaymentEnvironment31 {
	a.Environment = new(CardPaymentEnvironment31)
	return a.Environment
}

func (a *AcceptorDiagnosticResponse3) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}
