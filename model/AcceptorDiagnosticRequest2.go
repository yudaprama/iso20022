package model

// Diagnostic request from an acceptor.
type AcceptorDiagnosticRequest2 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment17 `xml:"Envt"`
}

func (a *AcceptorDiagnosticRequest2) AddEnvironment() *CardPaymentEnvironment17 {
	a.Environment = new(CardPaymentEnvironment17)
	return a.Environment
}
