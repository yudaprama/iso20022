package model

// Information related to the response of a diagnostic from an ATM manager.
type ATMDiagnosticResponse1 struct {

	// Environment of the ATM.
	Environment *ATMEnvironment6 `xml:"Envt"`

	// Maintenance command to perform on the ATM.
	Command []*ATMCommand1 `xml:"Cmd,omitempty"`
}

func (a *ATMDiagnosticResponse1) AddEnvironment() *ATMEnvironment6 {
	a.Environment = new(ATMEnvironment6)
	return a.Environment
}

func (a *ATMDiagnosticResponse1) AddCommand() *ATMCommand1 {
	newValue := new(ATMCommand1)
	a.Command = append(a.Command, newValue)
	return newValue
}
