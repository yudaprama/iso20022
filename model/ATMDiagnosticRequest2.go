package model

// Information related to the request of a diagnostic from an ATM..
type ATMDiagnosticRequest2 struct {

	// Environment of the ATM.
	Environment *ATMEnvironment9 `xml:"Envt"`

	// Global status of the ATM.
	ATMGlobalStatus *ATMStatus1 `xml:"ATMGblSts"`
}

func (a *ATMDiagnosticRequest2) AddEnvironment() *ATMEnvironment9 {
	a.Environment = new(ATMEnvironment9)
	return a.Environment
}

func (a *ATMDiagnosticRequest2) AddATMGlobalStatus() *ATMStatus1 {
	a.ATMGlobalStatus = new(ATMStatus1)
	return a.ATMGlobalStatus
}
