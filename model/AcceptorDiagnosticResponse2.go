package model

// Diagnostic response from the acquirer.
type AcceptorDiagnosticResponse2 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment17 `xml:"Envt"`

	// Instructions for contacting the terminal management host.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`
}

func (a *AcceptorDiagnosticResponse2) AddEnvironment() *CardPaymentEnvironment17 {
	a.Environment = new(CardPaymentEnvironment17)
	return a.Environment
}

func (a *AcceptorDiagnosticResponse2) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}
