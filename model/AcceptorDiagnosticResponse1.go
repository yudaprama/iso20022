package model

// Diagnostic response from the acquirer.
type AcceptorDiagnosticResponse1 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment8 `xml:"Envt"`

	// Instructions for contacting the terminal management host.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`
}

func (a *AcceptorDiagnosticResponse1) AddEnvironment() *CardPaymentEnvironment8 {
	a.Environment = new(CardPaymentEnvironment8)
	return a.Environment
}

func (a *AcceptorDiagnosticResponse1) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}
