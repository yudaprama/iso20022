package model

// Information related to the response of a diagnostic from an ATM manager.
type ATMDiagnosticResponse2 struct {

	// Environment of the ATM.
	Environment *ATMEnvironment6 `xml:"Envt"`

	// Maintenance command to perform on the ATM.
	Command []*ATMCommand7 `xml:"Cmd,omitempty"`
}

func (a *ATMDiagnosticResponse2) AddEnvironment() *ATMEnvironment6 {
	a.Environment = new(ATMEnvironment6)
	return a.Environment
}

func (a *ATMDiagnosticResponse2) AddCommand() *ATMCommand7 {
	newValue := new(ATMCommand7)
	a.Command = append(a.Command, newValue)
	return newValue
}
