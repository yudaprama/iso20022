package model

// Diagnostic response from the acquirer.
type AcceptorDiagnosticResponse4 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment43 `xml:"Envt"`

	// Instructions for contacting the terminal management host.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`
}

func (a *AcceptorDiagnosticResponse4) AddEnvironment() *CardPaymentEnvironment43 {
	a.Environment = new(CardPaymentEnvironment43)
	return a.Environment
}

func (a *AcceptorDiagnosticResponse4) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}
