package model

// Diagnostic request from an acceptor.
type AcceptorDiagnosticRequest1 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment8 `xml:"Envt"`
}

func (a *AcceptorDiagnosticRequest1) AddEnvironment() *CardPaymentEnvironment8 {
	a.Environment = new(CardPaymentEnvironment8)
	return a.Environment
}
