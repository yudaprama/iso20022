package model

// Diagnostic request from an acceptor.
type AcceptorDiagnosticRequest5 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment55 `xml:"Envt"`
}

func (a *AcceptorDiagnosticRequest5) AddEnvironment() *CardPaymentEnvironment55 {
	a.Environment = new(CardPaymentEnvironment55)
	return a.Environment
}
