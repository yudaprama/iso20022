package model

// Diagnostic request from an acceptor.
type AcceptorDiagnosticRequest4 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment42 `xml:"Envt"`
}

func (a *AcceptorDiagnosticRequest4) AddEnvironment() *CardPaymentEnvironment42 {
	a.Environment = new(CardPaymentEnvironment42)
	return a.Environment
}
