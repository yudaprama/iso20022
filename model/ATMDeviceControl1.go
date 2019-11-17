package model

// Information related to the control of an ATM device.
type ATMDeviceControl1 struct {

	// Environment of the transaction.
	Environment *ATMEnvironment7 `xml:"Envt"`

	// Maintenance command the ATM must perform.
	Command []*ATMCommand4 `xml:"Cmd,omitempty"`
}

func (a *ATMDeviceControl1) AddEnvironment() *ATMEnvironment7 {
	a.Environment = new(ATMEnvironment7)
	return a.Environment
}

func (a *ATMDeviceControl1) AddCommand() *ATMCommand4 {
	newValue := new(ATMCommand4)
	a.Command = append(a.Command, newValue)
	return newValue
}
