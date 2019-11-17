package model

// Information related to the request of a diagnostic from an ATM..
type ATMDiagnosticRequest1 struct {

	// Environment of the ATM.
	Environment *ATMEnvironment9 `xml:"Envt"`
}

func (a *ATMDiagnosticRequest1) AddEnvironment() *ATMEnvironment9 {
	a.Environment = new(ATMEnvironment9)
	return a.Environment
}
