package model

// Diagnostic request from an acceptor.
type AcceptorDiagnosticRequest3 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment29 `xml:"Envt"`
}

func (a *AcceptorDiagnosticRequest3) AddEnvironment() *CardPaymentEnvironment29 {
	a.Environment = new(CardPaymentEnvironment29)
	return a.Environment
}
